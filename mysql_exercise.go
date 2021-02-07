package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
)

type user struct {
	id   int
	age  int
	name string
}

var db *sql.DB

func InnDB() (err error) {
	db, err = sql.Open("mysql", "root:9090@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	db.SetMaxIdleConns(100) // 设置数据库最大连接数
	db.SetMaxIdleConns(10)  // 设置数据库做大闲置连接数
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return err
	}
	fmt.Println("connect success")
	//defer db.Close()
	return nil

}

func insertDB() {
	sqlStr := "insert into user(name, age) values (?,?)"
	for i := 0; i < 100; i++ {
		_, err := db.Exec(sqlStr, "王五"+strconv.Itoa(i), rand.Intn(100-10)+10)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func selectDB() {
	sqlStr := "select id ,name, age from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.id, u.age, u.name)

	sqlStr = "select id ,name, age from user where id >?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.id, u.name, u.age)
	}
}

func deleteRowDemo() {
	sqlStr := "delete from user where id > ?"
	ret, err := db.Exec(sqlStr, 95)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}

func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 88, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响行数
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
func main() {
	err := InnDB()
	if err != nil {
		fmt.Println("init db failed, err: %v \n", err)
	}
	// insertDB()
	// selectDB()
	// deleteRowDemo()
	updateRowDemo()
}
