#ifndef __GO_ONNX_TYPE_H__
#define __GO_ONNX_TYPE_H__

typedef struct OnnxParam{
    char* key;
    char* value;
    unsigned int vlen;
    struct OnnxParam* next;
}* OnnxParamList;     // 配置对复用该结构定义

#endif