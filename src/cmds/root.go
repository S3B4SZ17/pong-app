package cmds

import (
	"os"

	"github.com/S3B4SZ17/pong-app/src/app"
	"github.com/S3B4SZ17/pong-app/src/config"
	"github.com/S3B4SZ17/pong-app/src/management"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Execute() {

	err := config.LoadConfig()
	management.InitializeZapCustomLogger()

	if err != nil {
		management.Log.Error("An error occured:", zap.String("error", err.Error()))
		os.Exit(1)
	} else {
		management.Log.Info("Loaded config initial configuration", zap.String("configFile", viper.ConfigFileUsed()))
		app.StartApp()
	}

}
