# GoMigrate
GoMigrate 是一个用来进行数据库管理的小工具，主要提供了NewMigrate、Migrate、Rollback、RefreshSchema的功能。
## 如何开始？
你可以直接使用编译好的gomigrate，提供的命令有:

* install  第一次你需要install已生成必要的配置文件
* new 新建一个migrate
* rollback 回滚
* refresh 刷新schema文件
* 默认为进行migrate

核心代码在src/migrate/migrate.go。你可以用通过`import "migrate"`来在代码中使用
