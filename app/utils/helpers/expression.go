package helpers

type ExpressionName int

const (
	Equal ExpressionName = iota
	GreatThan
	LessThan
	EqualGreatThan
	EqualLessThan
)

func ExpressionNameToValue(e ExpressionName) string {
	switch e {
	case Equal:
		return "="
	case GreatThan:
		return ">"
	case LessThan:
		return "<"
	case EqualGreatThan:
		return ">="
	case EqualLessThan:
		return "<="
	default:
		return "Invalid"
	}
}
