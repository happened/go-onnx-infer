#ifndef __GO_ONNX_TYPE_H__
#define __GO_ONNX_TYPE_H__

typedef struct OnnxParam{
    char* key;
    void* value;
    unsigned int len;
    struct OnnxParam* next;
}* OnnxParamList;     // 配置对复用该结构定义


void* cLibOpen(const char* libName, char** err){
	void* hdl = dlopen(libName, RTLD_NOW);
	if (hdl == NULL){
		*err = (char*)dlerror();
	}
	return hdl;
}

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

struct OnnxParam* OnnxParamCreate(){
	struct OnnxParam* pParamPtr= (struct OnnxParam*)malloc(sizeof(struct OnnxParam));
	pParamPtr->key=NULL;
	pParamPtr->value=NULL;
	pParamPtr->len=0;
	pParamPtr->next=NULL;
	return pParamPtr;
}

struct OnnxParam* OnnxParamAppend(struct OnnxParam* ptr,char* key,void* data,unsigned int len){
	struct OnnxParam* head=ptr;
	if(ptr!=NULL && ptr->key==NULL){
		ptr->key=key;
		ptr->value=data;
		ptr->len=len;
		return ptr;
	}
	while(ptr->next!=NULL){
		ptr=ptr->next;
	}
	struct OnnxParam* pParamPtr= (struct OnnxParam*)malloc(sizeof(struct OnnxParam));
	pParamPtr->key=key;
	pParamPtr->value=data;
	pParamPtr->len=len;
	pParamPtr->next=NULL;
	ptr->next=pParamPtr;
	return head;
}
void OnnxParamFree(struct OnnxParam* ptr){
	struct OnnxParam* current=ptr;
	while(current!=NULL){
		ptr=ptr->next;
		free(current);
		current=ptr;
	}
	return ;
}

#endif