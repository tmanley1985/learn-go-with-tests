package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// If we have a pointer then of course we need to go get that
	// actual value to run the function over!
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			// Why do we do this .Interface() thing?
			// Well, think about it, we don't know what the type
			// is here, so we could either figure it out or pass the
			// value in as an empty interface. The walk function already
			// handles the figuring out of the type anyway, so we don't have to.
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
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