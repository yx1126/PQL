package logger

import (
	"sync"
	"time"
	"unicode"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var (
	instances = make(map[string]*Logger, 0)
	mu        sync.RWMutex
	ErrorPath = ".PQL/log/error.log"
	InfoPath  = ".PQL/log/PQL.log"
)

// 目标时间格式
const timeFmt = time.DateTime

type Logger struct {
	lumber *lumberjack.Logger
	log    zerolog.Logger
	path   string
}

func New(path string) *Logger {
	if path == "" {
		return nil
	}
	mu.RLock()
	inst, ok := instances[path]
	mu.RUnlock()
	if ok {
		return inst
	}

	mu.Lock()
	defer mu.Unlock()

	lumber := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    64,
		MaxBackups: 15,
		MaxAge:     0,
		Compress:   true,
		LocalTime:  true,
	}
	fileCW := zerolog.ConsoleWriter{
		Out:     lumber,
		NoColor: true,
		FormatTimestamp: func(v any) string {
			return "[" + v.(string) + "]"
		},
		FormatLevel: func(v any) string {
			runes := []rune(v.(string))
			runes[0] = unicode.ToUpper(runes[0])
			return "\t" + string(runes)
		},
		FormatCaller: func(v any) string {
			return "\t" + v.(string)
		},
		FormatMessage: func(v any) string {
			return "\t" + v.(string)
		},
	}
	multiWriter := zerolog.MultiLevelWriter(fileCW)

	zerolog.TimeFieldFormat = timeFmt
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log := zerolog.New(multiWriter).
		With().
		Timestamp().
		Caller().
		Logger()
	inst = &Logger{
		lumber: lumber,
		log:    log,
		path:   path,
	}
	instances[path] = inst
	return inst
}

func (l *Logger) Close() error {
	mu.Lock()
	defer mu.Unlock()
	delete(instances, l.path)
	return l.lumber.Close()
}

func (l *Logger) Debug(msg string, err error) {
	l.log.Debug().Err(err).Msg(msg)
}

func (l *Logger) Error(msg string, err error) {
	l.log.Error().Err(err).Msg(msg)
}

func (l *Logger) Info(msg string, err error) {
	l.log.Info().Err(err).Msg(msg)
}
