/*
 * @           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE                 :
 * @                     Version 2, December 2004                          :
 * @                                  -                                    :
 * @Copyright (C) 2008 Connard Lee <lovelacelee@gmail.com>                 :
 * @Everyone is permitted to copy and distribute verbatim or modified      :
 * @copies of this license document, and changing it is allowed as long    :
 * @as the name is changed.                                                :
 * @                                                                       :
 * @DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE                            :
 * @TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION        :
 * @                                 ---                                   :
 * @0. You just DO WHAT THE FUCK YOU WANT TO.                              :
 * @1. Just be happy every day.                                            :
 * @                                -----                                  :
 * @Author          : Connard
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-16 16:02:18
 * @LastEditTime    : 2022-07-16 16:06:05
 * @LastEditors     : Lee
 * @Description     :
 * @FilePath        : /internal/utils/util.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package utils

import (
	"os"

	"github.com/lovelacelee/clsgo/pkg/log"
)

var l = log.ClsLog()

func ExitWith(s string) {
	l.Info(s)
	os.Exit(1)
}

func CheckIfError(e error) {
	if e != nil {
		l.Error(e.Error())
		os.Exit(-1)
	}
}
