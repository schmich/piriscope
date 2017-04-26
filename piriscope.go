package main

import (
  "os"
  "os/exec"
  "strings"
  "strconv"
  "encoding/json"
  "io/ioutil"
  "github.com/jawher/mow.cli"
  log "github.com/Sirupsen/logrus"
)

var version string
var commit string

type configuration struct {
  Periscope periscope `json:"periscope"`
  Video video `json:"video"`
}

type periscope struct {
  Key string `json:"key"`
}

type video struct {
  Width int `json:"width"`
  Height int `json:"height"`
  Sharpness int `json:"sharpness"`
  Quality int `json:"quality"`
  Bitrate int `json:"bitrate"`
}

func joinProps(props map[string]string, kvSeparator string, fieldSeparator string) string {
  parts := []string{}
  for key, value := range props {
    parts = append(parts, key + kvSeparator + value)
  }
  return strings.Join(parts, fieldSeparator)
}

func mergeString(l string, r string) string {
  if r == "" {
    return l
  } else {
    return r
  }
}

func mergeInt(l int, r int) int {
  if r == 0 {
    return r
  } else {
    return l
  }
}

func mergePeriscope(l *periscope, r *periscope) *periscope {
  return &periscope{
    Key: mergeString(l.Key, r.Key),
  }
}

func mergeVideo(l *video, r *video) *video {
  return &video{
    Width: mergeInt(l.Width, r.Width),
    Height: mergeInt(l.Height, r.Height),
    Sharpness: mergeInt(l.Sharpness, r.Sharpness),
    Quality: mergeInt(l.Quality, r.Quality),
    Bitrate: mergeInt(l.Bitrate, r.Bitrate),
  }
}

func mergeConfig(l *configuration, r *configuration) *configuration {
  return &configuration{
    Periscope: *mergePeriscope(&l.Periscope, &r.Periscope),
    Video: *mergeVideo(&l.Video, &r.Video),
  }
}

func main() {
  app := cli.App("piriscope", "Piriscope - https://github.com/schmich/piriscope")
  key := app.StringOpt("k key", "", "Periscope stream key")
  conf := app.StringOpt("c conf", "", "Configuration file")

  app.Version("v version", "piriscope " + version + " " + commit)

  app.Action = func () {
    var fileConfig configuration
    if *conf != "" {
      content, err := ioutil.ReadFile(*conf)
      if err != nil {
        log.Fatal(err)
      }

      log.WithFields(log.Fields{ "file": *conf }).Info("Using configuration file")

      err = json.Unmarshal(content, &fileConfig)
      if err != nil {
        log.Fatal(err)
      }
    }

    defaultConfig := &configuration{
      Periscope: periscope{
        Key: "",
      },
      Video: video{
        Width: 960,
        Height: 540,
        Sharpness: 30,
        Quality: 80,
        Bitrate: 800000,
      },
    }

    cliConfig := &configuration{
      Periscope: periscope{
        Key: *key,
      },
    }

    config := mergeConfig(defaultConfig, mergeConfig(&fileConfig, cliConfig))

    if config.Periscope.Key == "" {
      log.Fatal("Periscope stream key (-k, --key) is required.")
    }

    streamUrl := "rtmp://va.pscp.tv:80/x/" + config.Periscope.Key

    width := strconv.Itoa(config.Video.Width)
    height := strconv.Itoa(config.Video.Height)
    sharpness := strconv.Itoa(config.Video.Sharpness)
    quality := strconv.Itoa(config.Video.Quality)
    bitrate := strconv.Itoa(config.Video.Bitrate)

    videoProps := map[string]string{
      "width": width,
      "height": height,
      "pixelformat": "4",
    }

    controlProps := map[string]string{
      "sharpness": sharpness,
      "compression_quality": quality,
      "video_bitrate_mode": "1",
      "video_bitrate": bitrate,
    }

    exec.Command("v4l2-ctl", "--set-fmt-video=" + joinProps(videoProps, "=", ","))
    exec.Command("v4l2-ctl", "--set-ctrl=" + joinProps(controlProps, "=", ","))

    raspivid := exec.Command(
      "raspivid",
      "-o", "-",
      "-t", "0",
      "-w", width,
      "-h", height,
      "-vf", "-hf",
      "-fps", "30",
      "-b", bitrate,
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
      streamUrl,
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
