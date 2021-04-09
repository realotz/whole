适用与kratos-go的protobuf快速生成service biz data ent 增删查改代码
ps proto格式有符合定义要求，例 test.proto 中有Test service 与 Test message 会根据Test message中的结构生成代码
```golang
go get -u github.com/realotz/whole/pkg/protoc-gen-kratos-server
```
```bigquery
使用本项目下kratos-cli生成 请指定到具体service目录上级既 biz data service的父级
kratos-cli proto server api/test/v1/test.proto /internal/services/test/
```
生成
```bigquery
protoc --proto_path=.  --proto_path=./third_party --kratos-server_out=path=./internal/services/test/,paths=source_relative:./internal/services/test/ api/test/v1/test.proto
```