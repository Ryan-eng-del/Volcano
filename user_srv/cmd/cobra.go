package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"volcano.user_srv/config"
	"volcano.user_srv/utils"
)


var mode string
var conf string

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run a volcano application",
	Long:  `Run a volcano application by parameter`,
	Args: func(cmd *cobra.Command, args []string) error {
		mode = cmd.Flag("mode").Value.String()
		conf = cmd.Flag("conf").Value.String()
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func CmdExecute() error {
	cmdRun.Flags().StringVarP(&mode, "mode", "m", "dev", "set mode for server")
	cmdRun.Flags().StringVarP(&conf, "conf", "f", "/Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/config", "set mode for server")

	var rootCmd = &cobra.Command{
		Use:               "",
		Short:             "Volcano Command Manager",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true, DisableNoDescFlag: true, DisableDescriptions: true},
	}
	
	rootCmd.AddCommand(cmdRun)
	return rootCmd.Execute()
}

func GetCmdPort() int {
	if config.BaseMapConfInstance.Base.Port == 0 {
		 p, err := utils.GetFreePort()
		 if err != nil {
			 log.Println("[ERROR] GetCmdPort failed: ", err)
			 panic("0 is not a valid port number")
		 }
		 config.BaseMapConfInstance.Base.Port = p	
		 return p
	} 
	return config.BaseMapConfInstance.Base.Port
}

// func GetCmdIp() string {
// 	if ip == "" {
// 		 ip = "127.0.0.1"
// 	}
// 	return ip
// }


func GetCmdMode() string {
	return mode
}


func GetCmdConf() string {
	return conf
}

