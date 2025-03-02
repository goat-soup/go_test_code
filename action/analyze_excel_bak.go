package action

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func Analyzd_excel1() {

	excelFile := "/Users/mengxiaoce/sgc/go_test/go_practise/action/data.xlsx" // 替换为你的 Excel 文件路径

	sheet1File := "/Users/mengxiaoce/sgc/go_test/go_practise/action/sheet1_output.txt"
	sheet2File := "/Users/mengxiaoce/sgc/go_test/go_practise/action/sheet2_output.txt"

	// 打开Excel文件
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer f.Close()

	// 处理Sheet1
	if err := writeSheet1ToFile1(f, "Sheet1", sheet1File); err != nil {
		log.Fatalf("写入 Sheet1 文件失败: %v", err)
	}

	// 处理Sheet2
	if err := writeSheet2ToFile1(f, "Sheet2", sheet2File); err != nil {
		log.Fatalf("写入 Sheet2 文件失败: %v", err)
	}

	fmt.Println("处理完成")
}

func writeSheet1ToFile1(f *excelize.File, sheetName, outputFile string) error {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("读取 sheet1 出错: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建 sheet1 文件出错: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 写入说明信息
	if len(rows) > 1 {
		for i := 0; i < 2 && i < len(rows); i++ {
			if _, err := writer.WriteString(fmt.Sprintf("%s\n", rows[i][0])); err != nil {
				return fmt.Errorf("写入 sheet1 说明信息出错: %w", err)
			}
		}
	}

	// 定义Key映射表
	keyMap := map[string]string{
		"域名":    "Domain",
		"ip":    "IP",
		"is_ha": "IsHA",
		"管理网段":  "NetSegment",
	}

	// 遍历每一行并写入内容
	for i, row := range rows {
		if i < 2 || len(row) < 2 {
			continue // 跳过说明行和列数不足的行
		}
		key := row[0]
		mappedKey, exists := keyMap[key]
		if !exists {
			mappedKey = key // 使用原始key作为默认值
		}
		value := row[1]
		if _, err := writer.WriteString(fmt.Sprintf("%s=%s\n", mappedKey, value)); err != nil {
			return fmt.Errorf("写入 sheet1 文件出错: %w", err)
		}
	}
	return nil
}

func writeSheet2ToFile1(f *excelize.File, sheetName, outputFile string) error {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("读取 sheet2 出错: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建 sheet2 文件出错: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 写入说明信息
	if len(rows) > 1 {
		for i := 0; i < 2 && i < len(rows); i++ {
			if _, err := writer.WriteString(fmt.Sprintf("%s\n", rows[i][0])); err != nil {
				return fmt.Errorf("写入 sheet2 说明信息出错: %w", err)
			}
		}
	}

	// 遍历每一行并写入内容
	for i, row := range rows {
		if i < 2 || len(row) < 3 {
			continue // 跳过说明行和列数不足的行
		}
		name := row[0]
		ip := row[1]
		password := row[2]
		if _, err := writer.WriteString(fmt.Sprintf("%s\t%s\t%s\n", name, ip, password)); err != nil {
			return fmt.Errorf("写入 sheet2 文件出错: %w", err)
		}
	}
	return nil
}
