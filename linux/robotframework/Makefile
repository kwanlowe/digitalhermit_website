component=website
subcomponent=testing-demo
username=deploy
webhost=webserver
docpath=/var/www/html/architecture/$(component)/Tests

.PHONY: test

test:
	rm -rf Results/*
	robot -d Results Tests/$(subcomponent).robot

push:
	PATH=/usr/bin:~/bin/Archi:$$PATH \
	rsync -rvv -a Results \
	    --delete \
	    --rsync-path="mkdir -p $(docpath)/$(component)/$(subcomponent) && rsync" \
	    $(username)@$(webhost):$(docpath)/$(component)/$(subcomponent)

install-chromedriver:
	./install_chromedriver
