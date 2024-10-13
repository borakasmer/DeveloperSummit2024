/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bkcli/Parser"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bkcli",
	Short: "Get Top X Articles from borakasmer.com",
	Long: `Get Top X Articles from borakasmer.com
    bkcli
	bkcli -t 5
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var top, err = cmd.Flags().GetInt("top")
		if err != nil {
			fmt.Println(err)
		}
		GetArticles(top)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bkcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().IntP("top", "t", 5, "Help message for toggle")
}

func GetArticles(top int) {
	var result = Parser.GetArticles(top)
	var tableHeader = make([]string, 0)
	tableHeader = append(tableHeader, []string{"Title", "Url"}...)
	tableRows := make([][]string, 0)
	var table = tablewriter.NewWriter(os.Stdout)
	for _, article := range result {
		tableRows = append(tableRows, []string{article.Title, article.Url})
	}
	table.SetHeader(tableHeader)
	table.AppendBulk(tableRows)
	table.SetCaption(true, "Top X Articles from borakasmer.com")
	table.Render()
}
