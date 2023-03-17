package main

import (
	"context"
	"database/sql"

	usecases "article/src/app/usecases"

	"article/src/infra/config"
	postgres "article/src/infra/persistence/postgres"
	pgArticle "article/src/infra/persistence/postgres/article"
	redis "article/src/infra/persistence/redis"
	articleRedis "article/src/infra/persistence/redis/article"

	"article/src/interface/rest"

	ms_log "article/src/infra/log"

	articleUC "article/src/app/usecases/article"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

//reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)
	redisClient, err := redis.NewRedisClient(conf.Redis, logger)

	artcleRdb := articleRedis.NewArticleRedis(redisClient)
	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	articleRepository := pgArticle.NewArticleRepository(postgresdb.Conn, artcleRdb)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{

			ArticleUC: articleUC.NewArticleUseCase(articleRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
