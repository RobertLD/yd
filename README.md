# Yank Directory (yd)

Yank Directory (yd) is a simple command-line utility that facilitates copying the current working directory to the OS clipboard. This tool is designed to streamline the process of copying directory paths, making it convenient for users who often work with terminal commands or need to share directory paths quickly.

## Features

- Copy the current working directory to the clipboard with a single command.
- Lightweight and easy to use.
- Cross-platform support (Linux, macOS, Windows).

## Installation

### Pre-built binary
Download the latest binaries from the release tab [Releases](https://github.com/RobertLD/yd/releases)


### Build from source
1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/yd.git
   ```
2. **Navigate to the directory**
   ```bash
   cd yd
   ```
2. **Build with Go v1.21**
   ```bash
   go build
   ```

## Usage
After installation, you can use yd in the terminal to copy the current working directory to the clipboard.
  ```bash
    $ cd /path/to/some/directory
    $ yd
  ```
