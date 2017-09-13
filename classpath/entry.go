package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	//multi jar path
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	//wildcard path
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	//archived file
	if strings.HasSuffix(strings.ToUpper(path), ".JAR") ||
		strings.HasSuffix(strings.ToUpper(path), ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
