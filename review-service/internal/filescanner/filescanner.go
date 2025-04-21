package filescanner

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type FileScanner struct {
	Dir string
}

func New(dir string) *FileScanner {
	return &FileScanner{Dir: dir}
}

func (f *FileScanner) ListJLFiles() ([]string, error) {
	files, err := ioutil.ReadDir(f.Dir)
	if err != nil {
		return nil, err
	}

	var jlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".jl") {
			jlFiles = append(jlFiles, filepath.Join(f.Dir, file.Name()))
		}
	}
	return jlFiles, nil
}

func (f *FileScanner) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
