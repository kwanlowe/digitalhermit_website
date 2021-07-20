component=website
subcomponent=testing-demo
username=deploy
webhost=webserver
chromeport=4444
docpath=/var/www/html/architecture/$(component)/Tests

WEBPORT=18080
.PHONY: test

create-dirs:
	mkdir -p data/Results data/Tests

test:
	rm -rf Results/*
	robot -d Results Tests/$(subcomponent).robot

build-container:
	podman build -t robotframework .

docker-based-test:	create-dirs
	podman run \
	  -v $(CURDIR)/data/Results:/Results:Z \
	  -v $(CURDIR)/data/Tests:/Tests:Z \
	  --pod pod-test \
	  --label=disable \
	  robotframework

web-server:
	cd $(CURDIR)/data/Results && python3 -m http.server $(WEBPORT)

run-chrome-container:
	podman run \
	  -d \
	  --pod new:pod-test \
	  -p 4444:4444 \
	  -e SE_EVENT_BUS_PUBLISH_PORT=4442 \
	  -e SE_EVENT_BUS_SUBSCRIBE_PORT=4443 \
	  -v /dev/shm:/dev/shm \
	  --name standalone-chrome \
	  selenium/standalone-chrome:latest

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

install-chromedriver:
	./install_chromedriver
