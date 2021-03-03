provider "google" {
  project = "kubespray-rccl"
  region  = "us-central1"
  zone    = "us-central1-c"
}

data "google_client_openid_userinfo" "me" {
    }

resource "google_os_login_ssh_public_key" "cache" {
    user =  data.google_client_openid_userinfo.me.email
    key = file("/home/kwan/.ssh/kubespray.pub")
}

resource "google_compute_instance" "vm_workers" {
  allow_stopping_for_update = true
  count = "3"
  name         = "vm-worker-${count.index + 1}"
  machine_type = "n2-standard-2"
  # machine_type = "f1-micro"
  tags = ["kubespray-vm"]

  boot_disk {
    initialize_params {
      # image = "debian-cloud/debian-9"
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  network_interface {
    # A default network is created for all GCP projects
    network = google_compute_network.vpc_network.self_link
    access_config {
    }
  }
}


resource "google_compute_instance" "jumpoff" {
  name         = "vm-bastion-001"
  machine_type = "f1-micro"
  tags = ["kubespray-vm"]

  boot_disk {
    initialize_params {
      # image = "debian-cloud/debian-9"
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  network_interface {
    # A default network is created for all GCP projects
    network = google_compute_network.vpc_network.self_link
    access_config {
    }
  }
}

resource "google_compute_network" "vpc_network" {
  name                    = "terraform-network"
  auto_create_subnetworks = "true"
}

resource "google_compute_firewall" "ssh-rule" {
  name = "allow-ssh"
  network = google_compute_network.vpc_network.name
  allow { protocol = "icmp" }
  allow {
    protocol = "tcp"
    ports = ["22", "80", "443", "6443","2379", "2380", "10250"]
  }
  target_tags = ["kubespray-vm"]
  source_ranges = ["73.46.216.205/32", "10.128.0.0/24"]
}
