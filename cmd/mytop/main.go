package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-bridget/mytop-go/db"
	"github.com/go-bridget/mytop-go/terminal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := db.NewOptions()
	flag.Parse()

	ctx := context.Background()

	monitor := db.NewMySQLMonitor(config)

	if err := monitor.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	app := terminal.NewApp(monitor)

	app.Init()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				app.Stop()
				fmt.Println("Recovered: ", r)
			}
		}()
		for {
			pl, err := monitor.ShowProcessList(ctx)
			if err != nil {
				app.Stop()
				log.Fatalf("Error retrieving process list: %v", err)
			}
			monitor.ProcessList = pl
			app.SetTableData(pl)
			app.Draw()
			select {
			case <-ctx.Done():
				app.Stop()
				log.Fatalf("context cancelled")
			case <-time.After(time.Second * time.Duration(config.Delay)):
			case <-app.Refresh:
				// do nothing and cause loop to restart
			}
		}
	}()
	// Run blocks until app.Stop is called
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
