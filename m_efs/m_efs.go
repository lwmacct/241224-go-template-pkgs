package m_efs

import (
	"embed"
	"io/fs"
)

type Ts struct {
	Embed embed.FS
	Error error
}

func New(efs embed.FS) *Ts {
	return &Ts{
		Embed: efs,
		Error: nil,
	}
}

func (t *Ts) GetContent(Path string) string {
	content, err := t.Embed.ReadFile(Path)
	if err != nil {
		t.Error = err
	}
	return string(content)
}

func (t *Ts) GetFsFs(Path string) fs.FS {
	fsfs, err := fs.Sub(t.Embed, Path)
	if err != nil {
		t.Error = err
	}
	return fsfs
}
