// package logger is responsible for providing functions for writing data on console and http server
package logger

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	types "example.com/types"

	"github.com/gin-gonic/gin"
)

func Println(str interface{}) {
	fmt.Println(str)
}

func Str(in interface{}) string {
	return fmt.Sprint(in)
}

func ConsoleLog(rmsg map[string]map[string]string) {
	i := 0 + 1
	for key, val := range rmsg {
		if key == types.Session+Str(i) {
			for key1, val1 := range val {
				Println(key1 + "\t" + val1)
			}
		}
		i++
	}
}

func RouteLogWriter() gin.HandlerFunc {
	write := strings.Builder{}
	return func(c *gin.Context) {
		i := 1
		for key, val := range types.FinalMsg {
			if key == types.Session+Str(i) {
				for key1, val1 := range val {
					write.WriteString(fmt.Sprintf("%s\t%s\n", key1, val1))
				}
			}
			i++
		}
		write.WriteString(fmt.Sprintf("%v\n", time.Now()))
		c.String(http.StatusOK, write.String())
	}
}

func RouteLog() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/log", RouteLogWriter())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, string("Welcome"))
	})
	router.Run(":" + types.HostAddress)
}
