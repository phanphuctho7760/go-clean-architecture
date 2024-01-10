package validators

type ValidatorItf interface {
	ValidateStructOneErr(i interface{}) (err error)
	ValidateStructMoreErr(i interface{}) (errs []error)
}
