package uuids

import (
	"fmt"

	"github.com/google/uuid"
)

type Uuid struct {
}

func newUuid() UuidItf {
	return &Uuid{}
}

func (receiver Uuid) GenerateUUIDString() (s string) {
	s = uuid.New().String()
	return
}
func (receiver Uuid) GenerateUUIDInt32() (i uint32) {
	i = uuid.New().ID()
	return
}

func (receiver Uuid) GenerateUUIDInt32String() (s string) {
	s = fmt.Sprintf("%d", receiver.GenerateUUIDInt32())
	return
}
