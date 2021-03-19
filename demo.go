/**
 * @Description 操作oracle和dameng
 * @Author $
 * @Date $ $
 **/
package main

import (
	"fmt"
	"github.com/qcozof/gorm-dameng/customdbtype"
	"log"
	"os"
	"time"

	"github.com/cengsin/oracle"
	"github.com/qcozof/gorm-dameng/dameng"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const dbType = "dameng"

func main() {

	var err error
	var GORM_DB *gorm.DB

	fmt.Println(fmt.Printf("连接%s...",dbType))
	gormConfig := config(true)

	if dbType == "dameng"{
		GORM_DB, err = gorm.Open(dameng.Open("dm://sysdba:dameng123!@193.100.100.221:5236?autoCommit=true"), gormConfig)
	}else if dbType == "oracle"{
		GORM_DB, err = gorm.Open(oracle.Open("ZTK/sirc1234@193.100.100.43:1521/ORCL"), gormConfig)
		//GORM_DB, err := gorm.Open(dameng.Open( "dm://sysdba:dameng123!@193.100.100.221:5236?autoCommit=true&ignoreCase=true&columnNameUpperCase=true&compatibleMode=oracle"), gormConfig)
	}else{
		fmt.Println("dbType不正确")
		os.Exit(0)
	}

	if err != nil {
		fmt.Println("连接失败：", err)
		os.Exit(0)
	}
	fmt.Println("连接成功：", GORM_DB)

	/*	err = GORM_DB.Exec("update  DW.MEMBER_INFO set MANAGERANGE='吸收公众存款;发放短期、中期和长期贷款;办理国内结算;办理票据承兑与贴现;代理发行、代理兑付、承销政府债券;买卖政府债券、金融债券;从事同业拆借;办理借记卡、贷记卡业务;代理收付款项及代理保险业务;办理外汇存款、外汇贷款、外汇汇款、外币兑换、国际结算、同业外汇存放、拆借业务及资信调查、咨询、见证业务;经银行业监督管理机构批准的其他业务。(依法须经批准的项目,经相关部门批准后方可开展经营活动)' WHERE MEMBER_ID='C3DDBD2F17554E8A838DB706C139B883' ").Error
		if err != nil{
			fmt.Errorf("err:",err)
		}*/

	/*	var memberInfo MemberInfo
		err = GORM_DB.Raw("SELECT MEMBER_ID,CNNAME,MANAGERANGE FROM DW.MEMBER_INFO WHERE MEMBER_ID='C3DDBD2F17554E8A838DB706C139B883' ").Scan(&memberInfo).Error //MANAGERANGE
		if err != nil{
			fmt.Errorf("err:",err)
		}


		fmt.Println("----------------------")
		fmt.Println(fmt.Printf("%v",memberInfo.ManageRange))
		fmt.Println("----------------------")

		tmp := 	memberInfo.ManageRange
		le,err := tmp.GetLength()
		manageRange,err := tmp.ReadString(1,int(le))

		fmt.Println("str:",manageRange)*/

	//-----------------------------------------
	var memberList MemberInfo
	//err = GORM_DB.Raw("SELECT MEMBER_ID,CNNAME,MANAGERANGE FROM DW.MEMBER_INFO WHERE MEMBER_ID in('C3DDBD2F17554E8A838DB706C139B883') ").Scan(&memberList).Error //MANAGERANGE ,'4A5A6EA9B47445D48CB30683BEE68C4A'
	//err = GORM_DB.Raw("select t.ID,t.Title,t.Content,t.rowid from dw.table1 t where id=1 ").Scan(&memberList).Error //MANAGERANGE ,'4A5A6EA9B47445D48CB30683BEE68C4A'

	// 读取
	GORM_DB.Select("MEMBER_ID,CNNAME").Where("MEMBER_ID = ?", "C3DDBD2F17554E8A838DB706C139B883").Find(&memberList) // 查询id为1的product
	//GORM_DB.First(&memberList, "MEMBER_ID = ?", "C3DDBD2F17554E8A838DB706C139B883").Select("ID") // 查询code为l1212的product
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}

/*	err = GORM_DB.Exec("update  dw.table1 set CONTENT=? where ID=2 ").Error
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}*/

/*	err = GORM_DB.Exec("insert into  dw.table1(tiTle,conTENt)values('标题6','内容6') ").Error
	//err = GORM_DB.Exec("update  dw.table1 set title='标题333',content='内容333' where id=3").Error
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}*/

	fmt.Println(memberList)
}

//column名与增删改查的字段名大小写要保持一致，否则取不到
/*type MemberInfo struct {
	Id      string              `gorm:"column:ID"`
	Title   string              `gorm:"column:TITLE"`
	Content customdbtype.MyClob `gorm:"column:CONTENT"`
}*/

func (MemberInfo) TableName() string {
	return "dw.MEMBER_INFO"
}
type MemberInfo struct {
	MemberId    string              `gorm:"column:MEMBER_ID"`
	CnName      string              `gorm:"column:CNNAME"`
	ManageRange customdbtype.MyClob `gorm:"column:MANAGERANGE"`
}

func config(logMode bool) (c *gorm.Config) {
	if logMode {
		c = &gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: 1 * time.Millisecond,
				LogLevel:      logger.Warn,
				Colorful:      true,
			}),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:"DW.",
				SingularTable: true, //表名后面不加s
			},

		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:"DW.",
				SingularTable: true, //表名后面不加s
			},
			//Namer.ColumnName : func() {
			//
			//},
		}
	}
	return
}