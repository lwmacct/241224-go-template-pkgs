package m_file

import "bytes"

type Ts struct {
}

var ts *Ts = nil

func New() *Ts {
	return &Ts{}
}

func (t *Ts) CreateEmptyFile(path string) error {
	return CreateEmptyFile(path)
}

func (t *Ts) CreateEmptyDir(path string) error {
	return CreateEmptyDir(path)
}

func (t *Ts) CreateDirPath(filePath string) error {
	return CreateDirPath(filePath)
}

func (t *Ts) IsEmptyDir(path string) (bool, error) {
	return IsEmptyDir(path)
}

func (t *Ts) ExtractLines(buffer *bytes.Buffer, numLines int, skipBlankLines bool, existingLines []string) ([]string, *bytes.Buffer, error) {
	return ExtractLines(buffer, numLines, skipBlankLines, existingLines)
}

func (t *Ts) TailN(filePath string, numLines int, skipBlankLines bool) ([]string, error) {
	return TailN(filePath, numLines, skipBlankLines)
}