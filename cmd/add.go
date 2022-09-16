/*
Copyright © 2022 Safal Neupane <safal.npane@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/safalnpane/todo/todoData"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add will create a new todo item to the list`,
	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todoData.ReadItems(dataFile)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, x := range args {
		item := todoData.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err = todoData.SaveItems(dataFile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
