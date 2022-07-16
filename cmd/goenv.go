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
 * @Date            : 2022-07-16 11:37:18
 * @LastEditTime    : 2022-07-16 14:56:26
 * @LastEditors     : Lee
 * @Description     :
 * @FilePath        : /cmd/goenv.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	goProxy = false
)

func init() {
	goCmd.AddCommand(goEnvCmd)

	goEnvCmd.Flags().BoolVarP(&goProxy, "proxy", "p", false, "Use env -w GOPROXY=https://goproxy.cn,direct.")
}

var goEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Go environment manager. ",
	Long:  "Go environment manager. ",
	Run: func(cmd *cobra.Command, args []string) {
		var shellCmd *exec.Cmd

		if goProxy {
			shellCmd = exec.Command("go", "env", "-w", "GOPROXY=https://goproxy.cn,direct")
			err := shellCmd.Run()
			if err != nil {
				fmt.Printf("[%s] set environment failed.", shellCmd)
			}
		}
		shellCmd = exec.Command("go", "env", "GOPROXY")
		output, _ := shellCmd.Output()
		fmt.Print(string(output))
	},
}
