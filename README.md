# whole
whole是kratos-go 2.0版本的微服务脚手架
本仓库目前是初期开发阶段，仅供参考，不推荐应用生产环境
目标
* **单体服务**: 微服务可合并单一服务也可拆分启动，方便小项目初始阶段单机部署与开发。
* **通用业务功能**: 提供基础通用的基础业务功能服务，例: 用户，权限，cms，简易oa等
* **代码生成**: 扩展官方kratos cli 提供基于proto的增删查改代码，openapi3文档 前端ts api等代码生成，方便快速开发。
* **admin中台** 提供基于ts + react + antd-pro的业务功能中台前端。
* **微服务** 基于kubernetes+istio实现微服务架构与治理。后续对istio与kube的kind做统一维护。

## Install
```
go get -u https://github.com/realotz/whole/tree/master/pkg/protoc-gen-kratos-server@latest
go get -u https://github.com/realotz/whole/tree/master/pkg/kratos-cli@latest
```

