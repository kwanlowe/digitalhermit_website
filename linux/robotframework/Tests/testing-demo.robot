*** Settings ***
Documentation          Demonstration of Robot Framework
Resource               ../Resources/network-resources.robot
Library                SSHLibrary
Library                String
Library	               OperatingSystem
# Suite Setup            Open Connection And Log In
# Suite Teardown         Close All Connections

*** Variables ***
${HOST}                mfcachebox01
${USERNAME}            admin
${PRIVKEY}             /home/klowe/.ssh/deploy

${IP_ADDRESS}          192.168.5.236
${DEFAULT_ROUTE}       192.168.5.1


*** Test Cases ***

Verify internal IP address
     Verify own IP address    ${IP_ADDRESS}

Verify default route
     Verify own default route - Linux Mint   ${DEFAULT_ROUTE}

#Retrieve URLs
#      ${CONTENTS}    | OperatingSystem.Get File  | Tests/urls
#      @{LINES}=      | Split to lines            | ${CONTENTS}
#     :FOR           | ${URL}  | IN | @{LINES}
#      \    Log   ${URL}
#      \    Run Keyword And Continue on Failure   Retrieve URL   ${URL}

# Check DNS resolves externally
#      ${CONTENTS}   OperatingSystem.Get File    Tests/hosts
#      @{LINES}=  Split to lines   ${CONTENTS}
#      :FOR    ${QUERY}    IN    @{LINES}
#      \    Log   ${QUERY}
#      \    Run Keyword And Continue on Failure   Check DNS resolves externally   ${QUERY}  ${IP_ADDRESS}
# 
# 
# Check connectivity to Appliansys management host
#         Check connectivity to Appliansys management host        ${APPLIANSYS_MGMT_HOST}  ${APPLIANSYS_MGMT_PORT}

*** Keywords ***
Open Connection And Log In
   Open Connection           ${HOST}
   Login With Public Key     ${USERNAME}     ${PRIVKEY}


Verify internal IP address
    [Documentation]    Verify that the IP address is ${IP_ADDRESS}
    [Arguments]        ${IP_ADDRESS}
    ${output}=         Execute Command  ip addr show eth0|awk '/inet /{print $2}'
    Should Contain     ${output}        ${IP_ADDRESS}


