/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:45:09
 * @LastEditTime    : 2022-06-16 18:16:36
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/git.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/go-git/go-git/v5"
	// "github.com/go-git/go-git/v5/plumbing/object"
	// "github.com/lovelacelee/mirror-tools/internal/emoji"
	. "github.com/lovelacelee/mirror-tools/internal/utils"
)

func init() {
	rootCmd.AddCommand(gitCmd)
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Wrapper of git command",
	Long:  "Wrapper of git command",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("must provide a git command")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ColorInfo("args: %v", args)
		// if err := someFunc(); err != nil {
		// 	return err
		// }
		return nil
	},
}
