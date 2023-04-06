package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/umasoya/er-uml/pkg/config"
	"github.com/umasoya/er-uml/pkg/connection"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "er-uml",
	Short: "ER-diagrams generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("configFile: %s\nconfig: %#v\n", cfgFile, config.Conf)

		// db connection open
		db, err := connection.Open(&config.Conf)
		if err != nil {
			// @todo error handling
			panic(err)
		}

		err = db.Ping()
		if err != nil {
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
		viper.AddConfigPath(".") // working dir
		viper.SetConfigType("yaml")
		viper.SetConfigName("er-uml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	setDefaultConf() // set default values in config

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Copy config file content to config.Conf
	if err := viper.Unmarshal(&config.Conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setDefaultConf() {
	viper.SetDefault("Driver", "mysql")
	viper.SetDefault("User", "root")
	viper.SetDefault("User", "root")
	viper.SetDefault("Host", "localhost")
	viper.SetDefault("Port", "3306")
}
