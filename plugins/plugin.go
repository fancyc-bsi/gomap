package plugins

import (
	"fmt"
)

type Plugin interface {
	Run(ip, port string) (string, error)
	GetName() string
}

type HostPlugin interface {
	Plugin
}

var registeredPlugins = make(map[string]Plugin)
var portToPluginsMap = make(map[string][]string)
var serviceToPluginsMap = make(map[string][]string)
var hostPlugins []HostPlugin

func RegisterPlugin(plugin Plugin, ports []string, services []string) {
	name := plugin.GetName()
	registeredPlugins[name] = plugin
	for _, port := range ports {
		portToPluginsMap[port] = append(portToPluginsMap[port], name)
	}
	for _, service := range services {
		serviceToPluginsMap[service] = append(serviceToPluginsMap[service], name)
	}
}

func RegisterHostPlugin(plugin HostPlugin) {
	hostPlugins = append(hostPlugins, plugin)
}

func GetHostPlugins() []HostPlugin {
	return hostPlugins
}

func GetPluginByPort(port, service string) ([]Plugin, error) {
	var pluginNames []string
	pluginNames = append(pluginNames, portToPluginsMap[port]...)
	pluginNames = append(pluginNames, serviceToPluginsMap[service]...)

	if len(pluginNames) == 0 {
		return nil, nil
	}

	var plugins []Plugin
	for _, pluginName := range pluginNames {
		plugin, exists := registeredPlugins[pluginName]
		if !exists {
			return nil, fmt.Errorf("plugin not found: %s", pluginName)
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}
