
VENV=p3_venv
ROLES_PATH=roles
CREDFILE := /home/kwan/.ssh/anthos-playground-286117-7fc0fc56b867.json
GCLOUD_AUTH_JSON := $(shell cat ${CREDFILE})


setup:
	virtualenv -p $$(which python3) $(VENV)
	$(VENV)/bin/pip install ansible robotframework robotframework-sshlibrary
	mkdir -p $(ROLES_PATH)
	$(VENV)/bin/ansible-galaxy install --roles-path $(ROLES_PATH) geerlingguy.docker
	vagrant plugin install vagrant-disksize 
	
build-box:
	vagrant up 

provision-box: 
	vagrant provision

run-test:
	mkdir -p Test_Results
	$(VENV)/bin/robot -d Test_Results Tests/client.robot


gcloud-login:
	 @docker login -u _json_key -p '$(GCLOUD_AUTH_JSON)' https://us.gcr.io

FOO:
	@echo $(GCLOUD_AUTH_JSON)

