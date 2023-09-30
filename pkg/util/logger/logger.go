package logger

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

type CustomLogger interface {
	info(msj string)
	debug(msj string)
	fatal(msj string)
	warn(msj string)
	error(msj string)
}

var Logger CustomLogger

type ZapFacade struct {
	logger *zap.SugaredLogger
}

func (zap *ZapFacade) info(msj string) {
	zap.logger.Info(msj)
}

func (zap *ZapFacade) debug(msj string) {
	zap.logger.Debug(msj)
}

func (zap *ZapFacade) fatal(msj string) {
	zap.logger.Fatal(msj)
}

func (zap *ZapFacade) warn(msj string) {
	zap.logger.Warn(msj)
}

func (zap *ZapFacade) error(msj string) {
	zap.logger.Error(msj)
}

func InitZapLogger() CustomLogger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	zapFacade := &ZapFacade{logger.Sugar()}
	defer logger.Sync()

	return zapFacade
}

func getLoggerInstance() CustomLogger {
	if Logger == nil {
		Logger = InitZapLogger()
	}

	return Logger
}

func Info(msj string) {
	getLoggerInstance().info(msj)
}

func Debug(msj string) {
	getLoggerInstance().debug(msj)
}

func Fatal(msj string) {
	getLoggerInstance().fatal(msj)
}

func Warn(msj string) {
	getLoggerInstance().warn(msj)
}

func Error(msj string) {
	getLoggerInstance().error(msj)
}
