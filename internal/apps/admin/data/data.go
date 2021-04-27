package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/realotz/whole/internal/apps/admin/data/ent"
	"github.com/realotz/whole/internal/apps/admin/data/ent/migrate"
	"github.com/realotz/whole/internal/conf"
	"github.com/realotz/whole/pkg/token"
)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
	tk  *token.Token
}

var ProviderSet = wire.NewSet(
	NewData,
	NewCasbinAdapter,
	NewEmployeeRepo,
)

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc *conf.UserConfig) (*Data, error) {
	logHelper := log.NewHelper("data", logger)
	tk, err := token.New(uc.Cert.Key, uc.Cert.Cert)
	if err != nil {
		logHelper.Errorf("failed create jwt ", err)
		return nil, err
	}
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		logHelper.Errorf("failed opening connection to mysql: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		logHelper.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	return &Data{
		db:  client,
		rdb: rdb,
		tk:  tk,
	}, nil
}
