package convert

func Base10To36(id int64) string {
	alphabets := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var res []byte
	var idx int64
	for id > 0 {
		idx, id = id%36, id/36

		res = append([]byte{alphabets[idx]}, res...)
	}
	if string(res) == "" {
		return string(alphabets[0])
	}
	return string(res)
}
