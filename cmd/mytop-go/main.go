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
	"fmt"
	"github.com/carmo-evan/mytop-go/db"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"runtime"
	"strings"
	"time"
)

var user, password, hostname string

func init() {
	flag.StringVar(&user, "u" ,"", "Username")
	flag.StringVar(&password, "p","", "Password")
	flag.StringVar(&hostname, "h","",  "Hostname")
}
func main() {
	flag.Parse()
	m := db.GetMySQLMonitor(user, password, hostname)
	pl := m.ShowProcessList()
	m.ShowGlobalStatus()
	Draw(pl)
	fmt.Fprint(os.Stdout,"\r \r")
	for range time.Tick(time.Second * 1) {
		Clear(len(pl) + 1)
		pl = m.ShowProcessList()
		m.ShowGlobalStatus()
		Draw(pl)
	}
}

func Clear(n int) {
	if runtime.GOOS == "windows" {
		clearString := "\r" + strings.Repeat(" ", n) + "\r"
		fmt.Fprint(os.Stdout, clearString)
		return
	}

	fmt.Fprintf(os.Stdout, "\033[%dA", n) //move cursor to the top

	for _, s := range []string{"\b", "\127", "\b", "\033[K", "\x0c", string(27)} { // "\033[K" for macOS Terminal
		os.Stdout.Write([]byte(strings.Repeat(s, n)))
	}
	os.Stdout.Write([]byte("\r\033[k")) // erases to end of line
}

func Draw(pl []db.ProcessList) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Host", "User", "Db", "Command", "Time", "State", "Info")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, r := range pl {
		tbl.AddRow(r.Id, r.Host, r.User, r.Db, r.Command, r.Time, r.State, r.Info)
	}
	tbl.Print()
}
