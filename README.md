# goframework-gorm-mysql

包装`gorm-mysql`对象，实现[`godb`](https://github.com/kordar/godb)接口。

## 安装
```go
go get github.com/kordar/goframework-gorm-mysql@latest
```

## 使用

- 初始化 slog（可选，默认使用 slog.Default）

```go
import (
	"log/slog"
	"os"
)

func init() {
	// JSON 输出
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	// 或文本输出
	// slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}
```

- 设置 GORM 日志等级（error | warn | info）

```go
import (
	goframeworkgormmysql "github.com/kordar/goframework-gorm-mysql"
)

logLevel := "info"  // error warn info
goframeworkgormmysql.SetDbLogLevel(logLevel)
```

- 添加实例

```go
// 1、通过配置添加
import (
	"log/slog"
	goframeworkgormmysql "github.com/kordar/goframework-gorm-mysql"
)

key := "default"
section := map[string]string{
	"host":     "127.0.0.1",
	"port":     "3306",
	"user":     "root",
	"password": "123456",
	"db":       "test",
	"charset":  "utf8",
}
if err := goframeworkgormmysql.AddMysqlInstance(key, section); err != nil {
    slog.Error("add mysql instance failed", "err", err)
}

// 2、通过dsn添加, user:password@tcp(host:port)/db?params
goframeworkgormmysql.AddMysqlInstanceWithDsn(key, dsn)
```

## 说明
- 已移除 gologger，GORM 日志通过标准库 slog 输出，需要 Go 1.21+
- 通过 SetDbLogLevel 控制 GORM SQL 日志级别；应用层可自由配置 slog 的 Handler/Level/Formatter

