# Piriscope

Livestream to [Periscope](https://www.periscope.tv/) from the [Raspberry Pi](https://www.raspberrypi.org/products/).

## Hardware Setup

You will need the following:

- Raspberry Pi
- [Raspberry Pi Camera Module v2](https://www.raspberrypi.org/products/camera-module-v2/) ([shop](https://www.adafruit.com/product/3099))
- Raspberry Pi Zero v1.3 Camera Cable ([shop](https://www.adafruit.com/product/3157))
- Storage

## Software Setup

Piriscope is designed to work with [Raspbian](https://www.raspberrypi.org/downloads/raspbian/).

- [Install Raspbian](https://www.raspberrypi.org/documentation/installation/installing-images/)

### As a Raspbian Package (.deb)

```
apt-get install x264 v4l-utils
sudo modprobe bcm2835-v4l2
echo bcm2835-v4l2 | sudo tee -a /etc/modules
curl -LO https://github.com/ccrisan/motioneye/wiki/precompiled/ffmpeg_3.1.1-1_armhf.deb
dpkg -i ffmpeg_3.1.1-1_armhf.deb
curl -LO ...
dpkg -i piriscope-0.0.1-1_armhf.deb
```

### As a Docker Container

[Install Docker](https://www.raspberrypi.org/blog/docker-comes-to-raspberry-pi/) and run the container.

```
curl -sSL https://get.docker.com | sh
docker run -d --privileged --restart always -v /dev/video0:/dev/video0 schmich/piriscope:1.0.0 -k <key>
```

### As a Standalone Program

```
apt-get install x264 v4l-utils
sudo modprobe bcm2835-v4l2
echo bcm2835-v4l2 | sudo tee -a /etc/modules
curl -LO https://github.com/ccrisan/motioneye/wiki/precompiled/ffmpeg_3.1.1-1_armhf.deb
dpkg -i ffmpeg_3.1.1-1_armhf.deb
curl -LO ...
piriscope -k ...
```

## License

Copyright &copy; 2017 Chris Schmich  
MIT License. See [LICENSE](LICENSE) for details.
