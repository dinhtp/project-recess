package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "project recess migrate command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("project recess migrate command called")
    },
}

func init() {
    rootCmd.AddCommand(migrateCmd)
}
