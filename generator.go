package logr

import (
	"github.com/gogap/logrus_mate"
	"os"
	"github.com/spf13/viper"
	"errors"
	_ "github.com/gogap/logrus_mate/hooks/file"
	"fmt"
	"upEletrcSign/config"
	"github.com/sirupsen/logrus"
)

var initFlg bool
var logger *logrus.Logger

func InitModules() error {
	if !config.HasModuleInit() {
		return errors.New("配置文件未初始化，请先初始化")
	}
	runmode := viper.GetString("MAIN.runmode")
	logFileName := viper.GetString("MAIN.logFileName")
	fmt.Println(runmode)
	fmt.Println(logFileName)
	if logFileName == "" || runmode == "" {
		return errors.New("日志配置文件未配置")
	}

	os.Setenv("RUN_MODE", runmode)

	mate, err := logrus_mate.NewLogrusMate(logrus_mate.ConfigFile(logFileName))
	if err != nil {
		return err
	}

	newLoger := logrus.New()

	if err = mate.Hijack(newLoger, "guft"); err != nil {
		fmt.Println(err)
		return err
	}
	logger = newLoger
	initFlg = true

	return nil
}
