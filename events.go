package main

import (
  "fmt"
  "reflect"
)

type callback interface{}

var events = make(map[string]callback)

func on(name string, callback callback) {
  events[name] = callback;
}

func emit(name string, params ...interface{}) {
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


func main() {
  on("connection", func() {
      fmt.Println("connection created!")
  })

  on("teste", func(name string) {
      fmt.Println("teste", name)
  })

  on("teste2", func(name string, number int) {
      fmt.Println("teste2", name, number)
  })

  on("teste3", func(name string, number int, a float64) {
    fmt.Println("teste3", name, number, a)
  })

  emit("connection")
  emit("teste", "jose donizetti")
  emit("teste2", "jose donizetti", 1)
  emit("teste3", "jose donizetti", 10, 4.3)
}
