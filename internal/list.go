package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func ListIgnoredFiles(repoPath string) {
	gitignorePath := filepath.Join(repoPath, GITIGNORE_FILE)

	data, err := ioutil.ReadFile(gitignorePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	ignoredPaths := []string{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			ignoredPaths = append(ignoredPaths, line)
		}
	}

	if len(ignoredPaths) == 0 {
		fmt.Println("No ignored files found.")
		return
	}

	for _, f := range ignoredPaths {
		fmt.Println(f)
	}

}
