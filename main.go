package main

import (
	"strings"

	"github.com/Zioyi/setop/internal"

	"github.com/spf13/cobra"
)

const (
	Intersect = iota + 1 // 交集
	Union                // 并集
	Subtract             // 差集
)

var desc = strings.Join([]string{
	"对两个csv文件数据进行集合运算，模式如下: ",
	"1: 交集 两个集合公共的数据",
	"2: 并集 连个集合合并的数据",
	"3: 补集 集合1中有且集合2中没有的数据",
}, "\n")

var mode int8
var setAFilePath string
var setBFilePath string
var resultFilePath string

var rootCmd = &cobra.Command{
	Use:     "",
	Short:   "集合运算",
	Long:    desc,
	Example: "setop -m 1 --setA A.csv --setB B.csv --result C.csv",
	Run: func(cmd *cobra.Command, args []string) {
		setA := internal.ReadSet(setAFilePath)
		setB := internal.ReadSet(setBFilePath)
		var setResult []string

		switch mode {
		case Intersect:
			setResult = internal.Intersect(setA, setB)
		case Union:
			setResult = internal.Union(setA, setB)
		case Subtract:
			setResult = internal.Subtract(setA, setB)
		}
		internal.WriteSet(setResult, resultFilePath)
	},
}

func init() {
	rootCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "模式")
	rootCmd.Flags().StringVarP(&setAFilePath, "setA", "a", "", "集合A文件路径")
	rootCmd.Flags().StringVarP(&setBFilePath, "setB", "b", "", "集合B文件路径")
	rootCmd.Flags().StringVarP(&resultFilePath, "result", "r", "", "运算结果文件路径")
}

func Execute() error {
	return rootCmd.Execute()
}

func main() {
	Execute()
}
