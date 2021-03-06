// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recipes.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d348b217594f5a3a, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type RecipeTitle struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecipeTitle) Reset()         { *m = RecipeTitle{} }
func (m *RecipeTitle) String() string { return proto.CompactTextString(m) }
func (*RecipeTitle) ProtoMessage()    {}
func (*RecipeTitle) Descriptor() ([]byte, []int) {
	return fileDescriptor_d348b217594f5a3a, []int{1}
}

func (m *RecipeTitle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipeTitle.Unmarshal(m, b)
}
func (m *RecipeTitle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipeTitle.Marshal(b, m, deterministic)
}
func (m *RecipeTitle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipeTitle.Merge(m, src)
}
func (m *RecipeTitle) XXX_Size() int {
	return xxx_messageInfo_RecipeTitle.Size(m)
}
func (m *RecipeTitle) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipeTitle.DiscardUnknown(m)
}

var xxx_messageInfo_RecipeTitle proto.InternalMessageInfo

func (m *RecipeTitle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type RecipeList struct {
	Title                []string `protobuf:"bytes,1,rep,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecipeList) Reset()         { *m = RecipeList{} }
func (m *RecipeList) String() string { return proto.CompactTextString(m) }
func (*RecipeList) ProtoMessage()    {}
func (*RecipeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d348b217594f5a3a, []int{2}
}

func (m *RecipeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipeList.Unmarshal(m, b)
}
func (m *RecipeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipeList.Marshal(b, m, deterministic)
}
func (m *RecipeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipeList.Merge(m, src)
}
func (m *RecipeList) XXX_Size() int {
	return xxx_messageInfo_RecipeList.Size(m)
}
func (m *RecipeList) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipeList.DiscardUnknown(m)
}

var xxx_messageInfo_RecipeList proto.InternalMessageInfo

func (m *RecipeList) GetTitle() []string {
	if m != nil {
		return m.Title
	}
	return nil
}

type Ingredient struct {
	Amount               string   `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ingredient) Reset()         { *m = Ingredient{} }
func (m *Ingredient) String() string { return proto.CompactTextString(m) }
func (*Ingredient) ProtoMessage()    {}
func (*Ingredient) Descriptor() ([]byte, []int) {
	return fileDescriptor_d348b217594f5a3a, []int{3}
}

func (m *Ingredient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingredient.Unmarshal(m, b)
}
func (m *Ingredient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingredient.Marshal(b, m, deterministic)
}
func (m *Ingredient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingredient.Merge(m, src)
}
func (m *Ingredient) XXX_Size() int {
	return xxx_messageInfo_Ingredient.Size(m)
}
func (m *Ingredient) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingredient.DiscardUnknown(m)
}

var xxx_messageInfo_Ingredient proto.InternalMessageInfo

func (m *Ingredient) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Ingredient) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Recipe struct {
	Title                string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Ingredients          []*Ingredient `protobuf:"bytes,2,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	Process              string        `protobuf:"bytes,3,opt,name=process,proto3" json:"process,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Recipe) Reset()         { *m = Recipe{} }
func (m *Recipe) String() string { return proto.CompactTextString(m) }
func (*Recipe) ProtoMessage()    {}
func (*Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_d348b217594f5a3a, []int{4}
}

func (m *Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipe.Unmarshal(m, b)
}
func (m *Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipe.Marshal(b, m, deterministic)
}
func (m *Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipe.Merge(m, src)
}
func (m *Recipe) XXX_Size() int {
	return xxx_messageInfo_Recipe.Size(m)
}
func (m *Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Recipe proto.InternalMessageInfo

func (m *Recipe) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Recipe) GetIngredients() []*Ingredient {
	if m != nil {
		return m.Ingredients
	}
	return nil
}

func (m *Recipe) GetProcess() string {
	if m != nil {
		return m.Process
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "protobuf.Empty")
	proto.RegisterType((*RecipeTitle)(nil), "protobuf.RecipeTitle")
	proto.RegisterType((*RecipeList)(nil), "protobuf.RecipeList")
	proto.RegisterType((*Ingredient)(nil), "protobuf.Ingredient")
	proto.RegisterType((*Recipe)(nil), "protobuf.Recipe")
}

func init() {
	proto.RegisterFile("recipes.proto", fileDescriptor_d348b217594f5a3a)
}

var fileDescriptor_d348b217594f5a3a = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x93, 0xc6, 0x26, 0x76, 0x82, 0x7f, 0x18, 0xaa, 0x2c, 0x39, 0x95, 0xf5, 0xd2, 0x53,
	0x90, 0x0a, 0x22, 0x78, 0xf2, 0xe0, 0x41, 0xf0, 0x14, 0x7c, 0x01, 0xdb, 0x8c, 0xb2, 0x60, 0x93,
	0x65, 0x77, 0x2a, 0xe4, 0xad, 0x7c, 0x44, 0xc9, 0x26, 0x71, 0x43, 0xec, 0x69, 0xe7, 0x1b, 0x7e,
	0xfb, 0x7d, 0x33, 0x03, 0x67, 0x86, 0x76, 0x4a, 0x93, 0xcd, 0xb5, 0xa9, 0xb9, 0xc6, 0x53, 0xf7,
	0x6c, 0x0f, 0x1f, 0x32, 0x81, 0xf9, 0xf3, 0x5e, 0x73, 0x23, 0x6f, 0x20, 0x2d, 0x1c, 0xf3, 0xa6,
	0xf8, 0x8b, 0x70, 0x09, 0x73, 0x6e, 0x0b, 0x11, 0xae, 0xc2, 0xf5, 0xa2, 0xe8, 0x84, 0x94, 0x00,
	0x1d, 0xf4, 0xaa, 0x2c, 0x8f, 0x99, 0xc8, 0x33, 0x0f, 0x00, 0x2f, 0xd5, 0xa7, 0xa1, 0x52, 0x51,
	0xc5, 0x78, 0x0d, 0xf1, 0xfb, 0xbe, 0x3e, 0x54, 0xdc, 0x1b, 0xf5, 0x0a, 0x11, 0x4e, 0xb8, 0xd1,
	0x24, 0x66, 0xae, 0xeb, 0x6a, 0xa9, 0x21, 0xee, 0xdc, 0x8f, 0xa7, 0xe3, 0x3d, 0xa4, 0xea, 0xcf,
	0xd9, 0x8a, 0xd9, 0x2a, 0x5a, 0xa7, 0x9b, 0x65, 0x3e, 0xec, 0x92, 0xfb, 0xd8, 0x62, 0x0c, 0xa2,
	0x80, 0x44, 0x9b, 0x7a, 0x47, 0xd6, 0x8a, 0xc8, 0xf9, 0x0d, 0x72, 0xf3, 0x13, 0x42, 0xd2, 0x45,
	0xda, 0xd6, 0xbd, 0xdd, 0x6a, 0x90, 0x17, 0xde, 0xd7, 0x1d, 0x28, 0x1b, 0x05, 0xf9, 0x1b, 0xc8,
	0x00, 0x1f, 0xe1, 0xbc, 0x20, 0x36, 0x8a, 0xbe, 0xa9, 0x9f, 0xfe, 0x6a, 0x4a, 0xba, 0x93, 0x66,
	0x97, 0xd3, 0xb6, 0x0c, 0xf0, 0x16, 0x16, 0x4f, 0x65, 0xd9, 0xff, 0xfb, 0x07, 0x64, 0xd3, 0x21,
	0x64, 0xb0, 0x8d, 0x5d, 0xe7, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xa5, 0x6c, 0x7d, 0xd2,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecipesClient is the client API for Recipes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecipesClient interface {
	ListRecipes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RecipeList, error)
	RetrieveRecipe(ctx context.Context, in *RecipeTitle, opts ...grpc.CallOption) (*Recipe, error)
	AddRecipe(ctx context.Context, in *Recipe, opts ...grpc.CallOption) (*Empty, error)
}

type recipesClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipesClient(cc grpc.ClientConnInterface) RecipesClient {
	return &recipesClient{cc}
}

func (c *recipesClient) ListRecipes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RecipeList, error) {
	out := new(RecipeList)
	err := c.cc.Invoke(ctx, "/protobuf.Recipes/ListRecipes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesClient) RetrieveRecipe(ctx context.Context, in *RecipeTitle, opts ...grpc.CallOption) (*Recipe, error) {
	out := new(Recipe)
	err := c.cc.Invoke(ctx, "/protobuf.Recipes/RetrieveRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesClient) AddRecipe(ctx context.Context, in *Recipe, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Recipes/AddRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipesServer is the server API for Recipes service.
type RecipesServer interface {
	ListRecipes(context.Context, *Empty) (*RecipeList, error)
	RetrieveRecipe(context.Context, *RecipeTitle) (*Recipe, error)
	AddRecipe(context.Context, *Recipe) (*Empty, error)
}

// UnimplementedRecipesServer can be embedded to have forward compatible implementations.
type UnimplementedRecipesServer struct {
}

func (*UnimplementedRecipesServer) ListRecipes(ctx context.Context, req *Empty) (*RecipeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipes not implemented")
}
func (*UnimplementedRecipesServer) RetrieveRecipe(ctx context.Context, req *RecipeTitle) (*Recipe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveRecipe not implemented")
}
func (*UnimplementedRecipesServer) AddRecipe(ctx context.Context, req *Recipe) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecipe not implemented")
}

func RegisterRecipesServer(s *grpc.Server, srv RecipesServer) {
	s.RegisterService(&_Recipes_serviceDesc, srv)
}

func _Recipes_ListRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).ListRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Recipes/ListRecipes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).ListRecipes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recipes_RetrieveRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecipeTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).RetrieveRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Recipes/RetrieveRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).RetrieveRecipe(ctx, req.(*RecipeTitle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recipes_AddRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Recipe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).AddRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Recipes/AddRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).AddRecipe(ctx, req.(*Recipe))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recipes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Recipes",
	HandlerType: (*RecipesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRecipes",
			Handler:    _Recipes_ListRecipes_Handler,
		},
		{
			MethodName: "RetrieveRecipe",
			Handler:    _Recipes_RetrieveRecipe_Handler,
		},
		{
			MethodName: "AddRecipe",
			Handler:    _Recipes_AddRecipe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recipes.proto",
}
