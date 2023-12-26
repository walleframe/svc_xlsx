package svc_xlsx

import (
	"github.com/walleframe/walle/app/bootstrap"
	"github.com/walleframe/walle/services/configcentra"
	"github.com/walleframe/walle/xlsxmgr"
	"github.com/walleframe/walle/xlsxmgr/plugin/local"
	"github.com/walleframe/walle/xlsxmgr/plugin/simple"
)

var (
	registry = pluginRegistry{
		plugins: make(map[string]xlsxmgr.XlsxLoadPlugin),
	}
	XlsxMgr = xlsxmgr.NewXlsxConfigManager()
)

// RegisterXlsxPlugin register xlsx plugin to load data
var RegisterXlsxPlugin func(plugin xlsxmgr.XlsxLoadPlugin) = registry.RegisterXlsxPlugin

// RegAutoConfig 注册自动加载配置
var RegAutoConfig func(basename, fromXlsx, fromSheet string, parser xlsxmgr.DataParser, loader xlsxmgr.ConfigLoader) = XlsxMgr.RegAutoConfig

// RegAppendConfig 追加数据解析
var RegAppendConfig func(basename, tag string, loader xlsxmgr.ConfigAppendLoader) = XlsxMgr.RegAppendConfig

// RegMixConfig 混合数据接口
var RegMixConfig func(tag string, mix func() error, basename ...string) = XlsxMgr.RegMixConfig

func init() {
	// register default plugin
	RegisterXlsxPlugin(simple.XlsxPlugin)
	RegisterXlsxPlugin(local.XlsxPlugin)

	svc := NewService(XlsxMgr)
	// regitster plugin type config.
	configcentra.String(&svc.cfgName, "xlsxcfg.plugin", simple.XlsxPlugin.Name(), "xlsx plugin name")

	// xlsx manager service
	bootstrap.RegisterServiceByPriority(20, svc)
}
