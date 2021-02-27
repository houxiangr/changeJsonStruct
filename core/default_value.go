package core

func defaultValue(source map[string]interface{}) (interface{}, error) {
	return source[OprDataKey], nil
}

