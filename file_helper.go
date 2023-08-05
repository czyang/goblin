package main

import (
	"io"
	"os"
	"path/filepath"
	"io/ioutil"
	"log"
)

// logFatalIfError logs the error and exits if there's an error.
func logFatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CleanFolder removes all files in a folder.
func CleanFolder(folderPath string) {
	err := os.RemoveAll(folderPath)
	logFatalIfError(err)
}

// CopyFolder makes a duplicate folder.
func CopyFolder(srcPath, dstPath string) {
	srcInfo, err := os.Stat(srcPath)
	logFatalIfError(err)

	err = os.MkdirAll(dstPath, srcInfo.Mode())
	logFatalIfError(err)

	objects, err := ioutil.ReadDir(srcPath)
	logFatalIfError(err)

	for _, obj := range objects {
		src := filepath.Join(srcPath, obj.Name())
		dst := filepath.Join(dstPath, obj.Name())
		if obj.IsDir() {
			CopyFolder(src, dst)
		} else {
			CopyFile(src, dst)
		}
	}
}

// CopyFile makes a duplicate file.
func CopyFile(srcPath, dstPath string) {
	src, err := os.Open(srcPath)
	logFatalIfError(err)
	defer src.Close()

	dst, err := os.Create(dstPath)
	logFatalIfError(err)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	logFatalIfError(err)

	srcInfo, err := os.Stat(srcPath)
	logFatalIfError(err)

	err = os.Chmod(dstPath, srcInfo.Mode())
	logFatalIfError(err)
}

func CopyAssetsAndPostMDToStaticFolder(templatePath, postsPath, outputPath string) {
	CopyFolder(filepath.Join(templatePath, "assets"), filepath.Join(outputPath, "posts", "assets"))
	CopyFolder(filepath.Join(postsPath, "assets"), filepath.Join(outputPath, "posts", "assets"))
}

// CreateFolder creates a new folder.
func CreateFolder(path string, mode os.FileMode) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, mode)
		logFatalIfError(err)
	}
}
