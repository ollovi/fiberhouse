package validatecustom

import (
	"github.com/lamxy/fiberhouse/example_application/validatecustom/tags"
	"github.com/lamxy/fiberhouse/frame/component/validate"
)

// GetValidatorTagFuncs 获取注册指定或自定义tag及翻译提示
func GetValidatorTagFuncs() []validate.RegisterValidatorTagFunc {
	return []validate.RegisterValidatorTagFunc{
		tags.StartswithRegisterTranslation,
		tags.HascoursesRegisterValidation,
		tags.HascoursesRegisterTranslation,
	}
}
