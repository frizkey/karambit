package flogger

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"os"
)

func init() {
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func SetLogLevel(level string) {
	parseLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		return
	}
	zerolog.SetGlobalLevel(parseLevel)
}

// Debug is
func Debug(msg string) {
	zlog.Debug().Msg(msg)
}

// Debugf is
func Debugf(msg string, data ...interface{}) {
	zlog.Debug().Msgf(msg, data...)
}

// Info is
func Info(msg string) {
	zlog.Info().Msg(msg)
}

// Infof is
func Infof(msg string, data ...interface{}) {
	zlog.Info().Msgf(msg, data...)
}

// Trace is
func Trace(msg string) {
	zlog.Trace().Msg(msg)
}

// Tracef is
func Tracef(msg string, data ...interface{}) {
	zlog.Trace().Msgf(msg, data...)
}

// Warnf is
func Warnf(msg string, data ...interface{}) {
	zlog.Warn().Stack().Msgf(msg, data...)
}

// Error is
func Error(msg string) {
	zlog.Error().Stack().Msg(msg)
}

// Errorf is
func Errorf(msg string, data ...interface{}) {
	zlog.Error().Stack().Msgf(msg, data...)
}

// Fatal is
func Fatal(msg string) {
	zlog.Fatal().Stack().Msg(msg)
}

// Fatalf is
func Fatalf(msg string, data ...interface{}) {
	zlog.Fatal().Stack().Msgf(msg, data...)
}

// Panic is
func Panic(msg string) {
	zlog.Panic().Stack().Msg(msg)
}

// Panicf is
func Panicf(msg string, data ...interface{}) {
	zlog.Panic().Stack().Msgf(msg, data...)
}
