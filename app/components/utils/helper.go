package utils

import (
	"errors"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//var (
//	//Logger     = pedestal.NewContextLogger(context.Background())
//	YAMLConfig *options.AppOptions
//)


// SegmentSlice split slice by step size
// elements    segmented string slice
// step    每分数量
// [][]string    the result of segmented string slice
func SegmentSlice(elements []int64, step int) [][]int64 {
	var (
		ret = make([][]int64, 0)
	)

	if step <= 0 {
		step = 10
	}

	l := len(elements)

	n := l / step

	for i := 0; i < n; i++ {
		start := i * step
		end := start + step
		col := elements[start:end]
		ret = append(ret, col)
	}

	// process less than step data
	r := l % step
	if r != 0 {
		start := n * step
		col := elements[start:]
		ret = append(ret, col)
	}

	return ret
}

func TrimAllBlankCharacter(originalStr string) string{
	//TODO
	return strings.Trim(originalStr, " ")
}

// 过滤map类型value为空的k-v
func FilterEmptyValue(data map[string]string) map[string]string {
	if data != nil {
		for k, v := range data {
			if v == "" || v == "0" {
				delete(data, k)
			}
		}
	}

	return data
}

// 类型转换 struct to map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj) // 获取 obj 的类型信息
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { // 如果是指针，则获取其所指向的元素
		t = t.Elem()
		v = v.Elem()
	}

	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}

	}
	return data
}


// 过去随机的数 min ～ max之间
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func BeautifyContentReplaceHtmlLabel(content string) (beautifiedContent string) {

	beautifiedContent = strings.Replace(content, "<p>", "", -1)
	beautifiedContent = strings.Replace(beautifiedContent, "</p>", "\r\n", -1)
	beautifiedContent = strings.Replace(beautifiedContent, "<br>", "\r\n", -1)
	beautifiedContent = strings.Replace(beautifiedContent, "</br>", "\r\n", -1)

	return
}

// []string 转化成 []int64
func SliceStringInt64(data []string) (out []int64, err error) {
	if len(data) < 1 {
		return nil, errors.New("参数错误，转化失败")
	}

	out = make([]int64, 0)
	for _, v := range data {
		d, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}

		out = append(out, d)
	}

	return out, nil
}

func ValueInArray(val string, array []string) (exists bool) {
	exists = false

	for _, v := range array {
		if val == v {
			exists = true
			break
		}
	}

	return exists
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