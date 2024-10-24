package scripture

import "fmt"

type ScriptureVersion string

func (s *ScriptureService) IsVersionValid(v ScriptureVersion) bool {
	for _, vd := range VersionsData {
		if vd.Name == string(v) {
			return true
		}
	}
	return false
}

type ScriptureVersionBasicInfo struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Language string `json:"language"`
}

func (s *ScriptureService) getVersionIndex(version ScriptureVersion) int {
	for i, vd := range VersionsData {
		if vd.Name == string(version) {
			return i
		}
	}
	s.assert(false, fmt.Sprintf("Invalid version (version: `%s`)", version))
	return -1
}

func (s *ScriptureService) getVersionBookIndex(version ScriptureVersion, book int) int {
	version_index := s.getVersionIndex(version)
	s.assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Invalid version index (version: `%s`, index: %d)", version, version_index))
	for i, v_book := range VersionsData[version_index].Books {
		if v_book.BookNumber == book {
			return i
		}
	}
	return -1
}

func (s *ScriptureService) getVersionBookData(version ScriptureVersion, book int) *VersionBookData {
	version_index := s.getVersionIndex(version)
	s.assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Invalid version index (version: `%s`, index: %d)", version, version_index))
	for _, v_book := range VersionsData[version_index].Books {
		if v_book.BookNumber == book {
			return &v_book
		}
	}
	return nil
}

func (s *ScriptureService) GetVersionRange(version ScriptureVersion) ScriptureRange {
	version_index := s.getVersionIndex(version)
	s.assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Invalid version index (version: `%s`, index: %d)", version, version_index))

	version_books := VersionsData[version_index].Books
	s.assert(len(version_books) > 0, fmt.Sprintf("No books found for version (version: `%s`)", version))

	start_book := version_books[0].BookNumber
	end_book := version_books[len(version_books)-1].BookNumber
	return ScriptureRange{
		Version: version,
		Start:   s.CreateFirstValidRef(version, start_book),
		End:     s.CreateLastValidRef(version, end_book),
	}
}

func (s *ScriptureService) GetVersionsBasicData() []ScriptureVersionBasicInfo {
	versions := make([]ScriptureVersionBasicInfo, 0, len(VersionsData))
	for _, vd := range VersionsData {
		versions = append(versions, ScriptureVersionBasicInfo{
			Name:     vd.Name,
			FullName: vd.FullName,
			Language: vd.Language,
		})
	}
	return versions
}

func (s *ScriptureService) GetVersionData(version string) VersionData {
	for _, vd := range VersionsData {
		if vd.Name == version {
			return vd
		}
	}
	s.assert(false, fmt.Sprintf("Invalid version (version: `%s`)", version))
	return VersionData{}
}

func (s *ScriptureService) GetVersionLanguage(version string) string {
	for _, vd := range VersionsData {
		if vd.Name == version {
			return vd.Language
		}
	}
	s.assert(false, fmt.Sprintf("Invalid version (version: `%s`)", version))
	return ""
}

func (s *ScriptureService) GetVersionLanguageHeadings(version string) string {
	for _, vd := range VersionsData {
		if vd.Name == version {
			return vd.LanguageHeadings
		}
	}
	s.assert(false, fmt.Sprintf("Invalid version (version: `%s`)", version))
	return ""
}
