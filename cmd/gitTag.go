/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-17 18:34:19
 * @LastEditTime    : 2022-07-16 16:17:12
 * @LastEditors     : Lee
 * @Description     :
 * @FilePath        : /cmd/gitTag.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func tagExists(tag string, r *git.Repository) bool {
	tagFoundErr := "tag was found"
	l.Info("git show-ref --tag")
	tags, err := r.TagObjects()
	if err != nil {
		l.Printf("get tags error: %s", err)
		return false
	}
	res := false
	err = tags.ForEach(func(t *object.Tag) error {
		if t.Name == tag {
			res = true
			return fmt.Errorf(tagFoundErr)
		}
		return nil
	})
	if err != nil && err.Error() != tagFoundErr {
		l.Printf("iterate tags error: %s", err)
		return false
	}
	return res
}

func setTag(r *git.Repository, tag string) (bool, error) {
	if tagExists(tag, r) {
		l.Printf("tag %s already exists", tag)
		return false, nil
	}
	l.Printf("Set tag %s", tag)
	h, err := r.Head()
	if err != nil {
		l.Printf("get HEAD error: %s", err)
		return false, err
	}
	l.Infof("git tag -a %s %s -m \"%s\"", tag, h.Hash(), tag)
	_, err = r.CreateTag(tag, h.Hash(), &git.CreateTagOptions{
		Message: tag,
	})

	if err != nil {
		l.Printf("create tag error: %s", err)
		return false, err
	}

	return true, nil
}
