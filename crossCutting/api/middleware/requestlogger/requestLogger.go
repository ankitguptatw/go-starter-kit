package requestlogger

import (
	"myapp/crossCutting/logger"
	"myapp/crossCutting/util"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	TimeFormat string
	UTC        bool
	SkipPaths  []string
}

func Handler(conf *Config) func(c *gin.Context) {

	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		lgr := logger.GetLogger(c)

		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					lgr.Error(e)
				}
			} else {
				format := "status:%d, method:%s, path:%s, query:%s, ip:%s, user-agent:%s, latency:%d"
				resp := util.Format(format, c.Writer.Status(), c.Request.Method, path, query,
					c.ClientIP(), c.Request.UserAgent(), latency)
				if conf.TimeFormat != "" {
					resp += ", time:" + end.Format(conf.TimeFormat)
				}
				lgr.Info(resp)
			}
		}
	}
}
