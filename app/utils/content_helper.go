package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func GetCtxRawData(ctx *gin.Context) string {
	buf, _ := ctx.GetRawData()
	// gin框架中的Request的Body只能取一次，所以我们取出来后还需要赋值进去才可以
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	ctx.Next()

	return string(buf)
}
func GetQueryInt(params *gin.Context, param string) int {
	paramValue, yn := params.GetQuery(param)
	if !yn {
		return 0
	}
	result, err := strconv.Atoi(paramValue)
	if err != nil {
		return 0
	}
	return result
}
func RequestBody2Json(ctx *gin.Context, result interface{}) error {
	buf := GetCtxRawData(ctx)
	err := json.Unmarshal([]byte(buf), result)
	return err
}
func GetStr2Int64(str string, def int64) int64 {
	var result int64
	var err error
	if str != "" {
		result, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			result = def
		}
	} else {
		result = def
	}
	return result
}
func GetStr2Float64(str string, def float64) float64 {
	var result float64
	var err error
	if str != "" {
		result, err = strconv.ParseFloat(str, 64)
		if err != nil {
			result = def
		}
	} else {
		result = def
	}
	return result
}
func GetStr2Int(str string, def int) int {
	var result int
	var err error
	if str != "" {
		result, err = strconv.Atoi(str)
		if err != nil {
			result = def
		}
	} else {
		result = def
	}
	return result
}
func GetStr2UInt64(str string, def uint64) uint64 {
	var result uint64
	var err error
	if str != "" {
		result, err = strconv.ParseUint(str, 10, 64)
		if err != nil {
			result = def
		}
	} else {
		result = def
	}
	return result
}
func GetInt642Str(num int64, def int64) string {
	return strconv.FormatInt(num, 10)
}
