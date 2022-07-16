/*
 * @           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE                 :
 * @                     Version 2, December 2004                          :
 * @                                  -                                    :
 * @Copyright (C) 2008 Connard Lee <lovelacelee@gmail.com>                 :
 * @Everyone is permitted to copy and distribute verbatim or modified      :
 * @copies of this license document, and changing it is allowed as long    :
 * @as the name is changed.                                                :
 * @                                                                       :
 * @DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE                            :
 * @TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION        :
 * @                                 ---                                   :
 * @0. You just DO WHAT THE FUCK YOU WANT TO.                              :
 * @1. Just be happy every day.                                            :
 * @                                -----                                  :
 * @Author          : Connard
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-16 11:21:12
 * @LastEditTime    : 2022-07-16 11:39:57
 * @LastEditors     : Lee
 * @Description     :
 * @FilePath        : /cmd/go.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-20 15:57:07
 * @LastEditTime    : 2022-07-16 11:21:15
 * @LastEditors     : Lee
 * @Description     :
 * @FilePath        : /cmd/py.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"fmt"

	. "github.com/lovelacelee/mirror-tools/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Wrapper of go command",
	Long:  "Wrapper of go command",

	RunE: func(cmd *cobra.Command, args []string) error {
		ColorInfo("Supported functions:")
		fmt.Printf("%-40s %v\n", "go env", "Go env manager")

		return nil
	},
}
