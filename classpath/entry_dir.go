package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func new_dir_entry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) read_class(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) to_string() string{
	return self.absDir
}