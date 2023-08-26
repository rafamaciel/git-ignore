package internal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func RemovePatternFromGitignore(repoPath, patternToRemove string) error {
	gitignorePath := filepath.Join(repoPath, GITIGNORE_FILE)

	data, err := ioutil.ReadFile(gitignorePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	var newLines []string

	for _, line := range lines {
		if strings.TrimSpace(line) != patternToRemove {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")
	err = ioutil.WriteFile(gitignorePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Removed pattern '%s' from .gitignore\n", patternToRemove)
	return nil
}
