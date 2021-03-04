package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yann0917/go-tour-book/tour/internal/word"
)

const (
	MODE_UPPER                         = iota + 1 // 全部单词转为大写    MOD
	MODE_LOWER                                    // 全部单词转为小写    MOD
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE            // 下划线单词转为大写驼峰单词    MOD
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE            // 下划线单词转为小写驼峰单词    MOD
	MODE_CAMELCASE_TO_UNDERSCORE                  // 驼峰单词转为下划线单词
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)

		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscopeToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档 \n")
		}
		log.Printf("output:%s", content)
	},
}

var str string
var mode int8

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4：下划线单词转为小写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
