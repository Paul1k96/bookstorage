// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: books.proto

package books

import (
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

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBooksByAuthorIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetBooksByAuthorIn) Reset() {
	*x = GetBooksByAuthorIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksByAuthorIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksByAuthorIn) ProtoMessage() {}

func (x *GetBooksByAuthorIn) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksByAuthorIn.ProtoReflect.Descriptor instead.
func (*GetBooksByAuthorIn) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{2}
}

func (x *GetBooksByAuthorIn) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type GetBooksByAuthorOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books     []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	ErrorCode int32   `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *GetBooksByAuthorOut) Reset() {
	*x = GetBooksByAuthorOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksByAuthorOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksByAuthorOut) ProtoMessage() {}

func (x *GetBooksByAuthorOut) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksByAuthorOut.ProtoReflect.Descriptor instead.
func (*GetBooksByAuthorOut) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{3}
}

func (x *GetBooksByAuthorOut) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *GetBooksByAuthorOut) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type GetAuthorsByBookIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book string `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetAuthorsByBookIn) Reset() {
	*x = GetAuthorsByBookIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsByBookIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsByBookIn) ProtoMessage() {}

func (x *GetAuthorsByBookIn) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsByBookIn.ProtoReflect.Descriptor instead.
func (*GetAuthorsByBookIn) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuthorsByBookIn) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

type GetAuthorsByBookOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors   []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	ErrorCode int32     `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *GetAuthorsByBookOut) Reset() {
	*x = GetAuthorsByBookOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsByBookOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsByBookOut) ProtoMessage() {}

func (x *GetAuthorsByBookOut) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsByBookOut.ProtoReflect.Descriptor instead.
func (*GetAuthorsByBookOut) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{5}
}

func (x *GetAuthorsByBookOut) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *GetAuthorsByBookOut) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

var File_books_proto protoreflect.FileDescriptor

var file_books_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x57, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x4f, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22,
	0x5d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42,
	0x6f, 0x6f, 0x6b, 0x4f, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xab,
	0x01, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x6e, 0x1a, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x1a, 0x1a,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x75, 0x6c, 0x31,
	0x6b, 0x39, 0x36, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x3b, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_proto_rawDescOnce sync.Once
	file_books_proto_rawDescData = file_books_proto_rawDesc
)

func file_books_proto_rawDescGZIP() []byte {
	file_books_proto_rawDescOnce.Do(func() {
		file_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_proto_rawDescData)
	})
	return file_books_proto_rawDescData
}

var file_books_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_books_proto_goTypes = []interface{}{
	(*Author)(nil),              // 0: books.Author
	(*Book)(nil),                // 1: books.Book
	(*GetBooksByAuthorIn)(nil),  // 2: books.GetBooksByAuthorIn
	(*GetBooksByAuthorOut)(nil), // 3: books.GetBooksByAuthorOut
	(*GetAuthorsByBookIn)(nil),  // 4: books.GetAuthorsByBookIn
	(*GetAuthorsByBookOut)(nil), // 5: books.GetAuthorsByBookOut
}
var file_books_proto_depIdxs = []int32{
	1, // 0: books.GetBooksByAuthorOut.books:type_name -> books.Book
	0, // 1: books.GetAuthorsByBookOut.authors:type_name -> books.Author
	2, // 2: books.BookServiceGRPC.GetBooksByAuthor:input_type -> books.GetBooksByAuthorIn
	4, // 3: books.BookServiceGRPC.GetAuthorsByBook:input_type -> books.GetAuthorsByBookIn
	3, // 4: books.BookServiceGRPC.GetBooksByAuthor:output_type -> books.GetBooksByAuthorOut
	5, // 5: books.BookServiceGRPC.GetAuthorsByBook:output_type -> books.GetAuthorsByBookOut
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_books_proto_init() }
func file_books_proto_init() {
	if File_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksByAuthorIn); i {
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
		file_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksByAuthorOut); i {
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
		file_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsByBookIn); i {
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
		file_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsByBookOut); i {
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
			RawDescriptor: file_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_proto_goTypes,
		DependencyIndexes: file_books_proto_depIdxs,
		MessageInfos:      file_books_proto_msgTypes,
	}.Build()
	File_books_proto = out.File
	file_books_proto_rawDesc = nil
	file_books_proto_goTypes = nil
	file_books_proto_depIdxs = nil
}
