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
./pala \
    --listen-host="0.0.0.0" \
    --listen-port=19529 \
    --program-path="./pala_program" \
    --run-program-log-path="./log" \
    --run-program-paraller=8 \
    --is-log-dir-prefix-date=true \
    --heartbeat-interval=60 \
    --pili-server="localhost:19528" \
    --pili-api-version="api/v1" \
    --pili-task-update-url="http://%s/%s/pili/tasks" \
    --pili-heartbeat-url="http://%s/%s/pili/hosts/heartbeat/%s" \
    --pili-task-success-url="http://%s/%s/pili/tasks/success/%s" \
    --pili-task-fail-url="http://%s/%s/pili/tasks/fail/%s" \
    --pili-download-program-url="http://%s/%s/pili/programs/download/%s"
`,
	Run: func(cmd *cobra.Command, args []string) {
		service.Start(sc)
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

func init() {
	sc = new(config.ServerConfig)

	rootCmd.PersistentFlags().StringVar(&sc.ListenHost, "listen-host",
		config.LISTEN_HOST, "启动Http服务监听host")
	rootCmd.PersistentFlags().IntVar(&sc.ListenPort, "listen-port",
		config.LISTEN_PORT, "启动Http服务监听port")
	rootCmd.PersistentFlags().StringVar(&sc.ProgramPath, "program-path",
		config.PROGRAM_PATH, "命令存放位置")
	rootCmd.PersistentFlags().StringVar(&sc.RunProgramLogPath, "run-program-log-path",
		config.RUN_PROGRAM_LOG_PATH, "命令输出信息存放位置")
	rootCmd.PersistentFlags().IntVar(&sc.RunProgramParaller, "run-program-paraller",
		config.RUN_PROGRAM_PARALLER, "运行命令的并发数")
	rootCmd.PersistentFlags().BoolVar(&sc.IsLogDirPrefixDate, "is-log-dir-prefix-date",
		config.IS_LOG_DIR_PREFIX_DATE, "日志目录是否需要使用日期作为上级目录")
	rootCmd.PersistentFlags().IntVar(&sc.HeartbeatInterval, "heartbeat-interval",
		config.HEARTBEAT_INTERVAL, "心跳检测间隔时间")
	rootCmd.PersistentFlags().StringVar(&sc.PiliServer, "pili-server",
		config.PILI_SERVER, "调度器(pili)服务")
	rootCmd.PersistentFlags().StringVar(&sc.PiliAPIVersion, "pili-api-version",
		config.PILI_API_VERSTION, "访问pili的api版本")
	rootCmd.PersistentFlags().StringVar(&sc.PiliTaskUpdateURL, "pili-task-update-url",
		config.PILI_TASK_UPDATE_URL, "更新task信息api")
	rootCmd.PersistentFlags().StringVar(&sc.PiliHeartbeatURL, "pili-heartbeat-url",
		config.PILI_HEARTBEAT_URL, "上报心跳api")
	rootCmd.PersistentFlags().StringVar(&sc.PiliTaskSuccessURL, "pili-task-success-url",
		config.PILI_TASK_SUCCESS_URL, "命令执行成功,通知api")
	rootCmd.PersistentFlags().StringVar(&sc.PiliTaskFailURL, "pili-task-fail-url",
		config.PILI_TASK_FAIL_URL, "命令执行失败,通知api")
	rootCmd.PersistentFlags().StringVar(&sc.PiliDownloadProgramURL, "pili-download-program-url",
		config.PILI_DOWNLOAD_PROGRAM_URL, "下载程序api")
}
