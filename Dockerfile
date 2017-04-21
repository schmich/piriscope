FROM debian:latest
MAINTAINER Chris Schmich <schmch@gmail.com>
RUN apt-get update
 && apt-get install -y git dh-make dpkg-dev debhelper devscripts
COPY . /src
WORKDIR /src
CMD ["/bin/bash", "/src/build-linux.sh"]
