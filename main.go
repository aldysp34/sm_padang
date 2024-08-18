package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aldysp34/sm_padang/config"
	"github.com/aldysp34/sm_padang/database"
	"github.com/aldysp34/sm_padang/handler"
	"github.com/aldysp34/sm_padang/logger"
	"github.com/aldysp34/sm_padang/repository"
	"github.com/aldysp34/sm_padang/server"
	"github.com/aldysp34/sm_padang/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logger.Log.Errorf("unable to load env")
	}

	log := logger.NewLogger()

	logger.SetLogger(log)

	dbConfig := database.InitDatabaseConfig()

	ctx := context.TODO()

	barangInRepo := repository.NewBarangInRepository(dbConfig)
	barangOutRepo := repository.NewBarangOutRepository(dbConfig)
	barangRepo := repository.NewBarangRepository(dbConfig)
	brandRepo := repository.NewBrandRepository(dbConfig)
	requestRepo := repository.NewRequestRepository(dbConfig)
	roleRepo := repository.NewRoleRepository(dbConfig)
	satuanRepo := repository.NewSatuanRepository(dbConfig)
	supplierRepo := repository.NewSupplierRepository(dbConfig)
	userRepo := repository.NewUserRepository(dbConfig)

	adminUsecase := usecase.NewAdminUsecase(barangInRepo, barangOutRepo, barangRepo, brandRepo, requestRepo, roleRepo, satuanRepo, supplierRepo, userRepo)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	userUsecase := usecase.NewUserUsecase(requestRepo, barangRepo)

	adminHandler := handler.NewAdminHandler(adminUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	opts := server.RouterOpts{
		AuthHandler:  authHandler,
		AdminHandler: adminHandler,
		UserHandler:  userHandler,
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
	})

	<-wait
}
