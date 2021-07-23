# Fiduswriter on CentOS

This is an unofficial and private-use build of Fiduswriter (https://github.com/fiduswriter).
Run at your own risk.

* Tested on:
  - CentOS8/podman
  - WSL2/podman
  - RHEL8/podman

# Usage

## Prerequisites

* Working podman installation.
* ~2G image storage (yes, portly image)
* ~5G for documents/images 

## Installation

* Clone the repo

    git clone <repo>

* Change to base directory

    cd digitalhermit_website/linux/fiduswriter

* Create base diretory

    make create-dirs

* Initialize the installation

    make init

* Edit the data/configuration.py as desired (See Fiduswriter docs)

    vi data/configuration.py

* Create base installation

    make setup

* Create SuperUser

	make console
    fiduswriter createsuperuser
    (Enter username and password as requested)
    exit    

* Run the server

    make run

* Open browser to http://localhost:8000


