package lib

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// VttToText will read a VTT file and convert to text
func VttToText(VttFileName string) (err error) {

	fp, err := os.Open(VttFileName)
	if err != nil {
		return
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var TextFileName string
	TextFileName = strings.Replace(VttFileName, ".VTT", ".TXT", 1)
	var LastLine string
	NewLine := true
	var ConvertedText string

	for scanner.Scan() {
		LastLine = scanner.Text()
		// Empty Line
		if regexp.MustCompile(`^$`).Match([]byte(LastLine)) {
			continue
		}
		// WEBVTT Line
		if regexp.MustCompile(`^WEBVTT$`).Match([]byte(LastLine)) {
			continue
		}
		// Control Line
		if regexp.MustCompile(`^[0-9]{2}:[0-9]{2}:[0-9]{2}`).Match([]byte(LastLine)) {
			ConvertedText = ConvertedText + "\n"
			NewLine = true
			continue
		}
		if NewLine {
			ConvertedText = ConvertedText + LastLine
		} else {
			ConvertedText = ConvertedText + " " + LastLine
		}
		NewLine = false
	}

	// Write TXT file
	ioutil.WriteFile(TextFileName, []byte(ConvertedText), os.ModePerm)

	if err = scanner.Err(); err != nil {
		return
	}
	return
}
