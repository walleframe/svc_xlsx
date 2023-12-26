package svc_xlsx

import (
	"github.com/walleframe/walle/app"
	"github.com/walleframe/walle/zaplog"

	"github.com/walleframe/walle/xlsxmgr"
)

// ExcelConfigService excel配置管理器
type ExcelConfigService struct {
	app.NoopService

	cfgName string

	xlsxCfg *xlsxmgr.XlsxConfig
}

func NewService(cfg *xlsxmgr.XlsxConfig) *ExcelConfigService {
	return &ExcelConfigService{
		xlsxCfg: cfg,
	}
}

func (svc *ExcelConfigService) Name() string {
	return "xlsx-config-svc"
}

func (svc *ExcelConfigService) Init(s app.Stoper) (err error) {
	// 根据配置决定加载插件
	plugin, err := xlsxmgr.GetPlugin(svc.cfgName)
	if err != nil {
		return err
	}

	// 新建管理器
	err = svc.xlsxCfg.Init(plugin, zaplog.GetFrameLogger())
	if err != nil {
		return err
	}

	// 初始化配置
	return svc.xlsxCfg.UpdateConfig(s)
}

func (svc *ExcelConfigService) Stop() {
	svc.xlsxCfg.Stop()
}
