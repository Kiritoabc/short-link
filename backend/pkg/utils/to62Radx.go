package utils

import "sort"

var ch = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func To62RadixString(seq int64) string {
	ans := make([]byte, 0)
	for {
		remainder := int(seq % 62)
		ans = append(ans, ch[remainder])
		seq /= 62
		if seq == 0 {
			break
		}
	}
	// 翻转
	sort.Slice(ans, func(i, j int) bool {
		return i > j
	})
	return string(ans)
}
