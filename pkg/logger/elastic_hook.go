package logger

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"time"
)

type ElasticHookConfig struct {
	ElasticAddr []string
	ServerName  string
	Host        string
	Index       string
	HookLevel   logrus.Level
}

type elasticHook struct {
	client     *elastic.Client
	serverName string
	host       string
	index      string
	levels     []logrus.Level
	msgCh      chan *message
}

type message struct {
	host      string
	server    string
	level     string
	timestamp string
	message   string
	fields    logrus.Fields
	file      string
	function  string
}

func NewElasticHook(c *ElasticHookConfig) *elasticHook {
	var levels []logrus.Level
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	} {
		if l <= c.HookLevel {
			levels = append(levels, l)
		}
	}
	client, err := elastic.NewClient(elastic.SetURL(c.ElasticAddr...))
	if err != nil {
		panic(fmt.Errorf("logger.NewElasticHook() error(%v)", err))
	}

	hook := &elasticHook{
		client:     client,
		serverName: c.ServerName,
		host:       c.Host,
		index:      c.Index,
		levels:     levels,
		msgCh:      make(chan *message, 1024),
	}

	go hook.fireFunc()
	return hook
}

//消息转换
func formatMessage(entry *logrus.Entry, hook *elasticHook) *message {
	if e, ok := entry.Data[logrus.ErrorKey]; ok && e != nil {
		if err, ok := e.(error); ok {
			entry.Data[logrus.ErrorKey] = err.Error()
		}
	}
	return &message{
		host:      hook.host,
		server:    hook.serverName,
		level:     entry.Level.String(),
		timestamp: entry.Time.Format("2006-01-02 15:04:05"),
		message:   entry.Message,
		fields:    entry.Data,
		file:      fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line),
		function:  entry.Caller.Function,
	}
}

func (hook *elasticHook) fireFunc() {
	messages := make([]*message, 0, 100)
	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case data := <-hook.msgCh:
			if len(messages) > 100 {
				hook.createBulk(messages)
				messages = messages[:0]
			}
			messages = append(messages, data)
		case <-t.C:
			if len(messages) != 0 {
				hook.createBulk(messages)
				messages = messages[:0]
			}
		}
	}
}

func (hook *elasticHook) createBulk(messages []*message) {
	begin := time.Now()
	var docs = make([]interface{}, len(messages))
	for i, v := range messages {
		docs[i] = v
	}
	index := fmt.Sprintf("%s_%s", hook.index, time.Now().Format("2006-01-02"))
	bulk := hook.client.Bulk().Index(index).Type("_doc")
	for _, doc := range docs {
		bulk.Add(elastic.NewBulkIndexRequest().Doc(doc))
	}

	res, err := bulk.Do(context.TODO())
	if err != nil {
		Logger.Error(err)
	}
	if res.Errors {
		Logger.Error("bulk commit failed")
	}
	dur := time.Since(begin).Seconds()
	pps := int64(float64(len(docs)) / dur)
	Logger.Info("%-30s %10d | %10d req/s | %02d:%02d\n", "Insert Error Log Data To ES", len(docs), pps, dur/60, int(dur)%60)
}

// 实现hook接口
func (hook *elasticHook) Fire(entry *logrus.Entry) error {
	hook.msgCh <- formatMessage(entry, hook)
	return nil
}

func (hook *elasticHook) Levels() []logrus.Level {
	return hook.levels
}
