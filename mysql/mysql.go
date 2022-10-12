/**
 * @Author: Sun
 * @Company: 苏州天和汇工业
 * @Description:
 * @File:  mysql
 * @Version: 1.0.0
 * @Date: 2022/10/11 16:30
 */

package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
	"xlsx2mysql/service"
)

func Connect() {
	dns := "root:password@tcp(127.0.0.1:3306)/fms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，err:", err)
		return
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	defer sqlDB.Close()
	rows, _ := db.Table("t_procurement_order").Rows()
	tmpColumns, _ := rows.Columns() // 获取表的所有字段
	columns := tmpColumns[1:14]     // 获取和表格对应的字段
	tt := strings.Join(columns, ",")
	excelRows := service.ReadExcelFile()

	for _, excelRow := range excelRows {
		tmpValues := []string{}
		for _, cell := range excelRow {
			cell = "'" + cell + "'" // 给表格的每个字段加上 ''
			tmpValues = append(tmpValues, cell)
		}

		tmpTime := time.Now().Format("2006-01-02 15:04:05") // created_at 字段的值
		time := "'" + tmpTime + "'"
		status := "'" + "0" + "'" // t_procurement_order表中 status 的值
		tmpValues = append(tmpValues, status, time)
		values := strings.Join(tmpValues, ",")

		sql := fmt.Sprintf("insert into t_procurement_order(%s) values (%s)", tt, values)
		fmt.Println(sql)
		//fmt.Println(nm)
		res, err := sqlDB.Exec(sql)
		if err != nil {
			fmt.Println(res, err)
		}
	}
}
