package cmd

import (
    "errors"
    "github.com/dinhtp/project-recess/database"
    "github.com/dinhtp/project-recess/server"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var muxCmd = &cobra.Command{
    Use:   "mux",
    Short: "Project recess mux command",
    Run:   RunMuxCommand,
}

func init() {
    serveCmd.AddCommand(muxCmd)

    serveCmd.Flags().StringP("address", "", ":8080", "service address")
    serveCmd.Flags().StringP("mysqlDsn", "", "", "mysql DSN connection string")

    _ = viper.BindPFlag("address", serveCmd.Flags().Lookup("address"))
    _ = viper.BindPFlag("mysqlDsn", serveCmd.Flags().Lookup("mysqlDsn"))
}

func RunMuxCommand(cmd *cobra.Command, args []string) {
    // init DB Connection
    connector := database.NewConnector(database.DbTypeMySql, viper.GetString("mysqlDsn"))
    if connector == nil {
        panic(errors.New("unsupported database"))
    }

    orm, err := connector.Connect()
    if err != nil {
        panic(err)
    }

    // init HTTP server
    muxServer := server.NewServer(orm, viper.GetString("address"), database.DbTypeMySql)
    if muxServer == nil {
        panic(errors.New("unsupported http server"))
    }

    muxServer.Serve()
}
