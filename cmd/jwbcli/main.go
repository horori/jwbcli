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
	dataE, err := jwapi.ParseLatestVideo("E")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Leatest VideoCnt (English)
	mediacnt := dataE.Pagination.TotalCount
	if mediacnt > 20 {
		mediacnt = 20
	}
	fmt.Println("LatestVideos :", mediacnt)

	// Show Menu
	for i := 0; i < mediacnt; i++ {
		fmt.Println("[", i, "]", dataE.Category.Media[i].Title)
	}

	// Select Media
	fmt.Printf("Choose media number : ")
	mediaNumber := clihelper.ChooseNom(0)

	for mediaNumber >= mediacnt {
		fmt.Printf("Number is wrong! Choose video media again : ")
		mediaNumber = clihelper.ChooseNom(0)
	}

	naturalKey := dataE.Category.Media[mediaNumber].LanguageAgnosticNaturalKey
	fmt.Println("No", mediaNumber, dataE.Category.Media[mediaNumber].Title, naturalKey, " is selected.")

	// input quality.
	fmt.Printf("[ 0 ] 240p [ 1 ] 360p [ 2 ] 480p (default) [ 3 ] 720p : ")
	quaNum := clihelper.ChooseNom(2)

	fmt.Println("Available Language: ", strings.Join(dataE.Category.Media[mediaNumber].AvailableLanguages, ", "))

	// input subtitle language
	fmt.Printf("For subtitle select one of available language (eg. English=E, Japanese=J, German=X) : ")
	subTitleLang := clihelper.StrStdin()

	// Search VTT file
	// Load Latest Videos
	dataJ, err := jwapi.ParseLatestVideo(subTitleLang)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	vttURL := ""
	for i := 0; i < dataJ.Pagination.TotalCount; i++ {
		if dataJ.Category.Media[i].LanguageAgnosticNaturalKey == naturalKey {
			vttURL = dataJ.Category.Media[i].Files[0].Subtitles.URL
			break
		}
	}
	if vttURL == "" {
		fmt.Println("No Textdata!")
	} else {

		// Download Language VTT
		if err := downloader.HTTPDownload(vttURL, naturalKey+".VTT"); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(naturalKey + ".VTT is saved as the subtitle file.")

		// Convert to Text
		if err := vtt.VttToText(naturalKey + ".VTT"); err != nil {
			log.Fatal("Failed to convert Japanese VTT file to Text...")
		} else {
			fmt.Println(naturalKey + ".TXT is saved as the text file.")
		}

		// Download English VTT
		vttURL = dataE.Category.Media[mediaNumber].Files[0].Subtitles.URL
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
	videoURL := dataE.Category.Media[mediaNumber].Files[quaNum].ProgressiveDownloadURL
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
	ans := clihelper.ChooseNom(0)
	if ans == 0 {
		if vttURL == "" {
			vlc.PlayNow(naturalKey+".MP4", "")
		} else {
			vlc.PlayNow(naturalKey+".MP4", naturalKey+".VTT")
		}
	}
}
