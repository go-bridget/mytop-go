package db

import "strconv"

func (p *Process) GetLabels() []string {
	return []string{"ID", "Host", "User", "Db", "Command", "Time", "State", "Info", "Sent", "Examined"}
}

func (p *Process) GetValueByLabel(label string) string {
	switch label {
	case "ID":
		return strconv.Itoa(p.Id)
	case "Host":
		return p.Host
	case "User":
		return p.User
	case "Db":
		return p.Db.String
	case "Command":
		return p.Command
	case "Time":
		return strconv.Itoa(p.Time)
	case "State":
		return p.State
	case "Info":
		return p.Info.String
	case "Sent":
		return strconv.Itoa(p.RowsSent)
	case "Examined":
		return strconv.Itoa(p.RowsExamined)
	}
	return ""
}
