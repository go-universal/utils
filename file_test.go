package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-universal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizePath(t *testing.T) {
	expected := "/a/b/c"
	result := utils.NormalizePath("/a", "b", "c")
	assert.Equal(t, expected, result)
}

func TestCreateDirectory(t *testing.T) {
	dir := "testdir/1/2"
	err := utils.CreateDirectory(dir)
	require.NoError(t, err)
	defer os.RemoveAll("testdir")

	info, err := os.Stat(dir)
	require.NoError(t, err)
	assert.True(t, info.IsDir(), "expected a directory, got a file")
}

func TestIsDirectory(t *testing.T) {
	dir := "testdir"
	err := os.Mkdir(dir, os.ModePerm)
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	isDir, err := utils.IsDirectory(dir)
	require.NoError(t, err)
	assert.True(t, isDir, "expected a directory, got a file")
}

func TestGetSubDirectory(t *testing.T) {
	dir := "testdir"
	subDir := filepath.Join(dir, "subdir")
	err := os.MkdirAll(subDir, os.ModePerm)
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	subDirs, err := utils.GetSubDirectory(dir)
	require.NoError(t, err)
	assert.Equal(t, []string{"subdir"}, subDirs)
}

func TestClearDirectory(t *testing.T) {
	dir := "testdir"
	err := os.MkdirAll(filepath.Join(dir, "subdir"), os.ModePerm)
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	err = utils.ClearDirectory(dir)
	require.NoError(t, err)

	files, err := os.ReadDir(dir)
	require.NoError(t, err)
	assert.Empty(t, files)
}

func TestFileExists(t *testing.T) {
	file := "testfile.txt"
	err := os.WriteFile(file, []byte("content"), os.ModePerm)
	require.NoError(t, err)
	defer os.Remove(file)

	exists, err := utils.FileExists(file)
	require.NoError(t, err)
	assert.True(t, exists, "expected file to exist")
}

func TestFindFile(t *testing.T) {
	dir := "testdir"
	file := filepath.Join(dir, "testfile.txt")
	err := os.MkdirAll(dir, os.ModePerm)
	require.NoError(t, err)
	err = os.WriteFile(file, []byte("content"), os.ModePerm)
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	result := utils.FindFile(dir, "testfile.txt")
	require.NotNil(t, result)
	assert.Equal(t, file, *result)
}

func TestFindFiles(t *testing.T) {
	dir := "testdir"
	file1 := filepath.Join(dir, "testfile1.txt")
	file2 := filepath.Join(dir, "testfile2.txt")
	err := os.MkdirAll(dir, os.ModePerm)
	require.NoError(t, err)
	err = os.WriteFile(file1, []byte("content"), os.ModePerm)
	require.NoError(t, err)
	err = os.WriteFile(file2, []byte("content"), os.ModePerm)
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	result := utils.FindFiles(dir, "testfile.*.txt")
	assert.ElementsMatch(t, []string{file1, file2}, result)
}

func TestGetMime(t *testing.T) {
	data := []byte("test content")
	expected := "text/plain; charset=utf-8"
	mime := utils.GetMime(data)
	assert.Equal(t, expected, mime.String())
}

func TestGetExtension(t *testing.T) {
	file := "testfile.txt"
	expected := "txt"
	result := utils.GetExtension(file)
	assert.Equal(t, expected, result)
}

func TestGetFilename(t *testing.T) {
	file := "testfile.txt"
	expected := "testfile"
	result := utils.GetFilename(file)
	assert.Equal(t, expected, result)
}

func TestTimestampedFile(t *testing.T) {
	file := "testfile.txt"
	result := utils.TimestampedFile(file)
	assert.True(t, filepath.HasPrefix(result, "testfile-"), `expected prefix "testfile-"`)
}

func TestNumberedFile(t *testing.T) {
	dir := "testdir"
	file := "testfile.txt"
	err := os.MkdirAll(dir, os.ModePerm)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(dir, file), []byte("TEST"), 0644)
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	result, err := utils.NumberedFile(dir, file)
	require.NoError(t, err)
	assert.Equal(t, "testfile-1.txt", result)
}
