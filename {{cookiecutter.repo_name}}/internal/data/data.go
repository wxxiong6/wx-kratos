package data

import (
	"context"
	"fmt"
	"{{cookiecutter.module_name}}/internal/biz"
	"time"

	"{{cookiecutter.module_name}}/pkg/gorm/gormlog"
	"{{cookiecutter.module_name}}/pkg/gorm/trace"

	"{{cookiecutter.module_name}}/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRdb, New{{cookiecutter.service_name}}Repo, NewTransaction)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
	rdb *redis.Client
}

type contextTxKey struct{}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/NewData"))

	cleanup := func() {
		l.Info("closing the data resources")
	}

	return &Data{
		db:  db,
		log: l,
		rdb: rdb,
	}, cleanup, nil
}

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "data/NewDB"))
	gormLog := gormlog.NewGormLog(log.NewHelper(logger), "mysql")

	gormConfig := &gorm.Config{
		Logger:      gormLog,
		QueryFields: true, //QueryFields 模式会根据当前 model 的所有字段名称进行 select。
	}

	db, err := gorm.Open(mysql.Open(c.Database.Source), gormConfig)

	if err != nil {
		l.Errorf("failed opening connection to mysql: %v", err)
	}
	err = db.Use(&trace.GormTrace{})
	if err != nil {
		l.Errorf("failed db.Use error: %v", err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(int(c.Database.MinIdleConns))
	sqlDb.SetMaxOpenConns(int(c.Database.MaxOpenConns))
	sqlDb.SetConnMaxLifetime(time.Hour * time.Duration(c.Database.ConMaxLeftTime))
	err = sqlDb.Ping()
	if err != nil {
		l.Errorf("failed pinging mysql: %v", err)
		panic(err)
	} else {
		l.Infow("kind", "mysql", "status", "enable")
	}
	return db
}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewRdb(c *conf.Data, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "data/NewRdb"))
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		PoolSize:     int(c.Redis.PoolSize),     // 连接池数量
		MinIdleConns: int(c.Redis.MinIdleConns), //好比最小连接数
		MaxRetries:   int(c.Redis.MaxRetries),   // 命令执行失败时，最多重试多少次，默认为0即不重试
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.NewTracingHook(
		redisotel.WithAttributes(
			attribute.String("db.type", "redis"),
			attribute.String("db.ip", c.Redis.Addr),
			attribute.String("db.instance", fmt.Sprintf("%s/%d", c.Redis.Addr, c.Redis.Db)),
		),
	))
	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		l.Errorf("failed to ping redis: %s", err)
	} else {
		l.Infof("connected to redis: %s", result)
	}
	return rdb
}
