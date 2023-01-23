package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ISqlClient interface {
	OpenConnection() (*gorm.DB, error)
	CloseConnection(*gorm.DB)
}

type SqlClient struct{}

func NewSQLClient() *SqlClient {
	return &SqlClient{}
}

func (c *SqlClient) OpenConnection(uriConn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(uriConn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GenerateMysqlURIConnection() string {
	uri := "remembrall_user:remembrall_pwd@tcp(127.0.0.1:3306)/remembrall_dev"

	return uri
}
