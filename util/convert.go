// 注意不能转换map，想转换自己写吧=。=
package util

import (
	"fmt"
	"reflect"
)

func isSlice(v reflect.Value) (isSlice bool) {
	defer func() {
		if err := recover(); err != nil {
			isSlice = false
		}
	}()
	_ = v.Cap()
	return true
}

func isNilPointer(v reflect.Value) (isNilPointer bool) {
	defer func() {
		if err := recover(); err != nil {
			isNilPointer = false
		}
	}()

	return v.IsNil()
}

func PoToVo(po interface{}, vo interface{}) {
	// fmt.Println("------------PoToVo-------------------")
	// fmt.Println(b)
	// vo := reflect.New(t).Elem()
	voPtrValue := reflect.ValueOf(vo)
	poValue := reflect.ValueOf(po)
	voValue := reflect.Indirect(voPtrValue)
	poValueToVoValue(poValue, voValue)
	fmt.Println()
}

func poValueToVoValue(poValue reflect.Value, voValue reflect.Value) {
	poValue = reflect.Indirect(poValue)

	if isNilPointer(voValue) {
		voElemType := voValue.Type().Elem()
		fmt.Printf("voValue: %v\n", voValue)
		fmt.Printf("voType: %v\n", voValue.Type())
		fmt.Printf("voElemType: %v\n", voElemType)
		voValue.Set(reflect.Indirect(reflect.New(voValue.Type())))
	}

	voValue = reflect.Indirect(voValue)

	voType := voValue.Type()

	fmt.Printf("\n%v %v \n", poValue.Type(), voType)

	if isSlice(voValue) {
		fmt.Printf("%v is slice and prepare to convert from %v\n", voValue, poValue.Interface())
		poListValueToVoListValue(poValue, voValue.Addr())
		return
	}

	for m := 0; m < voType.NumField(); m++ {
		field := voType.Field(m)

		poFieldvalue := poValue.FieldByName(field.Name)
		voFieldValue := voValue.FieldByName(field.Name)

		if !reflect.Indirect(poFieldvalue).IsValid() {
			continue
		}
		fmt.Printf("%v\n", field.Name)
		fmt.Printf("%v %v\n", poFieldvalue.Type(), voFieldValue.Type())
		fmt.Printf("%v %v\n", poFieldvalue, voFieldValue)
		if voFieldValue.Type() != poFieldvalue.Type() {
			poValueToVoValue(poFieldvalue, voFieldValue)
		} else {

			voValue.FieldByName(field.Name).Set(poFieldvalue)
		}
	}
}

func poListValueToVoListValue(poListValue reflect.Value, voListPtrValue reflect.Value) {
	voListType := voListPtrValue.Type().Elem()
	voListValue := reflect.Indirect(voListPtrValue)
	fmt.Printf("voListType: %v\n", voListType)
	fmt.Printf("volistValue: %v\n", voListValue)
	voType := voListType.Elem()
	fmt.Printf("voType: %v\n", voType)
	len := poListValue.Len()
	for i := 0; i < len; i++ {
		poValue := poListValue.Index(i)
		voValue := reflect.Indirect(reflect.New(voType))
		poValueToVoValue(poValue, voValue)
		voListValue = reflect.Append(voListValue, voValue)
	}

	if len == 0 {
		voListValue = reflect.MakeSlice(voListType, 0, 0)
	}

	voListPtrValue.Elem().Set(voListValue)
}

func PoListToVoList(poList interface{}, voList interface{}) {
	voListPtrValue := reflect.ValueOf(voList)
	poListValue := reflect.ValueOf(poList)

	poListValueToVoListValue(poListValue, voListPtrValue)
	fmt.Println()
}
