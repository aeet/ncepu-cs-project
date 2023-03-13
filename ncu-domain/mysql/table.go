package mysql

import (
	"context"

	"ncu-domain/domain/migrate"
)

func setupTable() error {
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
