/**
 * @Author: Sun
 * @Company: 苏州天和汇工业
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/10/11 15:49
 */

package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func ReadExcelFile() [][]string {
	f, err := excelize.OpenFile("./系统切换方案20221010-附件.xlsx")
	if err != nil {
		fmt.Println("打开表格失败，err:", err)
	}

	rows, err := f.GetRows("sheet1") // 读取表格的所有行
	if err != nil {
		fmt.Println("读取sheet1 err:", err)
	}

	return rows[1:] // 表格内容从第二行开始
}
