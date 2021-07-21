component=website
subcomponent=testing-demo
username=deploy
webhost=webserver
chromeport=4444
docpath=/var/www/html/architecture/$(component)/Tests
podname=robotframework-pod

WEBPORT=18080
.PHONY: test

create-dirs:
	mkdir -p data/Results data/Tests

test:
	rm -rf Results/*
	robot -d Results Tests/$(subcomponent).robot

build-container:
	podman build -t robotframework .

build-container-with-buildah:
	buildah bud -t robotframework .

create-robot-pod:
	podman pod create -n $(podname) \
	  -p 4444:4444 \
	  -p 4442:4442 \
	  -p 4443:4443 \

run-robot-test:	create-dirs
	podman run \
	  -v $(CURDIR)/data/robot/Results:/Results:Z \
	  -v $(CURDIR)/data/robot/Tests:/Tests:Z \
	  --pod $(podname) \
	  --label=disable \
	  robotframework

run-chrome-container:
	podman run \
	  -d \
	  --pod $(podname) \
	  -e SE_EVENT_BUS_PUBLISH_PORT=4442 \
	  -e SE_EVENT_BUS_SUBSCRIBE_PORT=4443 \
	  -v /dev/shm:/dev/shm \
	  --name standalone-chrome \
	  selenium/standalone-chrome:latest

run-firefox-container:
	podman run \
	  -d \
	  --pod $(podname) \
	  -e SE_EVENT_BUS_PUBLISH_PORT=4442 \
	  -e SE_EVENT_BUS_SUBSCRIBE_PORT=4443 \
	  -v /dev/shm:/dev/shm \
	  --name standalone-firefox \
	  selenium/standalone-firefox:latest


push:
	PATH=/usr/bin:~/bin/Archi:$$PATH \
	rsync -rvv -a Results \
	    --delete \
	    --rsync-path="mkdir -p $(docpath)/$(component)/$(subcomponent) && rsync" \
	    $(username)@$(webhost):$(docpath)/$(component)/$(subcomponent)

run-apache:
	podman run \
		-d \
		-p 8080:80 \
		-v $(CURDIR)/data/htdocs:/usr/local/apache2/htdocs \
		--name apache \
		httpd

run-fgallery:
	podman run \
	  --rm \
	  -v $(CURDIR)/data/fgallery:/images \
	  -p 9080:80 \
	  -e "GALLERY_TITLE=FLUX Demo" \
	  alterway/docker-image-gallery

create-secret:
	echo "FluxDemo"|podman secret create fluxdemo-secret -

run-multitool:
	podman run \
	  --secret fluxdemo-secret \
	  --detach \
	  --name multitool \
	  praqma/network-multitool

attach-multitool:
	podman exec -it multitool /bin/ash

web-server:
	cd $(CURDIR)/data/robot/Results && python3 -m http.server $(WEBPORT)

install-chromedriver:
	./install_chromedriver
