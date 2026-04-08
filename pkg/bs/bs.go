// Package bs формирование файлов для отчеты Бюджетные смета
package bs

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

func ProcessFile(filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows("Лист 2")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	currentDate := time.Now().Format("2006-01-02T15:04:05")
	// .Format("2006-01-02T15:04:05")
	oldYear := time.Now().AddDate(-1, 0, 0).Format("2006")

	position := Position{
		PositionID:    uuid.New().String(),
		ChangeDate:    currentDate,
		VersionNumber: "1",
		FinancialYear: currentDate[:4],
		ConfirmDate:   oldYear + "-12-31",
		Section:       "504",
	}

	start := 0
	for i, row := range rows { // Диапазон B85:S100

		if len(row) > 3 {
			if row[1] == "Наименование показателя" {
				start = i + 5
				break
			}
		}
	}

	for _, row := range rows[start:] { // Диапазон B85:S100
		if len(row) < 3 {
			break
		}
		if row[1] == "" {
			break
		}

		tmp := fmt.Sprintf("504%s%s%s%s", row[4], row[6], row[7], row[8])
		circumstance := BudgetaryCircumstance{
			KbkBudget: KbkBudget{
				Code: tmp,
				Name: row[1],
				Budget: Budget{
					Code: "52030407",
					Name: "Бюджет Горьковского муниципального района Омской области",
				},
			},
			Circumstance: row[10],
		}
		position.Circumstances = append(position.Circumstances, circumstance)
	}

	doc := BudgetaryCircumstances{
		Xmlns:    "http://bus.gov.ru/external/1",
		XmlnsNs2: "http://bus.gov.ru/types/1",
		Header: Header{
			ID:             uuid.New().String(),
			CreateDateTime: currentDate,
		},
		Body: Body{
			Position: position,
		},
	}

	outputFile := fmt.Sprintf("./out/%s.xml", strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath)))
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

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
