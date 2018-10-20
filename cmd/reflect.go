package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Flags
var (
	target string
)

// reflectCmd represents the reflect command
var reflectCmd = &cobra.Command{
	Use:   "reflect",
	Short: "Execute reflected XSS payload on a target",
	Long: `Executes reflected XSS payload on a target

This command can test for reflected XSS vulnerabilities on a target website

Examples:
	Single Target: 

	gfys reflect --target "www.fckd.site/index.php?vulnerable={{payload}}" --payload "<script>alert(1337)</script>"
	
	This generates the following request:

	www.fckd.site/index.php?vulnerable=<script>alert(1337)</script>
	
	Multiple Targets:

	gfys reflect --target fkcd_sites.txt --payload payloads.txt

	This generates the following requests based on your target template (will iterate through all sites and payloads provided):

	www.fckd.site/index.php?vulnerable=<script>alert(1337)</script>&pwned=<script>alert(1337)</script>
	www.fckd.site/index.php?vulnerable=<script>alert(1337)</script>&pwned=somedata
	www.fckd.site/index.php?vulnerable=somedata&pwned=<script>alert(1337)</script>
	`,

	Run: func(cmd *cobra.Command, args []string) {

		workingDir, err := os.Getwd()

		if err != nil {
			fmt.Println("An error occurred while getting you current working directory: ", err)
		}

		filePath := filepath.Join(workingDir, target)

		file, err := os.Stat(filePath)

		if err != nil {
			file, err = os.Stat(target)
			if err != nil {
				response := getRequest(target)
				fmt.Println(response)
			} else {
				fmt.Println("Do some file stuffz not in working dir (returning IsDir for now) ", file.IsDir())
			}
		} else {
			fmt.Println("Do some file stuffz (returning IsDir for now) ", file.IsDir())
		}

	},
}

func getRequest(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("An error occurred while making your request: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("An error occurred while reading the respone: ", err)
	}

	return string(body)
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
	reflectCmd.Flags().StringVarP(&target, "target", "t", "", "Target(s), can be a single templated URL or a File of templated URLs")
}
