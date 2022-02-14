DOWNLOAD_URL=http://sourceforge.net/projects/plantuml/files/plantuml.jar/download
ENGINE=docker
GALLERY_TYPE=clean
WEB_SERVER=rhldakll609
WEB_USER=deploy
GRAPHIC_TYPE=png
REPORT_DIR=Gallery
CLEAN=false
UML=template
WEBPORT=8080
GALLERY_WEBPORT=18080
YAML=install-config.yaml
SCRATCH=tmp
# Watermark options
# Gravity - NorthWest, North, NorthEast, West, Center, East, SouthWest, South, SouthEast
WATERMARK=DRAFT
DENSITY=100
POINTSIZE=200
GRAVITY=center
# yaml2json converter
Y2JURL=https://github.com/bronze1man/yaml2json/releases/download/v1.3/yaml2json_linux_amd64
YQURL=https://github.com/mikefarah/yq/releases/download/v4.12.0/yq_linux_amd64
VENV=p3_venv

install-python:
	virtualenv -p $$(which python3) $(VENV)
	$(VENV)/bin/pip install yq

setup:
	mkdir -p bin tmp $(REPORT_DIR)/images
	@if test -f bin/plantuml.jar ; \
            then echo "plantuml.jar exists. Not Downloading.";\
	    else wget $(DOWNLOAD_URL) -O bin/plantuml.jar ; fi
	@if test 0f bin/yaml2json; \
			then echo "yaml2json exists. Not Downloading.";\
		else wget $(Y2JURL) -O bin/yaml2json; fi
	@if test 0f bin/yq; \
			then echo "yq exists. Not Downloading.";\
		else wget $(YQURL) -O $(CURDIR)/bin/yq; \
			chmod +x $(CURDIR)/bin/yq; fi


album:	render gallery index 

create-dirs:
	@mkdir -p $(SCRATCH)
	@mkdir -p $(CURDIR)/$(REPORT_DIR)/images/input
	@mkdir -p $(CURDIR)/$(REPORT_DIR)/images/output
	@mkdir -p $(CURDIR)/$(REPORT_DIR)/webroot/

clean:
	rm -rf ./$(REPORT_DIR)

server:
	cd $(REPORT_DIR) && python3 -m http.server $(WEBPORT)

gallery-server:
	cd $(REPORT_DIR)/images/output && python3 -m http.server $(GALLERY_WEBPORT)

single:
	java -jar bin/plantuml.jar -t$(GRAPHIC_TYPE) uml/$(UML).pu -o $(CURDIR)/$(REPORT_DIR)/images/ 
	cp $(CURDIR)/$(REPORT_DIR)/images/$(UML).* $(CURDIR)/$(REPORT_DIR)/images/output/

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

create-thumbsup-gallery:	create-dirs	render
	cp $(CURDIR)/$(REPORT_DIR)/images/*.png $(CURDIR)/$(REPORT_DIR)/images/input/
	$(ENGINE) run -t \
	-v "$(CURDIR)/$(REPORT_DIR)/images/input:/work/input:Z" \
	-v "$(CURDIR)/$(REPORT_DIR)/images/output:/work/output:Z" \
	--label=disable \
	thumbsupgallery/thumbsup \
	thumbsup --input /work/input --output /work/output


index:
	awk 'NR==FNR { rep = (rep=="" ? "" : rep ORS) $$0; next }\
          /GRIDBOX/ { $$0 = rep }\
          { print } ' ./tmp/image_list ./gallery_template/$(GALLERY_TYPE).html_template| \
		m4 -DTITLE1=$(TITLE1) \
		   -DTITLE2=$(TITLE2) \
		   -DTITLE3=$(TITLE3) \
	 	>./$(REPORT_DIR)/index.html
	cp gallery_template/*.css $(REPORT_DIR)/
 
push: gallery index
	rsync -avh $(REPORT_DIR) $(WEB_USER)@$(WEB_SERVER):/var/www/html/architecture/hcm


generate-from-install-config:
	$(VENV)/bin/yq '.' tmp/$(YAML)> tmp/input.json
	m4 json_template.m4 > tmp/output.pu

watermark-image:
	convert -density $(DENSITY) -pointsize $(POINTSIZE) -font "Helvetica-Bold" \
	    label:$(WATERMARK) $(REPORT_DIR)/images/draft-watermark.gif
	composite -dissolve 10% \
		-gravity $(GRAVITY) \
		$(REPORT_DIR)/images/draft-watermark.gif \
		$(REPORT_DIR)/images/$(UML).png \
		$(REPORT_DIR)/images/$(UML)-$(WATERMARK).png

convert-yaml-to-json:
	$(CURDIR)/bin/yq eval -o=j $(SCRATCH)/in.yaml > $(SCRATCH)/input.json