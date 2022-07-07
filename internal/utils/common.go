/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:50:33
 * @LastEditTime    : 2022-06-22 15:28:52
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

var (
	ANSI_CYAN    = "\x1b[36;1m"
	ANSI_RESET   = "\x1b[0m"
	ANSI_DEFAULT = "\x1b[39;1m"
	ANSI_BLUE    = "\x1b[34;1m"
	ANSI_BLACK   = "\x1b[30;1m"
	ANSI_RED     = "\x1b[31;1m"
	ANSI_GREEN   = "\x1b[32;1m"
	ANSI_YELLOW  = "\x1b[33;1m"
	ANSI_WHITE   = "\x1b[37;1m"
	ANSI_MAGENTA = "\x1b[35;1m"
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
	fmt.Printf("%s%s%s\n", ANSI_GREEN, fmt.Sprintf(format, args...), ANSI_RESET)
}

// Warning should be used to display a warning
func ColorWarning(format string, args ...interface{}) {
	fmt.Printf("%s%s%s\n", ANSI_MAGENTA, fmt.Sprintf(format, args...), ANSI_RESET)
}

func ColorError(format string, args ...interface{}) {
	fmt.Printf("%s%s%s\n", ANSI_RED, fmt.Sprintf(format, args...), ANSI_RESET)
}

func ColorImportant(format string, args ...interface{}) {
	fmt.Printf("%s%s%s\n", ANSI_BLUE, fmt.Sprintf(format, args...), ANSI_RESET)
}

func RunInDir(dir string, cmd *exec.Cmd) {
	os.Chdir(dir)
	//
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
