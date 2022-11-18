package cmd

import (
    "errors"

    "github.com/dinhtp/project-recess/database"
    "github.com/dinhtp/project-recess/server"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var echoCmd = &cobra.Command{
    Use:   "echo",
    Short: "Project recess echo command",
    Run:   RunEchoCommand,
}

func init() {
    serveCmd.AddCommand(echoCmd)

    echoCmd.Flags().StringP("address", "", ":8080", "service address")
    echoCmd.Flags().StringP("sqliteDsn", "", "sqlite.db", "sqlite connection string")

    _ = viper.BindPFlag("address", echoCmd.Flags().Lookup("address"))
    _ = viper.BindPFlag("sqliteDsn", echoCmd.Flags().Lookup("sqliteDsn"))
}

func RunEchoCommand(cmd *cobra.Command, args []string) {
    // init DB Connection
    connector := database.NewConnector(database.DbTypeSqLite, viper.GetString("sqliteDsn"))
    if connector == nil {
        panic(errors.New("unsupported database"))
    }

    orm, err := connector.Connect()
    if err != nil {
        panic(err)
    }

    // init HTTP server
    echoServer := server.NewServer(orm, viper.GetString("address"), database.DbTypeSqLite)
    if echoServer == nil {
        panic(errors.New("unsupported http server"))
    }

    echoServer.Serve()
}
