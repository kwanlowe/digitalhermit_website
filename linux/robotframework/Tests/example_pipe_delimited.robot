| *** Settings ***
| Documentation         |  Demonstration of Robot Framework
| Resource              |  ../Resources/net-resources.robot
| Library               |  SSHLibrary
| Library               |  String
| Library	        |        OperatingSystem

| *** Variables ***
| ${HOST}               |  mfcachebox01
| ${USERNAME}           |  admin
| ${PRIVKEY}            |  /home/klowe/.ssh/deploy

| ${IP_ADDRESS}         |  192.168.5.236
| ${DEFAULT_ROUTE}      |  192.168.5.1


| *** Test Cases ***

| Verify internal IP address   |               |
|     | Verify own IP address  | ${IP_ADDRESS} |

| Verify own default route
|     | Verify own default route - Linux Mint  |   ${DEFAULT_ROUTE}

