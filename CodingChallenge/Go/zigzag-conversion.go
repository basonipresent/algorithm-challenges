package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	idx, d := 0, 1
	rows := make([]string, numRows)
	for _, char := range s {
		rows[idx] += string(char)
		if idx == 0 {
			d = 1
		} else if idx == numRows-1 {
			d = -1
		}
		idx += d
	}
	result := ""
	for _, row := range rows {
		result += row
	}
	return result
}
