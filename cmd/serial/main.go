package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jessevdk/go-flags"
	"github.com/nao1215/serial/pkg/fileutil"
)

const cmdName string = "serial"

var osExit = os.Exit

const version = "0.0.2"

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
	for org, dst := range newFileNames {
		fmt.Printf("Rename %s to %s\n", org, dst)
		if dryRun == true {
			continue
		}
		if err := os.Rename(org, dst); err != nil {
			fmt.Fprintf(os.Stderr, "Can't rename %s to %s\n", org, dst)
			osExit(1)
		}
	}
}

func copy(newFileNames map[string]string, dryRun bool) {
	for org, dst := range newFileNames {
		fmt.Printf("Copy %s to %s\n", org, dst)
		if dryRun == true {
			continue
		}
		if err := os.Link(org, dst); err != nil {
			fmt.Fprintf(os.Stderr, "Can't copy %s to %s\n", org, dst)
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

// getFilePathInDir returns the name of the file in the directory.
func getFilePathsInDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get file list.")
		osExit(ExitFailuer)
	}

	var paths []string
	for _, file := range files {
		if fileutil.IsFile(filepath.Join(dir, file.Name())) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}
	return paths
}

func newNames(opts options, path []string) map[string]string {
	newNames := make(map[string]string)
	destDir := filepath.Dir(opts.Name)
	var fileName string
	var format string

	format = fileNameFormat(opts, len(path))

	for i, file := range path {
		e := filepath.Ext(file)

		if opts.Prefix == true && opts.Suffix == false {
			fileName = fmt.Sprintf(format, i, e)
		} else {
			fileName = fmt.Sprintf(format, i, e)
		}

		if destDir == "." {
			newNames[file] = fileName
		} else {
			newNames[file] = destDir + "/" + fileName
		}

	}
	return newNames
}

func fileNameFormat(opts options, totalFileNr int) string {
	baseName := filepath.Base(opts.Name)
	serial := "%0" + strconv.Itoa(len(strconv.Itoa(totalFileNr))) + "d"
	ext := "%s"

	// Default format（e.x.：%s03d%s → test001.txt）
	format := baseName + "_" + serial + ext

	if opts.Prefix == true && opts.Suffix == false {
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
