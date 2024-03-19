package helper

import "reflect"

func InterateStructField[T comparable](payload T) interface{} {
  ref := reflect.ValueOf(payload)

  values := make([]interface{}, ref.NumField());

  for i := 0; i < ref.NumField(); i++ {
    values[i] = ref.Field(i).Interface()
  }

  return values
}
