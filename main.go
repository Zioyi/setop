package main

import (
	"strings"

	"github.com/Zioyi/setop/internal"

	"github.com/spf13/cobra"
)

var desc = strings.Join([]string{
	"对两个单列csv文件数据进行集合运算: ",
	"交集 两个集合公共的数据",
	"并集 两个集合合并的数据",
	"差集 集合1中有且集合2中没有的数据",
}, "\n")

var examples = strings.Join([]string{
	"求交集：setop intersect A.csv B.csv result.csv",
	"求并集：setop union A.csv B.csv result.csv",
	"求差集：setop subtract A.csv B.csv result.csv",
}, "\n")

var rootCmd = &cobra.Command{
	Use:     "setop",
	Short:   "集合运算",
	Long:    desc,
	Example: examples,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.MinimumNArgs(4)
		mode := args[0]
		setAFilePath := args[1]
		setBFilePath := args[2]
		resultFilePath := args[3]

		setA := internal.ReadSet(setAFilePath)
		setB := internal.ReadSet(setBFilePath)
		var setResult []string

		switch mode {
		case "intersect":
			setResult = internal.Intersect(setA, setB)
		case "union":
			setResult = internal.Union(setA, setB)
		case "subtract":
			setResult = internal.Subtract(setA, setB)
		}
		internal.WriteSet(setResult, resultFilePath)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func main() {
	Execute()
}
