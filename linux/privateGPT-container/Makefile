APPNAME=privategpt

REGISTRY=docker.io
NAMESPACE=kwanlowe
VERSION=0.1

build:
	podman build  -t $(APPNAME):$(VERSION) .

build-clean:
	podman build --no-cache  -t $(APPNAME):$(VERSION) .

push:
	podman tag $(APPNAME):$(VERSION) $(REGISTRY)/$(NAMESPACE)/$(APPNAME):$(VERSION) 
	podman push $(APPNAME):$(VERSION) $(REGISTRY)/$(NAMESPACE)/$(APPNAME):$(VERSION) --tls-verify=false

pull:
	podman pull $(REGISTRY)/$(NAMESPACE)/$(APPNAME):$(VERSION) --tls-verify=false

.ONESHELL:
run:
	@podman run \
		-it $(APPNAME):$(VERSION) \
		/bin/bash
