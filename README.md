# goframework-gorm-mysql

包装`gorm-mysql`对象，实现[`godb`](https://github.com/kordar/godb)接口。

## 安装
```go
go get github.com/kordar/goframework-gorm-mysql v1.0.8
```

## 使用

- 设置日志等级

```go
logLevel := "info"  // error warn info
goframeworkgormmysql.SetDbLogLevel(logLevel)
```

- 添加实例

```go
// 1、通过配置添加
section := map[string]interface{}{
	"host": "127.0.0.1",
	"port": 3306,
	"user": "root",
	"password": "123456",
	"db": "test",
	"charset": "utf8"
}
if err := goframeworkgormmysql.AddMysqlInstance(key, section); err != nil {
    log.Errorf("初始化mysql异常，err=%v", err)
}

// 2、通过dsn添加, user:password@tcp(host:port)/db?params
goframeworkgormmysql.AddMysqlInstanceWithDsn(key, dsn)
```


