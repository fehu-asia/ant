package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// MySQLSource : 要连接的数据库源；
	// 其中test:test 是用户名密码；
	// 127.0.0.1:3306 是ip及端口；
	// fileserver 是数据库名;
	// charset=utf8 指定了数据以utf8字符编码进行传输
	db, _ = sql.Open("mysql", "root:123qwe!@#@tcp(10.10.10.33:3306)/test?charset=utf8&parseTime=true&loc=Asia%2FChongqing")
	db.SetMaxOpenConns(100)
	err := db.Ping()
	if err != nil {
		fmt.Println("Connect to mysql  error : ", err)
		os.Exit(1)
	}
}

/**
获取链接对象
*/
func GetDBConn() *sql.DB {
	return db
}

/**
获取链接对象
*/
func ExecSql(sql string, params ...interface{}) bool {
	stmt, e := GetDBConn().Prepare(sql)
	if e != nil {
		fmt.Println("Sql Prepare error : ", e)
		return false
	}
	defer stmt.Close()
	result, e := stmt.Exec(params...)
	if e != nil {
		fmt.Println(" Exec sql error : ", e)
		return false
	}
	rf, e := result.RowsAffected()
	if nil != e {
		fmt.Println("err:", e, "rf:", rf)
		return false
	}
	if rf == 0 {
		return false
	}
	return true
}
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
