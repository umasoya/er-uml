package cmd

import (
	"fmt"
	"os"

	_ "github.com/caarlos0/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/umasoya/er-uml/pkg/config"
	"github.com/umasoya/er-uml/pkg/connection"
	// "github.com/umasoya/er-uml/pkg/mysql"
)

const (
	CONFIG_FILE_PATH string = "." // working dir
	CONFIG_FILE_EXT  string = "yaml"
	CONFIG_FILE_NAME string = "er-uml"
	PREFIX           string = "ER"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "er-uml",
	Short: "ER-diagrams generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			switch config.Conf.Driver {
			case "mysql":
				if err := mysql.Run(); err != nil {
					panic(err)
				}
			default:
				panic("unsupport driver: " + config.Conf.Driver)
			}
		*/
		fmt.Printf("%#v\n", config.Conf)
		if err := connection.Execute(&config.Conf); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "er-uml.yaml", "config file (default is er-uml.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in working directory with name "er-uml" (without extension).
		viper.AddConfigPath(CONFIG_FILE_PATH)
		viper.SetConfigType(CONFIG_FILE_EXT)
		viper.SetConfigName(CONFIG_FILE_NAME)
	}

	viper.SetEnvPrefix(PREFIX)

	// read in environment variables that match
	viper.AutomaticEnv()

	// set default values in config
	setDefaultConf()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Copy config file content to config.Conf
	if err := viper.Unmarshal(&config.Conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setDefaultConf() {
	viper.SetDefault("DRIVER", "mysql")
	viper.SetDefault("USER", "root")
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "3306")
	viper.SetDefault("PASSWORD", "")
	viper.SetDefault("DB", "")
}
