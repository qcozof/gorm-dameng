/**
 * @Description 达梦数据库 Clob (*dm.DmClob)数据类型处理
 * @Author $
 * @Date $ $
 **/
package customdbtype

import (
	"database/sql/driver"
	"dm"
	"errors"
	"fmt"
)

type MyDmClob string

// 写入数据库之前，对数据做类型转换
func (g MyDmClob) Value() (driver.Value, error) {
	if len(g) == 0 {
		return nil, nil
	}
	return g, nil
}

// 将数据库中取出的数据，赋值给目标类型
func (g *MyDmClob) Scan(v interface{}) error {
	switch v.(type) {
	case *dm.DmClob:
		tmp := v.(*dm.DmClob)
		le, err := tmp.GetLength()
		if err != nil {
			return errors.New(fmt.Sprint("err2", err))
		}

		str, err := tmp.ReadString(1, int(le))
		*g = MyDmClob(str)
		break

		//非clob，用当成字符串
	default:
		*g = MyDmClob(v.(string))
	}
	return nil
}
