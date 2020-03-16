package main

import (
	"accounting_software/server/handlers"
	"accounting_software/server/tcsctx"
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

	router.Get("/categories", handlers.CategoriesGet)
	router.Post("/categories/add", handlers.CategorySet)
	router.Post("/categories/remove", handlers.CategoryRemove)
	router.Post("/categories/update", handlers.CategoryUpdate)

	router.Get("/cashbox", handlers.CashBoxGet)
	router.Post("/cashbox/add", handlers.CashBoxSet)
	router.Post("/cashbox/remove", handlers.CashBoxRemove)

	router.Get("/warehouse", handlers.WarehouseGet)
	router.Post("/warehouse/add", handlers.WarehouseSet)
	router.Post("/warehouse/remove", handlers.WarehouseRemove)

	router.Get("/credit", handlers.CreditGet)
	router.Post("/credit/add", handlers.CreditSet)
	router.Post("/credit/report", handlers.ReportMonthGet)
	router.Post("/credit/remove", handlers.CreditRemove)
	router.Post("/credit/update", handlers.CreditUpdate)

	router.Get("/debit", handlers.DebitGet)
	router.Get("/debit/last_id", handlers.DebitGetLastId)
	router.Post("/debit/add", handlers.DebitSet)
	router.Post("/debit/report", handlers.ReportDebitMonthGet)

	router.Get("/debits_products", handlers.DebitsProductsGet)
	router.Post("/debits_products/add", handlers.DebitsProductsSet)

	router.Get("/supply", handlers.SupplyGet)
	router.Get("/supply/id", handlers.SupplyGetId)
	router.Post("/supply/add", handlers.SupplySet)
	router.Post("/supply/remove", handlers.SupplyRemove)
	router.Post("/supply/update", handlers.SupplyUpdate)
	router.Post("/supply/decommissioned", handlers.Decommissioned)
	router.Post("/supply/month_decommissioned", handlers.DecommissionedMonthGet)
	router.Post("/supply/update_amount", handlers.SupplyUpdateAmount)

	router.Get("/utility_get", handlers.UtilityGet)
	router.Post("/utility_set", handlers.UtilitySet)
	router.Post("/utility_remove", handlers.BillRemove)
	router.Post("/utility_bills", handlers.ReportUtilMonthGet)

	router.Get("/units", handlers.UnitsGet)

	router.Get("/utility/get", handlers.NewUtilityGet)
	router.Get("/utility/names", handlers.NewBillNamesGet)
	router.Post("/utility/remove", handlers.NewBillRemove)
	router.Post("/utility/set", handlers.NewUtilitySet)

	router.Post("/utility/month_bills", handlers.UtilBillsMonthGet)

	router.Post("/order", handlers.OrderSet)
	router.Post("/order/remove", handlers.OrderRemove)
	router.Get("/orders", handlers.OrdersGet)
	router.Post("/order/last", handlers.OrderLastGet)
	router.Post("/customer", handlers.CustomerIDGet)
	router.Post("/customer/add", handlers.CustomerSet)
	router.Post("/customer/get/chat_id", handlers.CustomerGetChatID)
	router.Get("/customer/get/last_id", handlers.CustomersGetLastId)

	//ReportMonthGet

	webserver.Start(router, config.MustString("port"))
	mainloop.Loop(stop, grace, config.DefaultConfig())

}

func stop() {
	webserver.Stop()
}

func grace() {
	webserver.Grace(configPath)
}
