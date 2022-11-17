package cmd

import (
    "errors"
    "fmt"
    "github.com/dinhtp/project-recess/database"
    "os"

    "github.com/dinhtp/project-recess/migration"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "project recess migrate command",
    Run:   runMigrateCommand,
}

func init() {
    migrationCmd.AddCommand(migrateCmd)
}

func runMigrateCommand(cmd *cobra.Command, args []string) {
    // init DB Connection
    connector := database.NewConnector(database.DbTypeSqLite, viper.GetString("sqliteDsn"))
    if connector == nil {
        panic(errors.New("unsupported database"))
    }

    orm, err := connector.Connect()
    if err != nil {
        panic(err)
    }

    err = migration.Migrate(orm)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("migration executed successfully")
    os.Exit(0)
}
