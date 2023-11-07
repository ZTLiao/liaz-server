package cmd

import (
	"core/application"
	"core/config"
	"core/utils"
	"core/web"
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
	setup()
	start()
}

func setup() {
	fmt.Printf("The profile active is %s\n", env)
	application.GetApp().SetEnv(env)
	config.Setup()
}

func start() {
	var server = config.SystemConfig.Server
	if server == nil {
		return
	}
	fmt.Println("Start server...")
	var engine = application.GetApp().GetEngine()
	//初始化
	web.InitRouter(engine)
	//端口
	var port = strconv.Itoa(server.Port)
	var serverAddr = utils.COLON + port
	engine.Run(serverAddr)
	fmt.Printf("The server port is %s\n", port)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Execute err : %s\n", err.Error())
		os.Exit(-1)
	}
}
