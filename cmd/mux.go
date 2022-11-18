package cmd

import (
    "errors"
    "github.com/dinhtp/project-recess/database"
    "github.com/dinhtp/project-recess/server"
    "github.com/sirupsen/logrus"
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

    muxCmd.Flags().StringP("address", "", "0.0.0.0:8050", "service address")
    muxCmd.Flags().StringP("mysqlDsn", "", "root:root@tcp(10.0.255.174:3306)/recess?charset=utf8mb4", "mysql DSN connection string")

    _ = viper.BindPFlag("address", muxCmd.Flags().Lookup("address"))
    _ = viper.BindPFlag("mysqlDsn", muxCmd.Flags().Lookup("mysqlDsn"))
}

func RunMuxCommand(cmd *cobra.Command, args []string) {
    serverAddress := viper.GetString("address")

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
    muxServer := server.NewServer(orm, serverAddress, database.DbTypeMySql)
    if muxServer == nil {
        panic(errors.New("unsupported http server"))
    }

    logrus.WithFields(logrus.Fields{
        "address": serverAddress,
        "type":    "mux",
    }).Info("mux http server started successfully")

    muxServer.Serve()
}
