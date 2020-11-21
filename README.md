{: align="center"}
![The serial command project logo was created by the [DesignEvo (online logo maker)](https://www.designevo.com/).](./docs/resources/image/serial_logo.jpg)

# serial: Add serial number to file name
![build](https://github.com/nao1215/serial/workflows/build/badge.svg?event=push)
![test](https://github.com/nao1215/serial/workflows/test/badge.svg)
![GitHub](https://img.shields.io/github/license/nao1215/serial)

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
  - [Usage](#usage)
  - [Development](#development)
  - [Contact](#contact)
  - [LICENSE](#license)
  - [Credits](#credits)

## Installation
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
| gobump   | Used for serial command version control |
| ronn   | Convert markdown files to manpages |

## Contact

If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/serial/issues)
- [mail@Naohiro CHIKAMATSU](n.chika156@gmail.com)
- [Twitter@ARC_AED](https://twitter.com/ARC_AED)

## LICENSE
This project is licensed under the terms of the [MIT license](./LICENSE).

##  Credits
The serial command project logo was created by the [DesignEvo (online logo maker)](https://www.designevo.com/).