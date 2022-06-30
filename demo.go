/**
 * @Description 操作oracle和dameng
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

	fmt.Println(fmt.Printf("连接%s...", dbType))
	gormConfig := config(true)

	if dbType == "dameng" {
		GORM_DB, err = gorm.Open(dameng.Open("dm://sysdba:dameng123456!@192.168.1.200:5236?autoCommit=true"), gormConfig)
	} else if dbType == "oracle" {
		GORM_DB, err = gorm.Open(oracle.Open("DBUSER1/ora123456@192.168.1.200:1521/ORCL"), gormConfig)
		//GORM_DB, err := gorm.Open(dameng.Open( "dm://sysdba:dameng123456!@192.168.1.200:5236?autoCommit=true&ignoreCase=true&columnNameUpperCase=true&compatibleMode=oracle"), gormConfig)
	} else {
		fmt.Println("dbType不正确")
		os.Exit(0)
	}

	if err != nil {
		fmt.Println("连接失败：", err)
		os.Exit(0)
	}
	fmt.Println("连接成功：", GORM_DB)

	/*	err = GORM_DB.Exec("update  DW.COMPANY_INFO set MANAGE_RANGE='吸收公众存款;发放短期、中期和长期贷款;办理国内结算;办理票据承兑与贴现;代理发行、代理兑付、承销政府债券;买卖政府债券、金融债券;从事同业拆借;办理借记卡、贷记卡业务;代理收付款项及代理保险业务;办理外汇存款、外汇贷款、外汇汇款、外币兑换、国际结算、同业外汇存放、拆借业务及资信调查、咨询、见证业务;经银行业监督管理机构批准的其他业务。(依法须经批准的项目,经相关部门批准后方可开展经营活动)' WHERE COMPANY_ID='C3DDBD2F17554E8A838DB706C139B883' ").Error
		if err != nil{
			fmt.Errorf("err:",err)
		}*/

	/*	var companyInfo CompanyInfo
		err = GORM_DB.Raw("SELECT COMPANY_ID,CN_NAME,MANAGE_RANGE FROM DW.COMPANY_INFO WHERE COMPANY_ID='C3DDBD2F17554E8A838DB706C139B883' ").Scan(&companyInfo).Error //MANAGE_RANGE
		if err != nil{
			fmt.Errorf("err:",err)
		}


		fmt.Println("----------------------")
		fmt.Println(fmt.Printf("%v",companyInfo.ManageRange))
		fmt.Println("----------------------")

		tmp := 	companyInfo.ManageRange
		le,err := tmp.GetLength()
		manageRange,err := tmp.ReadString(1,int(le))

		fmt.Println("str:",manageRange)*/

	//-----------------------------------------
	var companyList []CompanyInfo
	//err = GORM_DB.Raw("SELECT COMPANY_ID,CN_NAME,MANAGE_RANGE FROM DW.COMPANY_INFO WHERE COMPANY_ID in('C3DDBD2F17554E8A838DB706C139B883') ").Scan(&companyList).Error //MANAGE_RANGE ,'4A5A6EA9B47445D48CB30683BEE68C4A'
	//err = GORM_DB.Raw("select t.ID,t.Title,t.Content,t.rowid from dw.table1 t where id=1 ").Scan(&companyList).Error //MANAGE_RANGE ,'4A5A6EA9B47445D48CB30683BEE68C4A'

	// 读取
	GORM_DB.Select("COMPANY_ID,CN_NAME").Where("COMPANY_ID = ?", "C3DDBD2F17554E8A838DB706C139B888").Find(&companyList) // 查询id为1的product
	//GORM_DB.First(&companyList, "COMPANY_ID = ?", "C3DDBD2F17554E8A838DB706C139B883").Select("ID") // 查询code为l1212的product
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

	fmt.Println(companyList)
}

//column名与增删改查的字段名大小写要保持一致，否则取不到
/*type CompanyInfo struct {
	Id      string              `gorm:"column:ID"`
	Title   string              `gorm:"column:TITLE"`
	Content customdbtype.MyClob `gorm:"column:CONTENT"`
}*/

func (CompanyInfo) TableName() string {
	return "dd.COMPANY_INFO"
}

type CompanyInfo struct {
	CompanyId   string              `gorm:"column:COMPANY_ID"`
	CnName      string              `gorm:"column:CN_NAME"`
	ManageRange customdbtype.MyClob `gorm:"column:MANAGE_RANGE"`
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
				TablePrefix:   "DD.",
				SingularTable: true, //表名后面不加s
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "DD.",
				SingularTable: true, //表名后面不加s
			},
			//Namer.ColumnName : func() {
			//
			//},
		}
	}
	return
}
