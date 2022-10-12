/**
 * @Author: Sun
 * @Company: 苏州天和汇工业
 * @Description:
 * @File:  cmd
 * @Version: 1.0.0
 * @Date: 2022/10/11 15:48
 */

package cmd

import "xlsx2mysql/service"

func Run() {
	service.ReadExcelFile()
}
