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

// Exit code
const (
	ExitSuccess int = iota // 0
	ExitFailuer
)

type options struct {
	DryRun bool   `short:"d" long:"dry-run" description:"Output the file renaming result to standard output (do not update the file)"`
	Force  bool   `short:"f" long:"force" description:"Forcibly overwrite and save even if a file with the same name exists"`
	Keep   bool   `short:"k" long:"keep" description:"Keep the file before renaming"`
	Name   string `short:"n" long:"name" value-name:"<name>" description:"Base file name with/without directory path (assign a serial number to this file name)"`
	Prefix bool   `short:"p" long:"prefix" description:"Add a serial number to the beginning of the file name"`
	Suffix bool   `short:"s" long:"suffix" description:"Add a serial number to the end of the file name(default)"`
}

func main() {
	newFileNames := make(map[string]string)
	var opts options
	var args = args(&opts)

	var dirPath = args[0]

	if !fileutil.Exists(dirPath) {
		fmt.Fprintf(os.Stderr, "%s doesn't exist.\n", dirPath)
		os.Exit(ExitFailuer)
	}

	var files = getFilepathInDir(dirPath)
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No files in %s directory.\n", dirPath)
		os.Exit(ExitFailuer)
	}

	newFileNames = newNames(opts, files)
	dieIfExistSameNameFile(opts, newFileNames)
	makeDirIfNeeded(newFileNames[files[0]])

	if opts.Keep == true {
		copy(newFileNames, opts.DryRun)
	} else {
		rename(newFileNames, opts.DryRun)
	}
	os.Exit(ExitSuccess)
}

func rename(newFileNames map[string]string, dryRun bool) {
	for org, dst := range newFileNames {
		fmt.Printf("Rename %s to %s\n", org, dst)
		if dryRun == true {
			continue
		}
		if err := os.Rename(org, dst); err != nil {
			fmt.Fprintf(os.Stderr, "Can't rename %s to %s\n", org, dst)
			os.Exit(1)
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
			os.Exit(ExitFailuer)
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
		os.Exit(ExitFailuer)
	}
	if isValidArgNr(args) == false {
		p.WriteHelp(os.Stdout)
		os.Exit(ExitFailuer)
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

// GetFilepathInDir returns the name of the file in the directory.
func getFilepathInDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get file list.")
		os.Exit(ExitFailuer)
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
			newNames[file] = filepath.Dir(file) + "/" + destDir + "/" + fileName
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

	// デフォルトフォーマット（例：%s03d%s → test001.txt）
	format := baseName + serial + ext

	if opts.Prefix == true && opts.Suffix == false {
		format = serial + baseName + ext
	}
	return format
}

func dieIfExistSameNameFile(opts options, fileNames map[string]string) {
	if opts.Force == true {
		return
	}

	for _, file := range fileNames {
		if fileutil.Exists(file) {
			fmt.Fprintf(os.Stderr, "%s (file name which is after renaming) is already exists.\n", file)
			fmt.Fprintf(os.Stderr, "Renaming may erase the contents of the file. ")
			fmt.Fprintf(os.Stderr, "So, nothing to do.\n")
			os.Exit(1)
		}
	}
}

func makeDirIfNeeded(path string) {
	dirPath := filepath.Dir(path)

	if fileutil.Exists(dirPath) {
		return
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Can't make %s directory\n", dirPath)
		os.Exit(ExitFailuer)
	}
}
