package uuids

var UuidInstance UuidItf

func NewUuidGlobalInstance() {
	UuidInstance = newUuid()
}
