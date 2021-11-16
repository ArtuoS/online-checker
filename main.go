package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	paths = map[string]string{
		"websites": "websites.txt",
		"logs":     "logs.txt",
	}
)

func main() {
	if !logFileExists() {
		createLogFile()
	}

	readFile()
}

func createLogFile() {
	os.Create(paths["logs"])
}

func logFileExists() bool {
	if _, err := os.Stat(paths["logs"]); err != nil {
		return false
	}
	return true
}

func readFile() {
	content, err := os.Open(paths["websites"])
	if err != nil {
		log.Fatal(err.Error())
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		checkIfWebsiteIsOnline(scanner.Text())
	}
}

func checkIfWebsiteIsOnline(url string) {
	_, err := http.Head(url)
	if err != nil {
		logWebsite(url, false)
		return
	}
	logWebsite(url, true)
}

func logWebsite(url string, online bool) {
	file, err := os.OpenFile(paths["logs"], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, _ = file.WriteString(concatUrlWithStatus(url, online))
	file.Close()
}

func concatUrlWithStatus(url string, online bool) string {
	urlWithStatus := url + " is "
	if !online {
		return urlWithStatus + "offline\n"
	}
	return urlWithStatus + "online\n"
}
