package helpers

import (
	"reflect"
	"web/todo/pkg/common/logger"
)

type Module interface {
	Init()
}

func getModuleName(m Module) string {
	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}

func InitializeModule(modules []Module) {
	for _, module := range modules {
		logger.Log.Infof("%s Init", getModuleName(module))
		module.Init()
	}
}
