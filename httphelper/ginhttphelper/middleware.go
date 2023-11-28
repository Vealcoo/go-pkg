package ginhttphelper

import (
	"time"

	"github.com/Vealcoo/go-pkg/warninghelper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func CORSDefault() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		g.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		g.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		g.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(204)
			return
		}

		g.Next()
	}
}

type LogStruct struct {
	Ip     string      `json:"ip,omitempty"`
	Url    string      `json:"url,omitempty"`
	Method string      `json:"method,omitempty"`
	Host   string      `json:"host,omitempty"`
	Body   interface{} `json:"body,omitempty"`
	Time   time.Time   `json:"time,omitempty"`
	Tag    string      `json:"tag,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

func defaultLog(g *gin.Context) *LogStruct {
	return &LogStruct{
		Ip:     g.ClientIP(),
		Url:    g.Request.URL.Path,
		Method: g.Request.Method,
		Host:   g.Request.Host,
		Body:   g.Request.Body,
		Time:   time.Now(),
	}
}

func switchLogFunc(customLog func(ctx *gin.Context) *LogStruct) func(ctx *gin.Context) *LogStruct {
	switch customLog {
	case nil:
		return defaultLog
	default:
		return customLog
	}
}

func LogHandler(log zerolog.Logger, tag string, customLog func(ctx *gin.Context) *LogStruct) gin.HandlerFunc {
	return func(g *gin.Context) {
		l := switchLogFunc(customLog)(g)

		if tag != "" {
			l.Tag = tag
		}

		for _, err := range g.Errors {
			log.Error().Err(err.Err).Fields(l)
		}

		g.Next()
	}
}

func ErrWarning(warningClient warninghelper.WarningService, codeStartFrom int, tag string, customLog func(ctx *gin.Context) *LogStruct) gin.HandlerFunc {
	return func(g *gin.Context) {
		if g.Writer.Status() >= codeStartFrom {
			l := switchLogFunc(customLog)(g)

			if tag != "" {
				l.Tag = tag
			}

			go warningClient.Warning(l)
		}

		g.Next()
	}
}
