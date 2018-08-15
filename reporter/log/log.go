/*
Package log implements a reporter to send spans in V2 JSON format to the Go
standard Logger.
*/
package log

import (
	"encoding/json"
	"fmt"
	"github.com/qutoutiao/zipkin-go/model"
	"github.com/qutoutiao/zipkin-go/reporter"
	"github.com/Sirupsen/logrus"
)

// logReporter will send spans to the default Go Logger.
type logReporter struct {
	logger *logrus.Logger
}

type ZipkinFormatter struct {
}

// Format renders a single log entry
func (f *ZipkinFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := fmt.Sprintf("%s\n", entry.Message)

	return []byte(data), nil
}

// NewReporter returns a new log reporter.
func NewReporter(l *logrus.Logger) reporter.Reporter {
	return &logReporter{
		logger: l,
	}
}

// Send outputs a span to the Go logger.
func (r *logReporter) Send(s model.SpanModel) {
	var t []model.SpanModel
	t = append(t, s)
	if b, err := json.Marshal(t); err == nil {
		r.logger.Info(string(b))
	}
}

// Close closes the reporter
func (*logReporter) Close() error { return nil }
