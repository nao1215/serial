#!/bin/bash -eu
# [Description]
#  This shell script is the installer for the serial command.
#  This is created to be stored in tar.gz or zip for release files.
function errMsg() {
    local message="$1"
    echo -n -e "\033[31m\c"
    echo "${message}" >&2
    echo -n -e "\033[m\c"
}

function warnMsg() {
    local message="$1"
    echo -n -e "\033[33m\c"
    echo "${message}"
    echo -n -e "\033[m\c"
}

function isRoot() {
    if [ ${EUID:-${UID}} != 0 ]; then
        errMsg "[Usage]"
        errMsg " $ sudo ./installer.sh"
        exit 1
    fi
}

function installSerialCmd() {
    echo "Install serial command at /usr/local/bin/serial"
    install -m 0755 -D ./serial /usr/local/bin/.
}

function installManPages() {
    echo "Install man-pages"
    install -m 0644 -D man/en/serial.1.gz /usr/share/man/man1/serial.1.gz
    install -m 0644 -D man/ja/serial.1.gz /usr/share/man/ja/man1/serial.1.gz
}

isRoot
echo "Start install."
installSerialCmd
installManPages
echo ""
warnMsg "Now, you can use the serial command."
warnMsg "You can check the usage with \"$ serial --help\" or \"$ man serial\""