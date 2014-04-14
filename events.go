package events

import (
  "reflect"
  "fmt"
  //"unsafe"
)

type callback interface{}

var events = make(map[string][]callback)

func On(name string, c callback) {
  if events[name] == nil {
    events[name] = []callback{c}
  } else {
    events[name] = append(events[name], c)
  }
}

func emitCallback(c interface{}, params []interface{}) {
  callback := reflect.ValueOf(c)

  if length := len(params); length == 0 {
    callback.Call([]reflect.Value{})
  } else {
    values := make([]reflect.Value, len(params))

    for i := 0; i < len(params); i++ {
      values[i] = reflect.ValueOf(params[i])
    }

    //before executing the callback I need
    //to very the arity, and return an error
    fmt.Println("")
    callback.Call(values)
  }
}

func Emit(name string, params ...interface{}) {
  if events[name] != nil {

    callbacks := events[name]

    for i := 0; i < len(callbacks); i++ {
      emitCallback(callbacks[i], params)
    }
  }
}

func AddEventListener(name string, callback callback) {
  On(name, callback)
}
