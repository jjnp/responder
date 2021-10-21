package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.Any("/static", static)
	r.Any("/uniform", uniform)
	r.Any("/lognorm", lognorm)
	r.Run()
}

const sampleJSON = "{\n  \"first\": 0.12345,\n  \"second\": 0.12345,\n  \"third\": 0.12345,\n  \"forth\": 0.12345,\n  \"fifth\": 0.12345,\n  \"sixth\": 0.12345\n}"
const defaultStatic = 25
const defaultUniformFrom = 10
const defaultUniformTo = 30
const defaultLognormMean = 25
const defaultLognormStdev = 15

func static(c *gin.Context) {
	params := struct{
		Time float64 `form:"time"`
	}{
		Time: defaultStatic,
	}
	c.ShouldBindQuery(&params)
	sleepAndRespond(c, time.Duration(params.Time) * time.Millisecond)
}

func uniform(c *gin.Context) {
	params := struct {
		From int64 `form:"from"`
		To int64 `form:"to"`
	}{
		From: defaultUniformFrom,
		To: defaultUniformTo,
	}
	c.BindQuery(&params)
	delay := uniformRand(params.From, params.To)
	sleepAndRespond(c, time.Millisecond * time.Duration(delay))
}

func lognorm(c *gin.Context) {
	params := struct {
		Mean float64 `form:"mean"`
		Stdev float64 `form:"stdev"`
	}{
		Mean:  defaultLognormMean,
		Stdev: defaultLognormStdev,
	}
	c.BindQuery(&params)
	delay := lognormalRand(params.Mean, params.Stdev)
	sleepAndRespond(c, time.Millisecond * time.Duration(delay))
}

func sleepAndRespond(c *gin.Context, sleepTime time.Duration) {
	time.Sleep(sleepTime)
	c.JSON(http.StatusOK, gin.H{
		"first": 1234.1234,
		"second": 1234.1234,
		"third": 1234.1234,
		"fourth": 1234.1234,
		"fifth": 1234.1234,
	})
}

