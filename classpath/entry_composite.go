package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func new_composite_entry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := new_entry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) read_class(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.read_class(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("Not found class: " + className)
}

func (self CompositeEntry) to_string() string{
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.to_string()
	}

	return strings.Join(strs, pathListSeparator)
}