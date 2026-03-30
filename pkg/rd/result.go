package rd

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

func getFloat(s string) string {
	if s == "" {
		return "0.00"
	}
	result := strings.Replace(s, ",", "", -1)
	return result
}
func getFinLen(s []string, i int) string {
	if len(s) < i {
		return "0.00"
	}
	if s[i] == "" {
		return "0.00"
	}
	result := strings.Replace(s[i], ",", "", -1)
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
	rows, err := f.GetRows("Лист1")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	position := Position{
		PositionId: uuid.New().String(),
		ChangeDate: currentTime,
		// Placer            Placer
		// Initiator         Initiator
		VersionNumber: "1",
		// SignerDetails     SignerDetails
		ReportMonth:       "Январь",
		ReportYearShort:   year,
		AgencyTypeCode:    "02",
		FounderAgencyName: "Комитет по образованию администрации Горьковского района Омской области",
		GlavaCode:         "704",
		PpoName:           "",
		OktmoCode:         "52509000147",
		ReportYear:        year,
		ConfirmDate:       currentTime[:10],
		// Result            Result
		AssetsUse: AssetsUse{
			Xmlns: "http://bus.gov.ru/types/1",
		},
		EffectiveActivity: EffectiveActivity{
			Xmlns: "http://bus.gov.ru/types/1",
		},
	}

	position.Placer = Placer{
		Xmlns:    "http://bus.gov.ru/types/1",
		RegNum:   "523D2025",
		FullName: rows[16][15],
		Inn:      rows[17][86],
		Kpp:      rows[18][86],
	}

	position.Initiator = Initiator{
		Xmlns:    "http://bus.gov.ru/types/1",
		RegNum:   "523D1347",
		FullName: "Комитет по образованию администрации Горьковского района Омской области",
		Inn:      "5501290035",
		Kpp:      "550101001",
	}

	position.SignerDetails = SignerDetails{
		Xmlns:             "http://bus.gov.ru/types/1",
		ManagerName:       rows[48][53],
		ManagerPosition:   rows[48][14],
		ExecutorName:      rows[48][53],
		ExecutorPosition:  rows[48][14],
		Phone:             " ",
		SignDate:          currentTime[:10],
		SignerDetailsType: "PAYMENTS_AND_RECEIPTS",
	}

	receipts := make([]Receipts, 0)
	payments := make([]Payments, 0)

	rows, err = f.GetRows("Лист2-3")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	for i := 18; i < 68; i++ {
		if len(rows[i]) < 122 {
			continue
		}

		receipts = append(receipts, Receipts{
			Name:                   rows[i][0],
			LineCode:               rows[i][55],
			ReceiptsTotal:          "80",
			ReportingFinancialYear: getFloat(rows[i][61]),
			PrecedingFinancialYear: getFloat(rows[i][81]),
			Change:                 getFloat(rows[i][101]),
			TotalAmountShare:       getFloat(rows[i][121]),
		})
	}

	rows, err = f.GetRows("Лист4-5")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	// TODO заполнить в цикле
	for i := 14; i < 67; i++ {
		if len(rows[i]) < 19 {
			continue
		}

		// fmt.Println("LineCode:", getFinLen(rows[i], 31))
		payments = append(payments, Payments{
			Name:                                  rows[i][0],
			LineCode:                              rows[i][19],
			PaymentsTotal:                         getFinLen(rows[i], 24),
			PaymentsTotalShare:                    getFinLen(rows[i], 31),
			FinancialSupportSubsidiesStateTasks:   getFinLen(rows[i], 37),
			PaymentsTotalShare6:                   getFinLen(rows[i], 44),
			FinancialSupportSubsidiesOther:        getFinLen(rows[i], 50),
			PaymentsTotalShare8:                   getFinLen(rows[i], 57),
			GrantSubsidiesStateBudget:             getFinLen(rows[i], 63),
			PaymentsTotalShare10:                  getFinLen(rows[i], 70),
			GrantSubsidiesStateBudgetSubject:      getFinLen(rows[i], 76),
			PaymentsTotalShare12:                  getFinLen(rows[i], 83),
			FinancialSupportOms:                   getFinLen(rows[i], 89),
			PaymentsTotalShare14:                  getFinLen(rows[i], 96),
			FinancialSupportIncomeActivitiesTotal: getFinLen(rows[i], 102),
			PaymentsTotalShare16:                  getFinLen(rows[i], 109),
			FundsIncome:                           getFinLen(rows[i], 115),
			PaymentsTotalShare18:                  getFinLen(rows[i], 122),
			FundsIncomeGratuitousReceipts:         getFinLen(rows[i], 128),
			PaymentsTotalShare20:                  "0",
		})
	}

	receiptsAndPayments := ReceiptsAndPayments{
		Receipts: receipts,
		Payments: payments,
	}

	// TODO сделать перебор
	numberEmployeesRemunerationGroups := make([]NumberEmployeesRemunerationGroups, 3)

	rows, err = f.GetRows("Лист11")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	rows2, err := f.GetRows("Лист12")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	rows3, err := f.GetRows("Лист13")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	s11 := []int{26, 30, 34}
	s12 := []int{17, 21, 25}
	for i := range 3 {
		numberEmployeesRemunerationGroups[i] = NumberEmployeesRemunerationGroups{
			// Name: rows[s11[i]][0],
			GroupStaff: GroupStaff{
				Name:     rows[s11[i]][22],
				LineCode: rows[s11[i]][22],
			},
			Employees: Employees{
				StaffingStartYear: StaffingStartYear{
					StaffingTotal: getFinLen(rows[s11[i]], 27),
					StaffingBase:  getFinLen(rows[s11[i]], 35),
					Replaced:      getFinLen(rows[s11[i]], 35),
					Vacancy:       getFinLen(rows[s11[i]], 50),
				},
				AverageOfYear: AverageOfYear{
					Total:           getFinLen(rows[s11[i]], 57),
					Base:            getFinLen(rows[s11[i]], 65),
					MainPlace:       getFinLen(rows[s11[i]], 73),
					InsidePartTme:   getFinLen(rows[s11[i]], 81),
					OutsidePartTime: getFinLen(rows[s11[i]], 88),
				},
				Contracts: Contracts{
					Employees:    getFinLen(rows[s11[i]], 95),
					NotEmployees: getFinLen(rows[s11[i]], 103),
				},
				StaffingEndYear: StaffingEndYear{
					StaffingTotal: getFinLen(rows[s11[i]], 111),
					StaffingBase:  getFinLen(rows[s11[i]], 119),
					Replaced:      getFinLen(rows[s11[i]], 119),
					Vacancy:       getFinLen(rows[s11[i]], 134),
				},
			},
			Salary: Salary{
				SalaryFund: SalaryFund{
					Total:           getFinLen(rows2[s12[i]], 29),
					Base:            getFinLen(rows2[s12[i]], 37),
					BaseFullTime:    getFinLen(rows2[s12[i]], 45),
					MainPartTme:     getFinLen(rows2[s12[i]], 53),
					InsideCombining: getFinLen(rows2[s12[i]], 61),
					Outside:         getFinLen(rows2[s12[i]], 69),
				},
				Accrued: Accrued{
					Employees:    getFinLen(rows2[s12[i]], 77),
					NotEmployees: getFinLen(rows2[s12[i]], 85),
				},
				SalaryAnalytic: SalaryAnalytic{
					BaseWorkplace: BaseWorkplace{
						SubsidiesContract:      getFinLen(rows2[s12[i]], 93),
						SubsidiesOther:         getFinLen(rows2[s12[i]], 101),
						SubsidiesGrantFederal:  getFinLen(rows2[s12[i]], 109),
						SubsidiesGrantRegional: getFinLen(rows2[s12[i]], 117),
						Oms:                    getFinLen(rows2[s12[i]], 125),
						IncomeActivities:       getFinLen(rows2[s12[i]], 133),
					},
					InsideCombining: InsideCombining{
						SubsidiesContract:      getFinLen(rows3[17], 31),
						SubsidiesOther:         getFinLen(rows3[17], 41),
						SubsidiesGrantFederal:  getFinLen(rows3[17], 50),
						SubsidiesGrantRegional: getFinLen(rows3[17], 59),
						Oms:                    getFinLen(rows3[17], 68),
						IncomeActivities:       getFinLen(rows3[17], 86),
					},
					OutsideWorkplace: OutsideWorkplace{
						SubsidiesContract:      getFinLen(rows3[17], 96),
						SubsidiesOther:         getFinLen(rows3[17], 96),
						SubsidiesGrantFederal:  getFinLen(rows3[17], 105),
						SubsidiesGrantRegional: getFinLen(rows3[17], 114),
						Oms:                    getFinLen(rows3[17], 123),
						IncomeActivities:       getFinLen(rows3[17], 132),
					},
					EmployeesContract: EmployeesContract{
						SubsidiesContract:      "0",
						SubsidiesOther:         "0",
						SubsidiesGrantFederal:  "0",
						SubsidiesGrantRegional: "0",
						Oms:                    "0",
						IncomeActivities:       "0",
					},
					NotEmployeesContract: NotEmployeesContract{
						SubsidiesContract:      "0",
						SubsidiesOther:         "0",
						SubsidiesGrantFederal:  "0",
						SubsidiesGrantRegional: "0",
						Oms:                    "0",
						IncomeActivities:       "0",
					},
				},
			},
		}
	}
	rows, err = f.GetRows("Лист8")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	creditPayment := []CreditPayment{}
	fmt.Println(len(rows[25]))
	for i := 24; i < 47; i++ {
		if len(rows[i]) < 40 {
			continue
		}
		creditPayment = append(creditPayment, CreditPayment{
			Name:                        getFinLen(rows[i], 0),
			LineCode:                    getFinLen(rows[i], 38),
			CreditPaymentStartYearTotal: getFinLen(rows[i], 43),
			CreditPaymentStartYearDedlineReportingYear: getFinLen(rows[i], 50),
			CreditPaymentEndYearTotal:                  getFinLen(rows[i], 57),
			CreditPaymentEndYearTotalQuarter1:          getFinLen(rows[i], 64),
			CreditPaymentEndYearTotalJanuary:           getFinLen(rows[i], 71),
			CreditPaymentEndYearTotalQuarter2:          getFinLen(rows[i], 78),
			CreditPaymentEndYearTotalQuarter3:          getFinLen(rows[i], 85),
			CreditPaymentEndYearTotalQuarter4:          getFinLen(rows[i], 92),
			CreditPaymentEndYearTotalNextYear:          getFinLen(rows[i], 99),
			AmountDeferredTotal:                        getFinLen(rows[i], 106),
			AmountDeferredSalary:                       getFinLen(rows[i], 113),
			AmountDeferredClaims:                       getFinLen(rows[i], 120),
			AmountDeferredNotReceived:                  getFinLen(rows[i], 127),
			AmountDeferredOther:                        getFinLen(rows[i], 134),
		})
	}

	position.Result = Result{
		Xmlns:                    "http://bus.gov.ru/types/1",
		ReceiptsAndPayments:      receiptsAndPayments,
		NonFinancialAssetsChange: NonFinancialAssetsChange{},
		NumberEmployeesRemuneration: NumberEmployeesRemuneration{
			NumberEmployeesRemunerationGroups: numberEmployeesRemunerationGroups,
		},
		CreditPayment: creditPayment,
		OpenedCreditAccounts: OpenedCreditAccounts{
			CreditAccountsRF:       "",
			CreditAccountsCurrency: "",
		},
	}

	doc := ReportActivityResult{
		Xmlns:    "http://bus.gov.ru/external/1",
		XmlnsNs2: "http://bus.gov.ru/types/1",

		Header: Header{
			Xmlns:          "http://bus.gov.ru/types/1",
			Id:             uuid.New().String(),
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
