// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/daiguadaidai/poow/pala/config"
	"github.com/daiguadaidai/poow/pala/service"
	"github.com/spf13/cobra"
)

var sc *config.ServerConfig

// rootCmd represents the base program when called without any subprograms
var rootCmd = &cobra.Command{
	Use:   "pala",
	Short: "运行命令工具",
	Long: `
    监听并获取执行命令的通知, 让后启动一个任务.
./pala --config=./pala.toml
`,
	Run: func(cmd *cobra.Command, args []string) {
		service.Start(cfg)
	},
}

// Execute adds all child programs to the root program and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgPath string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", config.CONFIG_FILE_PATH,
		"指定的配置文件路径")
}

var cfg *config.Config

func initConfig() {
	var err error
	cfg, err = config.NewConfig(cfgPath)
	if err != nil {
		fmt.Println("解析配置文件错误: %s", err.Error())
		os.Exit(1)
	}
}
