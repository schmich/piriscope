# Piriscope

Stream to Periscope from the Raspberry Pi.

## Setup

Raspberry Pi:

- Download/install Etcher
- Download RPi image
- Flash image to SD card
- `touch /Volumes/boot/ssh`
- `cp wpa_supplicant.conf /Volumes/boot`
- `sudo raspi-config` (enable SSH, enable camera)
- (setup SSH keys)
- (remove password-based SSH)
- (update pi user password)
- `sudo apt-get install -y vlc`
- https://www.jeffgeerling.com/blogs/jeff-geerling/raspberry-pi-zero-conserve-energy
- `sudo update-rc.d -f bluetooth remove`

Video setup (maybe unnecessary with `raspivid`):

- `sudo modprobe bcm2835-v4l2`
- `sudo echo bcm2835-v4l2 >> /etc/modules`

Server URL: rtmp://va.pscp.tv:80/x  
Stream name/key: xxxxxxxxxxxxx

http://www.kianryan.co.uk/2015/10/buliding-a-youtube-live-streaming-camera-with-a-raspberry-pi/

Periscope video requirements
  - Audio stream is required
  - FPS, resolution, video codec, bitrate

Enable raspivid, drivers, etc.

```
wget https://github.com/ccrisan/motioneye/wiki/precompiled/ffmpeg_3.1.1-1_armhf.deb
sudo dpkg -i ffmpeg_3.1.1-1_armhf.deb
raspivid -o - -t 0 -w 960 -h 540 -vf -hf -fps 30 -b 800000 |\
  ffmpeg -re -f lavfi -i anullsrc -i - -acodec aac -b:a 0 -map 0:a -map 1:v -f h264 -vcodec copy -g 60 -f flv rtmp://va.pscp.tv:80/x/xxxxxxxxxxxx
```

Periscope:

- Steps to enable stream via Periscope app
- App/phone required?
- Screencast from phone

## Configuration

- Stream name/key
- Portrait vs. landscape video
- Horizontal/vertical flip
- Bitrate (Periscope guidelines)
- FPS (Periscope guidelines)
- Other/arbitrary `raspivid`/`ffmpeg` options

## Packaging

- Bare script, roll-your-own
- Debian service
- Debian package (.deb)
- Flashable Raspbian image, configuration via `/boot`
