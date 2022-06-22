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
#include "../include/onnxruntime_c_api.h"
#include "../include/type.h"
// @return library handle
void* cLibOpen(const char* libName, char** err){
	void* hdl = dlopen(libName, RTLD_NOW);
	if (hdl == NULL){
		*err = (char*)dlerror();
	}
	return hdl;
}
// @return symbol address
void* cLibLoad(void* hdl, const char* sym, char** err){
	void* addr = dlsym(hdl, sym);
	if (addr == NULL){
		*err = (char*)dlerror();
	}
	return addr;
}
int  cLibClose(void* hdl){
	int ret = dlclose(hdl);
	if (ret != 0)
		return -1;
	return 0;
}

const OrtApi* g_ort = NULL;

#define ORT_ABORT_ON_ERROR(expr)                             \
  do {                                                       \
    OrtStatus* onnx_status = (expr);                         \
    if (onnx_status != NULL) {                               \
      const char* msg = g_ort->GetErrorMessage(onnx_status); \
      fprintf(stderr, "%s\n", msg);                          \
      g_ort->ReleaseStatus(onnx_status);                     \
      abort();                                               \
    }                                                        \
  } while (0);


int Run() {
	 g_ort = OrtGetApiBase()->GetApi(ORT_API_VERSION);
	  if (!g_ort) {
		fprintf(stderr, "Failed to init ONNX Runtime engine.\n");
		return -1;
	  }
	OrtMemoryInfo* memory_info;
	ORT_ABORT_ON_ERROR(g_ort->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &memory_info));

	
	return 0;
}

*/
import "C"
import "fmt"

func Inference() {
	ret := C.Run()
	fmt.Println(int(ret))
}
