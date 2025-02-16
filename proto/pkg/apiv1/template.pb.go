// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/api/v1/template.proto

package apiv1

import (
	templatev1 "github.com/determined-ai/determined/proto/pkg/templatev1"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sorts templates by the given field.
type GetTemplatesRequest_SortBy int32

const (
	// Returns templates in an unsorted list.
	GetTemplatesRequest_SORT_BY_UNSPECIFIED GetTemplatesRequest_SortBy = 0
	// Returns templates sorted by name.
	GetTemplatesRequest_SORT_BY_NAME GetTemplatesRequest_SortBy = 1
)

// Enum value maps for GetTemplatesRequest_SortBy.
var (
	GetTemplatesRequest_SortBy_name = map[int32]string{
		0: "SORT_BY_UNSPECIFIED",
		1: "SORT_BY_NAME",
	}
	GetTemplatesRequest_SortBy_value = map[string]int32{
		"SORT_BY_UNSPECIFIED": 0,
		"SORT_BY_NAME":        1,
	}
)

func (x GetTemplatesRequest_SortBy) Enum() *GetTemplatesRequest_SortBy {
	p := new(GetTemplatesRequest_SortBy)
	*p = x
	return p
}

func (x GetTemplatesRequest_SortBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetTemplatesRequest_SortBy) Descriptor() protoreflect.EnumDescriptor {
	return file_determined_api_v1_template_proto_enumTypes[0].Descriptor()
}

func (GetTemplatesRequest_SortBy) Type() protoreflect.EnumType {
	return &file_determined_api_v1_template_proto_enumTypes[0]
}

func (x GetTemplatesRequest_SortBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetTemplatesRequest_SortBy.Descriptor instead.
func (GetTemplatesRequest_SortBy) EnumDescriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{0, 0}
}

// Get a list of templates.
type GetTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sort templates by the given field.
	SortBy GetTemplatesRequest_SortBy `protobuf:"varint,1,opt,name=sort_by,json=sortBy,proto3,enum=determined.api.v1.GetTemplatesRequest_SortBy" json:"sort_by,omitempty"`
	// Order templates in either ascending or descending order.
	OrderBy OrderBy `protobuf:"varint,2,opt,name=order_by,json=orderBy,proto3,enum=determined.api.v1.OrderBy" json:"order_by,omitempty"`
	// Skip the number of templates before returning results. Negative values
	// denote number of templates to skip from the end before returning results.
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit the number of templates. A value of 0 denotes no limit.
	Limit int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Limit templates to those that match the name.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTemplatesRequest) Reset() {
	*x = GetTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplatesRequest) ProtoMessage() {}

func (x *GetTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplatesRequest.ProtoReflect.Descriptor instead.
func (*GetTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{0}
}

func (x *GetTemplatesRequest) GetSortBy() GetTemplatesRequest_SortBy {
	if x != nil {
		return x.SortBy
	}
	return GetTemplatesRequest_SORT_BY_UNSPECIFIED
}

func (x *GetTemplatesRequest) GetOrderBy() OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return OrderBy_ORDER_BY_UNSPECIFIED
}

func (x *GetTemplatesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetTemplatesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetTemplatesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response to GetTemplatesRequest.
type GetTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the list of requested templates.
	Templates []*templatev1.Template `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
	// Pagination information of the full dataset.
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetTemplatesResponse) Reset() {
	*x = GetTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplatesResponse) ProtoMessage() {}

func (x *GetTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplatesResponse.ProtoReflect.Descriptor instead.
func (*GetTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{1}
}

func (x *GetTemplatesResponse) GetTemplates() []*templatev1.Template {
	if x != nil {
		return x.Templates
	}
	return nil
}

func (x *GetTemplatesResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// Get the requested template.
type GetTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the template.
	TemplateName string `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
}

func (x *GetTemplateRequest) Reset() {
	*x = GetTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateRequest) ProtoMessage() {}

func (x *GetTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{2}
}

func (x *GetTemplateRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

// Response to GetTemplateRequest.
type GetTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The requested template.
	Template *templatev1.Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *GetTemplateResponse) Reset() {
	*x = GetTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateResponse) ProtoMessage() {}

func (x *GetTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateResponse.ProtoReflect.Descriptor instead.
func (*GetTemplateResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{3}
}

func (x *GetTemplateResponse) GetTemplate() *templatev1.Template {
	if x != nil {
		return x.Template
	}
	return nil
}

// Update or create (upsert) the requested template.
type PutTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The template to put.
	Template *templatev1.Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *PutTemplateRequest) Reset() {
	*x = PutTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutTemplateRequest) ProtoMessage() {}

func (x *PutTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutTemplateRequest.ProtoReflect.Descriptor instead.
func (*PutTemplateRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{4}
}

func (x *PutTemplateRequest) GetTemplate() *templatev1.Template {
	if x != nil {
		return x.Template
	}
	return nil
}

// Response to PutTemplateRequest.
type PutTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The updated or created template.
	Template *templatev1.Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *PutTemplateResponse) Reset() {
	*x = PutTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutTemplateResponse) ProtoMessage() {}

func (x *PutTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutTemplateResponse.ProtoReflect.Descriptor instead.
func (*PutTemplateResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{5}
}

func (x *PutTemplateResponse) GetTemplate() *templatev1.Template {
	if x != nil {
		return x.Template
	}
	return nil
}

// Delete the template with the given id.
type DeleteTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the template.
	TemplateName string `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
}

func (x *DeleteTemplateRequest) Reset() {
	*x = DeleteTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateRequest) ProtoMessage() {}

func (x *DeleteTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteTemplateRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTemplateRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

// Response to DeleteTemplateRequest.
type DeleteTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTemplateResponse) Reset() {
	*x = DeleteTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateResponse) ProtoMessage() {}

func (x *DeleteTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateResponse.ProtoReflect.Descriptor instead.
func (*DeleteTemplateResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_template_proto_rawDescGZIP(), []int{7}
}

var File_determined_api_v1_template_proto protoreflect.FileDescriptor

var file_determined_api_v1_template_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b,
	0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x35,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x22, 0xb5, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x1e, 0x92, 0x41, 0x1b, 0x0a, 0x19, 0xd2, 0x01, 0x09, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0xd2, 0x01, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x65, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2, 0x01, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x50, 0x75, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x53, 0x0a, 0x13, 0x50, 0x75,
	0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22,
	0x3c, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_api_v1_template_proto_rawDescOnce sync.Once
	file_determined_api_v1_template_proto_rawDescData = file_determined_api_v1_template_proto_rawDesc
)

func file_determined_api_v1_template_proto_rawDescGZIP() []byte {
	file_determined_api_v1_template_proto_rawDescOnce.Do(func() {
		file_determined_api_v1_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_api_v1_template_proto_rawDescData)
	})
	return file_determined_api_v1_template_proto_rawDescData
}

var file_determined_api_v1_template_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_determined_api_v1_template_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_determined_api_v1_template_proto_goTypes = []interface{}{
	(GetTemplatesRequest_SortBy)(0), // 0: determined.api.v1.GetTemplatesRequest.SortBy
	(*GetTemplatesRequest)(nil),     // 1: determined.api.v1.GetTemplatesRequest
	(*GetTemplatesResponse)(nil),    // 2: determined.api.v1.GetTemplatesResponse
	(*GetTemplateRequest)(nil),      // 3: determined.api.v1.GetTemplateRequest
	(*GetTemplateResponse)(nil),     // 4: determined.api.v1.GetTemplateResponse
	(*PutTemplateRequest)(nil),      // 5: determined.api.v1.PutTemplateRequest
	(*PutTemplateResponse)(nil),     // 6: determined.api.v1.PutTemplateResponse
	(*DeleteTemplateRequest)(nil),   // 7: determined.api.v1.DeleteTemplateRequest
	(*DeleteTemplateResponse)(nil),  // 8: determined.api.v1.DeleteTemplateResponse
	(OrderBy)(0),                    // 9: determined.api.v1.OrderBy
	(*templatev1.Template)(nil),     // 10: determined.template.v1.Template
	(*Pagination)(nil),              // 11: determined.api.v1.Pagination
}
var file_determined_api_v1_template_proto_depIdxs = []int32{
	0,  // 0: determined.api.v1.GetTemplatesRequest.sort_by:type_name -> determined.api.v1.GetTemplatesRequest.SortBy
	9,  // 1: determined.api.v1.GetTemplatesRequest.order_by:type_name -> determined.api.v1.OrderBy
	10, // 2: determined.api.v1.GetTemplatesResponse.templates:type_name -> determined.template.v1.Template
	11, // 3: determined.api.v1.GetTemplatesResponse.pagination:type_name -> determined.api.v1.Pagination
	10, // 4: determined.api.v1.GetTemplateResponse.template:type_name -> determined.template.v1.Template
	10, // 5: determined.api.v1.PutTemplateRequest.template:type_name -> determined.template.v1.Template
	10, // 6: determined.api.v1.PutTemplateResponse.template:type_name -> determined.template.v1.Template
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_determined_api_v1_template_proto_init() }
func file_determined_api_v1_template_proto_init() {
	if File_determined_api_v1_template_proto != nil {
		return
	}
	file_determined_api_v1_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_determined_api_v1_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplatesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_api_v1_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplatesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_api_v1_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_api_v1_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_api_v1_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutTemplateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_api_v1_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutTemplateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_api_v1_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_api_v1_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_determined_api_v1_template_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_api_v1_template_proto_goTypes,
		DependencyIndexes: file_determined_api_v1_template_proto_depIdxs,
		EnumInfos:         file_determined_api_v1_template_proto_enumTypes,
		MessageInfos:      file_determined_api_v1_template_proto_msgTypes,
	}.Build()
	File_determined_api_v1_template_proto = out.File
	file_determined_api_v1_template_proto_rawDesc = nil
	file_determined_api_v1_template_proto_goTypes = nil
	file_determined_api_v1_template_proto_depIdxs = nil
}
