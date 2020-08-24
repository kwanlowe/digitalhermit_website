REPORT_DIR=Gallery
DOWNLOAD_URL=http://sourceforge.net/projects/plantuml/files/plantuml.jar/download
CURDIR := $(shell pwd)
GALLERY_TYPE=clean
WEB_SERVER=rhldakll609
WEB_USER=deploy
GRAPHIC_TYPE=png
REPORT_DIR=Gallery
CLEAN=false
UML=template
TITLE1="Container Orchestration"
TITLE2="Kubernetes Diagram"
TITLE3=""



setup:
	mkdir -p bin tmp $(REPORT_DIR)/images
	@if test -f bin/plantuml.jar ; \
		then echo "plantuml.jar exists. Not Downloading.";\
		else wget $(DOWNLOAD_URL) -O bin/plantuml.jar ; fi

single:
	java -jar bin/plantuml.jar -t$(GRAPHIC_TYPE) uml/$(UML).pu -o $(CURDIR)/$(REPORT_DIR)/images/


album:	render gallery index push

clean:
	rm -rf ./$(REPORT_DIR)

server:
	cd $(REPORT_DIR) && serve -p 8000

render:
ifeq ($(CLEAN),true)
	@echo "Cleaning up..."
	rm -rf ./$(REPORT_DIR)/images
endif
	@mkdir -p $(CURDIR)/$(REPORT_DIR)/images
	java -jar bin/plantuml.jar -t$(GRAPHIC_TYPE) uml/*.pu -o $(CURDIR)/$(REPORT_DIR)/images/

gallery:
	mkdir -p $(REPORT_DIR)
	rm -f ./tmp/image_list
	$(eval IMAGES=$(shell sh -c "ls $(CURDIR)/$(REPORT_DIR)/images/*.$(GRAPHIC_TYPE)| sed 's/.*\///'"))
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
 
local: render gallery index

push: gallery index
	rsync -avh $(REPORT_DIR) $(WEB_USER)@$(WEB_SERVER):/var/www/html/architecture/hcm

