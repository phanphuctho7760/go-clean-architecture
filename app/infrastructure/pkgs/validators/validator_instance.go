package validators

var ValidatorGlobalInstance ValidatorItf

func NewValidatorGlobalInstance() {
	ValidatorGlobalInstance = newPlayGroundValidator()
}
