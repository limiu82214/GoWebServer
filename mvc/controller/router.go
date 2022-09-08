package controller

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
)

var Router *router

func init() {
	Router = &router{}
	Router.controllerList = make(map[string]Controller)
}

type ControllerBase struct {
	name string
}

func (c ControllerBase) Name() string {
	return c.name
}

type Controller interface {
	Name() string
}
type router struct {
	controllerList map[string]Controller
}

func (r *router) AddController(c Controller) {
	r.controllerList[c.Name()] = c
}
func (r *router) IsNotActionExist(controller string, action string) bool {
	return !r.IsActionExist(controller, action)
}
func (r *router) IsActionExist(controller string, action string) bool {
	if _, ok := r.controllerList[controller]; !ok {
		return false
	}

	t := reflect.TypeOf(r.controllerList[controller])
	_, isExist := t.MethodByName(action)
	return isExist
}

func (r *router) Execute(controller string, action string, req *http.Request) (rst string, err error) {
	rst = ""
	err = nil
	if _, ok := r.controllerList[controller]; !ok {
		rst = "controller not found"
		err = errors.New(rst)
		return
	}
	t := reflect.TypeOf(r.controllerList[controller])
	reflectMethod, isExist := t.MethodByName(strings.ToUpper(action)[0:1] + action[1:])
	if !isExist {
		rst = "action not found"
		err = errors.New(rst)
		return
	}

	in := []reflect.Value{reflect.ValueOf(r.controllerList[controller]), reflect.ValueOf(req)}
	rv := reflectMethod.Func.Call(in)
	rst = rv[0].String()
	if rv[1].IsNil() {
		err = nil
	} else {
		err = rv[1].Interface().(error)
	}
	return

}
