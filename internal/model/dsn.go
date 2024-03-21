package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB2 sqlx.SqlConn
var DB1 *gorm.DB

func Gorm(dsn string) (err error) {
	DB1, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func SQLX(dsn string) {
	DB2 = sqlx.NewMysql(dsn)
}
