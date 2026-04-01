package rd

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
		// AssetsUse: AssetsUse{
		// 	Xmlns: "http://bus.gov.ru/types/1",
		// },
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
	s13 := []int{12, 15, 20}
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
						SubsidiesContract:      getFinLen(rows3[s13[i]], 31),
						SubsidiesOther:         getFinLen(rows3[s13[i]], 41),
						SubsidiesGrantFederal:  getFinLen(rows3[s13[i]], 50),
						SubsidiesGrantRegional: getFinLen(rows3[s13[i]], 59),
						Oms:                    getFinLen(rows3[s13[i]], 68),
						IncomeActivities:       getFinLen(rows3[s13[i]], 77),
					},
					OutsideWorkplace: OutsideWorkplace{
						SubsidiesContract:      getFinLen(rows3[s13[i]], 86),
						SubsidiesOther:         getFinLen(rows3[s13[i]], 96),
						SubsidiesGrantFederal:  getFinLen(rows3[s13[i]], 105),
						SubsidiesGrantRegional: getFinLen(rows3[s13[i]], 114),
						Oms:                    getFinLen(rows3[s13[i]], 123),
						IncomeActivities:       getFinLen(rows3[s13[i]], 132),
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
	// fmt.Println(len(rows[25]))
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
	// ===================================
	// Сведения о недвижимом имуществе
	// ===================================
	rows, err = f.GetRows("Лист15")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	rows2, err = f.GetRows("Лист16")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	actualExpenses := ActualExpenses{
		Utilities: Utilities{
			Total:             getFinLen(rows2[10], 66),
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyMaintenanceServices: PropertyMaintenanceServices{
			Total:             getFinLen(rows2[10], 91),
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyTax: PropertyTax{
			Total:             "0",
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		Total: getFinLen(rows2[10], 57),
	}
	actualExpensesNull := ActualExpenses{
		Utilities: Utilities{
			Total:             "0",
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyMaintenanceServices: PropertyMaintenanceServices{
			Total:             "0",
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyTax: PropertyTax{
			Total:             "0",
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		Total: "0",
	}
	estateObject := []EstateObject{}

	tmp := false

	for i := 25; i < 40; i++ {
		if len(rows[i]) < 87 {
			continue
		}
		if tmp {
			actualExpenses = actualExpensesNull
		}
		estateObject = append(estateObject, EstateObject{
			Name:      rows[i][0],
			Address:   rows[i][17],
			CadNumber: rows[i][32],
			Oktmo: Oktmo{
				Code: rows[i][41],
				Name: "Name",
			},
			UniqueObjectCode: strconv.Itoa(i),
			BuildingYear:     rows[i][55],
			Unit: Unit{
				Code:   rows[i][68],
				Symbol: rows[i][60],
			},
			TypeObject: TypeObject{
				Type:     "1000",
				LineCode: "1100",
			},
			Used: Used{
				Total:               "1",
				MainPurposeTask:     "1",
				MainPurposeOverTask: "0",
				OtherPurpose:        "0",
			},
			ActualExpenses: actualExpenses,
		})
		tmp = true
	}

	estateExceptLand := EstateExceptLand{
		EstateExceptLandObjects: EstateExceptLandObjects{
			TypeObject: TypeObject{
				Type:     "1000",
				LineCode: "1000",
			},
			EstateObject: estateObject,
		},
	}

	// ===================================
	// Сведения о земельных участках, предоставленных на праве постоянного
	// (бессрочного ) пользования
	// ===================================
	rows, err = f.GetRows("Лист17")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	lendObject := []LendObject{}

	for i := 26; i < 31; i++ {
		if len(rows[i]) < 136 {
			continue
		}

		lendObject = append(lendObject, LendObject{
			Name:    rows[i][0],
			Address: rows[i][14],
			Oktmo: Oktmo{
				Code: rows[i][23],
				Name: "name",
			},
			CadNumber: rows[i][28],
			Unit: Unit{
				Code:   rows[i][41],
				Symbol: rows[i][34],
			},
			LendObjectUse: LendObjectUse{
				LineCode: rows[i][46],
				Total:    rows[i][51],
				Used: UsedL{
					Total:               rows[i][57],
					MainPurposeTask:     rows[i][63],
					MainPurposeOverTask: getFinLen(rows[i], 69),
					UsedOther:           getFinLen(rows[i], 75),
				},
				UsedByEasement: getFinLen(rows[i], 81),
				NotUsed: NotUsed{
					Total:                    getFinLen(rows[i], 87),
					TemporaryUsedRent:        getFinLen(rows[i], 93),
					TemporaryUsedFree:        getFinLen(rows[i], 99),
					TemporaryUsedWithoutRigt: getFinLen(rows[i], 105),
					TemporaryUsedOther:       getFinLen(rows[i], 111),
				},
				ExpensesForLand: ExpensesForLand{
					Total:       getFinLen(rows[i], 117),
					CostsTotal:  getFinLen(rows[i], 123),
					CostsRefund: getFinLen(rows[i], 129),
					LandTax:     getFinLen(rows[i], 135),
				},
			},
		})
	}

	landPermanentUse := LandPermanentUse{
		LendObject: lendObject,
		LendObjectUse: LendObjectUse{
			LineCode: "9000",
			Total:    "0",
			Used: UsedL{
				Total:               "0",
				MainPurposeTask:     "0",
				MainPurposeOverTask: "0",
				UsedOther:           "0",
			},
			UsedByEasement: "0",
			NotUsed: NotUsed{
				Total:                    "0",
				TemporaryUsedRent:        "0",
				TemporaryUsedFree:        "0",
				TemporaryUsedWithoutRigt: "0",
				TemporaryUsedOther:       "0",
			},
			ExpensesForLand: ExpensesForLand{
				Total:       "0",
				CostsTotal:  "0",
				CostsRefund: "0",
				LandTax:     "0",
			},
		},
	}

	// ==============================
	// Транспортные средства
	// ==============================
	vehicles := Vehicles{
		UsedVehicles:            UsedVehicles{},
		NotUsedVehicles:         NotUsedVehicles{},
		DirectionOfUse:          DirectionOfUse{},
		CostMaintenanceVehicles: CostMaintenanceVehicles{},
	}

	position.AssetsUse = AssetsUse{
		Xmlns:            "http://bus.gov.ru/types/1",
		EstateExceptLand: estateExceptLand,
		LandPermanentUse: landPermanentUse,
		Vehicles:         vehicles,
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
