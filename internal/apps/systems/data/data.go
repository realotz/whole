package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/realotz/whole/internal/apps/users/data/ent"
	"github.com/realotz/whole/internal/apps/users/data/ent/migrate"
	"github.com/realotz/whole/internal/conf"
)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

var ProviderSet = wire.NewSet(
	NewData,
	NewMemberRepo,
)

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, error) {
	log := log.NewHelper("data", logger)
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
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
	}, nil
}
