/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-20 15:57:07
 * @LastEditTime    : 2022-07-16 16:18:51
 * @LastEditors     : Lee
 * @Description     :
 * @FilePath        : /cmd/go.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Wrapper of go command",
	Long:  "Wrapper of go command",

	RunE: func(cmd *cobra.Command, args []string) error {
		l.Warn("Supported functions:")
		fmt.Printf("%-40s %v\n", "go env", "Go env manager")

		return nil
	},
}
