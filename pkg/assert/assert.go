package assert

import (
	"reflect"
	"runtime"
)

func Nil(err error) {
	if err != nil {
		panic(err)
	}
}

func True(value bool, err error) {
	if !value {
		panic(err)
	}
}

func False(value bool, err error) {
	True(!value, err)
}

// NotNil 断言对象不为 nil 否则 panic，主要用于应用启动阶段单例对象创建方法返回前的检查
func NotNil(object interface{}) {
	True(!IsNil(object), nil)
}

// NotCircular 通过当前 go routine 调用栈判断是否有循环调用依赖，有则 panic。
// 主要用于应用启动阶段单例对象创建时循环依赖检查，作为创建方法的第一条语句。
func NotCircular() {
	// 最多取 100 个 frame，单例创建一般在应用启动阶段 100 应该足够
	pc := make([]uintptr, 100)
	// skip 取 2 去掉 runtime 和 NotCircular 的调用
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	current, more := frames.Next()
	var next runtime.Frame
	for more {
		next, more = frames.Next()
		// 方法名形如(github.com/eaglecloud-inc/user-center/app/biz/demo.DefaultDemoService)
		// 因此相等即认为存在循环调用，直接 panic
		if current.Function == next.Function {
			panic("found circular dependency")
		}
	}
}

// IsNil checks if a specified object is nil or not
// See:
// https://github.com/stretchr/testify/blob/master/assert/assertions.go#L520
// https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	switch value.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func:
		return value.IsNil()
	}
	return false
}
