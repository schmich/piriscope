#!/bin/bash

set -euo pipefail

while getopts "k:" opt; do
  case $opt in
    k) periscope_stream_key=$OPTARG;;
  esac
done

if [ -z ${periscope_stream_key+x} ]; then
  >&2 echo "You must specify a Periscope stream key with -k <key>."
  exit 1
fi

periscope_stream_url=rtmp://va.pscp.tv:80/x/$periscope_stream_key

video_width=960
video_height=540
video_bitrate=800000
video_sharpness=30
video_quality=80

# pixelformat=4 is for H.264 compressed video (see v4l2-ctl --list-formats)
v4l2-ctl --set-fmt-video=width=$video_width,height=$video_height,pixelformat=4
v4l2-ctl --set-ctrl=sharpness=$video_sharpness,compression_quality=$video_quality,video_bitrate_mode=1,video_bitrate=$video_bitrate

raspivid_cmd=(
  raspivid
  -o -                    # Write video to stdout.
  -t 0                    # Continuous capturing (no timeout).
  -w $video_width         # Output video width.
  -h $video_height        # Output video height.
  -vf -hf                 # Flip video vertically/horizontally.
  -fps 30                 # Capture video at 30 frames per second.
  -b $video_bitrate       # Capture video bitrate.
)

ffmpeg_cmd=(
  ffmpeg
  -re                     # Read from input at its native framerate. Best for real-time/streaming output.
  -f lavfi -i anullsrc    # No input audio.
  -i -                    # Use stdin for video (from raspivid).
  -acodec aac             # Use AAC codec for audio (Periscope requirement).
  -b:a 0                  # Zero audio bitrate since we have no input audio.
  -map 0:a                # Use stream 0 for audio (anullsrc).
  -map 1:v                # Use stream 1 for video (stdin).
  -f h264                 # Use H.264 codec for video (Periscope requirement).
  -vcodec copy            # Copy video data directly from input.
  -g 60                   # Keyframe interval: one keyframe every 60 frames (2 seconds for 30 fps video; Periscope requirement).
  -f flv                  # Package output in a Flash Video container (Periscope requirement).
  $periscope_stream_url   # RTMP streaming destination.
)

${raspivid_cmd[@]} | ${ffmpeg_cmd[@]}
