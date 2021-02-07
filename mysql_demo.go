package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// func Open(friverName, dataSourceName string) (*DB, error){

// }

// func main() {
// dsn := "root:9090@tcp(127.0.0.1:3306)/test"
// db, err = sql.Open("mysql", dsn)
// if err != nil {
// panic(err)
// }
// // fmt.Println("成功链接")
// defer db.Close()
// }

type user struct {
	id   int
	age  int
	name string
}

var db *sql.DB

// 单条查询数据db.QueryRow()
func queryRowDemo() {
	sqlstr := "select id ,name, age from user where id=?"
	var u user
	err := db.QueryRow(sqlstr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}

	fmt.Printf("id:%d name:%s, age:%d \n", u.id, u.name, u.age)
}

// 多行查询，方法 db.Query()
func queryMultiRowDemo() {
	sqlStr := "select id, age, name from user where id> ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v \n", err)
			return
		}
		fmt.Printf("id:%d name:%s, age:%d \n", u.id, u.name, u.age)
	}
}

// 插入数据, db.Exec
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err: %v \n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed , err : %v \n", err)
		return
	}
	fmt.Printf("inset sunccess, the id is %d. \n", theID)
}

func updateRowDemo() {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n ", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v \n", err)
		return
	}
	fmt.Printf("update sunccess, affected rows: %d \n ", n)
}

func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err: %v \n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v \n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d \n", n)
}
