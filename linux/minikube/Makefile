
VENV=p3_venv
ROLES_PATH=roles

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
