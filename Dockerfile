FROM resin/rpi-raspbian:jessie
WORKDIR /tmp
RUN apt-get update \
 && apt-get install curl x264 v4l-utils \
 && curl -sSLO https://github.com/ccrisan/motioneye/wiki/precompiled/ffmpeg_3.1.1-1_armhf.deb \
 && dpkg -i ffmpeg_3.1.1-1_armhf.deb
COPY piriscope_0.0.1-1_armhf.deb /tmp/piriscope_0.0.1-1_armhf.deb
RUN dpkg -i piriscope_0.0.1-1_armhf.deb \
 && rm /tmp/*deb
WORKDIR /
ENTRYPOINT ["/usr/bin/piriscope"]
