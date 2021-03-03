
VENV=p3_venv
ROLES_PATH=roles
KUBESPRAY_PATH=src/kubespray
KUBESPRAY_GIT=https://github.com/kubernetes-sigs/kubespray.git
CREDFILE := /home/kwan/.ssh/anthos-playground-286117-7fc0fc56b867.json
GCLOUD_AUTH_JSON := $(shell cat ${CREDFILE})
KUBEVERSION=v1.20.2

checkout:
	mkdir -p $(KUBESPRAY_PATH)
	git clone $(KUBESPRAY_GIT) $(KUBESPRAY_PATH)

install-kubectl:
	mkdir -p bin
	curl -o bin/kubectl -LO "https://dl.k8s.io/release/$(KUBEVERSION)/bin/linux/amd64/kubectl"	
	chmod +x bin/kubectl

install-ansible:
	virtualenv -p $$(which python3) $(VENV)
	$(VENV)/bin/pip install -r $(KUBESPRAY_PATH)/requirements.txt
	mkdir -p $(ROLES_PATH)
	$(VENV)/bin/ansible-galaxy install --roles-path $(ROLES_PATH) geerlingguy.docker
	
build-box:
	vagrant up 

provision-box: 
	vagrant provision

run-test:
	mkdir -p Test_Results
	$(VENV)/bin/robot -d Test_Results Tests/client.robot


gcloud-login:
	 @docker login -u _json_key -p '$(GCLOUD_AUTH_JSON)' https://us.gcr.io

generate-kubespray-inventory:
	@echo Run the following commands in shell:
	@echo -n "    "
	@echo cd src/kubespray
	@echo -n "    "
	@gcloud compute instances list|awk 'BEGIN {printf "declare -a IPS=("} NR>1{printf "%s ", $$5} END{print ")\n"}'
	@echo -n "    "
	@echo 'CONFIG_FILE=inventory/mycluster/hosts.yaml python3 contrib/inventory_builder/inventory.py $${IPS[@]}'

