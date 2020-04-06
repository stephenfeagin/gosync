package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

// Copy copies a single file into the provided destination file.
// If dst is a directory, Copy copies the source file into the
// directory with the same file name as the source file.
func Copy(src, dst string) (int64, error) {
	srcStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !srcStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	if srcStat.IsDir() {
		return 0, fmt.Errorf("src must be a file, not a directory")
	}

	dstStat, err := os.Stat(dst)
	if err == nil { // the file/dir already exists
		if dstStat.IsDir() { // check if it's a directory
			dst = path.Join(dst, srcStat.Name())
		}
		if dstStat.Name() == srcStat.Name() { // same filename
			if dstStat.Size() == srcStat.Size() { // mod
				fmt.Println("Identical file")
				return 0, nil // same file, not modified
			}
		}
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	nBytes, err := io.Copy(dstFile, srcFile)
	return nBytes, nil
}

// CopyDir copies an entire directory and its contents into the provided destination.
// TODO: If src is a bare dir name, copy that dir into the dst
// TODO: If src is a dir name followed by /*, copy the contents into dst but not the parent folder
func CopyDir(src, dst string) error {
	return nil
}
