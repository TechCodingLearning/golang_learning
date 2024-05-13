/*
@author: louris
@since: 2022/9/24
@desc: //TODO
*/

package mysqlSyntax

import (
	"database/sql"
	"strings"

	// mysql driver
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "xxxx"
	passWord = "xxxx"
	ip       = "xxxx"
	port     = "3306"
	dbName   = "go_db"
)

// Connect mysq数据库连接
func Connect() *sql.DB {
	// dataSourceName: "用户名:密码@tcp(IP:port)/数据库名?...."
	path := strings.Join([]string{
		userName, ":", passWord, "@tcp(", ip, ":", port,
		")/", dbName, "?charset=utf8&parseTime=True",
	}, "")
	log.Println(path)
	//db, err := sql.Open("mysql", "root:ROot1234567&@tcp(9.134.5.247:3306)/go_db?charset=utf8&parseTime=True")
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("database connect failed!", err)
	}

	// 设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	// 设置数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	// 验证连接
	if err := db.Ping(); err != nil {
		log.Fatal("database ping failed!", err)
	}
	return db
}

/*

CREATE TABLE user_tb(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(20),
    age INTEGER
);

*/

// User 数据库表user_tb的实体数据结构
type User struct {
	ID       int32
	UserName string
	Age      int32
}

// QueryOneRecord 查询单条数据
func QueryOneRecord(db *sql.DB, id int32) {
	user := User{}
	rows := db.QueryRow("select * from user_tb where id = ?", id)
	// rows.Scan 用于把读取的数据赋值到User对象的属性上，注意字段顺序，按照表定义顺序
	err := rows.Scan(&user.ID, &user.UserName, &user.Age)
	if err != nil {
		log.Fatal(err)
	}

	// 延迟到函数结束关闭连接
	// defer db.Close()
	log.Println("单条数据结果：", user)
}

// QueryMoreRecords @Description: 查询多条记录
func QueryMoreRecords(db *sql.DB) {
	rows, err := db.Query("select * from user_tb where id > ?", 0)
	if err != nil {
		log.Fatal(err)
	}
	users := []User{}
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.UserName, &user.Age); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	log.Println("查询多条数据：\n", users)
	// defer db.Close()
}

// UpdateRecords 更删改查
func UpdateRecords(db *sql.DB, sql string, args ...interface{}) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}

	rows, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	log.Println("受影响行数：", rows)
	log.Println("自增id：", id)
}

// UpdateDataWithTransaction 事务
func UpdateDataWithTransaction(db *sql.DB) {
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	UpdateRecords(db, "update user_tb set age=40 where id = ?", 1)
	UpdateRecords(db, "update user_tb set age=70 where id = ?", 2)
	// 提交事务
	if e := tx.Commit(); e != nil {
		// 报错则回滚
		err := tx.Rollback()
		if err != nil {
			log.Fatal("事务回滚失败!")
		}
	}
	// defer db.Close()
}

func TestMysql() {
	db := Connect()
	// QueryOneRecord(db, 1)
	QueryMoreRecords(db)
	UpdateDataWithTransaction(db)
	QueryMoreRecords(db)
	defer db.Close()
}
