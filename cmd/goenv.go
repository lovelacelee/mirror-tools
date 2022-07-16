/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-16 11:37:18
 * @LastEditTime    : 2022-07-16 21:01:45
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/goenv.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	goProxyQiniu  = false
	goProxyIocn   = false
	goProxyAliyun = false
	goProxyOff    = false
)

func init() {
	goCmd.AddCommand(goEnvCmd)

	goEnvCmd.Flags().BoolVarP(&goProxyQiniu, "qiniu", "", false, "Use env -w GOPROXY=https://goproxy.cn,direct.")
	goEnvCmd.Flags().BoolVarP(&goProxyIocn, "iocn", "", false, "Use env -w GOPROXY=https://proxy.golang.com.cn,direct.")
	goEnvCmd.Flags().BoolVarP(&goProxyAliyun, "aliyun", "", false, "Use env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct.")
	goEnvCmd.Flags().BoolVarP(&goProxyOff, "off", "", false, "Disable GOPROXY")
}

var goEnvCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Go PROXY environment manager. ",
	Long:  "Go PROXY environment manager. ",
	Run: func(cmd *cobra.Command, args []string) {
		var shellCmd *exec.Cmd
		var proxyServer string = "https://goproxy.io/zh/"

		if goProxyIocn {
			proxyServer = "https://proxy.golang.com.cn,direct"
		} else if goProxyQiniu {
			proxyServer = "https://goproxy.cn,direct"
		} else if goProxyAliyun {
			proxyServer = "https://mirrors.aliyun.com/goproxy"
		}

		if goProxyOff {
			shellCmd = exec.Command("go", "env", "-u", "GOPROXY")
			shellCmd.Run()
		}

		shellCmd = exec.Command("go", "env", "-w", "GOPROXY="+proxyServer)
		err := shellCmd.Run()
		if err != nil {
			fmt.Printf("[%s] set environment failed.", shellCmd)
		}
		shellCmd = exec.Command("go", "env", "GOPROXY")
		output, _ := shellCmd.Output()
		fmt.Print(string(output))
	},
}
