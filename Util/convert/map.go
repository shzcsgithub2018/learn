package convert

func MapToList(m map[interface{}]interface{}) []interface{} {
	var res []interface{}
	for key, val := range m {
		res = append(res, []interface{}{key, val}...)
	}
	return res
}
