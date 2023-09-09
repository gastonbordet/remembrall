package database

import (
	"fmt"

	"github.com/gastonbordet/remembrall/cmd/util"
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

func GenerateMysqlURIConnection(
	config *util.Config,
) string {
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DB_USER,
		config.DB_PWD,
		config.DB_ADDRESS,
		config.DB_PORT,
		config.DB_NAME,
	)

	return uri
}
