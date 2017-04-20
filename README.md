# Piriscope

Stream to Periscope from the Raspberry Pi.

## Setup

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

## Packaging

- Bare script, roll-your-own
- Raspbian (.deb) package
- Flashable image, configuration via `/boot`
