package scripture

import "fmt"

type ScriptureVersion string

func (v ScriptureVersion) IsValid() bool {
	for _, vd := range VersionsData {
		if vd.Name == string(v) {
			return true
		}
	}
	return false
}

func getVersionIndex(version ScriptureVersion) int {
	for i, vd := range VersionsData {
		if vd.Name == string(version) {
			return i
		}
	}
	assert(false, fmt.Sprintf("Invalid version (version: `%s`)", version))
	return -1
}

func getVersionBookIndex(version ScriptureVersion, book int) int {
	version_index := getVersionIndex(version)
	assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Invalid version index (version: `%s`, index: %d)", version, version_index))
	for i, v_book := range VersionsData[version_index].Books {
		if v_book.BookNumber == book {
			return i
		}
	}
	return -1
}

func getVersionBookData(version ScriptureVersion, book int) *VersionBookData {
	version_index := getVersionIndex(version)
	assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Invalid version index (version: `%s`, index: %d)", version, version_index))
	for _, v_book := range VersionsData[version_index].Books {
		if v_book.BookNumber == book {
			return &v_book
		}
	}
	return nil
}
