/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-16 21:01:40
 * @LastEditTime    : 2022-07-16 21:09:58
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/gomod.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	goModInit = false
)

func init() {
	goCmd.AddCommand(goModCmd)

	goModCmd.Flags().BoolVarP(&goModInit, "init", "i", true, "Init modules: go mod tidy; go mod vendor;")
}

var goModCmd = &cobra.Command{
	Use:   "mod",
	Short: "Go module manager. ",
	Long:  "Go module manager. ",
	Run: func(cmd *cobra.Command, args []string) {

		var shellCmd *exec.Cmd
		if goModInit {
			shellCmd = exec.Command("go", "mod", "tidy")
			shellCmd.Run()
			shellCmd = exec.Command("go", "mod", "vendor")
			shellCmd.Run()
		}
	},
}
