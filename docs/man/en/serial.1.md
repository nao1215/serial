% SERIAL(1)
% Naohiro CHIKAMATSU <n.chika156@gmail.com>
% November 2020

# NAME

serial – rename the file name to the name with a serial number.

# SYNOPSIS

**serial** [OPTIONS] DIRECTORY_PATH

# DESCRIPTION
**serial** is a CLI command that renames files under any directory to <br>
the format user specified file name with serial number.<br>
serial can specify the destination directory of the renamed file. Also, <br>
if you want to keep the original file, you can copy the file instead of <br>
renaming the file.

# EXAMPLES
**Rename the file name to the name with a serial number at current directory.**

    $ ls
      a.txt  b.txt  c.txt
    $ serial --name demo  .
      Rename a.txt to demo_1.txt
      Rename b.txt to demo_2.txt
      Rename c.txt to demo_3.txt

**Copy & Rename the file at specified directory.**

    $ serial -p -k -n ../../dir/demo .
      Copy a.txt to ../../dir/0_demo.txt
      Copy b.txt to ../../dir/1_demo.txt
      Copy c.txt to ../../dir/2_demo.txt


# OPTIONS
**-d**, **--dry-run**
:   Output the file renaming result to standard output (do not update the file).

**-f**, **--force**
:   Forcibly overwrite and save even if a file with the same name exists.

**-h**, **--help**
:   Show help message.

**-k**, **--keep**
:   Keep the file before renaming (not rename, do copy).

**-n new_filename**, **--name=new_filename**
:   Base file name with/without directory path (assign a serial number to this file name).

**-p**, **--prefix**
:   Add a serial number to the beginning of the file name.

**-s**, **--suffix**
:   Add a serial number to the end of the file name(default).

**-v**, **--version**
:   Show serial command version.

# EXIT VALUES
**0**
:   Success

**1**
:   There is an error in the argument of the serial command.

# BUGS
See GitHub Issues: https://github.com/nao1215/serial/issues

# LICENSE
serial command project is licensed under the terms of the MIT license.