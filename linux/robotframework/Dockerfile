FROM python:3
# COPY config/pip.conf /root/.config/pip/pip.conf
# COPY certs/ZscalerRootCertificate-2048-SHA256.crt /certs/
RUN python -m pip install --upgrade pip
RUN python3 -m pip install robotframework robotframework-selenium2library selenium robotframework-sshlibrary
# COPY data/hello.robot /Tests/hello.robot 

# CMD ["robot", "/Tests/hello.robot"]
CMD ["robot", "-d", "/Results", "/Tests/main.robot"]
