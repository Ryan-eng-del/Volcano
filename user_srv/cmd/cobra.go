package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
	"volcano.user_srv/config"
	"volcano.user_srv/initialize"
	"volcano.user_srv/lib"
	"volcano.user_srv/utils"
)

// todo(fix): 在 cobra Long option 中打印不全
// fmt.Println(`
// _     ____  _     ____  ____  _      ____
// / \ |\/  _ \/ \   /   _\/  _ \/ \  /|/  _ \
// | | //| / \|| |   |  /  | / \|| |\ ||| / \|
// | \// | \_/|| |_/\|  \_ | |-||| | \||| \_/|
// \__/  \____/\____/\____/\_/ \|\_/  \|\____/
// 																					`)

var mode string
var conf string
var migrateMode string
var migrateConf string
var migrateStr = []string{"up", "down", "force"}

// up | down | force 5 | up 5 | down 5
var cmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "migrate volcano sql table",
	Version: "1.0.0",
	Example: "volcano migrate up | down | force 5 | up 5 | down 5",
	Long:`
	/ \ |\/  _ \/ \   /   _\/  _ \/ \  /|/  _ \
	| | //| / \|| |   |  /  | / \|| |\ ||| / \|
	| \// | \_/|| |_/\|  \_ | |-||| | \||| \_/|
	\__/  \____/\____/\____/\_/ \|\_/  \|\____/`,
	Args: func(cmd *cobra.Command, args []string) error {
		conf := cmd.Flag("conf").Value.String()
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
        return err
    }
		if err := CheckPath(conf); err != nil {
			return err
		}
		// Run the custom validation logic
		if !utils.IsStringInArray(args[0], migrateStr) {
			return fmt.Errorf("invalid migration position argument %s", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		conf := cmd.Flag("conf").Value.String()
		mode := cmd.Flag("mode").Value.String()

		log.Println(mode, conf, "mode conf")


		// Init instance use viper
		path := fmt.Sprintf("%s/%s/%s", conf, mode, "migrate.toml")
    libViper := lib.NewViper();
		libViper.Init()
		if err := libViper.Unmarshal(path,config.MigrateMapConfInstance, "Go Migrate"); err != nil {
			panic(fmt.Errorf("cmd.cobra.cmdMigrate.Run.Unmarshal: %v", err))
		}
		confInstance := config.MigrateMapConfInstance.Base
		log.Printf("[INFO] viper unmarshal migrateConf %+v", confInstance)


		// Start Migrate
		libMigrate := lib.NewMigrate(confInstance.DatabaseName, confInstance.MigrateLink, confInstance.MigratePath)

		if len(args) == 1 {
			switch command := args[0]; command{
			case "up":
				if err := libMigrate.MigrateUpDatabase(); err != nil {
					panic(fmt.Errorf("cmd.cobra.cmdMigrate.Run.MigrateUpDatabase: %v", err))
				}

			case "down":
				if err := libMigrate.MigrateDownDatabase(); err != nil {
					panic(fmt.Errorf("cmd.cobra.cmdMigrate.Run.MigrateUpDatabase: %v", err))
				}
			default:
			}
			return
	} 

	if len(args) == 2 {
		step, err := strconv.Atoi(args[1]); 
		
	  if err != nil {
			panic(fmt.Errorf("%s is not a valid int", args[1]))
		}

		switch command := args[0]; command{
		case "up":
			if step < 0 {
				panic("when the operation is up, the step value must be greater than zero")
			}
			if err := libMigrate.MigrateStepDatabase(step); err != nil {
				panic(fmt.Errorf("cmd.cobra.cmdMigrate.Run.MigrateUpDatabase: %v", err))
			}
		case "down":
			if step < 0 {
				panic("when the operation is down, the step value must be greater than zero")
			}
			if err := libMigrate.MigrateStepDatabase(step * -1); err != nil {
				panic(fmt.Errorf("cmd.cobra.cmdMigrate.Run.MigrateUpDatabase: %v", err))
			}
		case "force":
			if err := libMigrate.MigrateForceDatabase(step); err != nil {
				panic(fmt.Errorf("cmd.cobra.cmdMigrate.Run.MigrateUpDatabase: %v", err))
			}
		default:
		}
	}
	},
}



var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run volcano application",
	Version: "1.0.0",
	Example: "volcano run -f /config -m dev",
	Long:`
	/ \ |\/  _ \/ \   /   _\/  _ \/ \  /|/  _ \
	| | //| / \|| |   |  /  | / \|| |\ ||| / \|
	| \// | \_/|| |_/\|  \_ | |-||| | \||| \_/|
	\__/  \____/\____/\____/\_/ \|\_/  \|\____/`,
	Args: func(cmd *cobra.Command, args []string) error {
		conf = cmd.Flag("conf").Value.String()
    mode = cmd.Flag("mode").Value.String()

    if err := CheckPath(conf); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	// Init modules such as gorm | viper | nacos | zap
	if err := initialize.InitModules(GetCmdMode(), GetCmdConf()); err != nil {
    panic(err)
	}

	// Using consul for service registration
	consulLib := lib.NewConsul()
	if err := consulLib.Register(); err != nil {
		panic(err)
	}

	if err := initialize.RegisterGrpc(); err != nil {
		panic(err)
	}
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	consulLib.DeRegister()
	},
}

func CmdExecute() error {
	cmdRun.Flags().StringVarP(&mode, "mode", "m", "dev", "set mode for server")
	// my value of conf  is /Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/config
	cmdRun.Flags().StringVarP(&conf, "conf", "f", "", "set configuration file for server") 

  cmdRun.MarkFlagRequired("conf")

	cmdMigrate.Flags().StringVarP(&migrateMode, "mode", "m", "dev", "set mode for migrate")
	// my value of conf  is /Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/config
	cmdMigrate.Flags().StringVarP(&migrateConf, "conf", "f", "", "set configuration file for migrate") 

	cmdMigrate.MarkFlagRequired("conf")


	var rootCmd = &cobra.Command{
		Use:               "volcano",
		Short:             "Volcano Command Line Cli!",
		Version: "1.0.0",
		Example: "volcano run -f /config -m dev",
		Long:`
	 / \ |\/  _ \/ \   /   _\/  _ \/ \  /|/  _ \
	 | | //| / \|| |   |  /  | / \|| |\ ||| / \|
	 | \// | \_/|| |_/\|  \_ | |-||| | \||| \_/|
	 \__/  \____/\____/\____/\_/ \|\_/  \|\____/`,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(cmdMigrate)
	return rootCmd.Execute()
}



func GetCmdMode() string {
	return mode
}


func GetCmdConf() string {
	return conf
}

