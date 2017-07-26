package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	read_class(className string) ([]byte, Entry, error)
	to_string() string
}

func new_entry(path string) Entry {
	//multi jar path
	if strings.Contains(path, pathListSeparator) {
		return new_composite_entry(path)
	}

	//ubprecious path
	if strings.HasSuffix(path, "*") {
		return new_wildcard_entry(path)
	}

	//archived file
	if strings.HasSuffix(strings.ToUpper(path), ".JAR") || 
	   strings.HasSuffix(strings.ToUpper(path), ".ZIP") {
		return new_zip_entry(path)
	}

	return new_dir_entry(path)
}