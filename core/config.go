package core

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/util"
	"net/http"
	"time"
)

type Config struct {
	AutoRetry         bool              `default:"false"`
	MaxRetryTime      int               `default:"3"`
	HttpTransport     *http.Transport   `default:""`
	Transport         http.RoundTripper `default:""`
	EnableAsync       bool              `default:"false"`
	MaxTaskQueueSize  int               `default:"1000"`
	GoRoutinePoolSize int               `default:"5"`
	Timeout           time.Duration
}

func NewConfig() (config *Config) {
	config = &Config{}
	util.InitStructWithDefaultTag(config)
	return
}
