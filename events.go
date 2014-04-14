package events

import (
  "reflect"
  "fmt"
  //"unsafe"
)

type callback interface{}

type event struct {
  name string
  callback callback
  once bool
  called bool
}

var events = make(map[string][]event)

func On(name string, callback callback) {

  if events[name] == nil {
    e := event{name,callback,false,false}
    events[name] = []event{e}
  } else {
    events[name] = append(events[name], event{name,callback,false,false})
  }
}

func emitCallback(event *event, params []interface{}) {
  callback := reflect.ValueOf(event.callback)

  if event.once && event.called {
    return
  }

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

  event.called = true

}

func Emit(name string, params ...interface{}) {
  if events[name] != nil {

    events := events[name]

    for i := 0; i < len(events); i++ {
      emitCallback(&events[i], params)
    }
  }
}

func AddEventListener(name string, callback callback) {
  On(name, callback)
}

func Once(name string, callback callback) {
  if events[name] == nil {
    e := event{name,callback,true,false}
    events[name] = []event{e}
  } else {
    e := event{name,callback,true,false}
    events[name] = append(events[name], e)
  }
}
