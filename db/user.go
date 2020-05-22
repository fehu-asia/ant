package db

import (
	"ant/db/mysql"
	"fmt"
)

var (
	inserUserSql            = "INSERT INTO `storage`.`user`(`username`, `password`) VALUES (?, ?);"
	selectUserByUserNameSql = "SELECT * FROM `storage`.`user` u WHERE u.username = ? LIMIT 1"
	tokenSql                = "replace into token (`username`,`token`) values(?,?)"
)

/**
用户注册
*/
func InserUser(username string, password string) bool {
	return mysql.ExecSql(inserUserSql, username, password)
}

/**
获取用户
*/
func SelectUserByUserNameAndPassword(username string, password string) bool {
	stmt, e := mysql.GetDBConn().Prepare(selectUserByUserNameSql)
	if e != nil {
		fmt.Println("Sql Prepare error : ", e)
		return false
	}

	defer stmt.Close()
	rows, e := stmt.Query(username)
	if e != nil {
		fmt.Println("Sql Query : ", e)
		return false
	}
	if rows == nil {
		fmt.Println("user not found ! ")
		return false
	}
	parseRows := mysql.ParseRows(rows)
	if len(parseRows) > 0 && parseRows[0]["password"] == password {
		return true
	}
	return false
}

func UpdateToken(username string, token string) bool {
	return mysql.ExecSql(tokenSql, username, token)
}
