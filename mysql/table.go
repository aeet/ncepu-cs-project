package mysql

import (
	"context"

	"github.com/devcui/ncepu-cs-project/config"
	"github.com/devcui/ncepu-cs-project/domain/migrate"
)

func setupTable() error {
	if !config.Value.MySQLOption.Init {
		return nil
	}
	client, err := Open()
	if err != nil {
		return err
	}
	ctx := context.Background()
	if err := client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		return err
	}
	return client.Close()
}
