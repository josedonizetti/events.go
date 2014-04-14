package events

import "reflect"

type listener interface{}

type event struct {
  name string
  listener listener
  once bool
  fired bool
}

var events = make(map[string][]event)

func newEvent(name string, listener listener, once bool) event {

  if len(name) == 0 {
    panic("Listener can't be nil")
  }

  if listener == nil {
    panic("Listener can't be nil")
  }

  return event{name, listener, once, false}
}

func On(name string, listener listener) {
  addEventListener(name, listener, false)
}

func callListener(event *event, params []interface{}) {
  listener := reflect.ValueOf(event.listener)

  if event.once && event.fired {
    return
  }

  if length := len(params); length == 0 {
      listener.Call([]reflect.Value{})
  } else {
    values := make([]reflect.Value, len(params))

    for i := 0; i < len(params); i++ {
      values[i] = reflect.ValueOf(params[i])
    }

    //before executing the callback I need
    //to very the arity, and return an error
    listener.Call(values)
  }

  event.fired = true
}

func Emit(name string, params ...interface{}) {
  if events[name] != nil {

    events := events[name]

    for i := 0; i < len(events); i++ {
      callListener(&events[i], params)
    }
  }
}

func AddEventListener(name string, listener listener) {
  addEventListener(name, listener, false)
}

func Once(name string, listener listener) {
  addEventListener(name, listener, true)
}


func addEventListener(name string, listener listener, once bool) {
  e := newEvent(name,listener,once)

  if events[name] == nil {
    events[name] = []event{e}
  } else {
    events[name] = append(events[name], e)
  }
}
