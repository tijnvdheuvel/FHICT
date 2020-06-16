package main

import (
	"bufio"
	"os/exec"
	"strings"
	"time"
)



func main() {
	cmd := exec.Command("ipconfig", "/all")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmd.Start()

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		tekst := scanner.Text()

		if strings.Contains(tekst, "IPv4 Address") {
			dubbelePuntIndex := strings.Index(tekst, ":")
			if dubbelePuntIndex == -1 {
				continue
			}
			geknipteTekst := tekst[dubbelePuntIndex+1:]
			geknipteTekst = strings.Trim(geknipteTekst, " \t\n")
			haakjeIndex := strings.Index(geknipteTekst, "(")
			if haakjeIndex != -1 {
				geknipteTekst = geknipteTekst[:haakjeIndex]
			}
			ipadres = geknipteTekst
		} else if strings.Contains(tekst, "Physical Address") {
			dubbelePuntIndex := strings.Index(tekst, ":")
			if dubbelePuntIndex == -1 {
				continue
			}
			geknipteTekst := tekst[dubbelePuntIndex+1:]
			geknipteTekst = strings.Trim(geknipteTekst, " \t\n")
			geknipteTekst = geknipteTekst[:macadreslengte]
			macadres = geknipteTekst
		}
	}
	cmd.Wait()

	duration := time.Duration(10)*time.Second
	time.Sleep(duration)
}
