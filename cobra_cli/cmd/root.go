package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	text  string
	count int
)

var rootCmd = &cobra.Command{
	Use:   "printcli",
	Short: "Prints text a specified number of times",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			fmt.Println(text)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&text, "text", "t", "Hello, world!", "Text to print")
	rootCmd.PersistentFlags().IntVarP(&count, "count", "c", 1, "Number of times to print the text")
}
