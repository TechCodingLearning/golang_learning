package main

import (
	"GolangProjects/gormSyntax"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strings"
	"time"
)

const (
	MaxDBRange    = 100
	MaxTableRange = 100
)

type dateOldError struct{}
type dateNewError struct{}

func (e *dateOldError) Error() string {
	return "detail date is too old"
}

func (e *dateNewError) Error() string {
	return "detail date is too new"
}

// ValidateDateRange 校验时间范围
func ValidateDateRange(date time.Time, forwardDays, backwardDays uint32) error {
	fmtTime := "2006-01-02 15:04:05"
	timeStr := "2022-09-16 00:06:53"
	localTimeDate, _ := time.ParseInLocation(fmtTime, timeStr, time.Local)
	startDate := localTimeDate.AddDate(0, 0, -int(forwardDays))
	endDate := localTimeDate.AddDate(0, 0, int(backwardDays))

	//startDate := time.Now().AddDate(0, 0, -int(forwardDays))
	//endDate := time.Now().AddDate(0, 0, int(backwardDays))

	fmt.Println("startDate:", startDate)
	fmt.Println("endDate:", endDate)
	if date.Before(startDate) {
		return &dateOldError{}
	}
	if date.After(endDate) {
		return &dateNewError{}
	}
	return nil
}

// CalStorageId 计算StorageID，返回XXYY, XX: 表示在哪套DB YY: 表示在哪张表
func CalStorageId(detailID string, dbCount uint32, tableCount uint32) string {
	// 先计算hash值，并把hash最后4位字节转化为uint32整数
	hashUint32 := calHashUint32(detailID)
	fmt.Println("hash int:", hashUint32)
	// Table最大范围, 10000 = 100表/DB * 100套DB
	tableRange := hashUint32 % (MaxDBRange * MaxTableRange)
	// DB计算：取table最大范围的前两位，根据最大DB数量取模，最后加1，因为db编号必须从1开始
	dbID := (tableRange/MaxDBRange)%dbCount + 1
	// Table计算：取table最大范围的后两位，根据最大表数量取模
	tableID := (tableRange % MaxTableRange) % tableCount

	return fmt.Sprintf("%02d%02d", dbID, tableID)
}

func calHashUint32(input string) uint32 {
	// 计算md5值, hash值是128bit = 16个字节
	hash := md5.Sum([]byte(input))
	// 讲hash值转化为int32，需要4个字节，uint32最大值是42亿
	return binary.BigEndian.Uint32(hash[12:])
}

// CalPageID 计算页表id
func CalPageID() string {
	// pageID：作用：快速识别同一批明细的重入
	// 要求：pageID不能重复，否则会被当作重入直接返回成功，导致重复明细无法拦截
	// 格式：md5(subbatchid) + md5(明细信息), 长度64位字符串
	// 说明：第一个hash范围缩小到同一subbatchid；
	// 第二个hash用于标识同一批明细；这样会大大降低不同subbatchid直接hash冲突概率

	// 1.subbatchid hash
	hash1 := md5.New()
	io.WriteString(hash1, "10001_1")
	hash1Str := hex.EncodeToString(hash1.Sum(nil))

	// 2.明细 hash

	hash2 := md5.New()

	io.WriteString(hash2, "1000112")
	io.WriteString(hash2, "123")
	io.WriteString(hash2, "1234")

	hash2Str := hex.EncodeToString(hash2.Sum(nil))

	return hash1Str + hash2Str
}

// StructToMapStr struct转map
func StructToMapStr(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		iv := v.Field(i)
		if field.PkgPath != "" || iv.IsNil() {
			continue // 跳过非Public成员和nil对象
		}

		if iv.Kind() == reflect.Ptr {
			iv = iv.Elem()
		}

		strValue := fmt.Sprint(iv.Interface())
		data[field.Name] = strValue

	}
	return data
}

type A struct {
	Listid      *string
	BusiListid  *string
	AccountType *string
}

// CamelToSnake 驼峰式命令转为小写+下划线命名
func CamelToSnake(s string) string {
	re := regexp.MustCompile("(?U)[^[:alnum:]]+")
	s = re.ReplaceAllString(s, "_")
	s = strings.Trim(s, "_")
	s = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}

// CompareAndUpdate src与dst比较field，如果不一致更新dst；
// 如果fields为nil，则全字段比较
// src,dst必须传指针
func CompareAndUpdate(src interface{}, dst interface{}, fields []string) (err error) {
	srcType := reflect.Indirect(reflect.ValueOf(src)).Type()
	dstType := reflect.Indirect(reflect.ValueOf(dst)).Type()

	if srcType != dstType { // src, dst必须同类型
		return errors.New("type disconsistent")
	}

	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcValue.Kind() != reflect.Ptr || dstValue.Kind() != reflect.Ptr { // src, dst必须同为指针
		return errors.New("src dst must be ptr.")
	}

	srcValue = srcValue.Elem()
	dstValue = dstValue.Elem()

	if fields == nil {
		for i := 0; i < srcType.NumField(); i++ {
			fields = append(fields, srcType.Field(i).Name)
		}
	}

	for _, field := range fields {
		srcField := srcValue.FieldByName(field)
		dstField := dstValue.FieldByName(field)
		if dstField.IsValid() && !reflect.DeepEqual(srcField.Interface(), dstField.Interface()) {
			dstField.Set(srcField)
		}
		//if srcField.Interface() != dstField.Interface() {
		//	if dstField.IsValid() {
		//		switch dstField.Kind() {
		//		case reflect.Ptr:
		//			if srcField.IsNil() {
		//				dstField.Set(reflect.New(dstField.Type().Elem()))
		//			} else {
		//				dstField.Set(srcField)
		//			}
		//		case reflect.Slice:
		//			if srcField.IsNil() {
		//				dstField.Set(reflect.New(dstField.Type().Elem()))
		//			} else {
		//				dstField.Set(srcField)
		//			}
		//		default:
		//			dstField.Set(srcField)
		//		}
		//	}
		//}
	}

	return nil
}

func CheckBatchValues(actField, expField interface{}) error {
	actRefField := reflect.Indirect(reflect.ValueOf(actField))
	expRefField := reflect.Indirect(reflect.ValueOf(expField))
	if expRefField.IsValid() && !reflect.DeepEqual(actRefField.Interface(), expRefField.Interface()) {
		return errors.New("check batch failed!")
	}
	fmt.Printf("%v check pass! expect value: %v, actual value: %v", expRefField.Type(), expField, actField)
	return nil
}

func CheckBatchFields(actField, expField interface{}, fieldNames []string) error {
	actRefType := reflect.Indirect(reflect.ValueOf(actField)).Type()
	expRefType := reflect.Indirect(reflect.ValueOf(expField)).Type()

	if actRefType != expRefType {
		return errors.New("type failed!")
	}

	actRefValue := reflect.ValueOf(actField)
	expRefValue := reflect.ValueOf(expField)

	if fieldNames == nil {
		for i := 0; i < expRefType.NumField(); i++ {
			fieldNames = append(fieldNames, expRefType.Field(i).Name)
		}
	}

	for _, fieldName := range fieldNames {
		actField := actRefValue.FieldByName(fieldName)
		expField := expRefValue.FieldByName(fieldName)
		if expField.IsValid() && !reflect.DeepEqual(actField.Interface(), expField.Interface()) {
			return errors.New(fmt.Sprintf("%v check failed! expect value: %v, actual value:%v", fieldName, expField, actField))
		}
	}
	return nil
}

func GetSucMsg(key string, actValue, expValue interface{}) string {
	return fmt.Sprintf("%v check succ! actual value: %v, expect value: %v", key, actValue, expValue)
}

type Person struct {
	Name    string
	Age     int
	House   []string
	Address *string
}

func main() {
	//baseSyntax.TestInterface()
	//
	//add1 := "add1"
	//a := Person{
	//	Name:    "tom",
	//	Age:     1,
	//	House:   []string{"a", "b"},
	//	Address: &add1,
	//}
	//aValue := reflect.ValueOf(a)
	//aType := reflect.TypeOf(a)
	//fmt.Println(aValue)
	//fmt.Println(aType)
	////var fields []reflect.Value
	////fmt.Println(aValue.Field())
	////for i := 0; i < aType.NumField(); i++ {
	////	fieldType := aType.Field(i)
	////	fmt.Printf("%v %v \n", fieldType.Name, fieldType.Tag)
	////}
	////fmt.Println(fields)
	////return
	//add2 := "add2"
	//b := Person{
	//	Name:    "tom",
	//	Age:     2,
	//	House:   []string{"c", "d"},
	//	Address: &add2,
	//}
	//fmt.Println(a)
	//fmt.Println(b)
	//
	////c := "a"
	////d := "b"
	////fmt.Println(GetSucMsg("aaa", c, d))
	//err := CheckBatchFields(a, b, []string{"House"})
	////err := CheckBatchFields(a, a)
	////err := CompareAndUpdate(&a, &b, []string{"Age", "House"})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(a)
	//fmt.Println(b)
	//ans := CamelToSnake("ListId")
	//log.Println(ans)
	//inputMap := map[string]string{
	//	"task_id": "aaa",
	//	"module":  "bbb",
	//}
	//inMarshal, _ := json.Marshal(inputMap)
	//str := string(inMarshal)
	//log.Println(str)
	//gormSyntax.TestGORM()
	//tstr := "x.x.x.x|y.y.y.y"
	//ipList := strings.Split(tstr, "|")
	//log.Println(ipList)
	//for _, ip := range ipList {
	//	log.Println(ip)
	//}
	//mysqlSyntax.TestMysql()
	//gormSyntax.TestGORM()
	gormSyntax.TestGormInterface()
	//log.Println("saassaa")
	//ginSyntax.TestGin()

	//defer fmt.Println("in main")
	//defer func() {
	//	defer func() {
	//		panic("panic again and again")
	//	}()
	//	panic("panic again")
	//}()
	//panic("panic once")

	//for i := 0; i < 5; i++ {
	//	defer fmt.Println(i)
	//}
}
