package svc_xlsx

import (
	"fmt"
	"log"
	"strings"

	"github.com/walleframe/walle/xlsxmgr"
	"go.uber.org/zap"
)

type pluginRegistry struct {
	plugins map[string]xlsxmgr.XlsxLoadPlugin
}

func (registry *pluginRegistry) RegisterXlsxPlugin(plugin xlsxmgr.XlsxLoadPlugin) {
	name := plugin.Name()
	name = strings.ToLower(name)
	if old, ok := registry.plugins[name]; ok {
		log.Println(fmt.Sprintf("xlsx plugin[%s] %#T replaced by %#T:\n%s", name, old, plugin, zap.StackSkip("", 1).String))
	}
	registry.plugins[name] = plugin
}

func (registry *pluginRegistry) GetPlugin(name string) (xlsxmgr.XlsxLoadPlugin, error) {
	name = strings.ToLower(name)
	plugin, ok := registry.plugins[name]
	if !ok {
		return nil, fmt.Errorf("xlsx plugin[%s] not register", name)
	}
	return plugin, nil
}
