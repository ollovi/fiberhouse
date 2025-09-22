// Copyright (c) 2025 lamxy and Contributors
// SPDX-License-Identifier: MIT
//
// Author: lamxy <pytho5170@hotmail.com>
// GitHub: https://github.com/lamxy

package frame

// CommandStarter 命令行脚本启动器接口，定义命令行程序启动流程
type CommandStarter interface {
	IStarter
	// InitCoreApp 初始化核心命令行应用
	InitCoreApp(...interface{})
	// RegisterGlobalErrHandler 注册全局错误处理器
	RegisterGlobalErrHandler()
	// RegisterCommands 收集命令列表并注册到核心应用
	RegisterCommands()
	// RegisterCoreGlobalOptional 注册应用核心的全局初始化，如果有必要
	RegisterCoreGlobalOptional()
	// RegisterApplicationGlobals 注册应用预定义的全局对象实例到全局管理器容器中
	RegisterApplicationGlobals()
	// RegisterLoggerWithOriginToContainer 将预定义带日志源日志器注册到全局管理器容器中
	RegisterLoggerWithOriginToContainer()
	// RegisterApplication 注册应用注册器对象到启动器的Application属性
	RegisterApplication(application ApplicationCmdRegister)
	// AppCoreRun 运行核心命令行应用
	AppCoreRun() error
}

// ApplicationCmdRegister 命令行应用注册器
// 由用户自定义的，在CMD应用启动器启动时，实现必要的应用逻辑，
// 并注册绑定到CommandStarter的application属性，由启动器调用完成应用初始化
type ApplicationCmdRegister interface {
	IRegister
	IApplication
	// GetContext 返回全局上下文
	GetContext() ContextCommander

	// RegisterGlobalErrHandler 注册全局错误处理器到核心应用
	RegisterGlobalErrHandler(core interface{})
	// RegisterCommands 收集命令列表并注册到核心应用
	RegisterCommands(core interface{})
	// RegisterCoreGlobalOptional 注册应用核心的全局可选项
	RegisterCoreGlobalOptional(core interface{})
	// RegisterApplicationGlobals 注册应用预定义的全局对象实例到全局管理器容器中
	RegisterApplicationGlobals()
}

// CommandGetter 命令获取器接口，定义了获取单个命令的方法
type CommandGetter interface {
	GetCommand() interface{}
}
