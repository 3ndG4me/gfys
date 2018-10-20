package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reflectCmd represents the reflect command
var reflectCmd = &cobra.Command{
	Use:   "reflect",
	Short: "Execute reflected XSS payloads on multiple targets",
	Long: `Executes reflected XSS payloads on multiple targets

This command can test for reflected XSS vulnerabilities on a target website

It accepts the following arguments:
--targets <File of targets or a single target URL>
--payload <File of payloads or a single XSS payload>
--params <The GET paramters used to inject the payload(s) into the target(s)>`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reflect called")
	},
}

func init() {
	rootCmd.AddCommand(reflectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reflectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reflectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
