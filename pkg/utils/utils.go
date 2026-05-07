// Package utils вспомогательные функции
package utils

import (
	"fmt"
	"strings"
)

// GetString получение значения ячейки из rows xlsx
func GetString(s [][]string, r, c int) string {

	if r == 0 || c == 0 {
		r, err := fmt.Printf("неверный диапазон ячеек row:%v col:%v", r, c)
		if err != nil {
			panic(err.Error())
		}
		panic(r)
	}
	r -= 1
	c -= 1

	l := len(s[r])
	if l == 0 || l < c+1 || s[r][c] == "" {
		return ""
	}

	return strings.Trim(s[r][c], " ")
}

// GetDigit получение значения ячейки из rows xlsx
// с удалением запятых в числовых значениях
func GetDigit(s [][]string, r, c int) string {
	result := GetString(s, r, c)
	if result == "" {
		return "0.00"
	}
	return strings.ReplaceAll(result, ",", "")
}

func Capitalize(s string) string {
	fio := strings.Split(strings.Trim(strings.ToLower(s), " "), " ")
	for i := range fio {
		res := []rune(fio[i])
		fio[i] = strings.ToUpper(string(res[0])) + string(res[1:])
	}
	result := strings.Join(fio, " ")
	return result
}
