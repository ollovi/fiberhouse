// Copyright (c) 2025 lamxy and Contributors
// SPDX-License-Identifier: MIT
//
// Author: lamxy <pytho5170@hotmail.com>
// GitHub: https://github.com/lamxy

package frame

import (
	"github.com/lamxy/fiberhouse/frame/appconfig"
	"github.com/lamxy/fiberhouse/frame/bootstrap"
	"github.com/lamxy/fiberhouse/frame/component"
	"github.com/lamxy/fiberhouse/frame/component/validate"
	"github.com/lamxy/fiberhouse/frame/globalmanager"
	"github.com/rs/zerolog"
)

// IContext 全局上下文接口
type IContext interface {
	// GetConfig 定义获取全局配置的方法
	GetConfig() appconfig.IAppConfig
	// GetLogger 定义获取全局日志器的方法
	GetLogger() bootstrap.LoggerWrapper
	// GetContainer 定义获取全局管理器的方法
	GetContainer() *globalmanager.GlobalManager
	// GetStarter 定义获取启动器实例的方法
	GetStarter() IStarter
	// GetLoggerWithOrigin 定义获取附加来源的子日志器单例的方法（从全局管理器获取）
	GetLoggerWithOrigin(originFormCfg appconfig.LogOrigin) (*zerolog.Logger, error)
	// GetMustLoggerWithOrigin 定义获取附加来源的日志器实例的方法，若获取失败则panic（从全局管理器获取）
	GetMustLoggerWithOrigin(originFormCfg appconfig.LogOrigin) *zerolog.Logger
	// GetValidateWrap 定义获取全局验证器包装器的方法
	GetValidateWrap() validate.ValidateWrapper
}

// ContextFramer 框架Web应用上下文接口
type ContextFramer interface {
	IContext
	// RegisterCoreApp 挂载框架核心app
	RegisterCoreApp(core interface{})
	// RegisterStarterApp 挂载框架启动器app
	RegisterStarterApp(sApp ApplicationStarter)
	// GetStarterApp 获取框架启动器实例(FrameApplication)
	GetStarterApp() ApplicationStarter
}

// ContextCommander 框架命令行应用上下文接口
type ContextCommander interface {
	IContext
	// GetDigContainer 获取依赖注入容器
	GetDigContainer() *component.DigContainer
	// RegisterCoreApp 挂载框架核心app
	RegisterCoreApp(core interface{})
	// RegisterStarterApp 挂载框架启动器app
	RegisterStarterApp(app CommandStarter)
	// GetStarterApp 获取框架启动器实例(CommandStarter)
	GetStarterApp() CommandStarter
}
