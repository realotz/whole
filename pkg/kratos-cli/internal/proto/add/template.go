package add

import (
	"bytes"
	"strings"
	"text/template"
)

const protoTemplate = `
syntax = "proto3";

package {{.Package}};
import "google/api/annotations.proto";
import "google/protobuf/timestamp.proto";
option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";

// {{.Service}}服务
service {{.Service}}Service {
	// 创建{{.Service}}
	rpc Create{{.Service}} ({{.Service}}) returns (Create{{.Service}}Reply);
	// 更新{{.Service}}
	rpc Update{{.Service}} ({{.Service}}) returns (Update{{.Service}}Reply);
	// 批量删除{{.Service}}
	rpc Delete{{.Service}} (Delete{{.Service}}Request) returns (Delete{{.Service}}Reply);
	// 获取{{.Service}}详情
	rpc Get{{.Service}} (Get{{.Service}}Request) returns ({{.Service}});
	// 查询{{.Service}}列表
	rpc List{{.Service}} (List{{.Service}}Request) returns (List{{.Service}}Reply);
}

//创建请求
message Create{{.Service}}Reply {
	{{.Service}} data=1;
}

//更新请求
message Update{{.Service}}Reply {
	{{.Service}} data=1;
}

//批量删除请求
message Delete{{.Service}}Request {
	repeated int64 ids = 1;
}

//删除结果
message Delete{{.Service}}Reply {}

//查询单个数据
message Get{{.Service}}Request {
	int64 id = 1;
}

//列表查询条件
message List{{.Service}}Request {

}

//列表查询返回
message List{{.Service}}Reply {
	repeated {{.Service}} list=1;
	//数据总计
	int64 total=2;
	//页码
	uint32 page=3;
	//分页大小
	uint32 page_size=4;
}

// {{.Service}}结构体(请在此处定义数据结构)
message {{.Service}}{
	int64  id = 1;
	google.protobuf.Timestamp update_time=22;
	google.protobuf.Timestamp create_time=23;
}
`

func (p *Proto) execute() ([]byte, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("proto").Parse(strings.TrimSpace(protoTemplate))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
