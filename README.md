<div align="center">
<img src="https://github.com/nao1215/serial/blob/main/docs/resources/images/serial_logo.jpg" alt="The serial command project logo was created by the [DesignEvo (online logo maker)" width="539" height="164">
</div>

# serial: Add serial number to file name
![build](https://github.com/nao1215/serial/workflows/build/badge.svg?event=push)
![test](https://github.com/nao1215/serial/workflows/test/badge.svg)
![GitHub](https://img.shields.io/github/license/nao1215/serial)
![GitHub issues](https://img.shields.io/github/issues-raw/nao1215/serial)
![GitHub last commit](https://img.shields.io/github/last-commit/nao1215/serial)

serial is a CLI command that renames files under any directory to the format "specified file name + serial number".

```Demo
$ ls
a.txt  b.txt  c.txt  d.txt  e.txt
$ serial --name demo .
Rename a.txt to demo_1.txt
Rename b.txt to demo_2.txt
Rename c.txt to demo_3.txt
Rename d.txt to demo_4.txt
Rename e.txt to demo_5.txt
```

## Contents
- [serial: Add serial number to file name](#serial-add-serial-number-to-file-name)
  - [Contents](#contents)
  - [Installation](#installation)
    - [Get serial command binary from zip or tarball.](#get-serial-command-binary-from-zip-or-tarball)
    - [Build by yourself](#build-by-yourself)
  - [Usage](#usage)
  - [Development](#development)
  - [Supported OS](#supported-os)
  - [Contact](#contact)
  - [LICENSE](#license)
  - [Credits](#credits)

## Installation
### Get serial command binary from zip or tarball.
The source code and binaries are distributed on the [Release Page](https://github.com/nao1215/serial/releases) in ZIP format or tar.gz format.
Choose the binary that suits your OS and CPU architecture.

### Build by yourself
If you don't have Golang installed on your system, install it first.  Please install according to the procedure of [Golang official document](https://golang.org/doc/install).
```Install
ʕ◔ϖ◔ʔ< This procedure can be outdated.
$ wget https://golang.org/dl/go1.15.4.linux-amd64.tar.gz
$ tar -C /usr/local -xzf https://golang.org/dl/go1.15.4.linux-amd64.tar.gz
```

If you have Golang installed, the serial command will be installed on your system by typing the following command:
```
ʕ◔ϖ◔ʔ< Install serial command.
$ go get github.com/nao1215/serial/cmd/serial

ʕ◔ϖ◔ʔ< Check if the installation was successful.
$ serial --version
serial version 0.0.3

ʕ◔ϖ◔ʔ< If you don't see the version of the serial, the installation has failed.
```

If you want to install the binary of the serial command on your system, or if you want to install Man-pages, follow the steps below.
```
ʕ◔ϖ◔ʔ< Install dependency software.
$ sudo apt install gzip pandoc
$ go get github.com/golang/dep/cmd/dep

ʕ◔ϖ◔ʔ< Get serial source code from GitHub.
$ git clone https://github.com/nao1215/serial.git

ʕ◔ϖ◔ʔ< Build serial command and man-pages, and install them.
$ cd serial
$ make
$ make doc
$ sudo make install
```

## Usage
```Usage
$ serial [OPTIONS] DIRECTORY_PATH
```

| short option | long option | description |
|:------|:-----|:------|
| -d    | --dry-run    | Output the file renaming result to standard output (do not update the file)　   |
| -f   | --force    | Forcibly overwrite and save even if a file with the same name exists　   |
| -h   | --help    | Show help message　   |
| -k   | --keep    | Keep the file before renaming　   |
| -n | --name   | Base file name with/without directory path (assign a serial number to this file name)   |
| -p | --prefix   | Add a serial number to the beginning of the file name  |
| -s | --suffix  | Add a serial number to the end of the file name(default) |
| -v | --version  | Show serial command version |

## Development

If you want to contribute to the serial command, get the source code with the following command.
```
$ git clone https://github.com/nao1215/serial.git
```

The table below shows the tools used when developing the serial command.
| Tool | description |
|:-----|:------|
| dep   | Used for managing dependencies for Go projects|
| gobump   | Used for serial command version control |
| pandoc   | Convert markdown files to manpages |
| make   | Used for build, run, test, etc |
| gzip   | Used for compress man pages |
| install   | Used for install serial binary and document in the system |

## Supported OS
Currently, developers are only testing in a Linux (Debian, amd64) environment.
The cross-compiled binaries for each OS are placed on the [Release Page](https://github.com/nao1215/serial/releases) .
So, if you want to try it, please use them. If you find a bug, feel free to report it.

| OS | Architecture|Binary| Test |
|:-----|:------|:------|:------|
| AIX  | ppc64 | Released| Untested |
| Android | i386 | Unreleased| Untested |
| Android | amd64 | Unreleased| Untested |
| Android | arm | Unreleased| Untested |
| Android | arm64 | Released| Untested |
| Mac | i386 | Unreleased| Untested |
| Mac | amd64 | Released| Untested |
| Mac | arm | Unreleased| Untested |
| Mac | arm64 | Unreleased| Untested |
| Dragonfly | amd64 | Released| Untested |
| FreeBSD | i386 | Released| Untested |
| FreeBSD | amd64 | Released| Untested |
| FreeBSD | arm | Released| Untested |
| Illumos | amd64 | Released| Untested |
| Js | wasm | Unreleased| Untested |
| Linux | i386 | Released| Untested |
| Linux | amd64 | Released| Tested |
| Linux | arm | Released| Untested |
| Linux | arm64 | Released| Untested |
| Linux | ppc64 | Released| Untested |
| Linux | ppc64le | Released| Untested |
| Linux | mips | Released| Untested |
| Linux | mipsle | Released| Untested |
| Linux | mips64 | Released| Untested |
| Linux | mips64le | Released| Untested |
| Linux | ms390xips | Released| Untested |
| NetBSD | i386 | Released| Untested |
| NetBSD | amd64 | Released| Untested |
| NetBSD | arm | Released| Untested |
| OpenBSD | i386 | Released| Untested |
| OpenBSD | amd64 | Released| Untested |
| OpenBSD | arm | Released| Untested |
| OpenBSD | amd64 | Released| Untested |
| Plan9 | i386 | Released| Untested |
| Plan9 | amd64 | Released| Untested |
| Plan9 | arm | Released| Untested |
| Solaris | amd64 | Released| Untested |
| Windows | i386 | Unreleased| Untested (Seems to not work) |
| Windows | amd64 | Unreleased| Untested (Seems to not work) |

## Contact

If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/serial/issues)
- [mail@Naohiro CHIKAMATSU](n.chika156@gmail.com)
- [Twitter@ARC_AED](https://twitter.com/ARC_AED)

## LICENSE
This project is licensed under the terms of the [MIT license](./LICENSE).

##  Credits
The serial command project logo was created by the [DesignEvo (online logo maker)](https://www.designevo.com/).