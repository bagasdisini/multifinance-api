package validate

type Preprocessor interface {
	Preprocess()
}

type Rule struct {
	Validate func(value any) bool
	Message  string
}

type Validator struct {
	Errors map[string]string
}

func (v *Validator) Check(field string, value any, rules ...Rule) {
	for _, rule := range rules {
		if !rule.Validate(value) {
			if v.Errors == nil {
				v.Errors = make(map[string]string)
			}
			v.Errors[field] = rule.Message
			break
		}
	}
}
