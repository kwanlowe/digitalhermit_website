
USER=kwan
USB_PORT=/dev/ttyUSB0
PYVENV=p36_venv
BOARD=esp8266
VERSION=20191220-v1.12
UPLOAD_BAUDRATE=57600
BAUDRATE=115200
FLASHSIZE=detect

setup:
	mkdir firmware 
	virtualenv -p python3.6 $(PYVENV)
	$(PYVENV)/bin/pip install esptool

add-user-to-group:
	sudo usermod -a -G dialout $(USER)

install-minicom:
	sudo dnf install minicom

download-firmware:
	wget -O firmware/$(BOARD)-$(VERSION).bin http://micropython.org/resources/firmware/$(BOARD)-$(VERSION).bin

erase-flash:
	$(PYVENV)/bin/esptool.py --port $(USB_PORT) erase_flash

deploy-firmware:
	$(PYVENV)/bin/esptool.py --port $(USB_PORT) --baud $(UPLOAD_BAUDRATE) write_flash -fm dio --flash_size=$(FLASHSIZE)  0 firmware/$(BOARD)-$(VERSION).bin


start-minicom:
	        minicom -D $(USB_PORT) -b $(BAUDRATE)

clone-webrepl:
	git clone https://github.com/micropython/webrepl.git

push-code:
	@echo Copying boot.py and main.py to $(BOARD) at $(USB_PORT)
	$(PYVENV)/bin/rshell -p $(USB_PORT) cp boot.py main.py /pyboard
