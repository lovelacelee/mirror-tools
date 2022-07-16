/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-20 15:57:07
 * @LastEditTime    : 2022-07-16 16:20:37
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/py.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PythonCmd string

func init() {
	rootCmd.AddCommand(pyCmd)
	pyCmd.PersistentFlags().StringVarP(&PythonCmd, "python", "p", "python", "Python command. Deafult is python, but you can use python3")
}

var pyCmd = &cobra.Command{
	Use:   "py",
	Short: "Wrapper of python command",
	Long:  "Wrapper of python command",
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return fmt.Errorf("must provide a git command")
	// 	}
	// 	return nil
	// },
	RunE: func(cmd *cobra.Command, args []string) error {
		l.Warn("Supported functions:")
		fmt.Printf("%-40s %v\n", "python venv", "Python virtual environment manager")

		return nil
	},
}
