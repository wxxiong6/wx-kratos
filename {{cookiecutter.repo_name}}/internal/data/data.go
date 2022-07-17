package data

import (
	"database/sql"
	"{{cookiecutter.module_name}}/internal/conf"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/wxxiong6/sqlhooks"
	"github.com/wxxiong6/sqlhooks/hooks/loghooks"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, New{{cookiecutter.service_name}}Repo)

// Data .
type Data struct {
	db *sqlx.DB
	rdb *redis.Client
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "{{cookiecutter.module_name}}-service/data"))

	sql.Register("mysqlLog", sqlhooks.Wrap(&mysql.MySQLDriver{}, loghooks.New()))
	db, err := sqlx.Open("mysqlLog", c.Database.Source)
	//db, err := sqlx.Open("mysql", c.Database.Source)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	d := &Data{
		db:  db,
		rdb: rdb,
		log: log,
	}
	cleanup := func() {
		log.Info("closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error("close redis error", err)
		}
	}

	return d, cleanup, nil
}
