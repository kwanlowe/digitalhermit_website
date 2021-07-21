FROM python:3
RUN python -m pip install --upgrade pip
RUN python3 -m pip install robotframework robotframework-selenium2library selenium robotframework-sshlibrary

CMD ["robot", "-d", "/Results", "/Tests/main.robot"]
