package path

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}

func Ls(path string, all bool, recursive bool) (paths []string) {
	if !IsDir(path) {
		paths = append(paths, path)
		return
	}
	if recursive {
		_ = filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
			if e == nil && !info.IsDir() {
				if all || !strings.HasPrefix(".", info.Name()) {
					paths = append(paths, p)
				}
			}
			return nil
		})
		return
	}
	if entries, err := os.ReadDir(path); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				name := entry.Name()
				if all || name[0] != '.' {
					paths = append(paths, path+string(os.PathSeparator)+name)
				}
			}
		}
	}
	return
}

func Overwrite(path, content string) (err error) {
	err = ioutil.WriteFile(path, []byte(content), 0644)
	return
}

func OverwriteWithPerms(path, content string, perm fs.FileMode) (err error) {
	err = ioutil.WriteFile(path, []byte(content), perm)
	return
}

// Diff returns a unified diff comparing to files
func Diff(src, dst string) (unified string, err error) {
	if !Exists(src) || IsDir(src) {
		err = fmt.Errorf(`"%v" not found or not a file`, src)
		return
	}
	if !Exists(dst) || IsDir(dst) {
		err = fmt.Errorf(`"%v" not found or not a file`, dst)
		return
	}
	var srcBytes, dstBytes []byte
	if srcBytes, err = ioutil.ReadFile(src); err != nil {
		return
	}
	if dstBytes, err = ioutil.ReadFile(dst); err != nil {
		return
	}
	srcString, dstString := string(srcBytes), string(dstBytes)
	edits := myers.ComputeEdits(span.URIFromPath(src), srcString, dstString)
	unified = fmt.Sprint(gotextdiff.ToUnified(src, dst, srcString, edits))
	return
}

func CopyFile(src, dst string) (err error) {
	if !Exists(src) || IsDir(src) {
		return fmt.Errorf(`"%v" not found or not a file`, src)
	}
	var srcFile, dstFile *os.File
	if srcFile, err = os.Open(src); err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	if dstFile, err = os.Create(dst); err != nil {
		srcFile.Close()
		return fmt.Errorf("error creating file: %s", err)
	}
	defer dstFile.Close()
	defer srcFile.Close()
	if _, e := io.Copy(dstFile, srcFile); e != nil {
		return fmt.Errorf("error copying file: %s", e)
	}
	return nil
}

func MoveFile(src, dst string) (err error) {
	if !Exists(src) || IsDir(src) {
		return fmt.Errorf(`"%v" not found or not a file`, src)
	}
	if err = os.Rename(src, dst); err == nil {
		// rename worked, no need to copy
		return
	}
	if err = CopyFile(src, dst); err != nil {
		return
	}
	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("error removing file: %s", err)
	}
	return nil
}
