package random_string

import (
	"github.com/rs/xid"
)

type Xid struct {
}

func newXid() RandomStringItf {
	return &Xid{}
}

func (receiver *Xid) Generate20Character() string {
	return xid.New().String()
}
