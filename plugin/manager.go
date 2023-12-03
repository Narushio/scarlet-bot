package plugin

type Manager struct {
	plugins map[string]Plugin
}

func NewManager() *Manager {
	return &Manager{
		plugins: make(map[string]Plugin),
	}
}

func (m *Manager) RegisterPlugins(plugs ...Plugin) {
	for _, plug := range plugs {
		m.plugins[plug.Name] = plug
	}
}
