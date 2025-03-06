package filelog

import (
	"log"
	"os"
	"sync"
)

type fileLogger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	logFile       *os.File
	// writer        io.Writer
}

func (f *fileLogger) LogInfo(data []byte) {
	f.infoLogger.Println(string(data))
}

func (f *fileLogger) LogWarning(data []byte) {
	f.warningLogger.Println(string(data))
}

func (f *fileLogger) LogError(data []byte) {
	f.errorLogger.Println(string(data))
}

func (f *fileLogger) CloseLogFile() {
	err := f.logFile.Sync()
	if err != nil {
		log.Println("sync log file error: ", err)
	}

	err = f.logFile.Close()
	if err != nil {
		log.Println("close log file error: ", err)
	}
}

var (
	logger      *fileLogger
	logFilePath string
	once        sync.Once
)

func SetLogFilePath(filePath string) {
	if filePath == "" {
		log.Panicln("file path can't be empty")
	}

	logFilePath = filePath
}

// Call SetLogFilePath() before Get instance
func Get() *fileLogger {
	once.Do(initLogger)
	return logger
}

func initLogger() {
	if logFilePath == "" {
		log.Panicln("file path can't be empty")
	}

	logger = new(fileLogger)

	flag := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(logFilePath, flag, 0644)
	if err != nil {
		log.Panicln("open log file '" + logFilePath + "' error: " + err.Error())
	}
	logger.logFile = file

	flag = log.LstdFlags | log.Lmsgprefix
	logger.infoLogger = log.New(logger.logFile, "info: ", flag)
	logger.warningLogger = log.New(logger.logFile, "warning: ", flag)
	logger.errorLogger = log.New(logger.logFile, "error: ", flag)
}
