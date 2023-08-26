package internal

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func AddFileToGitignore(repoPath, filePath string) (bool, error) {
	gitignorePath := filepath.Join(repoPath, GITIGNORE_FILE)

	data, err := ioutil.ReadFile(gitignorePath)
	if err != nil {
		return true, err
	}

	// Verifica se o arquivo já está no .gitignore
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == filePath {
			return true, nil
		}
	}

	newContent := string(data) + "\n" + filePath

	err = ioutil.WriteFile(gitignorePath, []byte(newContent), 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
