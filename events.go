package events

import "reflect"

type callback interface{}

type event struct {
  name string
  callback callback
  once bool
  called bool
}

var events = make(map[string][]event)

func newEvent(name string, callback callback, once bool) event {
  return event{name, callback, once, false}
}

func On(name string, callback callback) {

  if events[name] == nil {
    events[name] = []event{ newEvent(name,callback,false) }
  } else {
    events[name] = append(events[name], newEvent(name,callback,false))
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
    events[name] = []event{ newEvent(name,callback,true) }
  } else {
    events[name] = append(events[name], newEvent(name,callback,true))
  }
}
