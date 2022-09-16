/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/safalnpane/todo/todoData"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todo items",
	Long:  `List command will read the files and check if there are any todo items`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todoData.ReadItems(dataFile)
	if err != nil {
		fmt.Printf("%v", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		fmt.Fprintln(w, i.Label()+"\t"+i.PrettyP()+"\t"+i.Text+"\t")
	}
	w.Flush()
}
