package cmd

import (
    "errors"

    "github.com/dinhtp/project-recess/database/connection"
    "github.com/dinhtp/project-recess/server"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "A brief description of your command",
    Run:   RunServeCommand,
}

func init() {
    rootCmd.AddCommand(serveCmd)

    serveCmd.Flags().StringP("address", "", "service-address", "service address")
    serveCmd.Flags().StringP("sqliteDsn", "", "sqlite-dsn", "sqlite connection string")

    _ = viper.BindPFlag("address", serveCmd.Flags().Lookup("address"))
    _ = viper.BindPFlag("sqliteDsn", serveCmd.Flags().Lookup("sqliteDsn"))
}

func RunServeCommand(cmd *cobra.Command, args []string) {
    // init DB Connection
    connector := connection.NewConnector(connection.DbTypeSqLite, viper.GetString("sqliteDsn"))
    if connector == nil {
        panic(errors.New("unsupported database"))
    }

    orm, err := connector.Connect()
    if err != nil {
        panic(err)
    }

    // init HTTP server
    server.NewServer(orm, viper.GetString("address")).Serve()
}
