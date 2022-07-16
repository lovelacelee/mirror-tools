/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-17 10:47:10
 * @LastEditTime    : 2022-07-16 16:26:17
 * @LastEditors     : Lovelace
 * @Description     : Commit the changes to the local repository and push them to multiple remotes
 * @FilePath        : /cmd/gitpush.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/lovelacelee/clsgo/pkg/utils"
	"github.com/lovelacelee/mirror-tools/internal/emoji"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var (
	tagVersion    string
	commitMessage string
)

func init() {
	gitCmd.AddCommand(gitPushCmd)

	gitPushCmd.Flags().StringVarP(&tagVersion, "tag", "t", "", "Tag version")
	gitPushCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "Commit message")
}

func gitPush(w *git.Worktree, r *git.Repository, remote string) error {
	var cmd *exec.Cmd
	if tagVersion != "" {
		l.Infof("git push %s --tags", remote)
		cmd = exec.Command("git", "push", remote, "--tags")
	} else {
		l.Infof("git push %s", remote)
		cmd = exec.Command("git", "push", remote)
	}
	utils.RunInDir(w.Filesystem.Root(), cmd)

	if Verbose {
		// Print the latest commit that was just pulled
		ref, err := r.Head()
		utils.WarnIfError(err)
		latestcommit, err := r.CommitObject(ref.Hash())
		utils.WarnIfError(err)

		fmt.Println(latestcommit)
		return err
	}
	return nil
}

func gitPushList(w *git.Worktree, repo *git.Repository) error {
	l.Infof("git add %s", w.Filesystem.Root())
	cmd := exec.Command("git", "add", ".")
	utils.RunInDir(w.Filesystem.Root(), cmd)

	var commitInfo string
	if commitMessage != "" {
		commitInfo = emoji.Message(commitMessage)
	} else {
		commitInfo = emoji.Message("Update")
	}
	l.Infof("git commit %s -m \"%s\"", w.Filesystem.Root(), commitInfo)
	cmd = exec.Command("git", "commit", "-m", commitInfo)
	utils.RunInDir(w.Filesystem.Root(), cmd)

	if tagVersion != "" {
		setTag(repo, tagVersion)
	}

	// Push the latest changes to the remote
	list, err := repo.Remotes()
	for _, r := range list {
		gitPush(w, repo, r.Config().Name)
	}
	return err
}

var gitPushCmd = &cobra.Command{
	Use:     "push",
	Short:   "Push the latest code from multiple remote repositories.",
	Long:    "Commit, creat tag and push the latest code from multiple remote repositories.",
	Example: `  clsmt git push -t v1.0.0 -m "Release v1.0.0"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// We instantiate a new repository targeting the given path (the .git folder)
		path, _ := os.Getwd()
		r, err := git.PlainOpenWithOptions(path, &git.PlainOpenOptions{
			DetectDotGit: true,
		})
		utils.WarnIfError(err)

		// Get the working directory for the repository
		w, err := r.Worktree()
		utils.WarnIfError(err)
		gitPushList(w, r)

		return nil
	},
}
