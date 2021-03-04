package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yann0917/go-tour-book/tour/internal/timer"
)

var calculateTime string
var duration string

var timeLayout = "2006-01-02 15:04:05"
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var timeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format(timeLayout), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currTimer time.Time
		if calculateTime == "" {
			currTimer = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, "") {
				timeLayout = "2006-01-02"
			}
			currTimer, err = time.Parse(timeLayout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currTimer = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.GetCalculateTime(currTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果: %s, %d", calculateTime.Format(timeLayout), calculateTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(timeNowCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间有效单位为时间戳或已格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "μs"), "ms", "s", "m", "h"`)
}
