/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:45:09
 * @LastEditTime    : 2022-06-20 18:05:22
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
	. "github.com/lovelacelee/mirror-tools/internal/utils"
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(gitCmd)
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Wrapper of git command",
	Long:  "Wrapper of git command",
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return fmt.Errorf("must provide a git command")
	// 	}
	// 	return nil
	// },
	RunE: func(cmd *cobra.Command, args []string) error {
		ColorInfo("Git commands:")
		fmt.Printf("%-40s %v\n", "git submodule update --init --recursive", "初次更新带有submodule的git仓库")
		fmt.Printf("%-40s %v\n", "git rm --cached <file>", "停止跟踪文件但不删除")
		fmt.Printf("%-40s %v\n", "git log -p <file>", "查看指定文件的提交历史")
		fmt.Printf("%-40s %v\n", "git blame <file>", "以列表方式查看指定文件的提交历史")
		fmt.Printf("%-40s %v\n", "git reset --hard HEAD", "撒消工作目录中所有未提交文件的修改内容")
		fmt.Printf("%-40s %v\n", "git checkout HEAD <file>", "撤消指定的未提交文件的修改内容")
		fmt.Printf("%-40s %v\n", "git revert <commit>", "撤消指定的提交")
		fmt.Printf("%-40s %v\n", "git checkout <branch/tag>", "切换到指定分支或标签")
		fmt.Printf("%-40s %v\n", "git branch -d <branch>", "删除本地分支")
		fmt.Printf("%-40s %v\n", "git tag -d <tagname>", "删除标签")
		shellCmd := exec.Command("git", "status", "--porcelain")
		out, _ := shellCmd.Output()

		fmt.Println("\n", string(out))
		return nil
	},
}
