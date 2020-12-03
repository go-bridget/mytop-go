package terminal

import (
	"fmt"
	"github.com/carmo-evan/mytop-go/db"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"runtime"
	"strings"
)

func Clear(n int) {
	if runtime.GOOS == "windows" {
		clearString := "\r" + strings.Repeat(" ", n) + "\r"
		fmt.Fprint(os.Stdout, clearString)
		return
	}

	c := fmt.Sprintf("\033[%dA", n) //move cursor to the top
	os.Stdout.Write([]byte(c))

	for _, s := range []string{"\b", "\127", "\b", "\033[K", "\r"} { // "\033[K" for macOS Terminal
		os.Stdout.Write([]byte(strings.Repeat(s, n)))
	}
	os.Stdout.Write([]byte("\r\033[k")) // erases to end of line
}

func Draw(pl []db.ProcessList) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Host", "User", "Db", "Command", "Time", "State", "Info", "Rows Sent", "Rows Examined")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, r := range pl {
		var database, info string
		if r.Info != nil {
			info = *r.Info
		}
		if r.Db != nil {
			database = *r.Db
		}
		tbl.AddRow(r.Id, r.Host, r.User, database, r.Command, r.Time, r.State, info, r.RowsSent, r.RowsExamined)
	}
	tbl.Print()
}
