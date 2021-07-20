*** Settings ***
Library    Selenium2Library

Suite Teardown    Close All Browsers

*** Test Cases ***
Test CNN.com
    ${chrome_options} =     Evaluate    sys.modules['selenium.webdriver'].ChromeOptions()    sys, selenium.webdriver
    Call Method    ${chrome_options}   add_argument    headless
    Call Method    ${chrome_options}   add_argument    disable-gpu
    ${options}=     Call Method     ${chrome_options}    to_capabilities

    Create Webdriver    Remote   command_executor=http://localhost:4444/wd/hub    desired_capabilities=${options}

    Go to     http://cnn.com

    Maximize Browser Window
    Capture Page Screenshot

Test podman.io Page
    ${chrome_options} =     Evaluate    sys.modules['selenium.webdriver'].ChromeOptions()    sys, selenium.webdriver
    Call Method    ${chrome_options}   add_argument    headless
    Call Method    ${chrome_options}   add_argument    disable-gpu
    ${options}=     Call Method     ${chrome_options}    to_capabilities


	Go to 		https://podman.io/
	Capture Page Screenshot
	Click Link		xpath=/html/body/div/section/section/p[1]/a
	Capture Page Screenshot


Test Wikipedia
	Go to 		https://en.wikipedia.org/wiki/Linux
	Capture Page Screenshot
	Click Link		xpath=//*[@id="mw-content-text"]/div[1]/div[1]/a[1]
	Capture Page Screenshot
