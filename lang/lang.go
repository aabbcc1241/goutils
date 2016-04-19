package lang

import (
  "fmt"
  "reflect"
  "strconv"
)

/* reference http://stackoverflow.com/questions/26744873/converting-map-to-struct/26746461#26746461 (stackoverflow/dave) */
func SetField(obj interface{}, name string, value interface{}) error {
  structValue := reflect.ValueOf(obj).Elem()
  structFieldValue := structValue.FieldByName(name)

  if !structFieldValue.IsValid() {
    return fmt.Errorf("No such field: %s in obj %v", name, obj)
  }

  if !structFieldValue.CanSet() {
    return fmt.Errorf("Cannot set %s field value", name)
  }

  structFieldType := structFieldValue.Type()
  val := reflect.ValueOf(value)
  valType := val.Type()
  if structFieldType != valType {
    if val.Kind() == reflect.Float64 {
      switch structFieldValue.Kind() {
      case reflect.Int:
        i, err := strconv.Atoi(fmt.Sprintf("%0.f", value))
        if err == nil {
          structFieldValue.Set(reflect.ValueOf(i))
          return nil
        }
      }
    }
    return fmt.Errorf("Provided value type didn't match obj field type, structType:%v, valueType:%v", structFieldType, val.Type())
  } else {
    structFieldValue.Set(val)
    return nil
  }
}

type GenericStruct struct{}

/* reference http://stackoverflow.com/questions/26744873/converting-map-to-struct/26746461#26746461 (stackoverflow/dave) */
/* demo function to use SetField from interface (e.g. from json.Decoder) */
func (s *GenericStruct) FillStruct(i interface{}) error {
  m := i.(map[string]interface{})
  for k, v := range m {
    if err := SetField(s, k, v); err != nil {
      return err
    }
  }
  return nil
}
