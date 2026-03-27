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

func ProcessFile(filePath string) error {
	f, err := excelize.OpenFile(filePath) //, excelize.Options{RawCellValue: true})
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer f.Close()
	currentTime := time.Now().Format("2006-01-02T15:04:05")
	year := currentTime[0:4]
	_, err = f.GetRows("Лист1")
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
		PpoName:           "ppoName1",
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
		FullName: "МБОУ &#34;Чучкинская основная общеобразовательная школа&#34; Горьковского района Омской области",
		Inn:      "5512004688",
		Kpp:      "551201001",
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
		ManagerName:       "Агейкина Н. А.",
		ManagerPosition:   "Главный бухгалтер",
		ExecutorName:      "Агейкина Н. А.",
		ExecutorPosition:  "Главный бухгалтер",
		Phone:             "83815721153",
		SignDate:          currentTime[:10],
		SignerDetailsType: "PAYMENTS_AND_RECEIPTS",
	}
	// TODO сделать перебор
	receipts := make([]Receipts, 0)
	payments := make([]Payments, 0)

	receipts = append(receipts, Receipts{
		Name:                   "Субсидии на финансовое обеспечение выполнения государственного (муниципального) задания",
		LineCode:               "0100",
		ReceiptsTotal:          "80",
		ReportingFinancialYear: "40",
		PrecedingFinancialYear: "40",
		Change:                 "0",
		TotalAmountShare:       "100",
	})

	// TODO заполнить в цикле
	payments = append(payments, Payments{
		Name:                                  "Оплата труда и компенсационные выплаты работникам",
		LineCode:                              "0100",
		PaymentsTotal:                         "4870038.51",
		PaymentsTotalShare:                    "59.31",
		FinancialSupportSubsidiesStateTasks:   "4524799.82",
		PaymentsTotalShare6:                   "55.11",
		FinancialSupportSubsidiesOther:        "345238.69",
		PaymentsTotalShare8:                   "4.2",
		GrantSubsidiesStateBudget:             "0",
		PaymentsTotalShare10:                  "0",
		GrantSubsidiesStateBudgetSubject:      "0",
		PaymentsTotalShare12:                  "0",
		FinancialSupportOms:                   "0",
		PaymentsTotalShare14:                  "0",
		FinancialSupportIncomeActivitiesTotal: "0",
		PaymentsTotalShare16:                  "0",
		FundsIncome:                           "0",
		PaymentsTotalShare18:                  "0",
		FundsIncomeGratuitousReceipts:         "0",
		PaymentsTotalShare20:                  "0",
	})

	// TODO заполнить в цикле
	receiptsAndPayments := ReceiptsAndPayments{
		Receipts: receipts,
		Payments: payments,
	}

	// TODO сделать перебор
	numberEmployeesRemunerationGroup := make([]NumberEmployeesRemunerationGroup, 0)

	// TODO заполнить в цикле
	numberEmployeesRemunerationGroup = append(numberEmployeesRemunerationGroup, NumberEmployeesRemunerationGroup{
		Name: "Основные",
		GroupStaff: GroupStaff{
			Name:     "1000",
			LineCode: "1000",
		},
		Employees: Employees{
			StaffingStartYear: StaffingStartYear{
				StaffingTotal: "8.08",
				StaffingBase:  "8.08",
				Replaced:      "8.08",
				Vacancy:       "0",
			},
			AverageOfYear: AverageOfYear{
				Total:           "4",
				Base:            "4",
				MainPlace:       "4",
				InsidePartTme:   "0",
				OutsidePartTime: "0",
			},
			Contracts: Contracts{
				Employees:    "0",
				NotEmployees: "0",
			},
			StaffingEndYear: StaffingEndYear{
				StaffingTotal: "6.52",
				StaffingBase:  "6.52",
				Replaced:      "6.52",
				Vacancy:       "0",
			},
		},
		Salary: Salary{
			SalaryFund: SalaryFund{
				Total:           "3301675.3",
				Base:            "3301675.3",
				BaseFullTime:    "3301675.3",
				MainPartTme:     "0",
				InsideCombining: "0",
				Outside:         "0",
			},
			Accrued: Accrued{
				Employees:    "0",
				NotEmployees: "0",
			},
			SalaryAnalytic: SalaryAnalytic{
				BaseWorkplace: BaseWorkplace{
					SubsidiesContract:      "2568305.80",
					SubsidiesOther:         "733369.50",
					SubsidiesGrantFederal:  "0",
					SubsidiesGrantRegional: "0",
					Oms:                    "0",
					IncomeActivities:       "0",
				},
				InsideCombining: InsideCombining{
					SubsidiesContract:      "0",
					SubsidiesOther:         "0",
					SubsidiesGrantFederal:  "0",
					SubsidiesGrantRegional: "0",
					Oms:                    "0",
					IncomeActivities:       "0",
				},
				OutsideWorkplace: OutsideWorkplace{
					SubsidiesContract:      "0",
					SubsidiesOther:         "0",
					SubsidiesGrantFederal:  "0",
					SubsidiesGrantRegional: "0",
					Oms:                    "0",
					IncomeActivities:       "0",
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
	})

	position.Result = Result{
		Xmlns:                    "http://bus.gov.ru/types/1",
		ReceiptsAndPayments:      receiptsAndPayments,
		NonFinancialAssetsChange: NonFinancialAssetsChange{},
		NumberEmployeesRemuneration: NumberEmployeesRemuneration{
			NumberEmployeesRemunerationGroups: NumberEmployeesRemunerationGroups{
				GroupStaff: GroupStaff{
					Name:     "9000",
					LineCode: "9000",
				},
				NumberEmployeesRemunerationGroup: numberEmployeesRemunerationGroup,
				Employees:                        "",
				Salary:                           "",
			},
		},
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
