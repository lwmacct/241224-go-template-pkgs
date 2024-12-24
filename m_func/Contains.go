package m_func

import (
	"github.com/thoas/go-funk"
)

type Ts struct{}

func New() *Ts {
	return &Ts{}
}

func Contains(in interface{}, elem interface{}) bool {
	return funk.Contains(in, elem)
}
