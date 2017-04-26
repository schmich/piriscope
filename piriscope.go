package main

import (
  "os"
  "os/exec"
  "strings"
  "fmt"
  "github.com/jawher/mow.cli"
)

var version string
var commit string

func joinProps(props map[string]string, kvSeparator string, fieldSeparator string) string {
  parts := []string{}
  for key, value := range props {
    parts = append(parts, key + kvSeparator + value)
  }
  return strings.Join(parts, fieldSeparator)
}

func main() {
  app := cli.App("piriscope", "Piriscope - https://github.com/schmich/piriscope")
  key := app.StringOpt("k key", "", "Periscope stream key")

  app.Version("v version", "piriscope " + version + " " + commit)

  app.Action = func () {
    if *key == "" {
      fmt.Fprintln(os.Stderr, "Error: Periscope stream key (-k, --key) is required.")
      os.Exit(1)
    }

    periscopeStreamUrl := "rtmp://va.pscp.tv:80/x/" + *key

    videoWidth := "960"
    videoHeight := "540"
    videoSharpness := "30"
    videoQuality := "80"
    videoBitrate := "800000"

    videoProps := map[string]string{
      "width": videoWidth,
      "height": videoHeight,
      "pixelformat": "4",
    }

    controlProps := map[string]string{
      "sharpness": videoSharpness,
      "compression_quality": videoQuality,
      "video_bitrate_mode": "1",
      "video_bitrate": videoBitrate,
    }

    exec.Command("v4l2-ctl", "--set-fmt-video=" + joinProps(videoProps, "=", ","))
    exec.Command("v4l2-ctl", "--set-ctrl=" + joinProps(controlProps, "=", ","))

    raspivid := exec.Command(
      "raspivid",
      "-o", "-",
      "-t", "0",
      "-w", videoWidth,
      "-h", videoHeight,
      "-vf", "-hf",
      "-fps", "30",
      "-b", videoBitrate,
    )

    ffmpeg := exec.Command(
      "ffmpeg",
      "-re",
      "-f", "lavfi", "-i", "anullsrc",
      "-i", "-",
      "-acodec", "aac",
      "-b:a", "0",
      "-map", "0:a",
      "-map", "1:v",
      "-f", "h264",
      "-vcodec", "copy",
      "-g", "60",
      "-f", "flv",
      periscopeStreamUrl,
    )

    ffmpegStdin, err := ffmpeg.StdinPipe()
    if err != nil {
      panic(err)
    }

    raspivid.Stderr = os.Stderr
    raspivid.Stdout = ffmpegStdin
    ffmpeg.Stdout = os.Stdout
    ffmpeg.Stderr = os.Stderr

    defer func () {
      raspivid.Wait()
      ffmpegStdin.Close()
      ffmpeg.Wait()
    }()

    ffmpeg.Start()
    raspivid.Start()
  }

  app.Run(os.Args)
}
