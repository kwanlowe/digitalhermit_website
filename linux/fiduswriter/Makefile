VERSION=0.01

build:
	podman build -t fiduswriter-centos:$(VERSION) .

init:
	podman run -d \
	  --label=disable \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -v $(CURDIR)/data/:/fiduswriter:Z \
	  --rm \
	  --name fiduswriter-init \
	  fiduswriter-centos startproject

setup:
	podman run \
	  --label=disable \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -v $(CURDIR)/data/:/fiduswriter:Z \
	  --rm \
	  --name fiduswriter-setup \
	  fiduswriter-centos setup

run:
	podman run -d \
	  --label=disable \
	  -v $(CURDIR)/data/:/fiduswriter:Z \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -p 8000:8000 \
	  --name fiduswriter \
	  fiduswriter-centos



console:
	podman run -it \
	  --label=disable \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -v $(CURDIR)/data/:/fiduswriter:Z \
	  --rm \
	  --name fw-console \
	  --entrypoint=/bin/bash \
	  fiduswriter-centos
