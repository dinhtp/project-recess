package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
    Use:   "migration",
    Short: "project recess migration command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("project recess migration command called")
    },
}

func init() {
    rootCmd.AddCommand(migrationCmd)
}
