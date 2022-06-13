package logger

import (
	"time"
	"io"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/lestrrat-go/file-rotatelogs"
)

var Log *zap.SugaredLogger

func init(){
	config := zapcore.EncoderConfig{
		MessageKey:"msg",
		LevelKey:"level",
		TimeKey:"ts",
		CallerKey:"file",
		StacktraceKey:"trace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel:zapcore.LowercaseLevelEncoder,
		EncodeCaller:zapcore.ShortCallerEncoder,
		EncodeTime:func(t time.Time,enc zapcore.PrimitiveArrayEncoder){
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration:func(d time.Duration,enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 10000)
		},
		
		encoder := zapcore.NewConsoleEncoder(config)
		
		info := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			reurn level = zapcore.InfoLevel
		})
		
		err := zap.LevelEnablerFunc(func(level zapcore.Level)bool {
			return level >= zapcore.ErrorLevel
		})
		
		infoWriter := writer("./log/info","%Y%m%d")
		errorWriter := writer("./log/info","%Y%m%d")
		
		core := zapcore.NewTee(
			 zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), info),
      		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), info),
      		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), err),
      		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), err), 
 
		)
		Log = zap.New(core,zap.AddCaller(),zap.AddStacktrace(zap.ErrorLevel)).Sugar()
}

func writer(filename, fileFormat string) io.Writer {
	hook, err := rotatelogs.New(
		filename +"_"+fileFormat+".log",
		rotatelogs.WithLinkName(filename+".log"),
		rotatelogs.WithMaxAge(time.Hour * 24 * 30),
		rotatelogs.WithRotationTime(time.Hour * 24)
	)
	if err != nil {
		panic(err )
	}
	return hook
}