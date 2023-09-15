/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type SeriesStruct struct {
	Title   string `json:"title"`
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
}

// refacCmd represents the refac command
var refacCmd = &cobra.Command{
	Use:   "refac",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, _ := os.ReadDir(".")
		for _, v := range files {
			if v.IsDir() {
				continue
			} else {
				pythonScript := exec.Command("python", "main.py", v.Name())
				out, _ := pythonScript.CombinedOutput()
				var mySeries SeriesStruct
				_ = json.Unmarshal(out, &mySeries)
				folderName := mySeries.Title + "/" + strconv.Itoa(mySeries.Season)
				if mySeries.Title != "" {
					root, _ := os.Getwd()
					err := os.MkdirAll(folderName, 0750)
					if err != nil {
						log.Println("###Error: ", err)
						return
					}
					oldPath := root + "/" + v.Name()
					newPath := root + "/" + folderName + "/" + v.Name()
					_ = os.Rename(oldPath, newPath)
				} else {
					continue
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(refacCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refacCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refacCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
