FROM ubuntu:latest
MAINTAINER Justin Marney <gotascii@gmail.com>

ENV TERRAFORM_VERSION=0.7.0

RUN apt-get update
RUN apt-get install -y --force-yes wget unzip

ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip ./

RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /bin
RUN rm -f terraform_${TERRAFORM_VERSION}_linux_amd64.zip
