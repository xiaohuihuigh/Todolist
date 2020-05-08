package utils

import (
	"todoList/app/components/utils"
	"reflect"
	"strconv"
	"strings"
)

func String2Arr(s string)[]string{
	var strArr  []string
	for _,item:=range strings.Split(s,","){
		if item == ""{
			continue
		}
		strArr = append(strArr,item)
	}
	return strArr
}
func String2Int64Arr(s string)[]int64{
	var strArr  []int64
	for _,item:=range strings.Split(s,","){
		if item == ""{
			continue
		}
		strArr = append(strArr,utils.GetStr2Int64(item,0))
	}
	return strArr
}
func StringArr2String(strs []string)(result string){
	result = ""
	for i,str := range strs{
		if str  == ""{
			continue
		}
		if i == len(strs){
			result += str
		} else{
			result +=str+","
		}
	}
	if result == ""{
		result = ","
	}
	return
}
func Int64Arr2String(ints []int64)(result string){
	result = ""
	for i,str := range ints{
		if i == len(ints){
			result += strconv.FormatInt(str,10)
		} else{
			result +=strconv.FormatInt(str,10)+","
		}
	}
	return
}
func ExistItem(array interface{},item interface{})bool{
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i:=0;i<s.Len();i++{
			if reflect.DeepEqual(item,s.Index(i).Interface()){
				return true
			}
		}
	}
	return false
}
func DeleteItemInString(str string ,item string)( result string){
	oldStrArr:=String2Arr(str)
	var newStrArr []string
	for _,strItem:= range oldStrArr{
		if strItem == item{
			continue
		}
		newStrArr = append(newStrArr,strItem)
	}
	return StringArr2String(newStrArr)
}


func MergeNewAndOld(new string,old string) (delete string,add string){
	newArr := String2Arr(new)
	oldArr := String2Arr(old)

	var deleteArr []string
	var addArr []string
	for _,new:=range newArr{
		if ! ExistItem(oldArr,new){
			addArr = append(addArr,new)
		}
	}
	for _,old:=range oldArr{
		if ! ExistItem(newArr,old){
			deleteArr = append(deleteArr,old)
		}
	}
	return StringArr2String(deleteArr),StringArr2String(addArr)
}






