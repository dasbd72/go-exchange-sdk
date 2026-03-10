package main

import (
	"github.com/spf13/cobra"
)

var (
	useLog bool
)

func main() {
	root := &cobra.Command{
		Use:   "experimental",
		Short: "Experimental commands",
	}
	maxCmd := &cobra.Command{
		Use:   "max",
		Short: "Max commands",
		Run:   Max,
	}
	balanceCmd := &cobra.Command{
		Use:   "balance",
		Short: "Get balance",
		RunE:  Balance,
	}
	root.PersistentFlags().BoolVarP(&useLog, "log", "l", false, "Use log")
	root.AddCommand(maxCmd)
	root.AddCommand(balanceCmd)
	root.Execute()
}
