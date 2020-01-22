package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/horori/jwbcli/pkg/clihelper"
	"github.com/horori/jwbcli/pkg/downloader"
	"github.com/horori/jwbcli/pkg/jwapi"
	"github.com/horori/jwbcli/pkg/vlc"
	"github.com/horori/jwbcli/pkg/vtt"
)

func main() {

	// Load English Latest Videos
	dataE, err := jwapi.GetLatestVideo("E")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Show Menu (max 20)
	m := jwapi.GetLatestVideoTitles(dataE, 20)
	for i := 0; i < len(m); i++ {
		fmt.Println("[", i, "]", m[i])
	}
	fmt.Printf("Select one of available video : ")

	// Select Media
	var selectedNumber int
	for {
		selectedNumber = clihelper.ChooseNumber(0)
		if selectedNumber < len(m) {
			break
		}
		fmt.Printf("Number is wrong! Choose video media again : ")
	}

	naturalKey := jwapi.GetNaturalKey(dataE, selectedNumber)
	fmt.Println("No", selectedNumber, m[selectedNumber], naturalKey, " is selected.")

	// input quality.
	resolutionMap := jwapi.GetAvailableResolution(dataE, selectedNumber)
	for i := 0; i < len(resolutionMap); i++ {
		fmt.Println("[", i, "]", resolutionMap[i])
	}
	fmt.Printf("Select one of available resolution : ")
	resolutionNumber := clihelper.ChooseNumber(2)

	fmt.Println("Available Language: ", strings.Join(jwapi.GetAvailableLanguage(dataE, selectedNumber), ", "))

	// input subtitle language
	fmt.Printf("For subtitle select one of available language (eg. English=E, Japanese=J, German=X) : ")
	selectedLanguage := clihelper.StrStdin()

	// Search VTT file in the selected languarge
	vttURL, err := jwapi.GetVttURLByNaturalKey(naturalKey, selectedLanguage)
	if err != nil {
		log.Fatal(err)
	}

	if vttURL == "" {
		fmt.Println("No Textdata!")
	} else {

		// Download Language VTT
		if err := downloader.HTTPDownload(vttURL, naturalKey+".VTT"); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(naturalKey + ".VTT is saved as " + selectedLanguage + " subtitle file.")

		// Convert to Text
		if err := vtt.VttToText(naturalKey + ".VTT"); err != nil {
			log.Fatal("Failed to convert " + selectedLanguage + " VTT file to Text...")
		} else {
			fmt.Println(naturalKey + ".TXT is saved as " + selectedLanguage + " text file.")
		}

		// Download English VTT
		vttURL, err := jwapi.GetVttURLByNaturalKey(naturalKey, "E")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Download " + vttURL)
		if err := downloader.HTTPDownload(vttURL, naturalKey+"_E.VTT"); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(naturalKey + "_E.VTT is saved as the subtitle file.")

		// Convert to Text
		if err := vtt.VttToText(naturalKey + "_E.VTT"); err != nil {
			log.Fatal("Failed to convert English VTT file to Text...")
		} else {
			fmt.Println(naturalKey + "_E.TXT is saved as the text file.")
		}

	}

	// Find Video URL
	videoURL := jwapi.GetVideoDownloadURL(dataE, selectedNumber, resolutionNumber)
	if videoURL == "" {
		log.Fatal("No video data for this resolution!")
		os.Exit(1)
	}

	// Download Video
	fmt.Println("Videofile downloading...")
	if err := downloader.HTTPDownload(videoURL, naturalKey+".MP4"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Play Now
	fmt.Printf("Play with VLC? [ 0 ] Yes (default) [ 1 ] No : ")
	ans := clihelper.ChooseNumber(0)
	if ans == 0 {
		if vttURL == "" {
			vlc.PlayNow(naturalKey+".MP4", "")
		} else {
			vlc.PlayNow(naturalKey+".MP4", naturalKey+".VTT")
		}
	}
}
