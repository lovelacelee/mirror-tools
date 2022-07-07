/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 18:00:17
 * @LastEditTime    : 2022-06-22 15:34:06
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/gitupdate.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	. "github.com/lovelacelee/mirror-tools/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	gitCmd.AddCommand(gitUpdateCmd)
}

func gitPull(w *git.Worktree, r *git.Repository, remote string) error {
	gitAuth := &http.BasicAuth{Username: os.Getenv("GITUSER"), Password: os.Getenv("GITPASS")}
	// Pull the latest changes from the origin remote and merge into the current branch
	ColorInfo("git pull %s", remote)
	err := w.Pull(&git.PullOptions{RemoteName: remote, Auth: gitAuth})
	ShowIfError(err)

	if Verbose {
		// Print the latest commit that was just pulled
		ref, err := r.Head()
		CheckIfError(err)
		latestcommit, err := r.CommitObject(ref.Hash())
		CheckIfError(err)

		fmt.Println(latestcommit)
	}
	return err
}

func gitPullList(w *git.Worktree, repo *git.Repository) error {
	list, err := repo.Remotes()
	for _, r := range list {
		gitPull(w, repo, r.Config().Name)
	}
	return err
}

var gitUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Pull the latest code from multiple remote repositories.",
	Long:  "Pull the latest code from multiple remote repositories.",

	RunE: func(cmd *cobra.Command, args []string) error {
		// We instantiate a new repository targeting the given path (the .git folder)
		path, _ := os.Getwd()
		r, err := git.PlainOpen(path)
		CheckIfError(err)

		// Get the working directory for the repository
		w, err := r.Worktree()
		CheckIfError(err)
		gitPullList(w, r)

		return nil
	},
}
