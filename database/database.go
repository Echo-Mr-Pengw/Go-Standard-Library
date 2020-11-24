package main

import (
	"log"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	// 打开一个数据库驱动
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/beego_db")

	// 检查是否打开成功（不靠谱）
	checkError(err)

	// 检查是否成功打开驱动
	checkError(db.Ping())

	// 获取多条记录
	//rows, err := db.Query("select tagname from tbl_tag")
	//checkError(err)
	//
	//var tagName string
	//for rows.Next() {
	//	err := rows.Scan(&tagName)
	//	checkError(err)
	//	log.Println(tagName)
	//}
	//rows.Close()

	// 获取单条记录
	//var tagName string
	//rows := db.QueryRow("select tagname from tbl_tag where id=?", 1)
	//checkError(rows.Scan(&tagName))
	//log.Println(tagName)

	// 修改
	//result, err := db.Exec("update tbl_tag set tagname=? where id=?", "测试", 1)
	//checkError(err)
	//affectRow, _ := result.RowsAffected()
	//log.Println(affectRow)

	// 删除
	//result, err := db.Exec("delete from tbl_tag where id=?", 1)
	//checkError(err)
	//log.Println(result.RowsAffected())


	// 通过stmt，操作数据库
	// 插入数据
	//stmt, err := db.Prepare("insert into tbl_tag(tagname) values(?)")
	//checkError(err)
	//result, err := stmt.Exec("测试2")
	//checkError(err)
	//lastId, _ := result.LastInsertId()
	//log.Println(lastId)

}