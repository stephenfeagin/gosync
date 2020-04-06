package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var (
	dstDir     = path.Join(tmpDir, "dst")
	srcFile    = path.Join(srcDir, "src.txt")
	srcContent = []byte("This is my test data file\n")
	srcDir     = path.Join(tmpDir, "src")
	tmpDir     = "tmp"
)

func TestCopy(t *testing.T) {
	setUpTestData(t)
	defer os.RemoveAll(tmpDir)

	testCases := []struct {
		TestName, InputName, OutputName string
	}{
		{"Basic Copy", path.Join(dstDir, "dst.txt"), path.Join(dstDir, "dst.txt")},
		{"Destination directory", path.Join(tmpDir, "dst"), path.Join(dstDir, "src.txt")},
		{"Destination directory with trailing slash", path.Join(tmpDir, "dst/"), path.Join(dstDir, "src.txt")},
	}

	for i, testCase := range testCases {
		t.Logf("Case %d: %s", i, testCase.TestName)
		_, err := Copy(srcFile, testCase.InputName)
		if err != nil {
			t.Fatalf("Error returned by Copy(): %v", err)
		}

		fileEqual(t, testCase.OutputName)
		os.Remove(testCase.OutputName)
	}
}

// Helpers
func setUpTestData(t *testing.T) {
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("Could not create test source data dir: %v", err)
	}
	if err := os.Mkdir(dstDir, 0755); err != nil {
		t.Fatalf("Could not create test destination dir: %v", err)
	}
	f, _ := os.Create(srcFile)
	f.Write(srcContent)
	f.Close()
}

func fileEqual(t *testing.T, dstFile string) {
	dstContent, err := ioutil.ReadFile(dstFile)
	if err != nil {
		t.Errorf("Could not read dst file: %v", err)
	}
	if !bytes.Equal(dstContent, srcContent) {
		t.Errorf("dst file content != src file content: %v", err)
	}
}
