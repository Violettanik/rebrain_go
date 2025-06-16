package convertor

import (
	"errors"
	"reflect"
	"strings"
        "time"
)

// StructToMap преобразует структуру в map[string]interface{}
func StructToMap(item interface{}) map[string]interface{} {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Пропускаем неэкспортируемые поля
		if !fieldValue.CanInterface() {
			continue
		}

		// Получаем имя поля из тега или используем имя поля
		key := getFieldName(field)
                if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
                    result[key] = fieldValue.Interface()
                    continue
                }

                // Обрабатываем вложенные структуры
		if fieldValue.Kind() == reflect.Struct {
			result[key] = StructToMap(fieldValue.Interface())
			continue
		}

		// Обрабатываем указатели на структуры
		if fieldValue.Kind() == reflect.Ptr && fieldValue.Elem().Kind() == reflect.Struct {
			if !fieldValue.IsNil() {
				result[key] = StructToMap(fieldValue.Interface())
			} else {
				result[key] = nil
			}
			continue
		}

		result[key] = fieldValue.Interface()
	}

	return result
}

// MapToStruct заполняет структуру из map[string]interface{}
func MapToStruct(mp map[string]interface{}, item interface{}) error {
	v := reflect.ValueOf(item)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("item must be a pointer to a struct")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if !fieldValue.CanSet() {
			continue
		}

		key := getFieldName(field)

		mapValue, exists := mp[key]
		if !exists {
			continue
		}

		if mapValue == nil {
			continue
		}

                // Специальная обработка time.Time
                if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
                    if timeValue, ok := mapValue.(time.Time); ok {
                        fieldValue.Set(reflect.ValueOf(timeValue))
                    }
                    continue
                }

		// Обрабатываем вложенные структуры
		if fieldValue.Kind() == reflect.Struct {
			if nestedMap, ok := mapValue.(map[string]interface{}); ok {
				if err := MapToStruct(nestedMap, fieldValue.Addr().Interface()); err != nil {
					return err
				}
			}
			continue
		}

		// Обрабатываем указатели на структуры
		if fieldValue.Kind() == reflect.Ptr && fieldValue.Type().Elem().Kind() == reflect.Struct {
			if nestedMap, ok := mapValue.(map[string]interface{}); ok {
				if fieldValue.IsNil() {
					fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
				}
				if err := MapToStruct(nestedMap, fieldValue.Interface()); err != nil {
					return err
				}
			}
			continue
		}

		// Устанавливаем значение
		mapValueReflect := reflect.ValueOf(mapValue)
		if mapValueReflect.Type().AssignableTo(fieldValue.Type()) {
			fieldValue.Set(mapValueReflect)
		} else if mapValueReflect.Type().ConvertibleTo(fieldValue.Type()) {
			fieldValue.Set(mapValueReflect.Convert(fieldValue.Type()))
		}
	}

	return nil
}

// getFieldName извлекает имя поля из тега или возвращает имя поля
func getFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("keyname")
	if tag != "" {
		// Учитываем возможность нескольких тегов (например, `keyname:"name,omitempty"`)
		if commaIdx := strings.Index(tag, ","); commaIdx != -1 {
			tag = tag[:commaIdx]
		}
		return tag
	}
	return field.Name
}
