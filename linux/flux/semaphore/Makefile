BINPATH=$(CURDIR)/bin
SCRATCH=$(CURDIR)/tmp
VENV=$(BINPATH)/py_venv
CERTPATH=$(CURDIR)/certs/ZscalerRootCertificate-2048-SHA256.crt
ROLES_PATH=$(CURDIR)/playbooks/roles
TARGET=rhlsapcdeploy

ENGINE=podman
SEMAPHORE_STORAGE=$(CURDIR)/data/semaphore
ANSIBLE_STORAGE=$(CURDIR)/storage/files
DB_STORAGE=$(CURDIR)/data/mysql


VAULT_DOWNLOAD_URL=https://releases.hashicorp.com/vault/1.7.1/vault_1.7.1_linux_amd64.zip
VAULT_HOST=10.18.24.140
VAULT_PORT=8200
VAULT_ADDR="https://$(VAULT_HOST):$(VAULT_PORT)"


##################################################################
create-dirs:
	mkdir -p $(BINPATH) $(ROLES_PATH)

pre-config:
	sudo dnf install python38

install-python-venv:
	virtualenv -p $$(which python3.8) $(VENV)
	$(VENV)/bin/pip config set global.cert $(CERTPATH)


install-ansible:	create-dirs
	$(VENV)/bin/pip install -r requirements.txt

install-ansible-explicit-trust:
	$(VENV)/bin/pip install --trusted-host pypi.org --trusted-host files.pythonhosted.org --upgrade pip
	$(VENV)/bin/pip install --trusted-host pypi.org --trusted-host files.pythonhosted.org -r requirements.txt

##################################################################
##  Vault auth
##################################################################
get-vault-cert:
	echo -n| \
		openssl s_client -connect ${VAULT_HOST}:${VAULT_PORT} \
		-servername ${VAULT_ADDR} 2>/dev/null|\
		openssl x509 > certs/${VAULT_HOST}.crt

vault-login:
	VAULT_CAPATH=$(CURDIR)/certs VAULT_ADDR=$(VAULT_ADDR) $(BINPATH)/vault login -method=userpass username=deploy

.ONESHELL:
install-vault-cli:      create-dirs
	wget -P tmp/ $(VAULT_DOWNLOAD_URL)
	export VAULT_BASENAME=$(shell sh -c "echo $(VAULT_DOWNLOAD_URL)|sed -e 's:^.*/::'")
	unzip -p $(SCRATCH)/$(VAULT_BASENAME) vault >$(BINPATH)/vault
	chmod +x $(BINPATH)/vault


##################################################################
## Semaphore Container startup
##################################################################


