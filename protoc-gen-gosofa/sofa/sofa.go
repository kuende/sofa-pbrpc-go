package sofa

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

// generatedCodeVersion indicates a version of the generated code.
// It is incremented whenever an incompatibility between the generated code and
// the sofa pbrpc package is introduced; the generated code references
// a constant, sofa.SupportPackageIsVersionN (where N is generatedCodeVersion).
const generatedCodeVersion = 1

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	contextPkgPath = "context"
	sofaPkgPath    = "github.com/kuende/sofa-pbrpc-go/sofa"
)

func init() {
	generator.RegisterPlugin(new(sofa))
}

// sofa is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for sofa-pbrpc support.
type sofa struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "sofa".
func (g *sofa) Name() string {
	return "sofa"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.
var (
	contextPkg string
	sofaPkg    string
)

// Init initializes the plugin.
func (g *sofa) Init(gen *generator.Generator) {
	g.gen = gen
	contextPkg = generator.RegisterUniquePackageName("context", nil)
	sofaPkg = generator.RegisterUniquePackageName("sofa", nil)
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *sofa) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *sofa) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *sofa) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *sofa) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ ", contextPkg, ".Context")
	g.P("var _ ", sofaPkg, ".Conn")
	g.P()

	// Assert version compatibility.
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the sofa-pbrpc package it is being compiled against.")
	g.P("const _ = ", sofaPkg, ".SupportPackageIsVersion", generatedCodeVersion)
	g.P()

	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *sofa) GenerateImports(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	g.P("import (")
	g.P(contextPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, contextPkgPath)))
	g.P(sofaPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, sofaPkgPath)))
	g.P(")")
	g.P()
}

// reservedClientName records whether a client name is reserved on the client side.
var reservedClientName = map[string]bool{
// TODO: do we need any in sofa?
}

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

// generateService generates all the code for the named service.
func (g *sofa) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.

	origServName := service.GetName()
	fullServName := origServName
	if pkg := file.GetPackage(); pkg != "" {
		fullServName = pkg + "." + fullServName
	}
	servName := generator.CamelCase(origServName)

	g.P()
	g.P("// Client API for ", servName, " service")
	g.P()

	// Client interface.
	g.P("type ", servName, "Client interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateClientSignature(servName, method))
	}
	g.P("}")
	g.P()

	// Client structure.
	g.P("type ", unexport(servName), "Client struct {")
	g.P("cc ", sofaPkg, ".Conn")
	g.P("}")
	g.P()

	// NewClient factory.
	g.P("func New", servName, "Client (cc ", sofaPkg, ".Conn) ", servName, "Client {")
	g.P("return &", unexport(servName), "Client{cc}")
	g.P("}")
	g.P()

	serviceDescVar := "_" + servName + "_serviceDesc"
	// Client method implementations.
	for _, method := range service.Method {
		g.generateClientMethod(servName, fullServName, serviceDescVar, method)
	}
}

// generateClientSignature returns the client-side signature for a method.
func (g *sofa) generateClientSignature(servName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}
	reqArg := ", in *" + g.typeName(method.GetInputType())
	respName := "*" + g.typeName(method.GetOutputType())
	return fmt.Sprintf("%s(ctx %s.Context%s, opts ...%s.CallOption) (%s, error)", methName, contextPkg, reqArg, sofaPkg, respName)
}

func (g *sofa) generateClientMethod(servName, fullServName, serviceDescVar string, method *pb.MethodDescriptorProto) {
	sname := fmt.Sprintf("%s.%s", fullServName, method.GetName())
	outType := g.typeName(method.GetOutputType())

	g.P("func (c *", unexport(servName), "Client) ", g.generateClientSignature(servName, method), "{")
	g.P("out := new(", outType, ")")
	g.P("err := ", sofaPkg, `.Invoke(ctx, "`, sname, `", in, out, c.cc, opts...)`)
	g.P("if err != nil { return nil, err }")
	g.P("return out, nil")
	g.P("}")
	g.P()
}
