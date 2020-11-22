#!/bin/bash

TEST_DIR=$(pwd)/test
FILES="Executable.txt Writable.txt Readable.txt NonExecutable.txt \
NonWritable.txt NonReadable.txt AllZero.txt"
DIRS="NoWritableDir NonExecutableDir EmptyDir NoEmptyDir"

function makeFile {
    touch $1
    if [ $? -ne 0 ]; then
        exit 1
    fi
}

function makeDir {
    if [ -e $1 ]; then
        return
    fi
    
    mkdir -p $1
    if [ $? -ne 0 ]; then
        exit 1
    fi
}

for file in $FILES;
do
    makeFile ${TEST_DIR}/$file
done

for dir in $DIRS;
do
    makeDir ${TEST_DIR}/$dir
done

chmod a+x ${TEST_DIR}/Executable.txt
chmod a+w ${TEST_DIR}/Writable.txt
chmod a+r ${TEST_DIR}/Readable.txt
chmod a-x ${TEST_DIR}/NonExecutable.txt
chmod a-w ${TEST_DIR}/NonWritable.txt
chmod a-r ${TEST_DIR}/NonReadable.txt
chmod a-x ${TEST_DIR}/NonExecutable.txt
chmod a-w ${TEST_DIR}/NonWritable.txt
chmod a-r ${TEST_DIR}/NonReadable.txt
chmod 000 ${TEST_DIR}/AllZero.txt

chmod a-w ${TEST_DIR}/NoWritableDir
chmod a-x ${TEST_DIR}/NonExecutableDir

ln -sf ${TEST_DIR}/Executable.txt ${TEST_DIR}/symbolic.txt

touch ${TEST_DIR}/NoEmptyDir/aaa.txt
touch ${TEST_DIR}/NoEmptyDir/bbb.txt
touch ${TEST_DIR}/NoEmptyDir/ccc.txt