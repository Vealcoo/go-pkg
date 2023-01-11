package ginhttphelper

import (
	"time"

	"github.com/Vealcoo/go-pkg/conversion"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func CORSMiddleware() gin.HandlerFunc {
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

func LoggingMiddleware(log *zerolog.Logger) gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Next()

		for _, err := range g.Errors {
			log.Error().Err(err.Err).Fields(struct {
				Url    interface{}
				Method interface{}
				Host   interface{}
				Body   interface{}
				Time   interface{}
			}{
				Url:    g.Request.URL,
				Method: g.Request.Method,
				Host:   g.Request.Host,
				Body:   g.Request.Body,
				Time:   time.Now(),
			})
		}
	}
}

func ErrorMiddleware(errMap map[error]interface{}) gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Next()

		for _, err := range g.Errors {
			if e, ok := errMap[err.Err]; ok {
				m := map[string]interface{}{}

				res := conversion.Struct2Map(e)
				for k, v := range res {
					m[k] = v
				}

				g.JSON(-1, m)
				continue
			}

			g.JSON(-1, gin.H{"code": -1, "error": "SystemError"})
		}
	}
}
