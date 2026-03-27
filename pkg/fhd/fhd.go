package fhd

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func getFloat(s string) string {
	if s == "" {
		return "0.00"
	}
	result := strings.Replace(s, ",", "", -1)
	return result
}

func ProcessFile(filePath string) error {
	f, err := excelize.OpenFile(filePath) //, excelize.Options{RawCellValue: true})
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer f.Close()
	currentTime := time.Now().Format("2006-01-02T15:04:05")
	year := currentTime[0:4]
	rows, err := f.GetRows(year)
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	position := Position{
		PositionID: uuid.New().String(),
		ChangeDate: currentTime,
		// Placer             `xml:"ns2:placer"`
		// Initiator          Initiator          `xml:"ns2:initiator"`
		VersionNumber: "1",
		FinancialYear: currentTime[:4],
		// GeneralData        GeneralData        `xml:"ns3:generalData"`
		// PlanPaymentIndexes []PlanPaymentIndex `xml:"ns3:planPaymentIndex"`
		// PlanPaymentTRUs    []PlanPaymentTRU   `xml:"ns3:planPaymentTRU"`
	}

	position.Placer = Placer{
		RegNum:   rows[15][7],
		FullName: rows[17][2],
		INN:      rows[16][7],
		KPP:      rows[17][7],
	}

	position.Initiator = Initiator{
		RegNum:   rows[13][7],
		FullName: rows[14][2],
		INN:      "5512001782",
		KPP:      "551201001",
	}

	position.GeneralData = GeneralData{
		Date:         currentTime[:10],
		DateApprovel: currentTime[:10],
		FounderAuthority: FounderAuthority{
			RegNum:   rows[15][7],
			FullName: rows[17][2],
			INN:      rows[16][7],
			KPP:      rows[17][7],
		},
		OKEI: OKEI{
			Code:   "383",
			Symbol: "руб",
		},
		ManagerName:            rows[126][5],
		ManagerPosition:        rows[126][2],
		ExecutorName:           rows[129][4],
		ExecutorPosition:       rows[129][2],
		Phone:                  rows[129][5],
		SignDate:               currentTime[:10],
		FounderManagerName:     rows[4][7],
		FounderManagerPosition: rows[2][6],
		FounderSignDate:        currentTime[:10],
	}

	// start := 0
	for _, row := range rows[25:90] { // Диапазон B85:S100
		if len(row) < 8 {
			continue
		}
		kbk_, _ := strconv.Atoi(row[3])
		kbk := ""
		if kbk_ != 0 {
			kbk = fmt.Sprintf("%v", kbk_)
		}
		s1 := getFloat(row[5])
		s2 := getFloat(row[6])
		s3 := getFloat(row[7])
		if s1 == "0.00" && s2 == "0.00" && s3 == "0.00" {
			continue
		}

		planPaymentIndex := PlanPaymentIndex{
			Name:     row[0],
			LineCode: row[2],
			KBK:      kbk,
			Sum: Sum{
				FinancialYearSum: s1,
				PlanFirstYearSum: s2,
				PlanLastYearSum:  s3,
			},
		}
		position.PlanPaymentIndex = append(position.PlanPaymentIndex, planPaymentIndex)

	}

	for _, row := range rows[97:125] { // Диапазон B85:S100
		if len(row) < 6 {
			continue
		}

		s1 := getFloat(row[5])
		s2 := getFloat(row[6])
		s3 := getFloat(row[7])
		if s1 == "0.00" && s2 == "0.00" && s3 == "0.00" {
			continue
		}
		YearStart := ""
		if len(row[3]) > 1 {
			YearStart = currentTime[0:4]
		}

		planPaymentTRU := PlanPaymentTRU{
			Name:      row[1],
			LineCode:  row[2],
			YearStart: YearStart, //TODO
			Sum: Sum{
				FinancialYearSum: getFloat(row[5]),
				PlanFirstYearSum: getFloat(row[6]),
				PlanLastYearSum:  getFloat(row[7]),
			},
		}
		position.PlanPaymentTRU = append(position.PlanPaymentTRU, planPaymentTRU)

	}

	doc := FinancialActivityPlan2020{
		Xmlns:    "http://bus.gov.ru/external/1",
		XmlnsNs2: "http://bus.gov.ru/types/1",
		XmlnsNs3: "http://bus.gov.ru/types/3",
		Header: Header{
			ID:             uuid.New().String(),
			CreateDateTime: currentTime, //"2024-12-31T00:00:00",
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
