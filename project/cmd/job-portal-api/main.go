package main

import (
	"fmt"
	"net/http"
	"os"
	"project/config"
	"project/internal/auth"
	"project/internal/database"
	"project/internal/handlers"
	redisconn "project/internal/redisConn"
	"project/internal/repository"
	"project/internal/services"

	"github.com/golang-jwt/jwt/v5"

	"github.com/rs/zerolog/log"
)

func main() {
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}

}
func startApp() error {
	cfg := config.GetConfig()
	log.Info().Msg("started main")
	privatePEM, err := os.ReadFile(`private.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file private.pem %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}

	publicPEM, err := os.ReadFile(`pubkey.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file pubkey.pem %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("cannot create auth instance %w", err)
	}

	db, err := database.Connection()
	if err != nil {
		return err
	}
	repou, err := repository.NewUserRepo(db)
	if err != nil {
		return err
	}
	repoc, err := repository.NewCompanyRepo(db)
	if err != nil {
		return err
	}

	con := redisconn.ReddisConc()
	re := redisconn.NewRDBLayer(con)

	//se, err := services.NewService(repo, repo,re)

	sec, err := services.NewCompanyServiceImp(repoc, re)
	if err != nil {
		return err
	}

	seu, err := services.NewUserServiceImp(repou)
	if err != nil {
		return err
	}

	api := http.Server{ //server config and settimngs
		Addr:    fmt.Sprintf("%s:%s", cfg.AppConfig.Host, cfg.AppConfig.Port),
		Handler: handlers.Api(a, seu, sec),
	}
	api.ListenAndServe()

	return nil

}
