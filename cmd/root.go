/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:09:28
 * @LastEditTime    : 2022-06-20 15:09:57
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

var mtHelp string = `Mirror-Tools(clsmt) is a set of personal assistant of Lovelace(https://github.com/lovelacelee), 
Complete documentation is available at https://github.com/lovelacelee/mirror-tools.`
var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Show some details")
}

var rootCmd = &cobra.Command{
	Use:   "clsmt",
	Short: "clsmt is a short description of mirror-tools. ",
	Long:  mtHelp,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("\n%s\n", mtHelp)
		fmt.Println("Use \"clsmt [command] --help\" for more information about a command.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
