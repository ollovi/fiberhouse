package validatecustom

import (
	"github.com/lamxy/fiberhouse/example_application/validatecustom/validators"
	"github.com/lamxy/fiberhouse/frame/component/validate"
)

// GetValidateInitializers 获取自定义的验证器初始化器列表
func GetValidateInitializers() []validate.ValidateInitializer {
	return []validate.ValidateInitializer{
		validators.GetJaValidateInitializer(),
		validators.GetKoValidateInitializer(),
	}
}
