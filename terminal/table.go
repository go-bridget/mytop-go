package terminal

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"os/exec"
	"runtime"

	"github.com/carmo-evan/mytop-go/db"
)

func Clear() error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func Draw(pl db.ProcessList) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Host", "User", "Db", "Command", "Time", "State", "Info", "Sent", "Examined")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, r := range pl {
		tbl.AddRow(r.Id, r.Host, r.User, r.Db.String, r.Command, r.Time, r.State, r.Info.String, r.RowsSent, r.RowsExamined)
	}
	tbl.Print()
}
