package fs

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func CurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

func CurrentFile() string {
	_, file, _, _ := runtime.Caller(2)
	return filepath.Base(file)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)

	if err != nil {
		return false
	}
	return stat.IsDir()
}

func IsFile(path string) bool {
	stat, err := os.Stat(path)

	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func ResolveRelativePath(path string) string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(filename), path)
}

func DestructFileName(filename string) (string, string) {
	fileExt := filepath.Ext(filename)
	fileName := filename

	if fileExt != "" {
		fileName = fileName[:len(fileName)-len(fileExt)]
	}
	return fileName, fileExt
}

func DestructFSEntry(path string) []string {
	return strings.Split(path, "/")
}

type EntryOptions struct {
	Recursive     bool
	ExcludeFiles  bool
	ExcludeDirs   bool
	ExcludeRoot   bool
	IncludeHidden bool
	FileTypes     []string
}

func Entries(path string, options EntryOptions) []string {
	var result []string

	if options.ExcludeDirs && options.ExcludeFiles {
		return result
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		return result
	}

	for _, entry := range entries {

		if !entry.IsDir() && options.ExcludeFiles {
			continue
		}

		if entry.IsDir() && !options.IncludeHidden && entry.Name()[0] == uint8('.') {
			continue
		}

		if !entry.IsDir() && len(options.FileTypes) != 0 {
			skipEntry := false
			for _, filetype := range options.FileTypes {
				if !strings.HasSuffix(entry.Name(), "."+filetype) {
					skipEntry = true
					continue
				}
			}

			if skipEntry {
				continue
			}
		}

		pathEntry := entry.Name()
		if !options.ExcludeRoot {
			pathEntry = filepath.Join(path, entry.Name())
		}

		if (entry.IsDir() && !options.ExcludeDirs) || !entry.IsDir() {
			result = append(result, pathEntry)
		}

		if options.Recursive {
			excludeRootOption := options.ExcludeRoot
			options.ExcludeRoot = false
			result = append(result, Entries(pathEntry, options)...)
			options.ExcludeRoot = excludeRootOption
		}
	}
	return result
}
