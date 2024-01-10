package helpers

func ConvertToPointerSlice[T any](input []T) (output []*T) {
	output = make([]*T, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = &input[i]
	}
	return
}
