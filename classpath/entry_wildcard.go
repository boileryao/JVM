package classpath

import "strings"
import "os"
import "path/filepath"

func new_wildcard_entry(pathList string) CompositeEntry {
	baseDir := strings.TrimRight(pathList, "*")
	compositeEntry := []Entry{}

	walk_fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(strings.ToUpper(path), ".JAR") {
			jarEntry := new_zip_entry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walk_fn)
	return compositeEntry
}
