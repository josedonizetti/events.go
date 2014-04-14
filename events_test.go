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

  On("connection", func(param1 string, param2 int, param3 bool){
    name = param1
    age = param2
    flag = param3
  })

  Emit("connection", "jose", 26, false)

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

  On("event", func() {
    listener1 = false
  })

  AddEventListener("event", func() {
    listener2 = false
  })

  Emit("event")

  if listener1 {
    t.Errorf("Listerner1 should be false")
  }

  if listener2 {
    t.Errorf("Listerner2 should be false")
  }
}
