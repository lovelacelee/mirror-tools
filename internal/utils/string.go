/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-07 14:54:56
 * @LastEditTime    : 2022-07-07 15:42:35
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /internal/utils/string.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package utils

import (
	"strings"
)

func HiddenSuffix(s string, n uint) string {
	if len(s) <= 1 {
		return s
	}
	if n >= uint(len(s))-1 {
		n = uint(len(s)) / 2
	}
	suffix := s[n:]
	return strings.Replace(s, suffix, "*", -1)
}
