package action

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func Analyzd_excel() {
	excel_file := "/Users/mengxiaoce/sgc/go_test/go_practise/action/data.xlsx"
	sheet1_txt := "/Users/mengxiaoce/sgc/go_test/go_practise/action/sheet1_output.txt"
	sheet2_txt := "/Users/mengxiaoce/sgc/go_test/go_practise/action/sheet2_output.txt"
	f, err := excelize.OpenFile(excel_file)
	if err != nil {
		log.Fatalf("open %s failed, err :%v", excel_file, err)

	}
	defer f.Close()
	// 处理sheet1
	if err := writeSheet1ToFile(f, "Sheet1", sheet1_txt); err != nil {
		log.Fatalf("write sheet1 failed, err: %v", err)
	}
	// 处理sheet2
	if err := writeSheet2ToFile(f, "Sheet2", sheet2_txt); err != nil {
		log.Fatalf("write sheet2 failed, err: %v", err)
	}
	fmt.Printf("处理完成\n")
}

func writeSheet1ToFile(f *excelize.File, sheelName string, outputFile string) error {
	rows, err := f.GetRows(sheelName)
	if err != nil {
		return fmt.Errorf("读取%s出错: %v", sheelName, err)
	}
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建%s出错: %v", sheelName, err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	if len(rows) > 1 {
		for i := 0; i < 2 && i < len(rows); i++ {
			if _, err := writer.WriteString(fmt.Sprintf("%s\n", rows[i][0])); err != nil {
				return fmt.Errorf("写入注释失败: %v", err)
			}
		}
	}
	mapped := map[string]string{
		"域名":    "domain",
		"ip":    "IP",
		"is_ha": "is_ha",
		"管理网段":  "NetSegment",
	}
	for i, r := range rows {
		if i < 2 || len(r) < 2 {
			continue
		}
		mappedKey, exist := mapped[r[0]]
		if !exist {
			mappedKey = r[0]
		}
		if _, err := writer.WriteString(fmt.Sprintf("%s=%s\n", mappedKey, r[1])); err != nil {
			return fmt.Errorf("写入错误%w", err)
		}

	}
	return nil
}

func writeSheet2ToFile(f *excelize.File, sheelName string, outputFile string) error {
	rows, err := f.GetRows(sheelName)
	if err != nil {
		return fmt.Errorf("get %s sheet failed,err: %v", sheelName, err)
	}
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("create %s failed, err: %v", outputFile, err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString(fmt.Sprintln("#gcsun"))
	for i, r := range rows {
		if i < 2 && len(r) < 3 {
			continue
		}
		if _, err := writer.WriteString(fmt.Sprintf("%s\t%s\t%s\n", r[0], r[1], r[2])); err != nil {
			return fmt.Errorf("write file failed : %v", err)
		}
	}
	return nil
}
