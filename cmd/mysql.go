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

var mysqlCmd = &cobra.Command{
    Use:   "mysql",
    Short: "Project recess mysql command",
    Run:   RunMySqlCommand,
}

func init() {
    migrateCmd.AddCommand(mysqlCmd)
}

func RunMySqlCommand(cmd *cobra.Command, args []string) {
    connector := database.NewConnector(database.DbTypeMySql, viper.GetString("mysqlDSN"))
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
