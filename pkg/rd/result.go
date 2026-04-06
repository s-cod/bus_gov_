// Package rd формирование отчета о результатах деятельности
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

// func GD(s string) string {
// 	if s == "" {
// 		return "0.00"
// 	}
// 	result := strings.ReplaceAll(s, ",", "")
// 	return result
// }

func GD(s [][]string, r, c int) string {
	if r == 0 || c == 0 {
		r, err := fmt.Printf("неверный диапазон ячеек row:%v col:%v", r, c)
		if err != nil {
			panic(err.Error())
		}
		panic(r)
	}
	r -= 1
	c -= 1

	if len(s[r]) < c {
		return "0.00"
	}
	if s[r][c] == "" {
		return "0.00"
	}

	tmp := s[r][c]
	result := strings.ReplaceAll(tmp, ",", "")
	return result
}

func ProcessFile(filePath string) error {
	f, err := excelize.OpenFile(filePath) //, excelize.Options{RawCellValue: true})
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err.Error())
		}
	}()
	currentTime := time.Now().Format("2006-01-02T15:04:05")
	rows, err := f.GetRows("Состав сведений")
	year := GD(rows, 11, 58)
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	OKTMO := GD(rows, 18, 94)

	position := Position{
		PositionID: uuid.New().String(),
		ChangeDate: currentTime,
		// Placer            Placer
		// Initiator         Initiator
		VersionNumber: "1",
		// SignerDetails     SignerDetails
		ReportMonth:       GD(rows, 11, 44),
		ReportYearShort:   year,
		AgencyTypeCode:    GD(rows, 15, 28),
		FounderAgencyName: GD(rows, 17, 28),
		GlavaCode:         GD(rows, 16, 94),
		PpoName:           "",
		OktmoCode:         OKTMO,
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
		FullName: GD(rows, 14, 28),
		Inn:      GD(rows, 13, 94),
		Kpp:      GD(rows, 14, 94),
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

	rows, err = f.GetRows("Поступления и выплаты")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	for i := 19; i < 52; i++ {
		if GD(rows, i, 80) == "0.00" {
			continue
		}

		receipts = append(receipts, Receipts{
			Name:                   rows[i][0],
			LineCode:               rows[i][71],
			ReceiptsTotal:          "80",
			ReportingFinancialYear: GD(rows, i, 80),
			PrecedingFinancialYear: GD(rows, i, 102),
			Change:                 GD(rows, i, 1241),
			TotalAmountShare:       GD(rows, i, 146),
		})
	}

	// TODO заполнить в цикле
	for i := 61; i < 89; i++ {
		if GD(rows, i, 25) == "0.00" {
			continue
		}

		// fmt.Println("LineCode:", GD(rows[i], 31))
		payments = append(payments, Payments{
			Name:                                  GD(rows, i, 0),
			LineCode:                              GD(rows, i, 19),
			PaymentsTotal:                         GD(rows, i, 25),
			PaymentsTotalShare:                    GD(rows, i, 32),
			FinancialSupportSubsidiesStateTasks:   GD(rows, i, 38),
			PaymentsTotalShare6:                   GD(rows, i, 47),
			FinancialSupportSubsidiesOther:        GD(rows, i, 55),
			PaymentsTotalShare8:                   GD(rows, i, 62),
			GrantSubsidiesStateBudget:             GD(rows, i, 70),
			PaymentsTotalShare10:                  GD(rows, i, 78),
			GrantSubsidiesStateBudgetSubject:      GD(rows, i, 86),
			PaymentsTotalShare12:                  GD(rows, i, 94),
			FinancialSupportOms:                   GD(rows, i, 102),
			PaymentsTotalShare14:                  GD(rows, i, 111),
			FinancialSupportIncomeActivitiesTotal: GD(rows, i, 119),
			PaymentsTotalShare16:                  GD(rows, i, 128),
			FundsIncome:                           GD(rows, i, 144),
			PaymentsTotalShare18:                  GD(rows, i, 152),
			FundsIncomeGratuitousReceipts:         GD(rows, i, 160),
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
					StaffingTotal: GD(rows, s11[i], 27),
					StaffingBase:  GD(rows, s11[i], 35),
					Replaced:      GD(rows, s11[i], 35),
					Vacancy:       GD(rows, s11[i], 50),
				},
				AverageOfYear: AverageOfYear{
					Total:           GD(rows, s11[i], 57),
					Base:            GD(rows, s11[i], 65),
					MainPlace:       GD(rows, s11[i], 73),
					InsidePartTme:   GD(rows, s11[i], 81),
					OutsidePartTime: GD(rows, s11[i], 88),
				},
				Contracts: Contracts{
					Employees:    GD(rows, s11[i], 95),
					NotEmployees: GD(rows, s11[i], 103),
				},
				StaffingEndYear: StaffingEndYear{
					StaffingTotal: GD(rows, s11[i], 111),
					StaffingBase:  GD(rows, s11[i], 119),
					Replaced:      GD(rows, s11[i], 119),
					Vacancy:       GD(rows, s11[i], 134),
				},
			},
			Salary: Salary{
				SalaryFund: SalaryFund{
					Total:           GD(rows2, s12[i], 29),
					Base:            GD(rows2, s12[i], 37),
					BaseFullTime:    GD(rows2, s12[i], 45),
					MainPartTme:     GD(rows2, s12[i], 53),
					InsideCombining: GD(rows2, s12[i], 61),
					Outside:         GD(rows2, s12[i], 69),
				},
				Accrued: Accrued{
					Employees:    GD(rows2, s12[i], 77),
					NotEmployees: GD(rows2, s12[i], 85),
				},
				SalaryAnalytic: SalaryAnalytic{
					BaseWorkplace: BaseWorkplace{
						SubsidiesContract:      GD(rows2, s12[i], 93),
						SubsidiesOther:         GD(rows2, s12[i], 101),
						SubsidiesGrantFederal:  GD(rows2, s12[i], 109),
						SubsidiesGrantRegional: GD(rows2, s12[i], 117),
						Oms:                    GD(rows2, s12[i], 125),
						IncomeActivities:       GD(rows2, s12[i], 133),
					},
					InsideCombining: InsideCombining{
						SubsidiesContract:      GD(rows3, s13[i], 31),
						SubsidiesOther:         GD(rows3, s13[i], 41),
						SubsidiesGrantFederal:  GD(rows3, s13[i], 50),
						SubsidiesGrantRegional: GD(rows3, s13[i], 59),
						Oms:                    GD(rows3, s13[i], 68),
						IncomeActivities:       GD(rows3, s13[i], 77),
					},
					OutsideWorkplace: OutsideWorkplace{
						SubsidiesContract:      GD(rows3, s13[i], 86),
						SubsidiesOther:         GD(rows3, s13[i], 96),
						SubsidiesGrantFederal:  GD(rows3, s13[i], 105),
						SubsidiesGrantRegional: GD(rows3, s13[i], 114),
						Oms:                    GD(rows3, s13[i], 123),
						IncomeActivities:       GD(rows3, s13[i], 132),
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
			Name:                        GD(rows, i, 0),
			LineCode:                    GD(rows, i, 38),
			CreditPaymentStartYearTotal: GD(rows, i, 43),
			CreditPaymentStartYearDedlineReportingYear: GD(rows, i, 50),
			CreditPaymentEndYearTotal:                  GD(rows, i, 57),
			CreditPaymentEndYearTotalQuarter1:          GD(rows, i, 64),
			CreditPaymentEndYearTotalJanuary:           GD(rows, i, 71),
			CreditPaymentEndYearTotalQuarter2:          GD(rows, i, 78),
			CreditPaymentEndYearTotalQuarter3:          GD(rows, i, 85),
			CreditPaymentEndYearTotalQuarter4:          GD(rows, i, 92),
			CreditPaymentEndYearTotalNextYear:          GD(rows, i, 99),
			AmountDeferredTotal:                        GD(rows, i, 106),
			AmountDeferredSalary:                       GD(rows, i, 113),
			AmountDeferredClaims:                       GD(rows, i, 120),
			AmountDeferredNotReceived:                  GD(rows, i, 127),
			AmountDeferredOther:                        GD(rows, i, 134),
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
			Total:             GD(rows2, 10, 66),
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyMaintenanceServices: PropertyMaintenanceServices{
			Total:             GD(rows2, 10, 91),
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyTax: PropertyTax{
			Total:             "0",
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		Total: GD(rows2, 10, 57),
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
				Code: OKTMO,
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
				Code: OKTMO,
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
					MainPurposeOverTask: GD(rows, i, 69),
					UsedOther:           GD(rows, i, 75),
				},
				UsedByEasement: GD(rows, i, 81),
				NotUsed: NotUsed{
					Total:                    GD(rows, i, 87),
					TemporaryUsedRent:        GD(rows, i, 93),
					TemporaryUsedFree:        GD(rows, i, 99),
					TemporaryUsedWithoutRigt: GD(rows, i, 105),
					TemporaryUsedOther:       GD(rows, i, 111),
				},
				ExpensesForLand: ExpensesForLand{
					Total:       GD(rows, i, 117),
					CostsTotal:  GD(rows, i, 123),
					CostsRefund: GD(rows, i, 129),
					LandTax:     GD(rows, i, 135),
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

	// ======================================================================
	// Сведения об особо ценном движимом имуществе (за исключение
	// м транспортны х средств)
	// ======================================================================
	rows, err = f.GetRows("Лист21")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	rows2, err = f.GetRows("Лист22")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	r := []int{32, 34}
	r2 := []int{17, 19}

	valuableMovablePropertyObject := []ValuableMovablePropertyObject{}
	for i := range r {
		valuableMovablePropertyObject = append(valuableMovablePropertyObject, ValuableMovablePropertyObject{
			Name: "53151",
			TypeMovableObject: TypeMovableObject{
				Type:     "2000",
				LineCode: GD(rows, r[i], 31),
			},
			AvailabilityEndPeriod: AvailabilityEndPeriod{
				Total:                  GD(rows, r[i], 37),
				UsedByAgency:           GD(rows, r[i], 50),
				TransferredForUseTotal: "0",
				ForRent:                "0",
				ForFree:                "0",
				NeedRepair:             "0",
				WaitWriteOffTotal:      "0",
				WaitWriteOffReplace:    "0",
			},
			ActualTermOfUse: ActualTermOfUse{
				ActualTerm: "Over121",
				Count:      GD(rows2, r2[i], 37),
				Cost:       GD(rows2, r2[i], 46),
			},
		})
	}

	valuableMovableProperty := ValuableMovableProperty{
		ValuableMovablePropertyObjects: ValuableMovablePropertyObjects{
			TypeMovableObject: TypeMovableObject{
				Type:     "2000",
				LineCode: "2000",
			},
			ValuableMovablePropertyObject: valuableMovablePropertyObject,
		},
	}

	// ==============================
	// Транспортные средства
	// ==============================

	// TODO Нужножно нормально обработать в цикле по таблице
	// улучшить структуру в срезы

	rows, err = f.GetRows("Листы25-26")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	usedVehicles := UsedVehicles{
		UsedVehiclesTransports: UsedVehiclesTransports{
			TransportsType: TransportsType{
				TranspType: "1000",
				LineCode:   "1000",
			},
			VehiclesDetails: VehiclesDetails{
				Name: "Автобусы",
				TransportsType: TransportsType{
					TranspType: "1000",
					LineCode:   "1500",
				},
				Total: Total{
					EndDate:      GD(rows, 48, 37),
					MiddleOfYear: GD(rows, 48, 50),
				},
				OperationManagment: OperationManagment{
					EndDate:      GD(rows, 48, 63),
					MiddleOfYear: GD(rows, 48, 76),
				},
			},
		},
	}
	notUsedVehicles := NotUsedVehicles{
		NotUsedVehiclesTransports: NotUsedVehiclesTransports{
			TransportsType: TransportsType{
				TranspType: "1000",
				LineCode:   "1000",
			},
		},
	}
	rows, err = f.GetRows("Листы29-30")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	directionOfUse := DirectionOfUse{
		DirectionOfUseTransports: DirectionOfUseTransports{
			TransportsType: TransportsType{
				TranspType: "1000",
				LineCode:   "1000",
			},
			DirectionOfUseTransport: DirectionOfUseTransport{
				Name: "Автобусы",
				TransportsType: TransportsType{
					TranspType: "1000",
					LineCode:   "1500",
				},
				DirectlyUsedVehicles: DirectlyUsedVehicles{
					Total: Total{
						EndDate:      GD(rows, 56, 21),
						MiddleOfYear: GD(rows, 56, 26),
					},
					OperationManagment: OperationManagment{
						EndDate:      GD(rows, 56, 31),
						MiddleOfYear: GD(rows, 56, 36),
					},
					UnderLease: UnderLease{
						EndDate:      "0",
						MiddleOfYear: "0",
					},
					UnderGratuitous: UnderGratuitous{
						EndDate:      "0",
						MiddleOfYear: "0",
					},
				},
			},
		},
	}

	rows, err = f.GetRows("Листы31-32")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	costMaintenanceVehicles := CostMaintenanceVehicles{
		CostMaintenanceVehiclesTransports: CostMaintenanceVehiclesTransports{
			TransportsType: TransportsType{
				TranspType: "1000",
				LineCode:   "1000",
			},
			CostMaintenanceVehiclesTransport: CostMaintenanceVehiclesTransport{
				Name: "Автобусы",
				TransportsType: TransportsType{
					TranspType: "1000",
					LineCode:   "1500",
				},
				VehiclesExpensesMaintenance: VehiclesExpensesMaintenance{
					PeriodTotal:             GD(rows, 40, 36),
					TransportTax:            GD(rows, 40, 45),
					FuelCosts:               GD(rows, 40, 53),
					WheelCosts:              GD(rows, 40, 61),
					OsagoCosts:              GD(rows, 40, 69),
					VolunteerInsuranceCosts: GD(rows, 40, 77),
					RepairsCosts:            GD(rows, 40, 85),
					MaintenanceCosts:        GD(rows, 40, 93),
					GaragesRent:             GD(rows, 40, 101),
					GaragesMaintenance:      GD(rows, 40, 109),
					Drivers:                 GD(rows, 40, 117),
					ServisesPersonnel:       GD(rows, 40, 125),
					AdministrativePersonnel: GD(rows, 40, 133),
				},
			},
		},
	}

	vehicles := Vehicles{
		UsedVehicles:            usedVehicles,
		NotUsedVehicles:         notUsedVehicles,
		DirectionOfUse:          directionOfUse,
		CostMaintenanceVehicles: costMaintenanceVehicles,
	}

	position.AssetsUse = AssetsUse{
		Xmlns:                   "http://bus.gov.ru/types/1",
		EstateExceptLand:        estateExceptLand,
		LandPermanentUse:        landPermanentUse,
		ValuableMovableProperty: valuableMovableProperty,
		Vehicles:                vehicles,
	}

	doc := ReportActivityResult{
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

	outputFile := fmt.Sprintf("./out/%s.xml", strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath)))
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err.Error())
		}
	}()

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
