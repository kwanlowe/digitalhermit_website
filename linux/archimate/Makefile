# export PATH=~/bin/Archi:$$PATH
PROJECT=example
WEBSERVER=192.168.5.190


.PHONY: report

report:
	PATH=~/bin/Archi:$$PATH  \
	UBUNTU_MENUPROXY=0 Archi -application com.archimatetool.commandline.app \
	  -consoleLog -nosplash \
	  --loadModel $(PROJECT).archimate \
	  --html.createReport ./report/

push:
	PATH=/usr/bin:~/bin/Archi:$$PATH \
	rsync -a report kwan@$(WEBSERVER):/var/www/html/architecture/$(PROJECT)
