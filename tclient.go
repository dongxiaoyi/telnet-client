package main

import (
	"fmt"
	tc "github.com/reiver/go-telnet"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.SugaredLogger


func init() {
	Logger = NewLogger()
}


func main() {
	var rootCmd = &cobra.Command{Use: "telnet"}

	var cmdTelnetClient = &cobra.Command{
		Use:   "client",
		Short: "telnet的客户端",
		Long: `telnet的客户端.

示例：telnet client host port
`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("###################Telnet Client##################")
			if len(args) > 2 {
				Logger.Error("Args must only contains [host] [port] !")
				fmt.Println("###################Error Exist##################")
			}

			var caller tc.Caller = tc.StandardCaller

			arg := fmt.Sprintf(`%s:%s`, args[0], args[1])
			err := tc.DialToAndCall(arg, caller)
			if err != nil {
				Logger.Error(err)
				fmt.Println("###################Error Exist##################")
			}
		},
	}
	rootCmd.AddCommand(cmdTelnetClient)
	err := rootCmd.Execute()
	if err != nil {
		Logger.Error(err)
	}
}


// 日志器
func LogLevel() map[string]zapcore.Level {
	level := make(map[string]zapcore.Level)
	level["debug"] = zap.DebugLevel
	level["info"] = zap.InfoLevel
	level["warn"] = zap.WarnLevel
	level["error"] = zap.ErrorLevel
	level["dpanic"] = zap.DPanicLevel
	level["panic"] = zap.PanicLevel
	level["fatal"] = zap.FatalLevel
	return level
}

// 初始化日志
func NewLogger() *zap.SugaredLogger {
	logLevelOpt := "DEBUG" // 日志级别
	levelMap := LogLevel()
	logLevel, _ := levelMap[logLevelOpt]
	atomicLevel := zap.NewAtomicLevelAt(logLevel)

	encodingConfig := zapcore.EncoderConfig{
		TimeKey: "Time",
		LevelKey: "Level",
		NameKey: "Log",
		CallerKey: "Celler",
		MessageKey: "Message",
		StacktraceKey: "Stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}
	var outPath []string
	var errPath []string
	outPath = append(outPath, "stdout")
	errPath = append(outPath, "stderr")

	logCfg := zap.Config{
		Level: atomicLevel,
		Development: true,
		DisableCaller: true,
		DisableStacktrace: true,
		Encoding:"console",
		EncoderConfig: encodingConfig,
		// InitialFields: map[string]interface{}{filedKey: fieldValue},
		OutputPaths: outPath,
		ErrorOutputPaths: errPath,
	}

	logger, _ := logCfg.Build()
	Logger = logger.Sugar()
	return Logger
}
