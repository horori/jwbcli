package vlc

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

// PlayNow play video file with VLC
func PlayNow(videoFile string, vttFile string) (err error) {
	vlcpath := "vlc"

	switch osName := runtime.GOOS; osName {
	case "windows":
		// Play on VLC
		vlcpath = os.Getenv("ProgramFiles(x86)") + "\\VideoLAN\\VLC\\vlc.exe"
		if !FileExists(vlcpath) {
			vlcpath = os.Getenv("ProgramFiles") + "\\VideoLAN\\VLC\\vlc.exe"
			if !FileExists(vlcpath) {
				fmt.Printf("Cannot find VNC player. Automatic play won't work now. Play the video manually.")
				return
			}
		}
	case "darwin":
		vlcpath = "/Applications/VLC.app/Contents/MacOS/VLC"
	case "linux":
		vlcpath = "vlc"
	default:
		fmt.Printf("%s.\n", osName)
		return
	}

	if vttFile != "" {
		err = exec.Command(vlcpath, videoFile, "--sub-file="+vttFile).Start()
	} else {
		err = exec.Command(vlcpath, videoFile).Start()
	}
	return
}
