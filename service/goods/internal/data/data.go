package data

import (
	"goods/internal/conf"
	slog "log"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	// "github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewCategoryRepo,

	NewDB,
	NewRedis,
	NewElasticsearch,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB

	rdb      *redis.Client
	EsClient *elastic.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client, es *elastic.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:       db,
		rdb:      rdb,
		EsClient: es,
	}, cleanup, nil
}

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {
	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			// IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			// SingularTable: true, // 表名是否加 s
		},
	})

	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}

func NewElasticsearch(c *conf.Data) *elastic.Client {
	es, err := elastic.NewClient(
		elastic.SetURL(c.Elastic.Addr),
		elastic.SetSniff(false),
		elastic.SetTraceLog(
			slog.New(os.Stdout, "shop", slog.LstdFlags),
		),
	)
	if err != nil {
		panic(err)
	}

	return es
}
