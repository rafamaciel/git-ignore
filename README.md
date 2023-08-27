# Git Ignore

Welcome to Git Ignore! This is a command-line utility that extends Git's functionality to manage entries in the `.gitignore` file.

## Installation

You can easily install Git Ignore CLI using `curl`, and we provide direct download links for various operating systems and architectures to make it more convenient. Follow the instructions below to get started:

### Installation via Curl

#### Linux and macOS

To install Git Ignore CLI via `curl` on Linux or macOS, open the terminal and execute the following commands:

```sh
# MacOS amd64
curl -Lo git-ignore https://github.com/rafamaciel/git-ignore/releases/download/1.0.0/git-ignore-darwin-amd64 && chmod +x git-ignore && sudo mv git-ignore /usr/local/bin/
# MacOS arm64
curl -Lo git-ignore https://github.com/rafamaciel/git-ignore/releases/download/1.0.0/git-ignore-darwin-arm64 && chmod +x git-ignore && sudo mv git-ignore /usr/local/bin/
```

```sh
# Linux amd64
curl -Lo git-ignore https://github.com/rafamaciel/git-ignore/releases/download/1.0.0/git-ignore-linux-amd64 && chmod +x git-ignore && sudo mv git-ignore /usr/local/bin/
# Linux arm64
curl -Lo git-ignore https://github.com/rafamaciel/git-ignore/releases/download/1.0.0/git-ignore-linux-arm64 && chmod +x git-ignore && sudo mv git-ignore /usr/local/bin/
```

#### Windows (PowerShell)

To install Git Ignore CLI via `curl` on Windows using PowerShell, execute the following command:

```powershell
Invoke-WebRequest -Uri "https://github.com/rafamaciel/git-ignore/releases/latest/download/git-ignore-windows-amd64.exe" -OutFile "git-ignore.exe"
```

### Installation as Git Alias

After installing Git Ignore CLI, you can set it up as a Git alias to make the commands easily accessible in your repository. Execute the following command in the terminal:

```sh
git config --global alias.ignore '!git-ignore'
```

Now you can use `git ignore` to access Git Ignore CLI commands.

## Usage

Git Ignore CLI allows you to manage entries in the `.gitignore` file more conveniently. Here are some examples of how you can use Git Ignore CLI:

### Adding files to .gitignore

To add files to `.gitignore`, execute:

```sh
git ignore file1 file2 file3
```

### Removing files from .gitignore

To remove files from `.gitignore`, execute:

```sh
git ignore -f file3 file2
```

### Listing files in .gitignore

To list the files present in `.gitignore`, execute:

```sh
git ignore -l
```

## Contribution

If you wish to contribute to Git Ignore CLI, feel free to create pull requests or report issues on our [issue page](link-to-issue-page).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for more details.