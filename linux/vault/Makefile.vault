SCRATCH=tmp
BINPATH=bin
VAULT_DOWNLOAD_URL=https://releases.hashicorp.com/vault/1.7.1/vault_1.7.1_linux_amd64.zip
ENGINE=podman
TLS_HOSTNAME=cerberus.digitalhermit.com
VAULT_HOST=10.14.82.10
VAULT_PORT=8200
BASEDIR=$(CURDIR)/data/vault
SECRETS_DIR=$(BASEDIR)/secrets
CONF_DIR=$(CURDIR)/conf/vault

create-dirs:
	mkdir -p $(BASEDIR) $(SECRETS_DIR)
	mkdir -p $(BASEDIR)/{config,file,logs}
	mkdir -p secrets

create-config:	create-dirs
	openssl req -nodes -x509 -days 365 -keyout $(SECRETS_DIR)/vault.key -out $(SECRETS_DIR)/vault.crt -config $(CONF_DIR)/$(TLS_HOSTNAME)_cert.conf
	cp $(SECRETS_DIR)/vault.key $(SECRETS_DIR)/vault.crt $(BASEDIR)/config/
	cp $(CONF_DIR)/vault.json $(BASEDIR)/config/vault.json

start-vault:
	 $(ENGINE) run \
		--cap-add=IPC_LOCK \
		--detach \
		-p 8200:8200 \
		--name hashicorp-vault \
		--label=disable \
	 	-v $(BASEDIR)/file:/vault/file:Z \
		-v $(BASEDIR)/config:/vault/config:Z \
		-v $(BASEDIR)/logs:/vault/logs:Z \
		vault server

install-vault-cli:
	mkdir -p $(SCRATCH) $(BINPATH)
	wget -P tmp/ $(VAULT_DOWNLOAD_URL)
	@$(eval VAULT_BASENAME=$(shell sh -c "echo $(VAULT_DOWNLOAD_URL)|sed -e 's:^.*/::'"))
	unzip -p $(SCRATCH)/$(VAULT_BASENAME) vault >./$(BINPATH)/vault
	chmod +x bin/vault

.ONESHELL:
get-vault-certs:	create-dirs
	####################################################
	# Pre-stage the vault certificate
	####################################################
	export VAULT_ADDR=https://$(VAULT_HOST):${VAULT_PORT}

	echo -n| \
	openssl s_client -connect ${VAULT_HOST}:$(VAULT_PORT) \
	-servername ${VAULT_ADDR} 2>/dev/null|\
	openssl x509 > vaults/${VAULT_HOST}/ca.crt
	
vault-login:
	VAULT_CAPATH=$(CURDIR)/vaults/$(VAULT_HOST) \
	VAULT_ADDR=$(VAULT_ADDR) \
	vault login -method=userpass username=deploy


add-policies:
	vault policy write deployment-acl policies/deployment-acl.hcl
	vault policy write ocp-platform-admin policies/ocp-platform-admin.hcl
	vault policy write infrastructure-test policies/infrastructure-test.hcl
