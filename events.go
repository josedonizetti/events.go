package events

import "reflect"

type listener interface{}

type isNil interface {
  isNil()
}

type EventListener struct {
  id   int
  name string
  listener listener
  once bool
  fired bool
}

type EventEmitter struct {
  events map[string][]EventListener
  eventId int
  defaultMaxListeners int
}

func newEventListener(emitter *EventEmitter, name string, listener listener, once bool) EventListener {

  if len(name) == 0 {
    panic("Event name can't be nil")
  }

  if listener == nil {
    panic("Listener can't be nil")
  }

  if reflect.TypeOf(listener).Kind() != reflect.Func {
    panic("Listener must be a func")
  }

  emitter.eventId += 1
  return EventListener{emitter.eventId, name, listener, once, false}
}

func (listener *EventListener) isNil() bool {
  return listener == nil
}

func NewEventEmitter() EventEmitter {
  events := make(map[string][]EventListener)
  return EventEmitter{events, 0, 10}
}

func (emitter *EventEmitter) On(name string, listener listener) EventListener {
  return emitter.addEventListener(name, listener, false)
}

func (emitter *EventEmitter) Off(eventListener EventListener) {
  emitter.RemoveEventListener(eventListener)
}

func (emitter *EventEmitter) Once(name string, listener listener) EventListener {
  return emitter.addEventListener(name, listener, true)
}

func (emitter *EventEmitter) AddEventListener(name string, listener listener) EventListener {
  return emitter.addEventListener(name, listener, false)
}

func (emitter *EventEmitter) addEventListener(name string, listener listener, once bool) EventListener {

  if(emitter.events[name] == emitter.defaultMaxListeners) {
    //TODO: change it to be an error
    panic("Max limit of listener for the event %s was reached", name)
  }

  e := newEventListener(emitter, name, listener, once)

  if emitter.events[name] == nil {
    emitter.events[name] = []EventListener{e}
  } else {
    emitter.events[name] = append(emitter.events[name], e)
  }

  if (emitter.events["newEventListener"] != nil) && (name != "newEventListener") {
    emitter.Emit("newEventListener", name)
  }

  return e
}

func (emitter *EventEmitter) RemoveEventListener(eventListener EventListener) {
  if eventListener.isNil() {
    panic("eventListener should not be nil")
  }

  slice := emitter.events[eventListener.name]

  for i := 0; i < len(slice); i++ {
    if slice[i].id == eventListener.id {
      if len(slice) == 1 {
          emitter.events[eventListener.name] = []EventListener{}
      } else {
          emitter.events[eventListener.name] = append(slice[0:i], slice[i+1:len(slice)]...)
      }

      if (emitter.events["removeEventListener"] != nil) && (eventListener.name != "removeEventListener") {
        emitter.Emit("removeEventListener", eventListener.name)
      }

      return
    }
  }
}

func (emitter *EventEmitter) callListener(eventListener *EventListener, params []interface{}) {
  listenerFunc := reflect.ValueOf(eventListener.listener)

  if eventListener.once && eventListener.fired {
    return
  }

  if length := len(params); length == 0 {
      listenerFunc.Call([]reflect.Value{})
  } else {
    values := make([]reflect.Value, len(params))

    for i := 0; i < len(params); i++ {
      values[i] = reflect.ValueOf(params[i])
    }

    listenerFunc.Call(values)
  }

  eventListener.fired = true
}

func (emitter *EventEmitter) Emit(name string, params ...interface{}) {
  if emitter.events[name] != nil {

    events := emitter.events[name]

    for i := 0; i < len(events); i++ {
      emitter.callListener(&events[i], params)
    }
  }
}

func (emitter *EventEmitter) Send(name string, params ...interface{}) {
  emitter.Emit(name, params...)
}

func (emitter *EventEmitter) listenersCount(name string) int {

  if emitter.events[name] != nil {
    return len(emitter.events[name])
  }

  return 0
}

func (emitter *EventEmitter) RemoveAllEventListeners(name string) {

  if len(name) == 0 {
    panic("Event name can't be nil")
  }

  emitter.events[name] = []EventListener{}
}

func (emitter *EventEmitter) setMaxListeners(maxListeners int) {
  if maxListeners == nil {
    panic("MaxListener can't be nil")
  }

  if maxListeners < 1 {
    panic("MaxListener must be a positive number")
  }

  emitter.defaultMaxListeners = maxListeners
}
