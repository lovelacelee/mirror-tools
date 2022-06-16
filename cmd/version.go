/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:41:07
 * @LastEditTime    : 2022-06-16 17:43:35
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/version.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mt",
	Long:  `Print software version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mirror-Tools version:", Version)
	},
}
