package internal

import (
	"path/filepath"

	gitignore "github.com/sabhiram/go-gitignore"
)

func IsFileInGitignore(repoPath, filePath string) (bool, error) {
	gitignorePath := filepath.Join(repoPath, GITIGNORE_FILE)

	parser, err := gitignore.CompileIgnoreFile(gitignorePath)
	if err != nil {
		return false, err
	}

	isIgnored := parser.MatchesPath(filePath)
	return isIgnored, nil
}
