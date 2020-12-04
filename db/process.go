package db

import (
	"fmt"
)

func GetProcessListLabels() []string {
	return []string{"ID", "Host", "User", "Db", "Command", "Time", "State", "Info", "Read", "Sent", "Examined"}
}

func (p *Process) GetValueByLabel(label string) string {
	switch label {
	case "ID":
		return fmt.Sprint(p.Id)
	case "Host":
		return p.Host
	case "User":
		return p.User
	case "Db":
		return p.Db.String
	case "Command":
		return p.Command
	case "Time":
		if p.TimeMS > 0 {
			return fmt.Sprintf("%.3f", float64(p.TimeMS)/1000.0)
		}
		return fmt.Sprint(p.Time)
	case "State":
		return p.State.String
	case "Info":
		return p.Info.String
	case "Read":
		return fmt.Sprint(p.RowsRead)
	case "Sent":
		return fmt.Sprint(p.RowsSent)
	case "Examined":
		return fmt.Sprint(p.RowsExamined)
	}
	return ""
}
