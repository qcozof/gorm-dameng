/**
 * @Description 达梦数据库 Clob (*dm.DmClob)数据类型处理
 * @Author $
 * @Date $ $
 **/
package customdbtype

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"gitee.com/chunanyong/dm"
)

type MyClob string

// 写入数据库之前，对数据做类型转换
func (clob MyClob) Value() (driver.Value, error) {
	if len(clob) == 0 {
		return nil, nil
	}
	return string(clob), nil
}

// 将数据库中取出的数据，赋值给目标类型
func (clob *MyClob) Scan(v interface{}) error {
	switch v.(type) {
	case *dm.DmClob:
		tmp := v.(*dm.DmClob)
		le, err := tmp.GetLength()
		if err != nil {
			return errors.New(fmt.Sprint("err：", err))
		}

		str, err := tmp.ReadString(1, int(le))
		*clob = MyClob(str)
		break

	//非clob，当成字符串，兼容oracle
	default:
		*clob = MyClob(v.(string))
	}
	return nil
}
