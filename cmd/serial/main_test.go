package main

import (
	"os"
	"testing"

	"github.com/nao1215/serial/pkg/fileutil"
	"github.com/stretchr/testify/assert"
)

func TestIsValidArgNr(t *testing.T) {
	var args []string

	// No arguments
	assert.Equal(t, false, isValidArgNr(args))

	// Only one argument
	args = append(args, "one")
	assert.Equal(t, true, isValidArgNr(args))

	// Two argumnets
	args = append(args, "two")
	assert.Equal(t, false, isValidArgNr(args))
}

func TestArgs(t *testing.T) {
	var opts options
	backupArgs := os.Args

	// short option
	testArgs := []string{"serial", "test", "-n", "name", "-p", "-s", "-d", "-k"}
	os.Args = testArgs
	a := args(&opts)
	os.Args = backupArgs
	assert.Equal(t, "test", a[0])
	assert.Equal(t, "name", opts.Name)
	assert.Equal(t, true, opts.Prefix)
	assert.Equal(t, true, opts.Suffix)
	assert.Equal(t, true, opts.DryRun)
	assert.Equal(t, true, opts.Keep)

	// long option
	testArgs = []string{"serial", "test", "--name", "name", "--prefix", "--suffix", "--dry-run", "--keep"}
	os.Args = testArgs
	opts = options{}
	a = args(&opts)
	os.Args = backupArgs
	assert.Equal(t, "test", a[0])
	assert.Equal(t, "name", opts.Name)
	assert.Equal(t, true, opts.Prefix)
	assert.Equal(t, true, opts.Suffix)
	assert.Equal(t, true, opts.DryRun)
	assert.Equal(t, true, opts.Keep)

	// only one argument
	testArgs = []string{"serial", "test"}
	os.Args = testArgs
	opts = options{}
	a = args(&opts)
	os.Args = backupArgs
	assert.Equal(t, "test", a[0])
	assert.Equal(t, "", opts.Name)
	assert.Equal(t, false, opts.Prefix)
	assert.Equal(t, false, opts.Suffix)
	assert.Equal(t, false, opts.DryRun)
	assert.Equal(t, false, opts.Keep)
}

func TestArgs2(t *testing.T) {
	var opts options
	backupArgs := os.Args

	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	testArgs := []string{"serial", "test", "test"}
	os.Args = testArgs
	args(&opts)

	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
	os.Args = backupArgs
}

func TestArgs3(t *testing.T) {
	var opts options
	backupArgs := os.Args

	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	testArgs := []string{"serial", "--noexist-option"}
	os.Args = testArgs
	args(&opts)

	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
	os.Args = backupArgs
}

func TestDieIfExistSameNameFile(t *testing.T) {
	files := map[string]string{"test": "../../.gitignore"}

	// Save current function and restore at the end:
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	dieIfExistSameNameFile(false, files)
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
	dieIfExistSameNameFile(true, files) // Not die
}

func TestMakeDirIfNeeded(t *testing.T) {
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	makeDirIfNeeded("../../test/NoWritableDir/testDir/makeDirIfNeeded.txt")
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
	makeDirIfNeeded("../../test/no_make_dir.txt")
}

func TestFileNameFormat(t *testing.T) {
	// BUG: If opts.Name is %d or %s, make invalid format.
	var opts options
	opts.Name = "test"
	opts.Prefix = true
	opts.Suffix = false
	format := fileNameFormat(opts, 100)
	assert.Equal(t, "%03d_test%s", format)

	opts.Name = "漢字"
	opts.Prefix = false
	opts.Suffix = false
	format = fileNameFormat(opts, 1000)
	assert.Equal(t, "漢字_%04d%s", format)

	opts.Name = "test"
	opts.Prefix = false
	opts.Suffix = true
	format = fileNameFormat(opts, 10000)
	assert.Equal(t, "test_%05d%s", format)

	opts.Name = "test"
	opts.Prefix = true
	opts.Suffix = true
	format = fileNameFormat(opts, 100000)
	assert.Equal(t, "test_%06d%s", format)
}

func TestGetFilepathInDir(t *testing.T) {
	files := getFilePathsInDir("../../cmd/serial")
	assert.Equal(t, "../../cmd/serial/main.go", files[0])
	assert.Equal(t, "../../cmd/serial/main_test.go", files[1])
}

func TestGetFilepathInDir2(t *testing.T) {
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	getFilePathsInDir("./no_exist")
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}

func TestRun(t *testing.T) {
	args := []string{"../../test"}
	var opts options
	opts.Name = "../../test/make_TestRun/test_file"
	opts.Prefix = true
	opts.Suffix = false
	opts.Keep = true
	run(args, opts)
	// .gitkeep is not filename extension. However, serial command recognize so.
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/0_test_file.gitkeep"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/1_test_file.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/2_test_file.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/3_test_file.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/4_test_file.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/5_test_file.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/6_test_file.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/7_test_file.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/8_test_file.txt"))

	args = []string{"../../test/make_TestRun"}
	opts.Name = "../../test/make_TestRun/test_file"
	opts.Prefix = false
	opts.Suffix = true
	run(args, opts)
	// .gitkeep is not filename extension. However, serial command recognize so.
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_0.gitkeep"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_1.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_2.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_3.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_4.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_5.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_6.txt"))
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_7.txt"))
	//  symbolic link is broken.
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/test_file_8.txt"))

	args = []string{"../../scripts"}
	opts.Name = "../../test/make_TestRun/shell"
	opts.Prefix = false
	opts.Suffix = false
	opts.Keep = true
	run(args, opts)
	assert.Equal(t, true, fileutil.Exists("../../test/make_TestRun/shell_0.sh"))

	args = []string{"."}
	opts.Name = "no_copy"
	opts.Prefix = true
	opts.Suffix = true
	opts.Keep = true
	opts.DryRun = true
	opts.Force = true
	run(args, opts)
	// check only one file
	assert.Equal(t, false, fileutil.Exists("no_copy_0.sh"))

	args = []string{"."}
	opts.Name = "no_copy"
	opts.Prefix = true
	opts.Suffix = true
	opts.Keep = false
	opts.DryRun = true
	opts.Force = true
	run(args, opts)
	// check only one file
	assert.Equal(t, false, fileutil.Exists("no_copy_0.sh"))
}

func TestRun2(t *testing.T) {
	args := []string{"no_exists"}
	var opts options
	opts.DryRun = true

	actual := run(args, opts)
	assert.Equal(t, ExitFailuer, actual)

	args = []string{"../../test/EmptyDir"}
	opts.DryRun = true
	actual = run(args, opts)
	assert.Equal(t, ExitFailuer, actual)
}

func TestCopy(t *testing.T) {
	files := map[string]string{"no_exist": "./no_exist_dir/new_file"}
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	copy(files, false)
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}

func TestRename(t *testing.T) {
	files := map[string]string{"no_exist": "./no_exist_dir/new_file"}
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	rename(files, false)
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}
