package scripture

import "fmt"

type ScriptureVersion string

func GetVersionIndex(v ScriptureVersion) (int, error) {
	for i, d := range VersionsData {
		if d.Name == string(v) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Could not find version in GetVersionIndex. args=(%s)", v)
}

func GetVersionBookIndex(v ScriptureVersion, b int) (int, error) {
	v_i, err := GetVersionIndex(v)
	if err != nil {
		return -1, err
	}
	for i, vb := range VersionsData[v_i].Books {
		if vb.BookNumber == b {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Could not find book for version in GetVersionBookIndex. args=(%s, %d)", v, b)
}

func GetVersionBookData(v ScriptureVersion, b int) (*VersionBookData, error) {
	v_i, err := GetVersionIndex(v)
	if err != nil {
		return nil, err
	}
	vb_i, err := GetVersionBookIndex(v, b)
	if err != nil {
		return nil, err
	}
	return &VersionsData[v_i].Books[vb_i], nil
}
