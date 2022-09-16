/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/safalnpane/todo/todoData"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Done will tick the task as done",
	Long: `Done will find the task and tick it as Done so that it will not
be listed by the list command.`,
	Run: DoneRun, 
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func DoneRun(cmd *cobra.Command, args []string) {
	items, err := todoData.ReadItems(dataFile)
	if err != nil {
		fmt.Println(err)
	}
	for _, arg := range args {
		for i, _ := range items {
			arg_int, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
				break
			}
			if items[i].Position == arg_int {
				items[i].Complete()
				continue
			}
		} 
	}
	todoData.SaveItems(dataFile, items)
}
