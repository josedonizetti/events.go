package events

import (
  "reflect"
)

type callback interface{}

var events = make(map[string]callback)

func On(name string, callback callback) {
  events[name] = callback;
}

func Emit(name string, params ...interface{}) {
  if events[name] != nil {
    callback := reflect.ValueOf(events[name])

    if length := len(params); length == 0 {
      callback.Call(make([]reflect.Value, 0))
    } else {
      values := make([]reflect.Value, len(params))

      for i := 0; i < len(params); i++ {
        values[i] = reflect.ValueOf(params[i])
      }

      callback.Call(values)
    }
  }
}

func AddEventListener(name string, callback callback) {
  
}
