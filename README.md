# gorm-dameng
**达梦8、达梦接入 gorm**  

**使用方法**  
引入：

``` go script
	"github.com/qcozof/gorm-dameng/customdbtype"
	"github.com/qcozof/gorm-dameng/dameng"
```



比如有个表TABLE1

```sql
CREATE TABLE "DW"."TABLE1"
(
"ID" INT IDENTITY(1, 1) NOT NULL,
"TITLE" VARCHAR2(50),
"CONTENT" CLOB,
NOT CLUSTER PRIMARY KEY("ID")) STORAGE(ON "MAIN", CLUSTERBTR) ;
```


数据库中的Clob类型在结构体中用 ***customdbtype.MyClob*** 表示而不能用 string，否则会抛异常。

```go script

type MemberInfo struct {
	Id      string                `gorm:"column:ID"`
	Title   string                `gorm:"column:TITLE"`
	Content customdbtype.MyClob   `gorm:"column:CONTENT"`
}

```



**注意大小写问题**

- 经测试，查询时，要查询的字段名与结构体***MemberInfo***中column后面定义的字段名或与变量名大小写保持一致，否则取不到值。oracle则不存在这个问题。

```go script
    var memberInfo MemberInfo

    //正确。与变量名大小写保持一致
	err = GORM_DB.Raw("select Id,Title,Content from dw.table1 t where id=1 ").Scan(&memberInfo).Error

    //正确。与gorm:"column:字段名"中的字段名大小写保持一致
    err = GORM_DB.Raw("select ID,TITLE,CONTENT from dw.table1 t where id=1 ").Scan(&memberInfo).Error
    //where后面的字段，如id又可以不用保持大小写一致

    //错误。没有与变量名或与gorm:"column:字段名"中的大小写保持一致
	err = GORM_DB.Raw("select id,titlE,ContenT from dw.table1 t where id=1 ").Scan(&memberInfo).Error
```



- 插入、更新时字段不分大小写

```go
	err = GORM_DB.Exec("insert into  dw.table1(tiTle,conTENt)values('标题3','内容3') ").Error
	err = GORM_DB.Exec("update  dw.table1 set tiTLE='标题333',content='内容333' where id=3").Error
```



更多，参考demo.go  