package uuids

type UuidItf interface {
	GenerateUUIDString() (s string)
	GenerateUUIDInt32() (i uint32)
	GenerateUUIDInt32String() (s string)
}
