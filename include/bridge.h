/*
@author duoliSheep
@address penitente@126.com
@date 2022.06.25 17:12:33
@desc
*/
#ifndef __GO_ONNX_BRIDGE_H
#define __GO_ONNX_BRIDGE_H
#include <stdlib.h>
#include <dlfcn.h>
#include <stdio.h>
#include "./onnxruntime/onnxruntime_c_api.h"
#include "./type.h"
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
    char* loglvl;

	for (OnnxParamList p = params; p != NULL; p = p->next)
    {
         if (p->key != NULL && p->value != NULL) {
			printf("key:%s,length:%d \n",p->key,p->len);
            if (strcmp("log.level",p->key)==0)
            {
                loglvl = p->value;
                continue;
            }
        }
    }
 	//校验input
	return NULL;

  	OrtValue* input_tensor = NULL;

	assert(input_tensor != NULL);
  	int is_tensor;
  	ORT_ABORT_ON_ERROR(g_ort->IsTensor(input_tensor, &is_tensor));
  	assert(is_tensor);


	OrtMemoryInfo* memory_info;

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
	OrtSession* session;
	ORT_ABORT_ON_ERROR(g_ort->Run(session, NULL, input_names, (const OrtValue* const*)&input_tensor, 1, output_names, 1,
									&output_tensor));
	assert(output_tensor != NULL);
	ORT_ABORT_ON_ERROR(g_ort->IsTensor(output_tensor, &is_tensor));
	assert(is_tensor);

	return 0;
}

#endif
