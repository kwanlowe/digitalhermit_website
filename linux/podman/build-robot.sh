#!/bin/bash

id=$(buildah from python:3)
buildah run $id python3 -m pip install --upgrade pip
buildah run $id python3 -m pip install robotframework robotframework-selenium2library selenium robotframework-sshlibrary
buildah config --entrypoint ["robot", "-d", "/Results", "/Tests/main.robot"] $id
buildah commit $id podman-robotframework
