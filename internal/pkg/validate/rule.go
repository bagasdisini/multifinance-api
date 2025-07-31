package validate

var (
	EmptyStringRule = Rule{
		Validate: func(v any) bool {
			s, ok := v.(string)
			return ok && s != ""
		},
		Message: "Value must not be empty",
	}
	PositiveFloatRule = Rule{
		Validate: func(v any) bool {
			f, ok := v.(float64)
			return ok && f > 0
		},
		Message: "Value must be a positive number",
	}
	PositiveUint8Rule = Rule{
		Validate: func(v any) bool {
			i, ok := v.(uint8)
			return ok && i > 0
		},
		Message: "Value must be a positive integer",
	}
)
