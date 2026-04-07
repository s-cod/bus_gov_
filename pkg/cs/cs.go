package cs

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func getFinLen(s [][]string, r, c int) string {
	if r < 1 || c < 1 {
		panic("неверные параметры ячейки")
	}
	r -= 1
	c -= 1

	if len(s[r]) < c {
		return "0.00"
	}

	if s[r][c] == "" {
		return "0.00"
	}

	result := strings.ReplaceAll(s[r][c], ",", "")
	return result
}

func ProcessFile(filePath string) error {
	f, err := excelize.OpenFile(filePath) //, excelize.Options{RawCellValue: true})
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer f.Close()
	currentTime := time.Now().Format("2006-01-02T15:04:05")

	// year := currentTime[0:4]
	rows, err := f.GetRows("Результат")

	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	year := getFinLen(rows, 15, 16)[6:]

	outputFile := fmt.Sprintf("./out/%s.xml", strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath)))
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

	position := Position{
		PositionID:    uuid.New().String(),
		ChangeDate:    currentTime,
		VersionNumber: "1",
		FinancialYear: year,
		ConfirmDate:   currentTime[:10],
		PlanBudgetaryFunds: PlanBudgetaryFunds{
			CapitalRealAssets: "0.00",
			Total:             "0.00",
		},
	}

	for i := 28; i < 35; i++ {
		if getFinLen(rows, i, 1) == "0.00" {
			break
		}

		position.OtherGrantFunds = append(position.OtherGrantFunds, OtherGrantFunds{
			Name:  getFinLen(rows, i, 1),
			Funds: getFinLen(rows, i, 14),
			Code:  getFinLen(rows, i, 2),
			Kosgu: Kosgu{
				Code: getFinLen(rows, i, 3),
			},
		})
	}

	doc := ActionGrant{
		Xmlns:    "http://bus.gov.ru/external/1",
		XmlnsNs2: "http://bus.gov.ru/types/1",

		Header: Header{
			Xmlns:          "http://bus.gov.ru/types/1",
			ID:             uuid.New().String(),
			CreateDateTime: currentTime, //"2024-12-31T00:00:00",
		},
		Body: Body{
			Position: position,
		},
	}

	// Записываем заголовок XML вручную
	if _, err := file.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>` + "\n")); err != nil {
		return fmt.Errorf("не удалось записать заголовок XML: %w", err)
	}

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "	")

	if err := encoder.Encode(doc); err != nil {
		return fmt.Errorf("не удалось записать XML: %w", err)
	}

	fmt.Printf("Файл успешно сохранен: %s\n", outputFile)
	return nil
}
