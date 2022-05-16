# Fiduswriter on CentOS

This is an unofficial and private-use build of Fiduswriter (https://github.com/fiduswriter).
Run at your own risk.

* Tested on:
  - CentOS8/podman
  - WSL2/podman
  - RHEL8/podman

# Usage

## Prerequisites

* Working podman installation.
* ~2G image storage (yes, portly image)
* ~5G for documents/images 

## Installation

* Clone the repo
```shell
    git clone https://github.com/kwanlowe/digitalhermit_website.git
```
* Change to base directory
```shell
    cd digitalhermit_website/linux/fiduswriter
```
* Create base diretory
```shell
    make create-dirs
```
* Initialize the installation
```shell
    make init
```
* Edit the data/configuration.py as desired (See Fiduswriter docs)
```shell
    vi data/configuration.py
```
* Create base installation
```shell
    make setup
```
* Create SuperUser
```shell
	make console
    fiduswriter createsuperuser
    (Enter username and password as requested)
    exit    
```
* Run the server
```shell
    make run
```
## Login and Usage

* Open browser to http://localhost:8000

* Sign-In (varies depending on configuration)

* Note: If you didn't setup a mail server and using local auth, you'll need to view the
  podman logs to get the login URL:

```shell
podman logs fiduswriter
```
