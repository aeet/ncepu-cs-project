package mysql

import (
	"context"
	"fmt"

	"github.com/devcui/ncepu-cs-project/config"
	"github.com/devcui/ncepu-cs-project/domain"
)

func SetupData() error {
	client, err := Open()
	if err != nil {
		return err
	}
	ctx := context.Background()
	tx, err := client.Tx(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	setupAuthorization(tx, ctx)
	setupUser(tx, ctx)
	setupRole(tx, ctx)
	setupResource(tx, ctx)
	setupAuthorizationResource(tx, ctx)
	setupUserRole(tx, ctx)
	setupUserResource(tx, ctx)
	setupRoleResource(tx, ctx)
	defer tx.Client().Close()
	defer client.Close()
	return tx.Commit()
}

// entity data
func setupAuthorization(tx *domain.Tx, ctx context.Context) {
	auth := &domain.Authorization{
		ClientID:     "auth",
		ClientName:   "认证中心",
		ClientSecret: "auth",
		Domain:       fmt.Sprintf("http://127.0.0.1:8080%s", config.Value.Server.Path),
		RedirectURL:  fmt.Sprintf("http://127.0.0.1:8080%s%s", config.Value.Server.Path, "/callback/show"),
		GrantType:    []string{"refresh_token", "password", "authorization_code", "client_credentials", "implicit"},
		Scope:        []string{"read", "write", "update", "all"},
	}
	tx.Authorization.Delete().Exec(ctx)
	tx.Authorization.Create().SetClientID(auth.ClientID).SetClientName(auth.ClientName).SetClientSecret(auth.ClientSecret).SetDomain(auth.Domain).SetRedirectURL(auth.RedirectURL).SetGrantType(auth.GrantType).SetScope(auth.Scope).Save(ctx)
}

func setupUser(tx *domain.Tx, ctx context.Context) {
	user := &domain.User{
		Account:  "default",
		Passwd:   "default",
		Avatar:   nil,
		Email:    "devcui@outlook.com",
		Username: "默认用户",
	}
	tx.User.Delete().Exec(ctx)
	tx.User.Create().SetAccount(user.Account).SetPasswd(user.Passwd).SetAvatar([]byte{}).SetEmail(user.Email).SetUsername(user.Username).Save(ctx)
}

func setupRole(tx *domain.Tx, ctx context.Context) {
	role := &domain.Role{
		RoleName:  "管理员",
		RoleValue: "administrator",
	}
	tx.Role.Delete().Exec(ctx)
	tx.Role.Create().SetRoleName(role.RoleName).SetRoleValue(role.RoleValue).Save(ctx)
}

func setupResource(tx *domain.Tx, ctx context.Context) {
	resource := &domain.Resource{
		ResourceName:  "认证中心资源",
		ResourceValue: "auth",
	}
	tx.Resource.Delete().Exec(ctx)
	tx.Resource.Create().SetResourceName(resource.ResourceName).SetResourceValue(resource.ResourceValue).Save(ctx)
}

// relation data
func setupAuthorizationResource(tx *domain.Tx, ctx context.Context) error {
	authorization, err := tx.Authorization.Query().First(ctx)
	if err != nil {
		return err
	}
	resources, err := tx.Resource.Query().All(ctx)
	if err != nil {
		return err
	}
	for _, resource := range resources {
		_, e := resource.Update().AddAuthorizationIDs(authorization.ID).Save(ctx)
		if e != nil {
			return e
		}
	}
	return nil
}

func setupUserRole(tx *domain.Tx, ctx context.Context) error {
	user, err := tx.User.Query().First(ctx)
	if err != nil {
		return err
	}
	roles, err := tx.Role.Query().All(ctx)
	if err != nil {
		return err
	}
	for _, role := range roles {
		_, e := role.Update().AddUserIDs(user.ID).Save(ctx)
		if e != nil {
			return e
		}
	}
	return nil
}

func setupRoleResource(tx *domain.Tx, ctx context.Context) error {
	role, err := tx.Role.Query().First(ctx)
	if err != nil {
		return err
	}
	resources, err := tx.Resource.Query().All(ctx)
	if err != nil {
		return err
	}
	for _, resource := range resources {
		_, e := resource.Update().AddRoleIDs(role.ID).Save(ctx)
		if e != nil {
			return e
		}
	}
	return nil
}

func setupUserResource(tx *domain.Tx, ctx context.Context) error {
	user, err := tx.User.Query().First(ctx)
	if err != nil {
		return err
	}
	resources, err := tx.Resource.Query().All(ctx)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for _, resource := range resources {
		_, e := resource.Update().AddUserIDs(user.ID).Save(ctx)
		if e != nil {
			return e
		}
	}
	return nil
}
