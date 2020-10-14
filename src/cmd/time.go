package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/smilextay/az_command_go/src/internal/timer"
	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var timeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		now := timer.GetNowTime()
		log.Println("now:	",now.Format("2006-01-02 15:04:05"))
		log.Println("timestamp:	" ,now.Unix())
	},
}

var timeCalcCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算时间",
	Long:  "计算时间",
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var ctime time.Time
		if len(strings.TrimSpace(calculateTime)) == 0 {
			ctime = timer.GetNowTime()
		} else {
			ctime, err = time.Parse(calculateTime, "2006-01-02 15:04:05")
			if err != nil {
				//格式话错误，尝试时间戳
				t, err := strconv.ParseInt(calculateTime, 10, 64)
				if err == nil {
					log.Println("输入的时间有误")
					return
				}
				ctime = time.Unix(t, 0)
			}
		}
		rtime, err := timer.GetCalculateTime(ctime, duration)
		if err != nil {
			log.Fatalf("计算遇到错误 err:%v", err)
		} else {
			log.Println(rtime.Format("2006-01-02 15:04:05"))
			log.Println(rtime.Unix())
		}
	},
}

func init() {

	timeCmd.AddCommand(timeNowCmd)

	timeCmd.AddCommand(timeCalcCmd)

	timeCalcCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效值为'2006-01-02 15:04:05'格式化后的字符串或者时间戳`)

	timeCalcCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为 'ns','us','ms','s','m','h'`)

}
