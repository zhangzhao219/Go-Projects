package tracer_log

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/log"
)

const (
	LogPath    = "logs"
	FileSuffix = "%Y%m%d%H%M.log"
)

var TracerLog = logrus.New()

func InitTracerLog() {

	_, statErr := os.Stat(LogPath)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			mkdirTracerLog()
		} else {
			log.Base().Errorf("tracer log stat error, :%v", statErr)
		}
	}

	var logLevel logrus.Level
	err := logLevel.UnmarshalText([]byte("info"))
	if err != nil {
		log.Base().Errorf("failed to set log level, :%v", err)
	}
	TracerLog.SetOutput(writer(LogPath))
	TracerLog.SetFormatter(&TracerFormatter{})
	TracerLog.SetLevel(logLevel)
}

func writer(logPath string) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logPath, "ads-sort")
	logier, err := rotatelogs.New(
		logFullPath+"-"+FileSuffix,
		rotatelogs.WithLinkName(logFullPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Minute*2),       // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Minute*1), // 日志切割时间间隔
	)

	if err != nil {
		log.Base().Errorf("create rotatelogs err : %v", err)
	}
	return logier
}

func mkdirTracerLog() {
	mdErr := os.Mkdir(LogPath, os.ModePerm)
	if mdErr != nil {
		log.Base().Errorf("tracer log mkdir error, :%v", mdErr)
	}
}

type TracerFormatter struct {
}

func (m *TracerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	var newLog string
	newLog = fmt.Sprintf("%s\n", entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}
