package events

import "testing"


func TestEmitSimpleCallback(t *testing.T) {
  var emitted bool

  On("connection", func() {
      emitted = true
  })

  Emit("connection")

  if !emitted {
    t.Errorf("Emitted should be true")
  }
}

func TestEmitCallbackWithParameters(t *testing.T) {
  name := "dayane"
  age  := 20
  flag := true

  On("event1", func(param1 string, param2 int, param3 bool){
    name = param1
    age = param2
    flag = param3
  })

  Emit("event1", "jose", 26, false)

  if name != "jose" {
    t.Errorf("Name should be jose")
  }

  if age != 26 {
    t.Errorf("Age should be 26")
  }

  if flag {
    t.Errorf("Flag should be false")
  }
}


func TestAddTwoListenersForTheSameEvent(t *testing.T) {

  listener1 := true
  listener2 := true

  On("event2", func() {
    listener1 = false
  })

  AddEventListener("event2", func() {
    listener2 = false
  })

  Emit("event2")

  if listener1 {
    t.Errorf("Listerner1 should be false")
  }

  if listener2 {
    t.Errorf("Listerner2 should be false")
  }
}


func TestEmitCallbackOnlyOnce(t *testing.T) {
  count := 0

  Once("once", func() {
    count++
  })

  Emit("once")
  Emit("once")

  if count != 1 {
    t.Error("Count should  be 1")
  }
}
