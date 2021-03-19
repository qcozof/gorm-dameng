# gorm-dameng
达梦 gorm驱动  
达梦8

***使用方法***  
引入：
``` go script
	"github.com/qcozof/gorm-dameng/customdbtype"
	"github.com/qcozof/gorm-dameng/dameng"
```

***注意***  	
数据库中的Clob类型在结构体中使用 customdbtype.MyDmClob 表示而不能用 string,否则会报错。
```go script

type MemberInfo struct {
	Id      string                `gorm:"column:ID"`
	Title   string                `gorm:"column:TITLE"`
	Content customdbtype.MyDmClob `gorm:"column:CONTENT"`
}

```

增删改查时字段名与结构体MemberInfo中定义的大小写要保持一致，否则取不到值。oracle则不存在这个问题。
```go script
    ...
	var memberInfo MemberInfo
	err = GORM_DB.Raw("select Id,Title,Content from dw.table1 t where id=1 ").Scan(&memberInfo).Error
```

具体参考demo.go  