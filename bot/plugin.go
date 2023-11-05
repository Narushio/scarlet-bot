package bot

type Plugin struct {
	Name       string
	InvokeCmd  string
	CmdActions []*CmdAction
}

func NewPlugin(name string, invokeCmd string) *Plugin {
	p := &Plugin{
		Name:       name,
		InvokeCmd:  invokeCmd,
		CmdActions: make([]*CmdAction, 0),
	}
	return p
}

func (p *Plugin) AddCmdAction(ca *CmdAction) {
	p.CmdActions = append(p.CmdActions, ca)
}

func (p *Plugin) DeleteCmdAction(cmdActionName string) {
	for i, ca := range p.CmdActions {
		if ca.Name == cmdActionName {
			p.CmdActions = append(p.CmdActions[:i], p.CmdActions[i+1:]...)
		}
	}
}
