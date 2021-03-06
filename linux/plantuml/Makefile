DOWNLOAD_URL=http://sourceforge.net/projects/plantuml/files/plantuml.jar/download
CURDIR := $(shell pwd)
GALLERY_TYPE=clean
WEB_SERVER=vm-helios-006
WEB_USER=deploy
REPORT_DIR=Gallery
TITLE1="Package Landscape"
TITLE2="Package Landscape"
TITLE3="How can I install an application on Linux?"

setup:
	mkdir -p bin tmp $(REPORT_DIR)/images
	@if test -f bin/plantuml.jar ; \
            then echo "plantuml.jar exists. Not Downloading.";\
	    else wget $(DOWNLOAD_URL) -O bin/plantuml.jar ; fi


album:	render gallery index push
local:	render render-svg gallery index 

render-all:	render-png render-svg

render:
	mkdir -p $(CURDIR)/$(REPORT_DIR)/images
	java -jar bin/plantuml.jar uml/*.pu -o $(CURDIR)/$(REPORT_DIR)/images/

render-svg:
	mkdir -p $(CURDIR)/$(REPORT_DIR)/images
	java -jar bin/plantuml.jar -tsvg uml/*_svg.pu -o $(CURDIR)/$(REPORT_DIR)/images/


gallery:
	mkdir -p $(REPORT_DIR)
	rm -f ./tmp/image_list
	$(eval IMAGES=$(shell sh -c "ls $(CURDIR)/$(REPORT_DIR)/images/*.png $(CURDIR)/$(REPORT_DIR)/images/*.svg| sed 's/.*\///'"))
	$(foreach var,$(IMAGES), \
		m4 -DIMAGENAME=$(var) \
		   -DIMAGETAG=$(var) \
		   -DTITLE1=$(TITLE1) \
		   -DTITLE2=$(TITLE2) \
		gallery_template/gridbox.m4 >> ./tmp/image_list;) 

index:
	awk 'NR==FNR { rep = (rep=="" ? "" : rep ORS) $$0; next }\
          /GRIDBOX/ { $$0 = rep }\
          { print } ' ./tmp/image_list ./gallery_template/$(GALLERY_TYPE).html_template| \
		m4 -DTITLE1=$(TITLE1) \
		   -DTITLE2=$(TITLE2) \
		   -DTITLE3=$(TITLE3) \
	 	>./$(REPORT_DIR)/index.html
	cp gallery_template/*.css $(REPORT_DIR)/
 
push:
	rsync -a $(REPORT_DIR) $(WEB_USER)@$(WEB_SERVER):/var/www/html/architecture/hcm

