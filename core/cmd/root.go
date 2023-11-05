package cmd

import (
	"core/application"
	"core/config"
	"core/utils"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	env     string
	rootCmd = &cobra.Command{
		Use:          application.GetApp().GetName(),
		Short:        application.GetApp().GetName(),
		SilenceUsage: true,
		Long:         application.GetApp().GetName(),
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least one arg")
			}
			return nil
		},
	}
	StartCmd = &cobra.Command{
		Use:     "start",
		Short:   "Start application",
		Example: application.GetApp().GetName() + " start -e dev",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	rootCmd.AddCommand(StartCmd)
	StartCmd.PersistentFlags().StringVarP(&env, "start", "e", "dev", "Setting up the running environment")
}

func run() {
	fmt.Printf("The profile active is %s\n", env)
	application.GetApp().SetEnv(env)
	config.Setup()
	fmt.Println("Start server...")
	var engine = application.GetApp().GetEngine()
	var port = strconv.Itoa(config.SystemConfig.Server.Port)
	fmt.Printf("The server port is %s\n", port)
	var serverAddr = utils.COLON + port
	engine.Run(serverAddr)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Execute err : %s\n", err.Error())
		os.Exit(-1)
	}
}
