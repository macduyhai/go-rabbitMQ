package middleware

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/macduyhai/go-rabbitMQ/producer/utils"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Content-Type", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type CheckAPIkey struct {
	Apikey string
}

func (c *CheckAPIkey) Check(context *gin.Context) {
	api := context.Request.Header.Get("api-key")
	if api == c.Apikey {
		context.Next()
	} else {
		context.AbortWithStatus(401)
	}
}

func SetUserID(context *gin.Context) {
	err := utils.SetUserID(context)
	if err != nil {
		utils.Response(context, nil, 401, "parse userID error: "+err.Error())
		context.Abort()
	}
	context.Next()
}

// add

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		// rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		//log.Println(readBody(rdr2)) // Print request body

		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
