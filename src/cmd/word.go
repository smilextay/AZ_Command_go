package cmd

import (
	"log"
	"strings"

	"github.com/smilextay/az_command_go/src/internal/word"
	"github.com/spf13/cobra"
)

var mode int8
var wordStr string

//命令模式定义
const (
	MODE_UPPER                          = iota + 1 //转为大写
	MODE_LOWER                                     //转为小写
	MODE_UNDERSCORE_TO_UPPER_CAMERLCASE            //下划线单词转为大写驼峰单词
	MODE_UNDERSCORE_TO_LOWER_CAMERLCASE            //下划线单词转为小写驼峰单词
	MODE_CAMELCASE_TO_UNDERSCORE                   //驼峰单词转为下画线单词
)

var helpDoc = strings.Join([]string{
	"该命令支持各种单词格式转换，格式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下划线单词转为小写驼峰单词",
	"4：下划线单词转为大写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  helpDoc,
	Run: func(cmd *cobra.Command, args []string) {

		content := ""
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(wordStr)
		case MODE_LOWER:
			content = word.ToLower(wordStr)

		case MODE_UNDERSCORE_TO_UPPER_CAMERLCASE:
			content = word.UnderScoreToUpperCameCase(wordStr)

		case MODE_UNDERSCORE_TO_LOWER_CAMERLCASE:
			content = word.UnderScoreToLowerCameCase(wordStr)

		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderScore(wordStr)

		default:
			content = "暂时不支持模式，请执行’help word‘查看文档"
		}
		log.Println(content)
	},
}

func init() {

	wordCmd.Flags().StringVarP(&wordStr, "word", "w", "", "请输入单词")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
