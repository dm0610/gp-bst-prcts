package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	_ "path/filepath"
	"testing"
)

var (
//testdataTop = "testdata"
//testdata1   = filepath.Join(testdataTop, "dir1")
//testdata2   = filepath.Join(testdataTop, "dir2")
)

//func sliceToSet(stuff []string) map[string]bool {
//	set := make(map[string]bool)
//	for _, thing := range stuff {
//		set[thing] = true
//	}
//	return set
//}

//****************test1
//test without testify
func TestCliRequiresOneOrMoreDirectories1(t *testing.T) {
	emptySlice := make([]string, 0)
	if validateArgs(emptySlice) == nil {
		t.Error("no arguments should be an error")
	}
}

//test with testify assert.NotEqual
func TestCliRequiresOneOrMoreDirectories2(t *testing.T) {
	emptySlice := make([]string, 0)
	assert.NotEqual(t, validateArgs(emptySlice), nil, "no arguments should be an error")
}

//test with testify assert.NotNill
func TestCliRequiresOneOrMoreDirectories3(t *testing.T) {
	emptySlice := make([]string, 0)
	assert.NotNil(t, validateArgs(emptySlice), "no arguments should be an error")
}

//****************test2

//Test Count Words with stub
func TestCountWords(t *testing.T) {
	content := []byte("testdata dir1 file1_1")
	buf := bytes.NewBuffer(content)
	assert.Equal(t, 3, countWords(buf))
}

//****************test3
//Integration tests
//не придумал что делать.
