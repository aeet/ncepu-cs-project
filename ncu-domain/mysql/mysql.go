package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"ncu-auth/config"
	"ncu-domain/domain"
)

func Open() (client *domain.Client, err error) {
	option := config.Value.MySQLOption
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", option.Username, option.Password, option.Host, option.Port, option.Schema)
	return domain.Open("mysql", dsn)
}

func Setup() error {
	if err := setupTable(); err != nil {
		return err
	}
	if err := SetupData(); err != nil {
		return err
	}
	return nil
}
