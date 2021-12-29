package whitelist

import "github.com/df-mc/dragonfly/server/cmd"

func NewRunnable(w *WhiteList) Command { return Command{whitelist: w} }

type Command struct {
	whitelist *WhiteList
	Status    status
}

func (c Command) Run(src cmd.Source, o *cmd.Output) {
	switch string(c.Status) {
	case "enable", "on":
		c.whitelist.Enabled = true
		o.Print("server is no longer whitelisted")
	case "disable", "off":
		c.whitelist.Enabled = false
		o.Print("server is now whitelisted")
	default:
		o.Printf("'%s' is not a valid argument for this command!", c.Status)
	}
}

type status string

func (status) Type() string                    { return "status" }
func (status) Options(src cmd.Source) []string { return []string{"enable", "disable"} }
