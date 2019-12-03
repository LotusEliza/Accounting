package main

import (
	"accounting_software/handlers"
	"accounting_software/tcsctx"
	"flag"
	"fmt"
	"github.com/finalist736/gokit/config"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/mainloop"
	"github.com/finalist736/gokit/webserver"
	"github.com/gocraft/web"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"math/rand"
	"os"
	"time"
)

var configPath = flag.String("config", "./config.ini", "config file path")

func main() {
	rand.Seed(time.Now().UnixNano())
	var err = config.Init(config.NewFileProvider(configPath))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't load config")
		os.Exit(1)
	}

	logger.ReloadLogs(config.DefaultConfig())

	logger.StdOut().Infof("Accounting Admin Panel")

	dbconf := database.DBConfig{}
	dbconf.Dsn = config.MustString("dsn")
	dbconf.Driver = "postgres"
	dbconf.LifeTime = time.Minute * 5
	dbconf.Stream = logger.DatabaseStream()
	dbconf.MaxIdleConns = 1
	dbconf.MaxOpenConns = 10
	err = database.Add(&dbconf)
	if err != nil {
		logger.StdErr().Errorf("db open error: %s", err)
		os.Exit(1)
	}

	err = goose.Up(database.GetDefaultSession().DB, "./migrations")
	if err != nil {
		logger.StdErr().Errorf("migrations error: %s", err)
		os.Exit(1)
	}

	router := web.New(tcsctx.Ctx{})
	router.Middleware((*tcsctx.Ctx).Cors)

	router.Get("/suppliers", handlers.SuppliersGet)
	router.Post("/suppliers/add", handlers.SupplierSet)
	router.Post("/suppliers/remove", handlers.SupplierRemove)
	router.Post("/suppliers/update", handlers.SupplierUpdate)

	router.Get("/products", handlers.ProductsGet)
	router.Post("/products/add", handlers.ProductSet)
	router.Post("/products/remove", handlers.ProductRemove)
	router.Post("/products/update", handlers.ProductUpdate)
	router.Post("/products/decommissioned", handlers.Decommissioned)

	router.Get("/categories", handlers.CategoriesGet)
	router.Post("/categories/add", handlers.CategorySet)
	router.Post("/categories/remove", handlers.CategoryRemove)
	router.Post("/categories/update", handlers.CategoryUpdate)

	router.Get("/cashbox", handlers.CashBoxGet)
	router.Post("/cashbox/add", handlers.CashBoxSet)
	router.Post("/cashbox/remove", handlers.CashBoxRemove)

	webserver.Start(router, config.MustString("port"))
	mainloop.Loop(stop, grace, config.DefaultConfig())

}

func stop() {
	webserver.Stop()
}

func grace() {
	webserver.Grace(configPath)
}
