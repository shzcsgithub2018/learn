package convert

func MapToList(m map[any]any) []any {
	var res []any
	for key, val := range m {
		res = append(res, []any{key, val}...)
	}
	return res
}
