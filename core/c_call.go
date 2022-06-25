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

struct OnnxParamList* OnnxParamListCreate(){
	struct OnnxParamList* pParamPtr= (struct OnnxParamList*)malloc(sizeof(struct OnnxParamList));
	pParamPtr->key=NULL;
	pParamPtr->value=NULL;
	pParamPtr->vlen=0;
	pParamPtr->next=NULL;
	return pParamPtr;
}

struct OnnxParamList* OnnxParamListAppend(struct OnnxParamList* ptr,char* key,char* value,unsigned int vlen){
	struct OnnxParamList* head=ptr;
	if(ptr!=NULL && ptr->key==NULL){
		ptr->key=key;
		ptr->value=value;
		ptr->vlen=vlen;
		return ptr;
	}
	while(ptr->next!=NULL){
		ptr=ptr->next;
	}
	struct OnnxParamList* pParamPtr= (struct OnnxParamList*)malloc(sizeof(struct OnnxParamList));
	pParamPtr->key=key;
	pParamPtr->value=value;
	pParamPtr->vlen=vlen;
	pParamPtr->next=NULL;
	ptr->next=pParamPtr;
	return head;
}
void OnnxParamListfree(struct OnnxParamList* ptr){
	struct OnnxParamList* current=ptr;
	while(current!=NULL){
		ptr=ptr->next;
		free(current);
		current=ptr;
	}
	return ;
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


int AllOcEnv(){

	g_ort = OrtGetApiBase()->GetApi(ORT_API_VERSION);
	if (!g_ort) {
		fprintf(stderr, "Failed to init ONNX Runtime engine.\n");
		return -1;
	}
	OrtMemoryInfo* memory_info;
	ORT_ABORT_ON_ERROR(g_ort->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &memory_info));
	// OrtValue* input_tensor = NULL;
	// //1. 内存info 2. 具体的输入数据float* 3， 输入数据长度 4. 模型解析出来的输入形状比如[1,3,720,720] 5.input形态的长度 6， 数据类型 7.生成的tensor
    // ORT_ABORT_ON_ERROR(g_ort->CreateTensorWithDataAsOrtValue(memory_info, model_input, model_input_len, input_shape,
    //                                                        input_shape_len, ONNX_TENSOR_ELEMENT_DATA_TYPE_FLOAT,
    //                                                        &input_tensor));
	return 0;
}

OnnxParamList RunFloat(OnnxParamList params) {
	for (OnnxParamList p = params; p != NULL; p = p->next)
    {
         if (p->key != NULL && p->value != NULL) {
            if (std::string("log.level") ==std::string(p->key))
            {
                loglvl = p->value;
                continue;
            }
        }
    }
	return NULL;
	//校验input
	assert(input_tensor != NULL);
  	int is_tensor;
  	ORT_ABORT_ON_ERROR(g_ort->IsTensor(input_tensor, &is_tensor));
  	assert(is_tensor);
  	g_ort->ReleaseMemoryInfo(memory_info);

	const char* input_names[] = {"inputImage"};
	const char* output_names[] = {"outputImage"};
	OrtValue* output_tensor = NULL;

	// 1. session
	// 2. run_options
	// 3. 输入名称
	// 4. 输入的tensor
	// 5. 输入tensor的key的长度
	// 6. 输出名称集合
	// 7. 输出名称长度
	// 8. 输出的tensor
	ORT_ABORT_ON_ERROR(g_ort->Run(session, NULL, input_names, (const OrtValue* const*)&input_tensor, 1, output_names, 1,
									&output_tensor));
	assert(output_tensor != NULL);
	ORT_ABORT_ON_ERROR(g_ort->IsTensor(output_tensor, &is_tensor));
	assert(is_tensor);

	return 0;
}

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
	paramList := C.OnnxParamListCreate()
	defer C.OnnxParamListfree(paramList)

	for _, v := range inputs {
		key := C.CString(v.Name)
		defer C.free(unsafe.Pointer(key))

		val := C.CString(v.Data)
		defer C.free(unsafe.Pointer(val))

		valLen := C.uint(len(v.Data))
		paramList = C.paramListAppend(paramList, key, val, valLen)
	}
	output := C.RunFloat(paramList)
	if output == nil {
		fmt.Println("empty result")
	}
}
