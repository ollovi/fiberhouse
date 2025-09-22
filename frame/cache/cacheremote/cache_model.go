// Copyright (c) 2025 lamxy and Contributors
// SPDX-License-Identifier: MIT
//
// Author: lamxy <pytho5170@hotmail.com>
// GitHub: https://github.com/lamxy

package cacheremote

import (
	"github.com/lamxy/fiberhouse/frame"
	"github.com/lamxy/fiberhouse/frame/cache"
)

// CacheModel 用于被具体的Cache Proxy实例继承
type CacheModel struct {
	name  string
	Ctx   frame.IContext
	proxy frame.Locator
}

func NewCacheModel(ctx frame.IContext) *CacheModel {
	return &CacheModel{
		Ctx: ctx,
	}
}

// GetContext 获取应用上下文
func (mo *CacheModel) GetContext() frame.IContext {
	return mo.Ctx
}

// GetRemote 获取远程缓存实例
func (mo *CacheModel) GetRemote() cache.Cache {
	key := mo.GetContext().GetStarter().GetApplication().GetRemoteCacheKey()
	return frame.GetMustInstance[cache.Cache](key)
}

// GetRemote 获取本地缓存实例
func (mo *CacheModel) GetLocal() cache.Cache {
	key := mo.GetContext().GetStarter().GetApplication().GetLocalCacheKey()
	return frame.GetMustInstance[cache.Cache](key)
}

// GetLevel2 获取二级缓存实例
func (mo *CacheModel) GetLevel2() cache.Cache {
	key := mo.GetContext().GetStarter().GetApplication().GetLevel2CacheKey()
	return frame.GetMustInstance[cache.Cache](key)
}

// SetTarget 设置缓存实例定位器
func (mo *CacheModel) SetOrigin(locator frame.Locator) frame.Locator { // SetTarget
	mo.proxy = locator
	return mo
}

// GetTarget 获取缓存实例定位器
func (mo *CacheModel) GetOrigin() frame.Locator {
	return mo.proxy
}

// GetName 获取名称
func (mo *CacheModel) GetName() string {
	return mo.name
}

// SetName 设置名称
func (mo *CacheModel) SetName(name string) frame.Locator {
	mo.name = name
	return mo
}

// GetInstance 获取实例（从全局管理器获取具体的单例）
func (mo *CacheModel) GetInstance(namespaceKey string) (interface{}, error) {
	gm := mo.GetContext().GetContainer()
	return gm.Get(namespaceKey)
}
