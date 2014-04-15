package events

import "reflect"
import "fmt"

type listener interface{}

type Event struct {
  id   int
  name string
  listener listener
  once bool
  fired bool
}

var events = make(map[string][]Event)
var eventId = 0

func newEvent(name string, listener listener, once bool) Event {

  //TODO: verify if panic is the best thing to do here
  if len(name) == 0 {
    panic("Listener can't be nil")
  }

  if listener == nil {
    panic("Listener can't be nil")
  }

  eventId += 1
  return Event{eventId, name, listener, once, false}
}

func On(name string, listener listener) Event {
  return addEventListener(name, listener, false)
}

func callListener(event *Event, params []interface{}) {
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

func AddEventListener(name string, listener listener) Event {
  return addEventListener(name, listener, false)
}

func Once(name string, listener listener) Event {
  return addEventListener(name, listener, true)
}

func addEventListener(name string, listener listener, once bool) Event {
  e := newEvent(name,listener,once)

  if events[name] == nil {
    events[name] = []Event{e}
  } else {
    events[name] = append(events[name], e)
  }

  return e
}


func RemoveEventListener(event Event) {
  slice := events[event.name]

  for i := 0; i < len(slice); i++ {
    if slice[i].id == event.id {
      if len(slice) == 1 {
          events[event.name] = []Event{}
      } else {
          fmt.Println(i)
          events[event.name] = append(slice[0:i], slice[i+1:len(slice)]...)
      }
    }
  }
}
