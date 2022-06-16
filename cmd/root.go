/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:09:28
 * @LastEditTime    : 2022-06-16 17:24:05
 * @LastEditors     : Lovelace
 * @Description     : The root command of Cobra.
 * @FilePath        : /cmd/root.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var mtHelp string = `Mirror-Tools is a set of personal assistant of Lovelace(https://github.com/lovelacelee)
maintained since 2022, when I started to learn Go.
Complete documentation is available at https://github.com/lovelacelee/mirror-tools.`

var rootCmd = &cobra.Command{
	Use:   "mt",
	Short: "mt is a short description of mirror-tools. ",
	Long:  mtHelp,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Root CMD")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
