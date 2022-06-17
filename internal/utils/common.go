/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:50:33
 * @LastEditTime    : 2022-06-17 10:07:21
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /internal/utils/common.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// ExitIfError should be used to naively panics if an error is not nil.
func ExitIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("%s\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("%s\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func ShowIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("%s\n", fmt.Sprintf("%s", err))
}

// ColorInfo should be used to display messages
func ColorInfo(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func ColorWarning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func RunInDir(dir string, cmd *exec.Cmd) {
	os.Chdir(dir)
	//
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
