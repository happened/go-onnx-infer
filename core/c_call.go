package core

/*
* @Author: kejun.sheng
* @Email: kejun.sheng@cyclone-robotics.com
* @DateTime: 2022.06.22 11:32:10
 */

/*
#cgo linux CFLAGS: -I../include -Wno-attributes
#cgo LDFLAGS: -L../lib -lonnxruntime -ldl
#include <stdlib.h>
#include <dlfcn.h>
#include <stdio.h>
#include "../include/bridge.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type OnnxInput struct {
	Name  string
	Shape []int
	Data  []byte
}

func Inference(inputs map[string]OnnxInput) {
	paramList := C.OnnxParamCreate()
	defer C.OnnxParamFree(paramList)
	for _, v := range inputs {
		key := C.CString(v.Name)
		defer C.free(unsafe.Pointer(key))

		data := C.CBytes(v.Data)
		defer C.free(unsafe.Pointer(data))

		valLen := C.uint(len(v.Data))
		paramList = C.OnnxParamAppend(paramList, key, data, valLen)
	}
	output := C.RunFloat(paramList)
	if output == nil {
		fmt.Println("empty result")
	}
}
