package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/acehinnnqru/gin-server-example/config"
	"github.com/acehinnnqru/gin-server-example/handlers"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"

	ginI18n "github.com/gin-contrib/i18n"
)

func init() {
	config.Init()
	initLogger()
}

func engine() *gin.Engine {
	r := gin.Default()

	// i18n support
	r.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		DefaultLanguage: language.Make(config.AppConfig.DefaultLanguage),
		AcceptLanguage: []language.Tag{
			language.English, language.Chinese,
		},
		UnmarshalFunc:    toml.Unmarshal,
		RootPath:         "./_locales",
		FormatBundleFile: "toml",
	}), ginI18n.WithGetLngHandle(func(context *gin.Context, defaultLng string) string {
		lang := context.GetHeader("Accept-Language")
		if lang != "" {
			return lang
		}
		return defaultLng
	})))

	return r
}

func main() {
	r := engine()

	handlers.Register(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
	}

	log.Fatal("Server exiting")
}
