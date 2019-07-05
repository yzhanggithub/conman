package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/iximiuz/conman/server"
)

var optContName string

func init() {
	contBaseCmd.PersistentFlags().StringVarP(&optContName,
		"name", "n",
		"",
		"Container name (required)")
	contBaseCmd.MarkPersistentFlagRequired("name")
	contBaseCmd.AddCommand(contCreateCmd)
	rootCmd.AddCommand(contBaseCmd)
}

var contBaseCmd = &cobra.Command{
	Use:   "container",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal("action required")
	},
}

var contCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client, conn := connect()
		defer conn.Close()

		resp, err := client.CreateContainer(
			context.Background(),
			&server.CreateContainerRequest{
				Name: optContName,
			},
		)
		logrus.Info(resp, err)
	},
}

func connect() (server.ConmanClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("unix://"+optHost, grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	return server.NewConmanClient(conn), conn
}
