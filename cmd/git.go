/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 17:45:09
 * @LastEditTime    : 2022-07-07 15:52:33
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/git.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package cmd

import (
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5/plumbing/transport/http"
	. "github.com/lovelacelee/mirror-tools/internal/utils"
	"github.com/spf13/cobra"
)

var (
	GitUser       string
	GitPassword   string
	GithubPAToken string
)

func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.PersistentFlags().StringVarP(&GitUser, "user", "u", "lovelacelee@live.cn", "Git username")
	gitCmd.PersistentFlags().StringVarP(&GitPassword, "password", "p", "", "Git password")
	gitCmd.PersistentFlags().StringVarP(&GithubPAToken, "token", "T", "", "Github personal access token")
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Wrapper of git command",
	Long: `Useful git commands:
	git submodule update --init --recursive 初次更新带有submodule的git仓库
	git rm --cached <file> 停止跟踪文件但不删除
	git log -p <file> 查看指定文件的提交历史
	git blame <file> 以列表方式查看指定文件的提交历史
	git reset --hard HEAD 撒消工作目录中所有未提交文件的修改内容
	git checkout HEAD <file> 撤消指定的未提交文件的修改内容
	git revert <commit> 撤消指定的提交
	git checkout <branch/tag> 切换到指定分支或标签
	git branch -d <branch "删除本地分支
	git tag -d <tagname> 删除标签
`,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return fmt.Errorf("must provide a git command")
	// 	}
	// 	return nil
	// },
	RunE: func(cmd *cobra.Command, args []string) error {
		u := os.Getenv("GITUSER")
		p := os.Getenv("GITPASS")
		t := os.Getenv("GITHUB_PA_TOKEN")
		ColorImportant("GITUSER: %s \nGITPASS: %s \nGITHUB_PA_TOKEN: %s", u, HiddenSuffix(p, 4), HiddenSuffix(t, 16))
		shellCmd := exec.Command("git", "status")
		out, err := shellCmd.Output()
		ColorInfo("\n%s", string(out))
		return err
	},
}

// Fix github remote = github, if your repo has multiple branches
func GetAuth(remote string) (gitAuth *http.BasicAuth) {
	envUser := os.Getenv("GITUSER")
	envPass := os.Getenv("GITPASS")
	envToken := os.Getenv("GITHUB_PA_TOKEN")
	var user string
	var pass string
	if GitUser != "" {
		user = GitUser
	} else if envUser != "" {
		user = envUser
	} else {
		ExitWith("GITUSER must be set or use -u/--user.")
	}
	if remote == "github" {
		if GithubPAToken != "" {
			pass = GithubPAToken
		} else if envToken != "" {
			pass = envToken
		} else {
			ExitWith("GITHUB_PA_TOKEN must be set or use -T/--token.")
		}
	} else {
		if GitPassword != "" {
			pass = GitPassword
		} else if envPass != "" {
			pass = envPass
		} else {
			ExitWith("GITPASS must be set or use -p/--password.")
		}
	}

	gitAuth = &http.BasicAuth{Username: user, Password: pass}
	return gitAuth
}
