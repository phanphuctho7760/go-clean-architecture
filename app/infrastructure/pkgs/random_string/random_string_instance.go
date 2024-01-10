package random_string

var XidGlobalInstance RandomStringItf

func NewXidGlobalInstance() {
	XidGlobalInstance = newXid()
}
