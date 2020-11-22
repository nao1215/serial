package fileutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFile(t *testing.T) {
	assert.Equal(t, true, IsFile("../../test/Readable.txt"))
	assert.Equal(t, true, IsFile("../../test/symbolic.txt"))
	assert.Equal(t, false, IsFile("../../test"))
	assert.Equal(t, true, IsFile("../../test/AllZero.txt"))
	assert.Equal(t, false, IsFile("../../test/NoReadableDir"))
	assert.Equal(t, true, IsFile("../../.gitignore"))
	assert.Equal(t, false, IsFile("abcdef"))
}

func TestExists(t *testing.T) {
	assert.Equal(t, true, Exists("../../test/Readable.txt"))
	assert.Equal(t, true, Exists("../../test/symbolic.txt"))
	assert.Equal(t, true, Exists("../../test"))
	assert.Equal(t, true, Exists("../../test/AllZero.txt"))
	assert.Equal(t, true, Exists("/"))
	assert.Equal(t, true, Exists("/etc"))
	assert.Equal(t, false, Exists("abcdef"))
}

func TestIsDir(t *testing.T) {
	assert.Equal(t, false, IsDir("../../test/Readable.txt"))
	assert.Equal(t, false, IsDir("../../test/symbolic.txt"))
	assert.Equal(t, true, IsDir("../../test"))
	assert.Equal(t, false, IsDir("../../test/AllZero.txt"))
	assert.Equal(t, true, IsDir("/"))
	assert.Equal(t, true, IsDir("/etc"))
	assert.Equal(t, false, IsDir("abcdef"))
	assert.Equal(t, true, IsDir("../../test/NoWritableDir"))
}

func TestIsSymlink(t *testing.T) {
	assert.Equal(t, false, IsSymlink("../../test/Readable.txt"))
	assert.Equal(t, true, IsSymlink("../../test/symbolic.txt"))
	assert.Equal(t, false, IsSymlink("../../test"))
	assert.Equal(t, false, IsSymlink("../../test/AllZero.txt"))
	assert.Equal(t, false, IsSymlink("/"))
	assert.Equal(t, false, IsSymlink("/etc"))
	assert.Equal(t, false, IsSymlink("abcdef"))
}

func TestIsZero(t *testing.T) {
	assert.Equal(t, true, IsZero("../../test/Readable.txt"))
	assert.Equal(t, true, IsZero("../../test/symbolic.txt"))
	assert.Equal(t, false, IsZero("../../test"))
	assert.Equal(t, false, IsZero("../../cmd/serial/main.go"))
	assert.Equal(t, false, IsZero("abcdef"))
	assert.Equal(t, true, IsZero("../../test/AllZero.txt"))
}

func TestIsReadable(t *testing.T) {
	assert.Equal(t, true, IsReadable("../../test/Readable.txt"))
	assert.Equal(t, true, IsReadable("../../test/symbolic.txt"))
	assert.Equal(t, true, IsReadable("../../test"))
	assert.Equal(t, true, IsReadable("/etc"))
	assert.Equal(t, false, IsReadable("../../test/NonReadable.txt"))
	assert.Equal(t, false, IsReadable("abcdef"))
	assert.Equal(t, false, IsReadable("../../test/AllZero.txt"))
}

func TestIsWritable(t *testing.T) {
	assert.Equal(t, true, IsWritable("../../test/Writable.txt"))
	assert.Equal(t, true, IsWritable("../../test/symbolic.txt"))
	assert.Equal(t, true, IsWritable("../../test"))
	assert.Equal(t, true, IsWritable("/etc"))
	assert.Equal(t, false, IsWritable("../../test/NonWritable.txt"))
	assert.Equal(t, false, IsWritable("abcdef"))
	assert.Equal(t, false, IsWritable("../../test/AllZero.txt"))
}
func TestIsExecutable(t *testing.T) {
	assert.Equal(t, true, IsExecutable("../../test/Executable.txt"))
	assert.Equal(t, true, IsExecutable("../../test/symbolic.txt"))
	assert.Equal(t, true, IsExecutable("../../test"))
	assert.Equal(t, true, IsExecutable("/etc"))
	assert.Equal(t, false, IsExecutable("../../test/NonExecutable.txt"))
	assert.Equal(t, false, IsExecutable("abcdef"))
	assert.Equal(t, false, IsExecutable("../../test/AllZero.txt"))
}

func TestIsHiddenFile(t *testing.T) {
	assert.Equal(t, false, IsHiddenFile("../../test/Executable.txt"))
	assert.Equal(t, true, IsHiddenFile("../../.gitignore"))
	assert.Equal(t, false, IsHiddenFile("./file.go"))
	assert.Equal(t, false, IsHiddenFile("/etc"))
	assert.Equal(t, false, IsHiddenFile("../../test/"))
	assert.Equal(t, false, IsHiddenFile("../../test"))
	assert.Equal(t, false, IsHiddenFile("abcdef"))
	assert.Equal(t, false, IsHiddenFile(".abcdef"))
}

func TestBaseNameWithoutExt(t *testing.T) {
	assert.Equal(t, "Executable", BaseNameWithoutExt("../../test/Executable.txt"))
	assert.Equal(t, "", BaseNameWithoutExt("../../.gitignore"))
	assert.Equal(t, "file", BaseNameWithoutExt("./file.go"))
	assert.Equal(t, "etc", BaseNameWithoutExt("/etc"))
	assert.Equal(t, "", BaseNameWithoutExt("../../test/"))
	assert.Equal(t, "test", BaseNameWithoutExt("../../test"))
	assert.Equal(t, "abcdef", BaseNameWithoutExt("abcdef"))
	assert.Equal(t, "", BaseNameWithoutExt(".abcdef"))
}
