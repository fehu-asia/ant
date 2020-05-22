package db

import (
	"ant/db/mysql"
	"database/sql"
	"fmt"
)

var insertMetaDataSql = "INSERT INTO `storage`.`file`( `sha256`, `name`, `size`, `addr`, `status`) " +
	"VALUES (?, ?, ?, ?, 1);"

/**
保存文件元数据
*/
func SaveFileMetaData(fileHash string, fileName string, fileSize int64, fileAddr string) bool {
	return mysql.ExecSql(insertMetaDataSql, fileHash, fileName, fileSize, fileAddr)
}

type TableFile struct {
	Sha256 string
	Name   sql.NullString
	Size   sql.NullInt64
	Addr   sql.NullString
}

var getFileMetaSql = "SELECT `sha256`, `name`, `size`, `addr` from `storage`.`file` f WHERE f.sha256 = ? AND f.status = 1 limit 1 "

func GetFileMeta(fileHash string) (*TableFile, error) {
	stmt, e := mysql.GetDBConn().Prepare(getFileMetaSql)
	if e != nil {
		fmt.Println("Sql Prepare error : ", e)
		return nil, e
	}

	defer stmt.Close()

	file := &TableFile{}

	e = stmt.QueryRow(fileHash).Scan(&file.Sha256, &file.Name, &file.Size, &file.Addr)
	if e != nil {
		fmt.Println("stmt.QueryRow().Scan error : ", e)
		return file, e
	}

	return file, nil
}
