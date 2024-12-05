# Aris-url-gen

[ [English](README.md) | 简体中文 ]

## 介绍

一个高性能的短链接生成服务，基于Go语言开发。项目名称来源于游戏《碧蓝档案》中的角色Aris，如下图所示。

---

<p align="center">
  <img src="https://raw.githubusercontent.com/hcd233/Aris-AI/master/assets/110531412.jpg" width="50%">
  <br>Aris: Blue Archive 中的角色
</p>

---

## 功能特性

- 生成短链接：将长URL转换为短链接
- 支持自定义过期时间
- 双向缓存：使用Redis实现高性能缓存
- RESTful API：提供标准的HTTP接口
- 数据持久化：使用MySQL存储URL映射关系

## 技术栈

- **Web框架**: [Fiber](https://github.com/gofiber/fiber)
- **ORM**: [GORM](https://gorm.io)
- **缓存**: Redis
- **数据库**: MySQL
- **日志**: [Zap](https://github.com/uber-go/zap)

## API接口

### 1. 生成短链接

```http
POST /v1/shortURL
Content-Type: application/json

{
    "originalURL": "https://example.com/very/long/url",
    "expireDays": 7  // 可选，过期天数
}
```

### 2. 访问短链接

```http
GET /v1/shortURL/{shortURL}
```

## 项目结构

```
.
├── cmd/                # 命令行入口
├── internal/          
│   ├── api/           # API相关代码
│   │   ├── dao/       # 数据访问层
│   │   ├── dto/       # 数据传输对象
│   │   ├── handler/   # 请求处理器
│   │   └── service/   # 业务逻辑层
│   ├── config/        # 配置
│   ├── logger/        # 日志
│   ├── resource/      # 资源
│   └── util/          # 工具函数
└── main.go            # 主入口
```

## 安装部署

### 前置要求

- Go 1.20+
- MySQL 8.0+
- Redis 6.0+

### 本地开发

1. 克隆仓库

```bash
git clone https://github.com/hcd233/Aris-url-gen.git
cd Aris-url-gen
```

2. 安装依赖

```bash
go mod download
```

3. 配置环境变量

参考 `api.env.template` 配置相关环境变量

4. 运行服务

```bash
go run main.go
```

## 许可证

本项目采用 Apache License 2.0 许可证。详见 [LICENSE](LICENSE) 文件。
