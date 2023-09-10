package logs

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

type CustomLogger interface {
	Info(msj string)
	Debug(msj string)
	Fatal(msj string)
	Warn(msj string)
	Error(msj string)
}

var Logger CustomLogger

type ZapFacade struct {
	logger *zap.SugaredLogger
}

func (zap *ZapFacade) Info(msj string) {
	zap.logger.Info(msj)
}

func (zap *ZapFacade) Debug(msj string) {
	zap.logger.Debug(msj)
}

func (zap *ZapFacade) Fatal(msj string) {
	zap.logger.Fatal(msj)
}

func (zap *ZapFacade) Warn(msj string) {
	zap.logger.Warn(msj)
}

func (zap *ZapFacade) Error(msj string) {
	zap.logger.Error(msj)
}

func InitZapLogger() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	zapFacade := &ZapFacade{logger.Sugar()}
	Logger = zapFacade
	defer logger.Sync()
}
