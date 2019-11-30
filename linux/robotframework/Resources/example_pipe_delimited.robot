| *** Settings ***
| Library                | SSHLibrary
| Library                | BuiltIn
| Library                | OperatingSystem

| *** Keywords ***
| Retrieve URL
|     | [Documentation]    | Attempt to retrieve URL via wget
|     | [Arguments]        | ${URL}
|     | ${output}=         | Execute Command  | wget ${URL} -t 1 -T 5 -O /dev/null 2>&1
|     | Should Contain     | ${output}        | connected


| *** Keywords ***
| Verify own IP address
|     | [Documentation]   |  Verify that the IP address is ${IP_ADDRESS}
|     | [Arguments]       |  ${IP_ADDRESS}
|     | ${output}=        |  OperatingSystem.run   | ip addr show eth0|awk '/inet /{print $2}'
|     | Should Contain    |  ${output}             | ${IP_ADDRESS}

| Verify own default route - Linux Mint
|     | [Documentation]    | Verify that the default route is set to ${DEFAULT_ROUTE}
|     | [Arguments]        | ${DEFAULT_ROUTE}
|     | ${output}=         | OperatingSystem.run   | /sbin/ip route show
|     | Should Contain     | ${output}             | default via ${DEFAULT_ROUTE}


| Verify remote default route - Linux Mint
|     | [Documentation]    | Verify that the default route is set to ${DEFAULT_ROUTE}
|     | [Arguments]        | ${DEFAULT_ROUTE}
|     | ${output}=         | Execute Command       | /usr/sbin/ip route show
|     | Should Contain     | ${output}             | default via ${DEFAULT_ROUTE}


| Check connectivity to remote host and port
|     | [Documentation]    | Verify that host can connect to the remote host and port
|     | [Arguments]        | ${REMOTE_HOST}        | ${REMOTE_PORT}
|     | ${output}=         | Execute Command       | /usr/bin/ncat -zv ${REMOTE_HOST} ${REMOTE_PORT}
|     | Should Contain     | ${output}             | Connected



