package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/nao1215/serial/pkg/fileutil"
)

const cmdName string = "serial"

var osExit = os.Exit

const version = "1.0.2"

// Exit code
const (
	ExitSuccess int = iota // 0
	ExitFailuer
)

type options struct {
	DryRun  bool   `short:"d" long:"dry-run" description:"Output the file renaming result to standard output (do not update the file)"`
	Force   bool   `short:"f" long:"force" description:"Forcibly overwrite and save even if a file with the same name exists"`
	Keep    bool   `short:"k" long:"keep" description:"Keep the file before renaming"`
	Name    string `short:"n" long:"name" value-name:"<name>" description:"Base file name with/without directory path (assign a serial number to this file name)"`
	Prefix  bool   `short:"p" long:"prefix" description:"Add a serial number to the beginning of the file name"`
	Suffix  bool   `short:"s" long:"suffix" description:"Add a serial number to the end of the file name(default)"`
	Version bool   `short:"v" long:"version" description:"Show serial command version"`
}

func main() {
	var opts options
	var args = args(&opts)

	os.Exit(run(args, opts))
}

func run(args []string, opts options) int {
	newFileNames := make(map[string]string)
	var dirPath = args[0]

	if !fileutil.Exists(dirPath) {
		fmt.Fprintf(os.Stderr, "%s doesn't exist.\n", dirPath)
		return ExitFailuer
	}

	var files = getFilePathsInDir(dirPath)
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No files in %s directory.\n", dirPath)
		return ExitFailuer
	}

	newFileNames = newNames(opts, files)
	dieIfExistSameNameFile(opts.Force, newFileNames)
	makeDirIfNeeded(newFileNames[files[0]])

	if opts.Keep == true {
		copy(newFileNames, opts.DryRun)
	} else {
		rename(newFileNames, opts.DryRun)
	}
	return ExitSuccess
}

func rename(newFileNames map[string]string, dryRun bool) {
	keys := make([]string, 0, len(newFileNames))
	for k := range newFileNames {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, org := range keys {
		fmt.Printf("Rename %s to %s\n", org, newFileNames[org])
		if dryRun == true {
			continue
		}
		if err := os.Rename(org, newFileNames[org]); err != nil {
			fmt.Fprintf(os.Stderr, "Can't rename %s to %s\n", org, newFileNames[org])
			osExit(1)
		}
	}
}

func copy(newFileNames map[string]string, dryRun bool) {
	var dest string
	keys := make([]string, 0, len(newFileNames))

	for k := range newFileNames {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, org := range keys {
		dest = newFileNames[org]
		fmt.Printf("Copy %s to %s\n", org, dest)
		if dryRun == true {
			continue
		}
		// In the case of renaming, even the same file name can be overwritten.
		// On the other hand, in the case of copying, an error will occur
		// if serial command try to overwrite with the same file name.
		if org == dest {
			continue
		}

		// If this function is running, it will force the file to be overwritten.
		// If there is the file with the same name in the copy destination,
		// delete it before copy the file.
		if fileutil.Exists(dest) {
			if err := os.Remove(dest); err != nil {
				fmt.Fprintf(os.Stderr, "Can't copy %s to %s\n", org, dest)
				osExit(ExitFailuer)
			}
		}

		if err := os.Link(org, dest); err != nil {
			fmt.Fprintf(os.Stderr, "Can't copy %s to %s\n", org, dest)
			osExit(ExitFailuer)
		}
	}
}

func args(opts *options) []string {
	return parseArgs(opts)
}

func parseArgs(opts *options) []string {
	p := initParser(opts)

	args, err := p.Parse()
	if err != nil {
		osExit(ExitFailuer)
	}

	if opts.Version {
		showVersion()
		osExit(ExitSuccess)
	}

	if len(opts.Name) != 0 && existFilenameInPath(opts.Name) == false {
		p.WriteHelp(os.Stdout)
		osExit(ExitFailuer)
	}

	if isValidArgNr(args) == false {
		p.WriteHelp(os.Stdout)
		osExit(ExitFailuer)
	}

	return args
}

func initParser(opts *options) *flags.Parser {
	parser := flags.NewParser(opts, flags.Default)
	parser.Name = cmdName
	parser.Usage = "[OPTIONS] DIRECTORY_PATH"

	return parser
}

func isValidArgNr(args []string) bool {
	if len(args) != 1 {
		return false
	}
	return true
}

func showVersion() {
	fmt.Printf("serial version %s\n", version)
}

func existFilenameInPath(path string) bool {
	if strings.HasSuffix(path, "/") {
		return false
	}
	return true
}

// getFilePathsInDir returns the paths of the file in the directory.
func getFilePathsInDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get file list.")
		osExit(ExitFailuer)
	}

	var path string
	var paths []string
	for _, file := range files {
		path = filepath.Join(dir, file.Name())
		if fileutil.IsFile(path) && fileutil.IsHiddenFile(path) == false {
			paths = append(paths, filepath.Clean(path))
		}
	}
	sort.Strings(paths)
	return paths
}

func newNames(opts options, path []string) map[string]string {
	newNames := make(map[string]string)
	destDir := filepath.Dir(opts.Name)

	var fileName string
	var format string
	// TODO: Refactor for a simpler implementation
	for i, file := range path {
		ext := filepath.Ext(file)

		if len(opts.Name) == 0 {
			format = fileNameFormat(opts.Prefix, opts.Suffix, fileutil.BaseNameWithoutExt(file), len(path))
		} else {
			format = fileNameFormat(opts.Prefix, opts.Suffix, opts.Name, len(path))
		}

		if opts.Prefix == true && opts.Suffix == false {
			fileName = fmt.Sprintf(format, i, ext)
		} else {
			fileName = fmt.Sprintf(format, i, ext)
		}

		if destDir == "." {
			newNames[file] = filepath.Clean(fileName)
		} else {
			newNames[file] = filepath.Clean(destDir + "/" + fileName)
		}
	}
	return newNames
}

func fileNameFormat(prefix bool, suffix bool, name string, totalFileNr int) string {
	baseName := filepath.Base(name)
	serial := "%0" + strconv.Itoa(len(strconv.Itoa(totalFileNr))) + "d"
	ext := "%s"

	// Default format（e.x.：%s03d%s → test001.txt）
	format := baseName + "_" + serial + ext

	if prefix == true && suffix == false {
		format = serial + "_" + baseName + ext
	}
	return format
}

func dieIfExistSameNameFile(force bool, fileNames map[string]string) {
	if force == true {
		return
	}

	for _, file := range fileNames {
		if fileutil.Exists(file) {
			fmt.Fprintf(os.Stderr, "%s (file name which is after renaming) is already exists.\n", file)
			fmt.Fprintf(os.Stderr, "Renaming may erase the contents of the file. ")
			fmt.Fprintf(os.Stderr, "So, nothing to do.\n")
			osExit(ExitFailuer)
		}
	}
}

func makeDirIfNeeded(filePath string) {
	dirPath := filepath.Dir(filePath)

	if fileutil.Exists(dirPath) {
		return
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Can't make %s directory\n", dirPath)
		osExit(ExitFailuer)
	}
}
