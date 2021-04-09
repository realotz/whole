package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

type BaseInfo struct {
	filename, basePackage, path, serviceNameC, serviceName string
}

func (b *BaseInfo) ImportPath(str string) string {
	strList := strings.Split(fmt.Sprintf("%s/%s/%s", b.basePackage, b.path, str), "/")
	newList := []string{}
	for _, v := range strList {
		if v != "" {
			newList = append(newList, v)
		}
	}
	return fmt.Sprintf("\"%s\"", strings.Join(newList, "/"))
}

// generateFile generates a _http.pb.go file containing kratos errors definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File, params map[string]string) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	path, _ := params["path"]
	path = strings.ReplaceAll(path, ".", "")
	basePackage := strings.ReplaceAll(file.GoImportPath.String(), "\"", "")
	packages := strings.Split(basePackage, "/")
	if len(packages) > 2 {
		basePackage = strings.Join(packages[0:3], "/")
		// 根据api定位包目录
		for k, v := range packages {
			if v == "api" {
				basePackage = strings.Join(packages[0:k], "/")
			}
		}
	}
	for _, s := range file.Services {
		serviceName := strings.ReplaceAll(strings.ToLower(string(s.Desc.Name())), "service", "")
		serviceNameC := Capitalize(serviceName)
		filename := serviceName + ".go"
		info := &BaseInfo{
			filename, basePackage, path, serviceNameC, serviceName,
		}
		generateService(gen, file, s, info)
		generateBiz(gen, file, s, info)
		generateEntSchema(gen, file, s, info)
		generateData(gen, file, s, info)
	}
	return nil
}

func generateData(gen *protogen.Plugin, file *protogen.File, s *protogen.Service, info *BaseInfo) {
	var objmsg *protogen.Message
	for _, msg := range file.Messages {
		if string(msg.Desc.Name()) == info.serviceNameC {
			objmsg = msg
		}
	}
	g := gen.NewGeneratedFile("/data/"+info.filename, file.GoImportPath)
	g.P("package data")
	g.P("import (")
	g.P("\"context\"")
	g.P("\"github.com/go-kratos/kratos/v2/log\"")
	g.P(fmt.Sprintf("pb %s", file.GoImportPath.String()))
	g.P(info.ImportPath("biz"))
	g.P(info.ImportPath("data/ent"))
	g.P(")")
	g.P("type ent" + info.serviceNameC + " ent." + info.serviceNameC)
	g.P("// ent转换biz结构")
	g.P("func (p ent" + info.serviceNameC + ") BizStruct() *biz." + info.serviceNameC + " {")
	g.P("return &biz." + info.serviceNameC + "{")
	if objmsg != nil {
		for _, f := range objmsg.Fields {
			name := MarshalStr(string(f.Desc.Name()))
			if name == "Id" {
				g.P(fmt.Sprintf("%s:p.%s,", name, "ID"))
			} else {
				g.P(fmt.Sprintf("%s:p.%s,", name, name))
			}
		}
	}
	g.P("}}")
	g.P("type " + info.serviceName + "Repo struct {\n\tdata *Data \n\t log  *log.Helper\n}")
	g.P("func New" + info.serviceNameC + "Repo(data *Data, logger log.Logger) biz." + info.serviceNameC + "Repo {")
	g.P("return &" + info.serviceName + "Repo{data: data,log:  log.NewHelper(\"" + info.serviceName + "_repo\", logger),}}")

	g.P("//查询")
	g.P("func (ar *" + info.serviceName + "Repo) Get" + info.serviceNameC + "(ctx context.Context, id int64) (*biz." + info.serviceNameC + ", error) {")
	g.P("p, err := ar.data.db." + info.serviceNameC + ".Get(ctx, id)")
	g.P("if err != nil {return nil, err}")
	g.P("return ent" + info.serviceNameC + "(*p).BizStruct(), nil")
	g.P("}")
	g.P("")
	g.P("//新增")
	g.P("func (ar *" + info.serviceName + "Repo) Create" + info.serviceNameC + "(ctx context.Context, m *biz." + info.serviceNameC + ") (*biz." + info.serviceNameC + ", error) {")
	g.P("modCreate := ar.data.db." + info.serviceNameC + ".Create()")
	g.P("p, err := modCreate.")
	if objmsg != nil {
		for _, f := range objmsg.Fields {
			name := MarshalStr(string(f.Desc.Name()))
			if name != "Id" && name != "UpdateTime" && name != "CreateTime" {
				g.P("Set" + name + "(m." + name + ").")
			}

		}
	}
	g.P("Save(ctx)")
	g.P("if err != nil  { return nil, err} \n return ent" + info.serviceNameC + "(*p).BizStruct(), err  }")
	g.P("")
	g.P("//更新")
	g.P("func (ar *" + info.serviceName + "Repo) Update" + info.serviceNameC + "(ctx context.Context, id int64, m *biz." + info.serviceNameC + ") (*biz." + info.serviceNameC + ", error){")
	g.P("p, err := ar.data.db." + info.serviceNameC + ".Get(ctx, id)")
	g.P("if err != nil  { return nil, err} ")
	g.P("modUp := ar.data.db." + info.serviceNameC + ".Update()")
	if objmsg != nil {
		for _, f := range objmsg.Fields {
			name := MarshalStr(string(f.Desc.Name()))
			if name != "Id" && name != "UpdateTime" && name != "CreateTime" {
				t := FieldType(f.Desc)
				if t == "string" {
					g.P("if m." + name + " != \"\" && m." + name + " != p." + name + " {\n modUp.Set" + name + "(m." + name + ")\n}")
				} else if t == "bool" {
					g.P("if m." + name + " != p." + name + " {\n modUp.Set" + name + "(m." + name + ")\n}")
				} else if t == "time.Time" {
					g.P("if !m." + name + ".IsZero() && m." + name + " != p." + name + " {\n modUp.Set" + name + "(m." + name + ")\n}")
				}
			}
		}
	}
	g.P("_, err = modUp.Save(ctx)")
	g.P("return ent" + info.serviceNameC + "(*p).BizStruct(), err")
	g.P("}")
	g.P("")
	g.P("//删除")
	g.P("func (ar *" + info.serviceName + "Repo) Delete" + info.serviceNameC + "(ctx context.Context, ids []int64) error {")
	g.P("tx, err := ar.data.db.Tx(ctx)")
	g.P("if err != nil {return err}")
	g.P("for _, id := range ids {")
	g.P("err = tx." + info.serviceNameC + ".DeleteOneID(id).Exec(ctx)")
	g.P("if err != nil {\n_ = tx.Rollback()\nreturn err\n}")
	g.P("}\nreturn tx.Commit()")
	g.P("}")
	g.P("")
	g.P("// 列表搜索")
	g.P("func (ar *" + info.serviceName + "Repo) List" + info.serviceNameC + "(ctx context.Context, op *pb.List" + info.serviceNameC + "Request) ([]*biz." + info.serviceNameC + ", error) {")
	g.P("query := ar.data.db." + info.serviceNameC + ".Query()")
	g.P("// todo 搜索条件")
	g.P("ps, err := query.Order(ent.Desc(\"id\")).All(ctx)")
	g.P("if err != nil {return nil, err}")
	g.P("rv := make([]*biz." + info.serviceNameC + ", 0)")
	g.P("for _, p := range ps {\nrv = append(rv, ent" + info.serviceNameC + "(*p).BizStruct())\n}")
	g.P("return rv, nil")
	g.P("}")
}

func generateEntSchema(gen *protogen.Plugin, file *protogen.File, s *protogen.Service, info *BaseInfo) {
	var objmsg *protogen.Message
	for _, msg := range file.Messages {
		if string(msg.Desc.Name()) == info.serviceNameC {
			objmsg = msg
		}
	}
	// biz
	g := gen.NewGeneratedFile("/data/ent/schema/"+info.filename, file.GoImportPath)
	g.P("package schema")
	g.P("import (\"entgo.io/ent\"\n\"entgo.io/ent/schema/field\"\n\"entgo.io/ent/schema/mixin\")")
	g.P("type " + info.serviceNameC + " struct {ent.Schema}")
	g.P("// todo 请检查数据格式，并且去除必要参数的Optional方法")
	g.P("func (" + info.serviceNameC + ") Fields() []ent.Field { return []ent.Field{")
	if objmsg != nil {
		for _, f := range objmsg.Fields {
			t := Capitalize(FieldType(f.Desc))
			if t == "Time.Time" {
				t = "Time"
			}
			if string(f.Desc.Name()) != "update_time" && string(f.Desc.Name()) != "create_time" {
				if string(f.Desc.Name()) != "id" {
					g.P("field." + t + "(\"" + string(f.Desc.Name()) + "\").Optional().Comment(\"" + commentDesc(f.Comments) + "\"),")
				} else {
					g.P("field." + t + "(\"" + string(f.Desc.Name()) + "\").Comment(\"" + commentDesc(f.Comments) + "\"),")
				}
			}
		}
	}
	g.P("}}")
	g.P("")
	g.P("func (" + info.serviceNameC + ") Mixin() []ent.Mixin { return []ent.Mixin{mixin.Time{}}}")
}

func generateBiz(gen *protogen.Plugin, file *protogen.File, s *protogen.Service, info *BaseInfo) {
	var objmsg *protogen.Message
	for _, msg := range file.Messages {
		if string(msg.Desc.Name()) == info.serviceNameC {
			objmsg = msg
		}
	}
	// biz
	g := gen.NewGeneratedFile("/biz/"+info.filename, file.GoImportPath)
	g.P("package biz")
	g.P("import (")
	g.P("\"context\"")
	g.P("\"time\"")
	g.P("\"github.com/go-kratos/kratos/v2/log\"")
	g.P(fmt.Sprintf("pb %s", file.GoImportPath.String()))
	g.P(")")
	g.P("")
	g.P("type " + info.serviceNameC + " struct {")
	if objmsg != nil {
		for _, f := range objmsg.Fields {
			name := MarshalStr(string(f.Desc.Name()))
			g.P(fmt.Sprintf("%s %s", name, FieldType(f.Desc)))
		}
	}
	g.P("}")
	g.P("")
	g.P("type " + info.serviceNameC + "Repo interface{")
	g.P("List" + info.serviceNameC + "(ctx context.Context, op *pb.List" + info.serviceNameC + "Request) ([]*" + info.serviceNameC + ", error)")
	g.P("Get" + info.serviceNameC + "(ctx context.Context, id int64) (*" + info.serviceNameC + ", error)")
	g.P("Create" + info.serviceNameC + "(ctx context.Context, mod *" + info.serviceNameC + ") (*" + info.serviceNameC + ", error)")
	g.P("Update" + info.serviceNameC + "(ctx context.Context, id int64, mod *" + info.serviceNameC + ") (*" + info.serviceNameC + ", error)")
	g.P("Delete" + info.serviceNameC + "(ctx context.Context, ids []int64) error")
	g.P("}")
	g.P("type " + info.serviceNameC + "Usecase struct {")
	g.P("repo " + info.serviceNameC + "Repo")
	g.P("log  *log.Helper")
	g.P("}")
	g.P("func New" + info.serviceNameC + "Usecase(repo " + info.serviceNameC + "Repo, logger log.Logger) *" + info.serviceNameC + "Usecase {")
	g.P("return &" + info.serviceNameC + "Usecase{repo: repo, log: log.NewHelper(\"biz/" + info.serviceName + "\", logger)}")
	g.P("}")
	g.P("")
	g.P("//查询")
	g.P("func (uc *" + info.serviceNameC + "Usecase) Get(ctx context.Context, id int64) (p *" + info.serviceNameC + ", err error) {")
	g.P("return uc.repo.Get" + info.serviceNameC + "(ctx, id)")
	g.P("}")
	g.P("//列表")
	g.P("func (uc *" + info.serviceNameC + "Usecase) List(ctx context.Context,  op *pb.List" + info.serviceNameC + "Request) ([]*" + info.serviceNameC + ",  error) {")
	g.P("return uc.repo.List" + info.serviceNameC + "(ctx, op)")
	g.P("}")
	g.P("//创建")
	g.P("func (uc *" + info.serviceNameC + "Usecase) Create(ctx context.Context, mod *" + info.serviceNameC + ") (p *" + info.serviceNameC + ", err error) {")
	g.P("return uc.repo.Create" + info.serviceNameC + "(ctx, mod)")
	g.P("}")
	g.P("//创建")
	g.P("func (uc *" + info.serviceNameC + "Usecase) Update(ctx context.Context, mod *" + info.serviceNameC + ") (p *" + info.serviceNameC + ", err error) {")
	g.P("return uc.repo.Update" + info.serviceNameC + "(ctx, mod.Id, mod)")
	g.P("}")
	g.P("//删除")
	g.P("func (uc *" + info.serviceNameC + "Usecase) Delete(ctx context.Context, ids []int64) error {")
	g.P("return uc.repo.Delete" + info.serviceNameC + "(ctx, ids)")
	g.P("}")
}

func generateService(gen *protogen.Plugin, file *protogen.File, s *protogen.Service, info *BaseInfo) {
	var objmsg *protogen.Message
	for _, msg := range file.Messages {
		if string(msg.Desc.Name()) == info.serviceNameC {
			objmsg = msg
		}
	}
	// service
	g := gen.NewGeneratedFile("/service/"+info.filename, file.GoImportPath)
	g.P("package service")
	g.P("import (")
	g.P("\"context\"")
	g.P(fmt.Sprintf("pb %s", file.GoImportPath.String()))
	g.P(info.ImportPath("biz"))
	g.P("\"google.golang.org/protobuf/types/known/timestamppb\"")
	g.P(")")
	g.P("")
	g.P("func New" + info.serviceNameC + "Service(bz *biz." + info.serviceNameC + "Usecase) pb." + info.serviceNameC + "ServiceServer {")
	g.P("return &" + info.serviceNameC + "Service{")
	g.P(info.serviceName + ": bz,")
	g.P("}")
	g.P("}")
	g.P("type " + info.serviceNameC + "Service struct {")
	g.P("pb.Unimplemented" + info.serviceNameC + "ServiceServer")
	g.P("" + info.serviceName + " *biz." + info.serviceNameC + "Usecase")
	g.P("}")
	for _, m := range s.Methods {
		g.P("//" + commentDesc(m.Comments))
		g.P(fmt.Sprintf("func (s *%sService)%s(ctx context.Context, req *pb.%s)(*pb.%s, error){",
			info.serviceNameC, m.Desc.Name(), m.Input.GoIdent.GoName, m.Output.GoIdent.GoName))
		ms := strings.ReplaceAll(string(m.Desc.Name()), info.serviceNameC, "")
		switch ms {
		case "Create", "Update":
			g.P("m, err := s." + info.serviceName + "." + ms + "(ctx, &biz." + info.serviceNameC + "{")
			if objmsg != nil {
				for _, f := range objmsg.Fields {
					name := MarshalStr(string(f.Desc.Name()))
					if name != "UpdateTime" && name != "CreateTime" {
						if ms != "Create" || name != "Id" {
							if f.Desc.Enum() != nil {
								g.P(fmt.Sprintf("%s:uint32(req.%s),", name, name))
							} else {
								g.P(fmt.Sprintf("%s:req.%s,", name, name))
							}
						}
					}
				}
			}
			g.P("})")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("return &pb." + ms + info.serviceNameC + "Reply{Data: s.protoMsg(m)}, nil")
		case "Delete":
			g.P("err := s." + info.serviceName + "." + ms + "(ctx, req.Ids)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("return &pb." + ms + info.serviceNameC + "Reply{}, nil")
		case "Get":
			g.P("m, err := s." + info.serviceName + "." + ms + "(ctx, req.Id)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("return s.protoMsg(m), nil")
		case "List":
			g.P("ms, err := s." + info.serviceName + "." + ms + "(ctx, req)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("var resp = &pb.List" + info.serviceNameC + "Reply{List: make([]*pb." + info.serviceNameC + ", 0, len(ms))}")
			g.P("for _, v := range ms {")
			g.P("resp.List = append(resp.List, s.protoMsg(v))")
			g.P("}")
			g.P("return resp, nil")
		default:
			g.P("panic(\"implement me\")")
		}
		g.P("}")
		g.P("")
	}
	g.P("//biz转换pb")
	g.P("func (uc *" + info.serviceNameC + "Service) protoMsg(m *biz." + info.serviceNameC + ") *pb." + info.serviceNameC + "{")
	g.P("msg := &pb." + info.serviceNameC + "{")
	if objmsg != nil {
		for _, f := range objmsg.Fields {
			name := MarshalStr(string(f.Desc.Name()))
			if f.Desc.Message() == nil || string(f.Desc.Message().Name()) != "Timestamp" {
				if f.Desc.Enum() != nil {
					g.P(fmt.Sprintf("%s:pb.%s(m.%s),", name, name, name))
				} else {
					g.P(fmt.Sprintf("%s:m.%s,", name, name))
				}
			}
		}
	}
	g.P("}")
	g.P("if !m.CreateTime.IsZero() {")
	g.P("msg.CreateTime = timestamppb.New(m.CreateTime)")
	g.P("}")
	g.P("if !m.UpdateTime.IsZero() {")
	g.P("msg.UpdateTime = timestamppb.New(m.UpdateTime)")
	g.P("}")
	g.P("return msg")
	g.P("}")
}

func KindToType(k protoreflect.Kind) string {
	switch k {
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		return "int32"
	case protoreflect.Uint32Kind:
		return "uint32"
	case protoreflect.Uint64Kind:
		return "uint64"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		return "int64"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.EnumKind:
		return "uint32"
	case protoreflect.StringKind:
		return "string"
	default:
		return "interface{}"
	}
}

func FieldType(desc protoreflect.FieldDescriptor) string {
	if desc.IsMap() {
		return fmt.Sprintf("map[%s][%s]", KindToType(desc.MapKey().Kind()), KindToType(desc.MapValue().Kind()))
	}
	if desc.IsList() {
		return fmt.Sprintf("[]%s", KindToType(desc.Kind()))
	}
	if desc.Kind() == protoreflect.MessageKind {
		if string(desc.Message().Name()) == "Timestamp" {
			return "time.Time"
		}
		return string(desc.Message().Parent().Name()) + string(desc.Message().Name())
	}
	return KindToType(desc.Kind())
}
