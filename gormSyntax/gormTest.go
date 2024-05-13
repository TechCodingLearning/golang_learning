/*
Package gormSyntax test
@author: louris
@since: 2022/9/25
@desc: //TODO
*/
package gormSyntax

import (
	"errors"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	userName  = "xxxx"
	passWord  = "xxxx"
	ip        = "xxxx"
	port      = "3306"
	dbName    = "go_db"
	tableName = "user_tb"
)

// User 数据库表user_tb的实体数据结构
type User struct {
	ID       int32  `gorm:"column:id;primaryKey;<-:create"`
	UserName string `gorm:"column:username;"`
	Age      int32  `gorm:"column:age;default:NULL"`
}

// Connect 数据库链接
// @return *gorm.DB
func Connect() *gorm.DB {
	// dataSourceName: "用户名:密码@tcp(IP:port)/数据库名?...."
	path := strings.Join([]string{
		userName, ":", passWord, "@tcp(", ip, ":", port,
		")/", dbName, "?charset=utf8&parseTime=True",
	}, "")
	log.Println(path)

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database!", err)
	}
	return db
}

// InsertOneRecord 插入记录
// @param db
func InsertOneRecord(db *gorm.DB, user User) {

	result := db.Table(tableName).Create(&user)
	//result := db.Raw("INSERT INTO ? (username, age) VALUES (?, ?)",
	//	clause.Table{Name: tableName}, "lourisxu", 27).Scan(&user)

	if result.Error != nil {
		log.Fatal("return the error：", result.Error)
	}
	log.Println("return the primary key of inserted data：", user.ID)
	log.Println("return the rows of affected：", result.RowsAffected)
}

// InsertMoreRecords 批量插入数据
// @param db
// @param users
func InsertMoreRecords(db *gorm.DB, users []User) {
	result := db.Table(tableName).CreateInBatches(users, len(users))

	if result.Error != nil {
		log.Fatal("return the error：", result.Error)
	}
	log.Println("return the rows of affected", result.RowsAffected)
}

// InsertMoreRecordsByMap 通过map批量插入数据
// association不会被调用，且主键也不会自动填充
// @param db
// @param userMap
func InsertMoreRecordsByMap(db *gorm.DB, userMap []map[string]interface{}) {
	result := db.Model(&User{}).Create(userMap)
	if result.Error != nil {
		log.Fatal("return the error：", result.Error)
	}
	log.Println("return the rows of affected", result.RowsAffected)
}

// BeforeCreate 创建记录前钩子示例
// @receiver u
// @param db
// @return err
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("before create....")
	if u.UserName == "louris" {
		log.Println("good luck!")
	}
	return nil
}

// QueryOneRecord 查询一条记录
// @param db
// @return User
func QueryOneRecord(db *gorm.DB) User {
	user := User{}
	// SELECT * FROM `user_tb` ORDER BY `user_tb`.`id` LIMIT 1
	//result := db.First(&user)
	// SELECT * FROM `user_tb` WHERE `username` = "tom" LIMIT 1
	result := db.Table(tableName).Where("username = ?", "louris").Scan(&user)
	//result := db.Raw("SELECT * FROM ? WHERE username = ?", clause.Table{Name: tableName}, "louris").Scan(&user)

	// 行锁
	//result := db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("username = ?", "tom").Scan(&user)

	if result.Error != nil {
		log.Fatal("return the error：", result.Error)
	}
	log.Println("return the rows of affected", result.RowsAffected)
	log.Println("result: ", user)
	return user
}

// QueryMoreRecords 查询所有数据
// @param db
func QueryMoreRecords(db *gorm.DB) {
	var users []User
	qryUser := User{
		UserName: "yy",
		//Age: 22,
	}
	// find查询所有数据
	// result := db.Table("user_tb").Select("age").Limit(2).Find(&users)
	// result := db.Table("user_tb").Limit(10).Find(&users)
	//result := db.Raw("SELECT * FROM ? LIMIT 10", clause.Table{Name: tableName}).Scan(&users)
	result := db.Table("user_tb").Where(qryUser).Scan(&users)
	if result.Error != nil {
		log.Fatal("return the error：", result.Error)
	}
	log.Println("return the rows of affected", result.RowsAffected)
	for _, user := range users {
		log.Println(user)
	}
}

// UpdateRecord 更新记录
// @param db
func UpdateRecord(db *gorm.DB) {
	user := User{UserName: "tom"}
	result := db.Table(tableName).Where("username = ?", "tom").Update("username", "Tom")
	if result.Error != nil {
		log.Fatal("return the error：", result.Error)
	}
	log.Println("return the rows of affected", result.RowsAffected)
	log.Println("result: ", user)
}

// DeleteRecord 删除记录
// @param db
func DeleteRecord(db *gorm.DB) {
	var users []User
	result := db.Table(tableName).Clauses(clause.Returning{}).Where("username = ?", "louris").Delete(&users)
	if result.Error != nil {
		log.Fatal("return the error：", result.Error)
	}
	log.Println("return the rows of delete affected", result.RowsAffected)
	log.Println("result: ")
	for _, user := range users { // 并没有返回任何删除的数据???
		log.Println(user)
	}
}

// UpdateRecordsWithTransaction 事务更新记录
// @param db
func UpdateRecordsWithTransaction(db *gorm.DB, user User) {
	err := db.Transaction(func(tx *gorm.DB) error {
		//result := db.Table(tableName).Where("username = ?", "louris").Scan(&user)
		result := db.Raw("SELECT * FROM ? WHERE username = ? LIMIT 2", clause.Table{Name: tableName}, user.UserName).Scan(&user)
		if result.Error != nil {
			return result.Error
		}
		log.Println("return the rows of delete affected", result.RowsAffected)
		if result.RowsAffected > 1 {
			return errors.New("affected row more than 1")
		}
		res := db.Table(tableName).Create(&user)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("insert sql affected row is not equal 1")
		}
		return nil
	})
	if err != nil {
		log.Fatal("transaction failed!", err)
	}
}

func TestGORM() {
	db := Connect()
	user := User{UserName: "yy"}
	InsertOneRecord(db, user)
	//InsertMoreRecords(db, []User{user, user})
	QueryMoreRecords(db)
	//DeleteRecord(db)
	//QueryMoreRecords(db)
	//UpdateRecordsWithTransaction(db, user)
	//QueryMoreRecords(db)
}
