package webhook

func valuesToMap(values map[string][]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for key, value := range values {
		if len(value) > 0 {
			ret[key] = value[0]
		}
	}
	return ret
}