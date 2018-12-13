package lib

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// FileExists Check file existance
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// IsLinuxCommandAvailable check
func IsLinuxCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// PlayNow play video file with VLC
func PlayNow(videoFile string, vttFile string) (err error) {
	vlcpath := "vlc"
	if runtime.GOOS == "windows" {
		// Play on VLC
		vlcpath = os.Getenv("ProgramFiles(x86)") + "\\VideoLAN\\VLC\\vlc.exe"
		if !FileExists(vlcpath) {
			vlcpath = os.Getenv("ProgramFiles") + "\\VideoLAN\\VLC\\vlc.exe"
			if !FileExists(vlcpath) {
				fmt.Printf("Cannot find VNC player. Automatic play won't work now. Play the video manually.")
				return
			}
		}
	} else {
		if !IsLinuxCommandAvailable("vlc") {
			fmt.Printf("Cannot find VNC player on the PATH. Automatic play won't work now. Play the video manually.")
			return
		}
	}
	if vttFile != "" {
		err = exec.Command(vlcpath, videoFile, "--sub-file="+vttFile).Start()
	} else {
		err = exec.Command(vlcpath, videoFile).Start()
	}
	return
}
