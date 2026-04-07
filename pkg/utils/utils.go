package utils

import (
	"fmt"
	"strings"
)

// GetData получение значения ячейки из rows xlsx
func GetData(s [][]string, r, c int) string {
	if r == 0 || c == 0 {
		r, err := fmt.Printf("неверный диапазон ячеек row:%v col:%v", r, c)
		if err != nil {
			panic(err.Error())
		}
		panic(r)
	}
	r -= 1
	c -= 1

	if len(s[r]) < c {
		return "0.00"
	}
	if s[r][c] == "" {
		return "0.00"
	}

	tmp := strings.Trim(s[r][c], " ")
	result := strings.ReplaceAll(tmp, ",", "")
	return result
}
