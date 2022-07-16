/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-20 15:29:02
 * @LastEditTime    : 2022-07-16 16:21:20
 * @LastEditors     : Lovelace
 * @Description     : Create or delete a virtual environment.
 * @FilePath        : /cmd/pyvenv.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	pyVenvCreate = false
	pyVenvDelete = false
)

func init() {
	pyCmd.AddCommand(pyVenvCmd)
	pyVenvCmd.Flags().BoolVarP(&pyVenvCreate, "create", "c", false, "Create a virtual environment, fixed as venv.")
	pyVenvCmd.Flags().BoolVarP(&pyVenvDelete, "delete", "d", false, "Delete a virtual environment, fixed as venv.")
}

var pyVenvCmd = &cobra.Command{
	Use:   "venv",
	Short: "Python virtual environment manager. ",
	Long:  "Python virtual environment manager. ",
	Run: func(cmd *cobra.Command, args []string) {
		var shellCmd *exec.Cmd
		cwd, _ := os.Getwd()
		if pyVenvCreate {
			l.Infof("[%s] Create a virtual environment.", PythonCmd)
			shellCmd = exec.Command(PythonCmd, "-m", "venv", "venv")

			// fmt.Println(shellCmd.Output())
			err := shellCmd.Run()
			if err != nil {
				l.Errorf("[%s] Create a virtual environment failed.", PythonCmd)
			}
		}
		if pyVenvDelete {
			l.Info("Delete a virtual environment.")
			os.RemoveAll(path.Join(cwd, "venv"))
		} else {
			sysType := runtime.GOOS

			if sysType == "windows" {
				fmt.Printf("%-20s %v%v\n", "cmd.exe", cwd, "> venv\\Scripts\\Activate.bat")
				fmt.Printf("%-20s %v%v\n", "PowerShell", cwd, "> venv\\Scripts\\Activate.ps1")
			} else {
				fmt.Printf("%-20s %v%v\n", "bash/zsh", cwd, "$ source venv/bin/activate")
				fmt.Printf("%-20s %v%v\n", "fish", cwd, "$ source venv/bin/activate.fish")
				fmt.Printf("%-20s %v%v\n", "csh/tcsh", cwd, "$ source venv/bin/activate.csh")
			}
		}
	},
}
