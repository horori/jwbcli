# JWBCli

CLI tool which enables you to download one of latest videos and a different language subtitles files.
You can play the videos on VLC with English audio and own language subtitle. (e.g. English audio + Japanese subtitles)

The process goes something like this:

* Download binary executable file to match your environment (Mac OS/Linux/Windows)
* Select one of latest videos
* Select resolution (240p, 360p, 480p and 720p)
* Select subtitle language

## Installation

[Download](https://github.com/horori/jwbcli/releases) latest release for your environment.

* Mac (darwin_amd64)
* Linux (linux_amd64)
* Windows (windows_amd64)

For ARM (e.g. Raspberry Pi) or 32bit OS, use this source code and compile.

```shell
env GOOS=linux GOARCH=arm GOARM=5 go build
```

Extract `.tar.gz` file or `.zip` and execute.

### Linux installation

Please replace the latest download URL

```shell
sudo curl -o /usr/local/bin/jwbcli -L <URL for tar.gz>
sudo chmod +x /usr/local/bin/jwbcli
```

### Windows installation

Extract the zip and put jwbcli.exe in a folder (e.g. desktop)
MP4 and VTT file will be downloaded to the same folder.
