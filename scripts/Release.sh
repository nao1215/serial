#!/bin/bash
# Support OS and Arch info: https://golang.org/doc/install/source#environment
PROJECT="serial"
CWD=$(pwd)
RELEASE=${CWD}/release
BIN_INFO_TXT=${RELEASE}/binary_info.txt
# Not build windows binary.
OS="linux darwin freebsd netbsd openbsd plan9 solaris aix dragonfly illumos js android"
ARCH="386 amd64 arm arm64 ppc64 ppc64le mips mipsle mips64 mips64le s390x wasm"
VERSION="v$(gobump show ${CWD}/cmd/serial | sed -e "s/{\"version\":\"\(.*\)\"}/\1/g")"
MANPAGES="${CWD}/docs/man"

function mkReleaseDir() {
    cd  ${CWD}
    for os in $OS;
    do
        mkdir -p ${RELEASE}/$os
        for arch in $ARCH;
        do
            mkdir -p ${RELEASE}/$os/${PROJECT}-${VERSION}-${os}-${arch}
        done
    done
}

function cpManpages() {
    release="$1"

    cd  ${CWD}
    cp -r ${MANPAGES} ${release}
    markdown=$(find ${release} -name "*.md")
    for md in ${markdown};
    do
        rm -f ${md}
    done
}

function cpLicense() {
    release="$1"

    cd  ${CWD}
    cp -f LICENSE ${release}
}

function mkRelease() {
    cd ${CWD}
    os="$1"
    arch="$2"
    release_dir=${PROJECT}-${VERSION}-${os}-${arch}
    release_path="${RELEASE}/$os/${release_dir}"
    tarball="${release_dir}.tar.gz"
    zip_file="${release_dir}.zip"

    echo "---Make release files for OS=$os Architecture=$arch" >> ${BIN_INFO_TXT}
    GOOS=$os GOARCH=$arch make
    echo -n "   " >> ${BIN_INFO_TXT}
    file ${CWD}/${PROJECT} >> ${BIN_INFO_TXT}
    echo "" >> ${BIN_INFO_TXT}
    mv ${CWD}/${PROJECT} ${release_path}/.
    cpLicense ${release_path}/.
    cpManpages ${release_path}/.

    cd ${RELEASE}/$os/
    tar cvfz "${tarball}" "${release_dir}"
    zip "${zip_file}" -r "${release_dir}"

    mv "${tarball}" "${RELEASE}/."
    mv "${zip_file}" "${RELEASE}/."
    rm -rf ${release_dir}
    cd ${CWD}
}

function mkLinuxRelease() {
    cd ${CWD}
    os="linux"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkMacOsRelease() {
    cd ${CWD}
    os="darwin"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkFreeBsdRelease() {
    cd ${CWD}
    os="freebsd"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkOpenBsdRelease() {
    cd ${CWD}
    os="openbsd"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkNetBsdRelease() {
    cd ${CWD}
    os="netbsd"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkPlan9Release() {
    cd ${CWD}
    os="plan9"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkSolarisRelease() {
    cd ${CWD}
    os="solaris"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkAixRelease() {
    cd ${CWD}
    os="aix"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkAndroidRelease() {
    cd ${CWD}
    os="android"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkDragonflyRelease() {
    cd ${CWD}
    os="dragonfly"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkIllumosRelease() {
    cd ${CWD}
    os="illumos"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkJsRelease() {
    cd ${CWD}
    os="js"
    arch="$1"

    mkRelease ${os} ${arch}
}

function mkLinuxAllRelease() {
    arch="386 amd64 arm arm64 ppc64 ppc64le mips mipsle mips64 mips64le s390x"
    for i in ${arch};
    do
        mkLinuxRelease "$i"
    done
}

function mkMacOsAllRelease() {
    # Not support 386, arm, arm64.
    arch="amd64"
    for i in ${arch};
    do
        mkMacOsRelease "$i"
    done
}

function mkFreeBsdAllRelease() {
    arch="386 amd64 arm"
    for i in ${arch};
    do
        mkFreeBsdRelease "$i"
    done
}

function mkOpenBsdAllRelease() {
    arch="386 amd64 arm arm64"
    for i in ${arch};
    do
        mkOpenBsdRelease "$i"
    done
}

function mkNetBsdAllRelease() {
    arch="386 amd64 arm"
    for i in ${arch};
    do
        mkNetBsdRelease "$i"
    done
}

function mkPlan9AllRelease() {
    arch="386 amd64 arm"
    for i in ${arch};
    do
        mkPlan9Release "$i"
    done
}

function mkAixAllRelease() {
    # Can't build ppc64, 2020/11/22
    arch="ppc64"
    for i in ${arch};
    do
        :
        # mkAixRelease "$i"
    done
}

function mkAndroidAllRelease() {
    # Can't build 386, amd64, arm, 2020/11/22
    arch="arm64"
    for i in ${arch};
    do
        mkAndroidRelease "$i"
    done
}

function mkDragonflyAllRelease() {
    arch="amd64"
    for i in ${arch};
    do
        mkDragonflyRelease "$i"
    done
}

function mkIllumosAllRelease() {
    arch="amd64"
    for i in ${arch};
    do
        mkIllumosRelease "$i"
    done
}

function mkJsAllRelease() {
    # Can't build wasm
    arch="wasm"
    for i in ${arch};
    do
        :
        #mkJsRelease "$i"
    done
}

function mkSolarisAllRelease() {
    arch="amd64"
    for i in ${arch};
    do
        mkSolarisRelease "$i"
    done
}

function mkSrcRelease() {
    TMP=$(mktemp -d)
    code="${PROJECT}-${VERSION}-src"
    tarball="$code.tar.gz"
    zip_file="$code.zip"

    cd  ${CWD}
    cp -r ${CWD}/../${PROJECT} ${TMP}/.
    mkdir -p ${RELEASE}
    mv ${TMP}/${PROJECT} ${RELEASE}/.

    cd  ${RELEASE}
    tar cvfz ${tarball} "${PROJECT}"
    zip ${zip_file} -r "${PROJECT}"
    rm -rf "${RELEASE}/${PROJECT}"
    cd  ${CWD}
}

function rmOsDirInRelease() {
    for os in ${OS};
    do
        rm -rf ${RELEASE}/$os
    done
}

function main() {
    cd ${CWD}
    make clean
    mkSrcRelease

    touch ${BIN_INFO_TXT}
    mkReleaseDir
    make doc
    mkAixAllRelease
    mkAndroidAllRelease
    mkMacOsAllRelease
    mkDragonflyAllRelease
    mkFreeBsdAllRelease
    mkIllumosAllRelease
    mkJsAllRelease
    mkLinuxAllRelease
    mkNetBsdAllRelease
    mkOpenBsdAllRelease
    mkPlan9AllRelease
    mkSolarisAllRelease
    rmOsDirInRelease
}

main