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

func TestRemoveEventWithOneListener(t *testing.T) {
  count := 0

  event := On("testRemoveListener1", func() {
      count++
  })

  Emit("testRemoveListener1")
  RemoveEventListener(event)
  Emit("testRemoveListener1")

  if count != 1 {
    t.Error("Count should be 1")
  }
}


func TestRemoveEventWithTwoListener(t *testing.T) {
  count := 0

  event1 := On("testRemoveListener2", func() {
      count++
  })

  count2 := 0

  On("testRemoveListener3", func() {
      count2++
  })

  Emit("testRemoveListener2")
  Emit("testRemoveListener3")

  RemoveEventListener(event1)

  Emit("testRemoveListener2")
  Emit("testRemoveListener3")

  if count != 1 {
    t.Error("Count should be 1")
  }

  if count2 != 2 {
    t.Error("Count2 should be 2")
  }
}


func TestRemoveEventWithThreeListener(t *testing.T) {
  count := 0

  On("testRemoveListener4", func() {
      count++
  })

  count2 := 0

  event5 := On("testRemoveListener5", func() {
      count2++
  })

  On("testRemoveListener6", func() {
      count++
  })

  Emit("testRemoveListener4")
  Emit("testRemoveListener5")
  Emit("testRemoveListener6")

  RemoveEventListener(event5)

  Emit("testRemoveListener4")
  Emit("testRemoveListener5")
  Emit("testRemoveListener6")

  if count != 4 {
    t.Error("Count should be 4")
  }

  if count2 != 1 {
    t.Error("Count2 should be 1")
  }
}
