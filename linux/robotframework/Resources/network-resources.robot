*** Settings ***
Library                SSHLibrary
Library                BuiltIn
Library                OperatingSystem

*** Keywords ***
Retrieve URL
    [Documentation]    Attempt to retrieve URL via wget
    [Arguments]        ${URL}
    ${output}=         Execute Command  wget ${URL} -t 1 -T 5 -O /dev/null 2>&1
    Should Contain     ${output}        connected

Verify own IP address
    [Documentation]    Verify that the IP address is ${IP_ADDRESS}
    [Arguments]        ${IP_ADDRESS}
    ${output}=         OperatingSystem.run  ip addr show eth0|awk '/inet /{print $2}'
    Should Contain     ${output}        ${IP_ADDRESS}

Verify remote IP address
    [Documentation]    Verify that the IP address is ${IP_ADDRESS}
    [Arguments]        ${IP_ADDRESS}
    ${output}=         Execute Command  ip addr show eth0|awk '/inet /{print $2}'
    Should Contain     ${output}        ${IP_ADDRESS}

Verify own default route - Linux Mint
    [Documentation]    Verify that the default route is set to ${DEFAULT_ROUTE}
    [Arguments]        ${DEFAULT_ROUTE}
    ${output}=         OperatingSystem.run	/sbin/ip route show
    Should Contain     ${output}        default via ${DEFAULT_ROUTE}


Verify remote default route - Linux Mint
    [Documentation]    Verify that the default route is set to ${DEFAULT_ROUTE}
    [Arguments]        ${DEFAULT_ROUTE}
    ${output}=         Execute Command  /usr/sbin/ip route show
    Should Contain     ${output}        default via ${DEFAULT_ROUTE}


Check connectivity to remote host and port
    [Documentation]    Verify that host can connect to the remote host and port
    [Arguments]        ${REMOTE_HOST}  ${REMOTE_PORT}
    ${output}=         Execute Command  /usr/bin/ncat -zv ${REMOTE_HOST} ${REMOTE_PORT}
    Should Contain     ${output}        Connected

Check DNS query resolves to specific IP Address
    [Documentation]    Verify that DNS queries resolve to an external address
    [Arguments]        ${QUERY}  ${IP_ADDRESS}
    ${output}=         Execute Command  /usr/bin/dig +noall +short +answer ${QUERY}|tail -1
    Should Not Contain     ${output}    ${IP_ADDRESS}

Execute remote command
    [Documentation]     Execute a login with public key
    ${output}=          Login With Public Key           ${USERNAME}          ${PUBKEY}



