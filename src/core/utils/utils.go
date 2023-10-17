package utils

func GetNullableValue[T any](i interface{}) *T {
	switch v := i.(type) {
	case T:
		return &v
	default:
		return nil
	}
}
