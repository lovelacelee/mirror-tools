/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-17 10:47:10
 * @LastEditTime    : 2022-06-17 17:21:11
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
	"github.com/lovelacelee/mirror-tools/internal/emoji"
	. "github.com/lovelacelee/mirror-tools/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var (
	tagVersion    string
	commitMessage string
	verbose       bool
)

func init() {
	gitCmd.AddCommand(gitPushCmd)
	gitPushCmd.Flags().BoolVarP(&verbose, "version", "v", false, "Show some details")
	gitPushCmd.Flags().StringVarP(&tagVersion, "tag", "t", "", "Tag version")
	gitPushCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "Commit message")
}

func gitPush(w *git.Worktree, r *git.Repository, remote string) error {

	// ColorInfo("git push %s", remote)
	// err := r.Push(&git.PushOptions{RemoteName: remote})
	// ShowIfError(err)

	cmd := exec.Command("git", "push", remote)
	RunInDir(w.Filesystem.Root(), cmd)

	if verbose {
		// Print the latest commit that was just pulled
		ref, err := r.Head()
		CheckIfError(err)
		latestcommit, err := r.CommitObject(ref.Hash())
		CheckIfError(err)

		fmt.Println(latestcommit)
		return err
	}
	return nil
}

func gitPushList(w *git.Worktree, repo *git.Repository) error {
	ColorInfo("git add %s", w.Filesystem.Root())
	cmd := exec.Command("git", "add", ".")
	RunInDir(w.Filesystem.Root(), cmd)

	ColorInfo("git commit %s -m %s", w.Filesystem.Root(), emoji.Message(commitMessage))
	cmd = exec.Command("git", "commit", "-m", emoji.Message(commitMessage))
	RunInDir(w.Filesystem.Root(), cmd)

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
	Example: `  mt git push -t v1.0.0 -m "Release v1.0.0"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// We instantiate a new repository targeting the given path (the .git folder)
		path, _ := os.Getwd()
		r, err := git.PlainOpen(path)
		CheckIfError(err)

		// Get the working directory for the repository
		w, err := r.Worktree()
		CheckIfError(err)
		gitPushList(w, r)

		return nil
	},
}
