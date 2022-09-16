/*
Copyright © 2022 Safal Neupane <safal.npane@gmail.com>

*/
package cmd

import (
	"os"
	"log"

	"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
)

var dataFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a simple cli-based todo application",
	Long: `Todo will help you get more done in less time.
It's designed to be as simple as possible to help you
accomplish your goals.`,
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	home, err := homedir.Dir()
	if err != nil {
		log.Printf("Unable to detect home directory. Please set data file using --datafile.")
	}
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile",
	home+string(os.PathSeparator)+".todos.json",
	"data file to store todos")
}


