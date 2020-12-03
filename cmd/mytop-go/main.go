/*
Copyright © 2020 Evan do Carmo carmo.evan@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/carmo-evan/mytop-go/db"
	"github.com/carmo-evan/mytop-go/terminal"

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

	for {
		pl, err := monitor.ShowProcessList(ctx)
		if err != nil {
			log.Fatalf("Error retrieving process list: %v", err)
		}

		if err := terminal.Clear(); err != nil {
			log.Fatalf("Error clearing terminal: %v", err)
		}
		terminal.Draw(pl)

		select {
		case <-ctx.Done():
			log.Fatalf("context cancelled")
		case <-time.After(time.Second * time.Duration(config.Delay)):
		}
	}
}
