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
	"flag"
	"github.com/carmo-evan/mytop-go/db"
	"github.com/carmo-evan/mytop-go/terminal"
	"log"
	"time"
)

var user, password, hostname, port, database string
var delay int
var noIdle bool
func init() {
	flag.StringVar(&user, "u" ,"", "Username")
	flag.StringVar(&password, "p","", "Password")
	flag.StringVar(&hostname, "h","",  "Hostname")
	flag.StringVar(&port, "P","3306",  "Port")
	flag.StringVar(&database, "d","mysql",  "Database")
	flag.IntVar(&delay, "s", 5, "Delay")
	flag.BoolVar(&noIdle, "i", false, "Hide Idle (sleeping) threads")
}
func main() {
	flag.Parse()
	m, err := db.GetMySQLMonitor(user, password, hostname, port, database)
	if err  != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	pl, err := m.ShowProcessList(noIdle)
	if err != nil {
		log.Fatalf("Error retrieving process list: %v", err)
	}

	terminal.Draw(pl)
	for range time.Tick(time.Second * time.Duration(delay)) {
		terminal.Clear(len(pl) + 1)
		pl, err = m.ShowProcessList(noIdle)
		if err != nil {
			log.Fatalf("Error retrieving process list: %v", err)
		}
		terminal.Draw(pl)
	}
}
