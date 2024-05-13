package gormSyntax

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"reflect"
)

// CommDaoAPI 通用db操作类
type CommDaoAPI struct {
	db *gorm.DB
}

// NewCommDaoAPI .
func NewCommDaoAPI() CommDaoAPI {
	return CommDaoAPI{
		db: Connect(),
	}
}

func (impl *CommDaoAPI) QueryRecord(tableName string, model interface{}) ([]interface{}, error) {
	modelRefType := reflect.TypeOf(model)
	resModels := reflect.New(reflect.SliceOf(modelRefType)).Interface()
	result := impl.db.WithContext(context.Background()).
		Table(tableName).
		Debug().
		Where(model).
		Find(resModels)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("db failed! %v", result.Error))
	}

	// 不对查询行数进行判断，由业务调用后自行判断是否符合预期
	// if result.RowsAffected != 1 {
	// 	return nil, ferror.RowsAffectedError
	// }

	resModulesValue := reflect.ValueOf(resModels).Elem()
	numRecords := resModulesValue.Len()
	records := make([]interface{}, numRecords)
	for i := 0; i < numRecords; i++ {
		records[i] = resModulesValue.Index(i).Interface()
	}
	return records, nil
}

// DeleteRecord 删除记录
func (impl *CommDaoAPI) DeleteRecord(tableName string, model interface{}) ([]interface{}, error) {
	modelRefType := reflect.TypeOf(model)
	resModels := reflect.New(reflect.SliceOf(modelRefType)).Interface()
	delModels := reflect.New(reflect.SliceOf(modelRefType)).Interface()
	db := impl.db.WithContext(context.Background()).
		Table(tableName).
		Debug()

	if result := db.Where(model).Find(resModels).Error; result != nil {
		return nil, errors.New(fmt.Sprintf("db failed! %v", result.Error()))
	}

	if result := db.Where(model).Delete(delModels).Error; result != nil {
		return nil, errors.New(fmt.Sprintf("db failed! %v", result.Error()))
	}

	resModulesValue := reflect.ValueOf(resModels).Elem()
	numRecords := resModulesValue.Len()
	records := make([]interface{}, numRecords)
	for i := 0; i < numRecords; i++ {
		records[i] = resModulesValue.Index(i).Interface()
	}
	return records, nil
}

// UpdateRecord 更新记录
func (impl *CommDaoAPI) UpdateRecord(tableName string, qryModel interface{}, uptModel interface{}) error {
	result := impl.db.WithContext(context.Background()).
		Table(tableName).
		Debug().
		Where(qryModel).
		Updates(uptModel)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("db failed! %v", result.Error))
	}
	return nil
}

// AddRecord 插入记录
func (impl *CommDaoAPI) AddRecord(tableName string, models interface{}) error {
	result := impl.db.WithContext(context.Background()).
		Table(tableName).
		Debug().
		Create(models)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("db failed! %v", result.Error))
	}
	return nil
}

//func testOther() {
//
//	qryUser := User{UserName: "lourisxu"}
//	modelRefType := reflect.Indirect(reflect.ValueOf(qryUser)).Type()
//	resModels := reflect.MakeSlice(reflect.SliceOf(modelRefType), 10, 10)
//	log.Printf("modelType: %v", modelRefType)
//	log.Printf("resModels: %v", resModels)
//
//	//var exampleUsers []User
//	exampleUsers := make([]User, 10)
//	log.Printf("exampleUsers: %v", exampleUsers)
//	type1 := reflect.Indirect(reflect.ValueOf(resModels)).Type()
//	type2 := reflect.Indirect(reflect.ValueOf(exampleUsers)).Type()
//	log.Printf("type1: %v", type1)
//	log.Printf("type2: %v", type2)
//	if type1 == type2 {
//		log.Printf("equal !!!")
//	}
//
//	urs := resModels.Index(0).Interface().(User)
//	type3 := reflect.Indirect(reflect.ValueOf(urs)).Type()
//	log.Printf("resModel: %v", type3)
//}

func TestGormInterface() {
	qryUser := User{UserName: "kity", Age: 27}
	//uptUser := User{Age: 27}
	//addUsers := []User{qryUser, qryUser, qryUser}
	log.Println("aaaa")
	daoApi := NewCommDaoAPI()
	tName := fmt.Sprintf("%s.%s", dbName, tableName)
	//retUsers, err := daoApi.QueryRecord(tName, qryUser)
	retUsers, err := daoApi.DeleteRecord(tName, qryUser)
	//err := daoApi.UpdateRecord(tName, qryUser, uptUser)
	//err := daoApi.AddRecord(tName, addUsers)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("retUsers: %v", retUsers)
	log.Println("resUsers type: %v", reflect.TypeOf(retUsers))

}
