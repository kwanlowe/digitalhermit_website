VERSION=0.01
IMAGE=quay.io/digitalhermit/fiduswriter-centos:latest

build:
	podman build -t fiduswriter-centos:$(VERSION) .

create-dirs:
	mkdir -p data/fidus
	mkdir -p data/nginx

init:	create-dirs
	podman run -d \
	  --label=disable \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -v $(CURDIR)/data/fidus/:/fiduswriter:Z \
	  --rm \
	  --name fiduswriter-init \
	  $(IMAGE) startproject

setup:	create-dirs
	podman run \
	  --label=disable \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -v $(CURDIR)/data/fidus/:/fiduswriter:Z \
	  --rm \
	  --name fiduswriter-setup \
	  $(IMAGE) setup

run:	create-dirs
	podman run -d \
	  --label=disable \
	  -v $(CURDIR)/data/fidus/:/fiduswriter:Z \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -p 8000:8000 \
	  --name fiduswriter \
	  $(IMAGE)

config-proxy:	create-dirs
	m4 -DSERVER_ADDRESS=$(SERVER_ADDRESS) fidus.nginx.m4 > data/nginx/fidus

proxy:
	podman run \
	  -p 6001:80 \
	  -v $(CURDIR)/data/nginx/fidus:/etc/nginx/sites-available/fidus:Z \
	  -v $(CURDIR)/data/nginx/fidus:/etc/nginx/sites-enabled/fidus:Z \
	  --rm \
	  --label=disable \
	  --detach \
	  --name nginx-server nginx:latest




console:	create-dirs
	podman run -it \
	  --label=disable \
	  -e PROJECT_PATH=/fiduswriter \
	  -e SRC_PATH=/usr/local/lib/python3.9/site-packages/fiduswriter \
	  -v $(CURDIR)/data/fidus/:/fiduswriter:Z \
	  --rm \
	  --name fw-console \
	  --entrypoint=/bin/bash \
	  fiduswriter-centos
