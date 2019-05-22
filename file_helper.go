package main

import (
	"io"
	"os"
	"path/filepath"
)

// CleanFolder remove all files in a folder.
func CleanFolder(folderPath string) {
	d, err := os.Open(folderPath)
	checkError(err)
	defer d.Close()
	names, err := d.Readdirnames(-1)
	checkError(err)
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(folderPath, name))
		checkError(err)
	}
}

// CopyFolder make a duplicate folder.
func CopyFolder(srcPath string, dstPath string) (err error) {
	srcInfo, err := os.Stat(srcPath)
	checkError(err)

	err = os.MkdirAll(dstPath, srcInfo.Mode())
	checkError(err)

	directory, err := os.Open(srcPath)
	checkError(err)

	objects, err := directory.Readdir(-1)
	checkError(err)

	for _, obj := range objects {
		src := srcPath + "/" + obj.Name()
		dst := dstPath + "/" + obj.Name()
		if obj.IsDir() {
			err = CopyFolder(src, dst)
			checkError(err)
		} else {
			err = CopyFile(src, dst)
			checkError(err)
		}
	}
	return
}

// CopyFile make a duplicate file.
func CopyFile(srcPath string, dstPath string) (err error) {
	src, err := os.Open(srcPath)
	checkError(err)
	defer src.Close()

	dst, err := os.Create(dstPath)
	checkError(err)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	checkError(err)

	srcInfo, err := os.Stat(srcPath)
	checkError(err)

	err = os.Chmod(dstPath, srcInfo.Mode())
	checkError(err)

	return
}

// CopyAssetsToStaticFolder copy assets folder to static folder.
func CopyAssetsToStaticFolder(inputPath string, outputPath string) {
	CopyFolder(inputPath+"/tmpl/assets/", outputPath+"/assets/")
	CopyFolder(inputPath+"/attachment/", outputPath+"/attachment/")
}

// CreateFolder create a new folder.
func CreateFolder(path string, mode os.FileMode) {
	_, err := os.Stat(path)
	os.IsNotExist(err)

	err = os.Mkdir(path, mode)
	checkError(err)
}
