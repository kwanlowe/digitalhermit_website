#!/bin/bash

# Utility script to install latest chromedriver binary for currently installed Chrome
# http://www.digitalhermit.com/linux/

CHROMEDRIVER_URL=https://chromedriver.storage.googleapis.com
CHROME_VERSION=$(google-chrome --version)
CHROME_VERSION_MAJOR=$(echo $CHROME_VERSION|awk -F '[. /]*'  '{print $3}' )
CHROMEDRIVER_VERSION=$(curl -s ${CHROMEDRIVER_URL}/LATEST_RELEASE_${CHROME_VERSION_MAJOR})
CHROMEDRIVER_BINARY=${CHROMEDRIVER_URL}/${CHROMEDRIVER_VERSION}/chromedriver_linux64.zip
curl -o bin/chromedriver.zip "${CHROMEDRIVER_BINARY}"
unzip -d bin/ bin/chromedriver.zip
chmod +x bin/chromedriver 
