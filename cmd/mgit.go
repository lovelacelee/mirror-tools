/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-14 17:26:30
 * @LastEditTime    : 2022-06-16 13:51:09
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/mgit.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// CheckIfError should be used to naively panics if an error is not nil.
func ExitIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("%s\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// ColorInfo should be used to display messages
func ColorInfo(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func ColorWarning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func main() {

	directory, _ := os.Getwd()

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	ExitIfError(err)
	ColorInfo("Worktree: %s\n", directory)
	w, err := r.Worktree()
	ExitIfError(err)

	// Verify the current status of the worktree using the method Status.
	ColorInfo("git status --porcelain")
	status, err := w.Status()
	ExitIfError(err)

	fmt.Println(status)

	author := &object.Signature{
		Name:  "lovelacelee",
		Email: "lovelaelee@gmail.com",
		When:  time.Now(),
	}
	// Commits the current staging area to the repository.
	// We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.
	ColorInfo("git commit")
	commit, err := w.Commit("Self committed", &git.CommitOptions{
		Author: author,
	})
	ExitIfError(err)

	// Prints the current HEAD to verify that all worked well.
	ColorInfo("git show -s")
	obj, err := r.CommitObject(commit)
	ExitIfError(err)

	fmt.Println(obj)

	// push using default options
	// err = r.Push(&git.PushOptions{})
	// ExitIfError(err)
}
