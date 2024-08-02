package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logger.Log.Errorf("unable to load env")
	}

	log := logger.NewLogger()

	logger.SetLogger(log)

	dbConfig, err := database.InitDatabaseConfig()
	if err != nil {
		logger.Log.Errorf(err.Error())
		panic(err)
	}

	ctx := context.TODO()

	if err := dbConfig.ConnectMongoDB(ctx); err != nil {
		logger.Log.Errorf(err.Error())
		panic(err)
	}

	dbDatabase := dbConfig.MongoClient.Database("maison")
	smtpConfig := config.GetSMTPConfig()
	sr := repository.NewSMTP(smtpConfig)

	resetRepo := repository.NewResetPasswordRepository(dbDatabase)

	authRepo := repository.NewAuthRepository(dbDatabase)
	authUsecase := usecase.NewAuthUsecase(authRepo, sr, resetRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	productRepo := repository.NewProductRepository(dbDatabase)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	articleRepo := repository.NewArticleRepository(dbDatabase)
	articleUsecase := usecase.NewArticleUsecase(articleRepo)
	articleHandler := handler.NewArticleHandler(articleUsecase)

	opts := server.RouterOpts{
		AuthHandler:    authHandler,
		ProductHandler: productHandler,
		ArticleHandler: articleHandler,
	}
	r := server.NewRouter(opts)
	appPort := os.Getenv("BASE_PORT")
	srv := &http.Server{
		Addr:    appPort,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Errorf(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	logger.Log.Info(fmt.Sprintf("running on port: %s\n", appPort))

	wait := config.GracefulShutdown(ctx, 2*time.Second, map[string]config.Operation{
		"http-server": func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
		"mongo-db": func(ctx context.Context) error {
			return dbConfig.MongoClient.Disconnect(ctx)
		},
	})

	<-wait
}
