BINPATH=bin
SHELL=/bin/bash
PLATFORM=x86_64
OSTYPE=linux
OSName=Linux

ARCH=x86_64
SCRATCH=$(CURDIR)/tmp/$(CLUSTER_NAME)

create-dirs:
	@mkdir -p $(BINPATH)/
	@mkdir -p $(SCRATCH)
	@mkdir -p $(CURDIR)/bin/npm-global


DOWNLOAD_URL=http://sourceforge.net/projects/plantuml/files/plantuml.jar/download
CURDIR := $(shell pwd)
GALLERY_TYPE=clean
WEB_SERVER=rhldakll609
WEB_USER=deploy
GRAPHIC_TYPE=png
REPORT_DIR=Gallery
CLEAN=false
UML=template
TITLE1="HCM Project"
TITLE2="Cachebox"
TITLE3=""

setup:
	npm config set prefix '$(CURDIR)/bin/npm-global' &&  npm install serve -g
	mkdir -p bin tmp $(REPORT_DIR)/images
	@if test -f bin/plantuml.jar ; \
		then echo "plantuml.jar exists. Not Downloading.";\
	else wget $(DOWNLOAD_URL) -O bin/plantuml.jar ; fi


build-updated-base:
	podman build -t dhrl-centos:stream-20220211 -f Containerfile-stream_updated

build-plantuml-image:
	podman build -t dhrl-plantuml:0.1 -f Containerfile
