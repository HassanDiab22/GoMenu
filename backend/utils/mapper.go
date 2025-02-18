package utils

import (
	"fmt"
	"reflect"
)

func AutoMap(source interface{}, destination interface{}) error {
	srcVal := reflect.ValueOf(source)
	destVal := reflect.ValueOf(destination)

	// Ensure destination is a pointer
	if destVal.Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be a pointer")
	}

	destElem := destVal.Elem()

	// Handle slice mapping
	if srcVal.Kind() == reflect.Slice {
		destSliceType := destElem.Type().Elem()

		// Ensure the destination is a slice
		if destElem.Kind() != reflect.Slice {
			return fmt.Errorf("destination must be a slice")
		}

		// Create a new slice of the destination type
		newSlice := reflect.MakeSlice(destElem.Type(), srcVal.Len(), srcVal.Len())

		// Iterate over the source slice and map each element
		for i := 0; i < srcVal.Len(); i++ {
			srcItem := srcVal.Index(i)

			// Create a new destination struct
			destItem := reflect.New(destSliceType).Elem()

			// Map fields from source item to destination item
			mapStruct(srcItem, destItem)

			// Assign to new slice
			newSlice.Index(i).Set(destItem)
		}

		// Set the destination slice
		destElem.Set(newSlice)
		return nil
	}

	// Handle single struct mapping
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if destElem.Kind() == reflect.Ptr {
		destElem = destElem.Elem()
	}

	mapStruct(srcVal, destElem)
	return nil
}

// mapStruct maps fields between two structs with matching names
func mapStruct(srcVal, destVal reflect.Value) {
	srcType := srcVal.Type()
	destType := destVal.Type()

	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		srcValue := srcVal.Field(i)

		if destField, ok := destType.FieldByName(srcField.Name); ok {
			destFieldValue := destVal.FieldByName(destField.Name)
			if destFieldValue.CanSet() && destFieldValue.Type() == srcValue.Type() {
				destFieldValue.Set(srcValue)
			}
		}
	}
}
