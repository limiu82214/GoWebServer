package controller

import (
	"net/http"
	"reflect"
	"strings"
)

type ControllerBase struct {
	name string
}

func (c ControllerBase) Name() string {
	return c.name
}

type Controller interface {
	Name() string
}
type Router struct {
	controllerList []Controller
}

func (r *Router) AddController(c Controller) {
	r.controllerList = append(r.controllerList, c)
}
func (r *Router) IsNotActionExist(controller string, action string) bool {
	return !r.IsActionExist(controller, action)
}
func (r *Router) IsActionExist(controller string, action string) bool {
	for _, c := range r.controllerList {
		if c.Name() == controller {
			t := reflect.TypeOf(c)
			_, isExist := t.MethodByName(action)
			if isExist {
				return true
			}
			// in := []reflect.Value{reflect.ValueOf(&controller.Member{}), reflect.ValueOf(w), reflect.ValueOf(r)}
			// rv := reflectMethod.Func.Call(in)
		}
	}

	return false
}

func (r *Router) Execute(controller string, action string, w http.ResponseWriter, req *http.Request) (rst string, err error) {
	for _, c := range r.controllerList {
		if c.Name() == controller {
			t := reflect.TypeOf(c)
			reflectMethod, isExist := t.MethodByName(strings.ToUpper(action)[0:1] + action[1:])
			if isExist {
				// panic(reflectMethod)

				in := []reflect.Value{reflect.ValueOf(c), reflect.ValueOf(w), reflect.ValueOf(req)}
				rv := reflectMethod.Func.Call(in)
				rst = rv[0].String()
				if rv[1].IsNil() {
					err = nil
				} else {
					err = rv[1].Interface().(error)
				}
				return
			}
		}
	}

	rst = ""
	err = nil
	return
}
