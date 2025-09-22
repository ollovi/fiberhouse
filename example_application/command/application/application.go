package application

import (
	"errors"
	"github.com/lamxy/fiberhouse/example_application"
	"github.com/lamxy/fiberhouse/example_application/command/application/commands"
	"github.com/lamxy/fiberhouse/frame"
	"github.com/lamxy/fiberhouse/frame/cache/cachelocal"
	"github.com/lamxy/fiberhouse/frame/cache/cacheremote"
	"github.com/lamxy/fiberhouse/frame/component/jsoncodec"
	"github.com/lamxy/fiberhouse/frame/database/dbmongo"
	"github.com/lamxy/fiberhouse/frame/database/dbmysql"
	"github.com/lamxy/fiberhouse/frame/globalmanager"
	"github.com/urfave/cli/v2"
	"reflect"
	"strconv"
)

// Application 定义应用对象，实现 frame.ApplicationCmdRegister 接口
type Application struct {
	Ctx             frame.ContextCommander
	name            string
	instanceFlagMap map[frame.InstanceKeyFlag]frame.InstanceKey // 预定义实例标识的key映射
}

func NewApplication(ctx frame.ContextCommander) frame.ApplicationCmdRegister {
	return &Application{
		name: "application",
		Ctx:  ctx,
	}
}

// GetName 获取应用注册器名称
func (app *Application) GetName() string {
	return app.name
}

// SetName 设置应用注册器名称
func (app *Application) SetName(name string) {
	app.name = name
}

// GetContext 获取应用上下文对象
func (app *Application) GetContext() frame.ContextCommander {
	return app.Ctx
}

// RegisterGlobalErrHandler 注册应用全局错误处理器
func (app *Application) RegisterGlobalErrHandler(core interface{}) {
	coreApp := core.(*cli.App)
	coreApp.ExitErrHandler = func(cCtx *cli.Context, err error) {
		if err != nil {
			// 记录错误、错误类型、命令行参数、堆栈和命令全名
			app.GetContext().GetLogger().Error(app.GetContext().GetConfig().LogOriginCMD()).Err(err).Str("errType", reflect.TypeOf(err).String()).
				Strs("Args", cCtx.Args().Slice()).Stack().Msg("Command: " + cCtx.Command.FullName())
			// 断言err 是否是 cli.ExitCoder接口类型 interface
			var errExit cli.ExitCoder
			if errors.As(err, &errExit) {
				app.GetContext().GetLogger().Error(app.GetContext().GetConfig().LogOriginCMD()).Err(errExit).Msg("exit: " + strconv.Itoa(errExit.ExitCode()))
				cli.OsExiter(errExit.ExitCode())
			}
			// 默认退出码为 1
			app.GetContext().GetLogger().Error(app.GetContext().GetConfig().LogOriginCMD()).Err(err).Msg("exit: 1")
			cli.OsExiter(1)
		}
	}
}

// RegisterCoreGlobalOptional 注册应用全局可选处理逻辑
func (app *Application) RegisterCoreGlobalOptional(core interface{}) {
	// 转换为具体核心应用类型
	coreApp := core.(*cli.App)

	// 注册全局命令行选项和Action
	coreApp.Flags = app.GetFlagsGlobalOption()
	coreApp.Action = app.GetActionGlobalAction()
}

func (app *Application) GetFlagsGlobalOption() []cli.Flag {
	return FlagsGlobalOption()
}

func (app *Application) GetActionGlobalAction() func(*cli.Context) error {
	return func(c *cli.Context) error {
		// 全局命令的统一Action处理逻辑...

		return nil
	}
}

// RegisterCommands 注册应用命令行命令列表
func (app *Application) RegisterCommands(core interface{}) {
	// 转换为具体核心应用类型
	coreApp := core.(*cli.App)

	cmdGetters := []frame.CommandGetter{
		commands.NewTestOrmCMD(app.Ctx),
		// TODO 收集更多的实现了CommandGetter接口的命令...

	}

	commonCMDs := app.getCommonCommands()

	cliCommands := make([]*cli.Command, 0, len(commonCMDs)+len(cmdGetters))
	copy(cliCommands, commonCMDs)

	for i := range cmdGetters {
		cliCommands = append(cliCommands, cmdGetters[i].GetCommand().(*cli.Command))
	}
	coreApp.Commands = cliCommands
}

// RegisterApplicationGlobals 注册应用全局相关处理逻辑
func (app *Application) RegisterApplicationGlobals() {
	// 注册全局对象初始化器
	initializers := globalmanager.InitializerMap{
		example_application.KEY_MONGODB: func() (interface{}, error) {
			confPath := "database.mongodb"
			return dbmongo.NewMongoDb(app.Ctx, confPath)
		},
		example_application.KEY_MYSQL: func() (interface{}, error) {
			confPath := "database.mysql"
			return dbmysql.NewMysqlDb(app.Ctx, confPath)
		},
		example_application.KEY_REDIS: func() (interface{}, error) {
			confPath := "cache.redis"
			return cacheremote.NewRedisDb(app.Ctx, confPath)
		},
		example_application.KEY_JSON_SONIC_ESCAPE: func() (interface{}, error) {
			return jsoncodec.SonicJsonEscape(), nil
		},
		example_application.KEY_JSON_SONIC_FAST: func() (interface{}, error) {
			return jsoncodec.SonicJsonFastest(), nil
		},
		example_application.KEY_LOCAL_CACHE: func() (interface{}, error) {
			return cachelocal.NewLocalCache(app.Ctx)
		},
		example_application.KEY_REMOTE_CACHE: func() (interface{}, error) {
			return app.GetContext().GetContainer().Get(example_application.KEY_REDIS)
		},
	}
	app.GetContext().GetContainer().Registers(initializers)
	// 预先初始化部分必要的全局对象实例
	requiredInitializers := []globalmanager.KeyName{
		example_application.KEY_MONGODB,
		example_application.KEY_REDIS,
		example_application.KEY_JSON_SONIC_ESCAPE,
		example_application.KEY_JSON_SONIC_FAST,
		example_application.KEY_MYSQL,
	}
	for _, key := range requiredInitializers {
		_, _ = app.GetContext().GetContainer().Get(key)
	}
}

/**
统一定义"获取部分必要对象在容器中的实例Key"
*/

func (app *Application) GetDBKey() string {
	return example_application.KEY_MONGODB
}
func (app *Application) GetDBMongoKey() string {
	return example_application.KEY_MONGODB
}
func (app *Application) GetDBMysqlKey() string {
	return example_application.KEY_MYSQL
}
func (app *Application) GetCacheKey() string {
	return example_application.KEY_REDIS
}
func (app *Application) GetRedisKey() string {
	return example_application.KEY_REDIS
}
func (app *Application) GetFastJsonCodecKey() string {
	return example_application.KEY_JSON_SONIC_FAST
}
func (app *Application) GetDefaultJsonCodecKey() string {
	return example_application.KEY_JSON_SONIC_ESCAPE
}
func (app *Application) GetTaskDispatcherKey() string {
	return example_application.KEY_JSON_SONIC_FAST
}
func (app *Application) GetTaskServerKey() string {
	return example_application.KEY_JSON_SONIC_ESCAPE
}
func (app *Application) GetLocalCacheKey() string {
	return example_application.KEY_LOCAL_CACHE
}
func (app *Application) GetRemoteCacheKey() string {
	return example_application.KEY_REMOTE_CACHE
}
func (app *Application) GetLevel2CacheKey() string {
	return example_application.KEY_LEVEL2_CACHE
}

func (app *Application) GetInstanceKey(flag frame.InstanceKeyFlag) frame.InstanceKey {
	if ik, ok := app.instanceFlagMap[flag]; ok {
		return ik
	}
	return ""
}

// getCommonCommands 获取应用通用命令行命令
func (app *Application) getCommonCommands() []*cli.Command {
	return []*cli.Command{}
}
