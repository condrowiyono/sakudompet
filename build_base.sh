#!/usr/bin/env bash
username=${1}
password=${2}
registry=${3}

requirements_hash="$(md5sum glide.lock | cut -d' ' -f1)"
curl -u ${username}:${password} https://${registry}/v2/${CI_PROJECT_NAME}/tags/list | grep ${requirements_hash}

if [ $? -ne 0 ]; then
    tag=${registry}/${CI_PROJECT_NAME}:${CI_BUILD_REF_NAME}-builder-${requirements_hash}
    echo "going to build new base image ${tag}"
    docker build -t ${tag} -f Dockerfile .
    docker push ${tag}
fi