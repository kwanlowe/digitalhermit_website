component=website
subcomponent=testing-demo
username=deploy
webhost=webserver
chromeport=4444
docpath=/var/www/html/architecture/$(component)/Tests

.PHONY: test

test:
	rm -rf Results/*
	robot -d Results Tests/$(subcomponent).robot


run-chrome-container:
	podman run \
	  -d \
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

install-chromedriver:
	./install_chromedriver
