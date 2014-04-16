package events

import "testing"
//import "fmt"


func TestEmitSimpleCallback(t *testing.T) {
  var emitted bool

  eventEmitter := NewEventEmitter()
  listener := func() { emitted = true }
  eventEmitter.On("connection", listener)
  eventEmitter.Emit("connection")

  if !emitted {
    t.Errorf("Emitted should be true but was %t", emitted)
  }
}

func TestEmitCallbackWithParameters(t *testing.T) {

  eventEmitter := NewEventEmitter()

  name := "dayane"
  age  := 20
  flag := true

  listener := func(param1 string, param2 int, param3 bool){
    name = param1
    age = param2
    flag = param3
  }

  eventEmitter.On("event", listener)
  eventEmitter.Send("event", "jose", 26, false)

  if name != "jose" {
    t.Errorf("Name should be jose but was %s", name)
  }

  if age != 26 {
    t.Errorf("Age should be 26b but was %d", age)
  }

  if flag {
    t.Errorf("Flag should be false but was %t", flag)
  }
}


func TestAddTwoListenersForTheSameEvent(t *testing.T) {

  eventEmitter := NewEventEmitter()

  count := 0
  listener := func() { count++ }

  eventEmitter.AddEventListener("event", listener)
  eventEmitter.AddEventListener("event", listener)

  eventEmitter.Emit("event")

  if count != 2 {
    t.Errorf("Count should be 2 but was %d", count)
  }
}


func TestEmitCallbackOnlyOnce(t *testing.T) {
  count := 0

  eventEmitter := NewEventEmitter()

  eventEmitter.Once("once", func() {
    count++
  })

  eventEmitter.Emit("once")
  eventEmitter.Emit("once")

  if count != 1 {
    t.Errorf("Count should  be 1 but was %d", count)
  }
}

func TestRemoveEventWithOneListener(t *testing.T) {
  count := 0

  eventEmitter := NewEventEmitter()

  eventListener := eventEmitter.On("event", func() {
      count++
  })

  eventEmitter.Emit("event")
  eventEmitter.RemoveEventListener(eventListener)
  eventEmitter.Emit("event")

  if count != 1 {
    t.Errorf("Count should be 1 but was %d", count)
  }
}


func TestRemoveEventWithTwoListener(t *testing.T) {
  count := 0

  eventEmitter := NewEventEmitter()
  listener := func() { count++ }

  eventListener := eventEmitter.On("event", listener)
  eventEmitter.On("event", listener)

  eventEmitter.Emit("event")
  eventEmitter.Off(eventListener)
  eventEmitter.Emit("event")

  if count != 3 {
    t.Errorf("Count should be 3 but was %d", count)
  }
}


func TestRemoveEventWithThreeListener(t *testing.T) {
  count := 0

  eventEmitter := NewEventEmitter()
  listener := func() { count++ }

  eventEmitter.AddEventListener("event", listener)
  eventListener := eventEmitter.AddEventListener("event", listener)
  eventEmitter.AddEventListener("event", listener)

  eventEmitter.Emit("event")
  eventEmitter.RemoveEventListener(eventListener)
  eventEmitter.Emit("event")


  if count != 5 {
    t.Errorf("Count should be 5 but was %d", count)
  }
}


func TestListenerCount(t *testing.T) {
  listener := func() { }

  eventEmitter := NewEventEmitter()

  eventEmitter.AddEventListener("event", listener)
  eventEmitter.AddEventListener("event", listener)
  eventEmitter.AddEventListener("event", listener)

  eventEmitter.AddEventListener("event1", listener)
  eventEmitter.AddEventListener("event1", listener)

  if (eventEmitter.listenersCount("event") != 3) {
    t.Errorf("Listeners count should be 3 for event")
  }

  if (eventEmitter.listenersCount("event1") != 2) {
    t.Errorf("Listeners count should be 2 for event")
  }

  if (eventEmitter.listenersCount("event3") != 0) {
    t.Errorf("Listeners count should be 0 for event")
  }
}
