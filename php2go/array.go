package php2go

func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key := range elements {
		keys[i] = key
		i++
	}
	return keys
}

func ArrayValues(elements map[interface{}]interface{}) []interface{} {
	i, values := 0, make([]interface{}, len(elements))
	for _, val := range elements {
		values[i] = val
		i++
	}
	return values
}
