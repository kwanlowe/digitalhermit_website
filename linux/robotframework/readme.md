Tips:
* White space is important

# Chrome in Docker Container

  The ```run-chrome-container``` entry pulls and runs a Chrome in a container. It exposes several ports to the local machine
including 4442, 4443 and 4444. The tests then connect to this port using the WebDriver:

```
Headless Chrome - Create Webdriver
    ${chrome_options} =     Evaluate    sys.modules['selenium.webdriver'].ChromeOptions()    sys, selenium.webdriver
    Call Method    ${chrome_options}   add_argument    headless
    Call Method    ${chrome_options}   add_argument    disable-gpu
    ${options}=     Call Method     ${chrome_options}    to_capabilities

    Create Webdriver    Remote   command_executor=http://localhost:4444/wd/hub    desired_capabilities=${options}
```

 
