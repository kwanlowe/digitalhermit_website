FROM quay.io/centos/centos:latest
RUN dnf install -y python39 python3-pip libjpeg-turbo-devel python39-devel gettext \
    zlib-devel git wget gcc-c++ make curl
RUN curl -sL https://rpm.nodesource.com/setup_14.x | bash - 
RUN dnf install -y nodejs 
RUN npm install --global typescript@3.9.10
RUN npm install --global --save-dev tap@15.0.9  
RUN python3.9 -m pip install --upgrade pip wheel setuptools
RUN mkdir fiduswriter && pip install fiduswriter fiduswriter-books fiduswriter-ojs
# RUN mkdir fiduswriter
RUN cd /u sr/local/lib/python3.9/site-packages/fiduswriter/ && python3.9 manage.py migrate
RUN dnf clean all 
ENTRYPOINT ["fiduswriter"]
CMD ["runserver"]
LABEL org.opencontainers.image.authors="kwan@digitalhermit.com"
