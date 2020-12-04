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
	"github.com/carmo-evan/mytop-go/db"
	"github.com/carmo-evan/mytop-go/terminal"
	"github.com/gdamore/tcell/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	config := db.NewOptions()
	flag.Parse()

	ctx := context.Background()

	monitor := db.NewMySQLMonitor(config)

	if err := monitor.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	app := terminal.NewApp()
	t := terminal.NewTable()

	// TODO: abstract out into a package
	app.SetInputCapture(func (event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			app.Stop()
		}
		switch event.Rune() {
			case 's':
				// invert sort order
			case 'k':
				// kill by id
			case 'K':
				// kill all with confirmation
			case 'f':
			// filter by query
			case 'u':
				// filter by user
		}
		return event
	})

	go func() {
		for {
			pl, err := monitor.ShowProcessList(ctx)
			if err != nil {
				log.Fatalf("Error retrieving process list: %v", err)
			}

			t = terminal.SetTableData(t, pl)
			app.Draw()
			select {
			case <-ctx.Done():
				log.Fatalf("context cancelled")
			case <-time.After(time.Second * time.Duration(config.Delay)):
			}
		}
	}()

	if err := app.SetRoot(t, true).SetFocus(t).Run(); err != nil {
		log.Fatal(err)
	}
}
