FROM golang:1.19
RUN apt-get update && apt-get -y install wget unzip

ENV TERRAFORM_VERSION 0.12.31
ADD install-terraform.sh /install-terraform.sh
RUN ["chmod", "+x", "/install-terraform.sh"]
ENTRYPOINT /install-terraform.sh

ENV GO111MODULE=on
WORKDIR /terraform-google-conversion
COPY . .
