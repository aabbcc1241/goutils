package lang

import (
	"fmt"
	"github.com/aabbcc1241/goutils/log"
	"reflect"
	"strconv"
)

/* reference http://stackoverflow.com/questions/26744873/converting-map-to-struct/26746461#26746461 (stackoverflow/dave) */
func SetField(obj interface{}, name string, value interface{}) error {
	//log.Debug.Println("name",name,"value",value)
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
		switch val.Kind() {
		case reflect.Float64:
			switch structFieldValue.Kind() {
			case reflect.Int:
				i, err := strconv.Atoi(fmt.Sprintf("%0.f", value))
				if err == nil {
					structFieldValue.Set(reflect.ValueOf(i))
					return nil
				}
				break
			case reflect.Int64:
				i, err := strconv.ParseInt(fmt.Sprintf("%0.f", value), 10, 64)
				if err == nil {
					structFieldValue.Set(reflect.ValueOf(i))
					return nil
				}
				break
			case reflect.Uint64:
				i, err := strconv.ParseUint(fmt.Sprintf("%0.f", value), 10, 64)
				if err == nil {
					structFieldValue.Set(reflect.ValueOf(i))
					return nil
				}
				break
			}
			break
		case reflect.String:
			log.Debug.Println("struct type kind", structFieldType.Kind())
			switch structFieldValue.Kind() {
			}
		}
		return fmt.Errorf("Provided value type didn't match obj field type, name:%v, value:%v, structType:%v, valueType:%v", name, value, structFieldType, val.Type())
	} else {
		structFieldValue.Set(val)
		return nil
	}
}

type demo_s struct{}

/* reference http://stackoverflow.com/questions/26744873/converting-map-to-struct/26746461#26746461 (stackoverflow/dave) */
/* demo function to use SetField from interface (e.g. from json.Decoder) */
func (s *demo_s) FillStruct(i interface{}) error {
	m := i.(map[string]interface{})
	for k, v := range m {
		if err := SetField(s, k, v); err != nil {
			return err
		}
	}
	return nil
}

func BytesToInterfaces(bs []byte) []interface{} {
	xs := make([]interface{}, len(bs))
	for i, v := range bs {
		xs[i] = v
	}
	return xs
}
