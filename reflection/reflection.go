package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// If we have a pointer then of course we need to go get that
	// actual value to run the function over!
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		// Why do we do this .Interface() thing?
		// Well, think about it, we don't know what the type
		// is here, so we could either figure it out or pass the
		// value in as an empty interface. The walk function already
		// handles the figuring out of the type anyway, so we don't have to.
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		// We just iterate forever until the channel is closed.
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {

		val = val.Elem()
	}
	return val
}