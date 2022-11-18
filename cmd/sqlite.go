package cmd

import (
    "errors"
    "fmt"
    "os"

    "github.com/dinhtp/project-recess/database"
    "github.com/dinhtp/project-recess/migration"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var sqliteCmd = &cobra.Command{
    Use:   "sqlite",
    Short: "A brief description of your command",
    Run:   runSqLiteCommand,
}

func init() {
    migrateCmd.AddCommand(sqliteCmd)
}

func runSqLiteCommand(cmd *cobra.Command, args []string) {
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
