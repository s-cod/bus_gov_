// Package rd формирование отчета о результатах деятельности
package rd

import (
	"bus_gov_go/pkg/utils"
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

func ProcessFile(filePath string) error {
	Gd := utils.GetDigit
	Gs := utils.GetString
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
	year := currentTime[0:4]
	R, err := f.GetRows("Лист1")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	OKTMO := Gs(R, 25, 87)

	position := Position{
		PositionID: uuid.New().String(),
		ChangeDate: currentTime,
		// Placer            Placer
		// Initiator         Initiator
		VersionNumber: "1",
		// SignerDetails     SignerDetails
		ReportMonth:       Gs(R, 15, 42),
		ReportYearShort:   year,
		AgencyTypeCode:    Gs(R, 20, 87),
		FounderAgencyName: "Комитет по образованию администрации Горьковского района Омской области",
		GlavaCode:         Gs(R, 21, 87),
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
		FullName: Gs(R, 17, 16),
		Inn:      Gs(R, 18, 87),
		Kpp:      Gs(R, 19, 87),
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
		ManagerName:       Gs(R, 49, 54),
		ManagerPosition:   Gs(R, 49, 15),
		ExecutorName:      Gs(R, 49, 54),
		ExecutorPosition:  Gs(R, 49, 15),
		Phone:             " ",
		SignDate:          currentTime[:10],
		SignerDetailsType: "PAYMENTS_AND_RECEIPTS",
	}

	receipts := make([]Receipts, 0)
	payments := make([]Payments, 0)

	R, err = f.GetRows("Лист2-3")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	for i := 19; i < 69; i++ {
		if len(R[i-1]) < 122 {
			continue
		}

		receipts = append(receipts, Receipts{
			Name:                   Gs(R, i, 1),
			LineCode:               Gs(R, i, 56),
			ReceiptsTotal:          "0",
			ReportingFinancialYear: Gd(R, i, 62),
			PrecedingFinancialYear: Gd(R, i, 82),
			Change:                 Gd(R, i, 102),
			TotalAmountShare:       Gd(R, i, 122),
		})
	}

	R, err = f.GetRows("Лист4-5")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	for i := 15; i < 69; i++ {
		if Gs(R, i, 25) == "" {
			continue
		}

		// fmt.Println("LineCode:", getFinLen(rows[i], 31))
		payments = append(payments, Payments{
			Name:                                  Gs(R, i, 1),
			LineCode:                              Gs(R, i, 20),
			PaymentsTotal:                         Gd(R, i, 25),
			PaymentsTotalShare:                    Gd(R, i, 32),
			FinancialSupportSubsidiesStateTasks:   Gd(R, i, 38),
			PaymentsTotalShare6:                   Gd(R, i, 45),
			FinancialSupportSubsidiesOther:        Gd(R, i, 51),
			PaymentsTotalShare8:                   Gd(R, i, 58),
			GrantSubsidiesStateBudget:             Gd(R, i, 64),
			PaymentsTotalShare10:                  Gd(R, i, 71),
			GrantSubsidiesStateBudgetSubject:      Gd(R, i, 77),
			PaymentsTotalShare12:                  Gd(R, i, 84),
			FinancialSupportOms:                   Gd(R, i, 90),
			PaymentsTotalShare14:                  Gd(R, i, 97),
			FinancialSupportIncomeActivitiesTotal: Gd(R, i, 103),
			PaymentsTotalShare16:                  Gd(R, i, 110),
			FundsIncome:                           Gd(R, i, 116),
			PaymentsTotalShare18:                  Gd(R, i, 123),
			FundsIncomeGratuitousReceipts:         Gd(R, i, 129),
			PaymentsTotalShare20:                  "0",
		})
	}

	// ===============================================
	// Сведения о численности и оплате труда
	// ===============================================

	receiptsAndPayments := ReceiptsAndPayments{

		Receipts: receipts,
		Payments: payments,
	}

	// TODO сделать перебор
	numberEmployeesRemunerationGroups := make([]NumberEmployeesRemunerationGroups, 3)

	R, err = f.GetRows("Лист11")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	R2, err := f.GetRows("Лист12")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	R3, err := f.GetRows("Лист13")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	s11 := []int{27, 31, 35}
	s12 := []int{18, 22, 26}
	s13 := []int{13, 16, 21}
	for i := range 3 {
		numberEmployeesRemunerationGroups[i] = NumberEmployeesRemunerationGroups{
			// Name: rows[s11[i]][0],
			GroupStaff: GroupStaff{
				Name:     Gs(R, s11[i], 23),
				LineCode: Gs(R, s11[i], 23),
			},
			Employees: Employees{
				StaffingStartYear: StaffingStartYear{
					StaffingTotal: Gd(R, s11[i], 28),
					StaffingBase:  Gd(R, s11[i], 36),
					Replaced:      Gd(R, s11[i], 44),
					Vacancy:       Gd(R, s11[i], 51),
				},
				AverageOfYear: AverageOfYear{
					Total:           Gd(R, s11[i], 58),
					Base:            Gd(R, s11[i], 66),
					MainPlace:       Gd(R, s11[i], 74),
					InsidePartTme:   Gd(R, s11[i], 82),
					OutsidePartTime: Gd(R, s11[i], 89),
				},
				Contracts: Contracts{
					Employees:    Gd(R, s11[i], 96),
					NotEmployees: Gd(R, s11[i], 104),
				},
				StaffingEndYear: StaffingEndYear{
					StaffingTotal: Gd(R, s11[i], 112),
					StaffingBase:  Gd(R, s11[i], 120),
					Replaced:      Gd(R, s11[i], 128),
					Vacancy:       Gd(R, s11[i], 135),
				},
			},
			Salary: Salary{
				SalaryFund: SalaryFund{
					Total:           Gd(R2, s12[i], 30),
					Base:            Gd(R2, s12[i], 38),
					BaseFullTime:    Gd(R2, s12[i], 46),
					MainPartTme:     Gd(R2, s12[i], 54),
					InsideCombining: Gd(R2, s12[i], 62),
					Outside:         Gd(R2, s12[i], 70),
				},
				Accrued: Accrued{
					Employees:    Gd(R2, s12[i], 78),
					NotEmployees: Gd(R2, s12[i], 86),
				},
				SalaryAnalytic: SalaryAnalytic{
					BaseWorkplace: BaseWorkplace{
						SubsidiesContract:      Gd(R2, s12[i], 94),
						SubsidiesOther:         Gd(R2, s12[i], 102),
						SubsidiesGrantFederal:  Gd(R2, s12[i], 110),
						SubsidiesGrantRegional: Gd(R2, s12[i], 118),
						Oms:                    Gd(R2, s12[i], 126),
						IncomeActivities:       Gd(R2, s12[i], 134),
					},
					InsideCombining: InsideCombining{
						SubsidiesContract:      Gd(R3, s13[i], 32),
						SubsidiesOther:         Gd(R3, s13[i], 42),
						SubsidiesGrantFederal:  Gd(R3, s13[i], 51),
						SubsidiesGrantRegional: Gd(R3, s13[i], 60),
						Oms:                    Gd(R3, s13[i], 69),
						IncomeActivities:       Gd(R3, s13[i], 78),
					},
					OutsideWorkplace: OutsideWorkplace{
						SubsidiesContract:      Gd(R3, s13[i], 87),
						SubsidiesOther:         Gd(R3, s13[i], 97),
						SubsidiesGrantFederal:  Gd(R3, s13[i], 106),
						SubsidiesGrantRegional: Gd(R3, s13[i], 115),
						Oms:                    Gd(R3, s13[i], 124),
						IncomeActivities:       Gd(R3, s13[i], 133),
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

	// ===============================================
	// Сведения о кредиторской задолженности
	// ===============================================
	R, err = f.GetRows("Лист8")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	creditPayment := []CreditPayment{}
	// fmt.Println(len(rows[25]))
	for i := 25; i < 48; i++ {
		if Gd(R, i, 44) == "0.00" && Gd(R, i, 51) == "0.00" && Gd(R, i, 58) == "0.00" && Gd(R, i, 65) == "0.00" && Gd(R, i, 72) == "0.00" {
			continue
		}

		creditPayment = append(creditPayment, CreditPayment{
			Name:                        Gs(R, i, 1),
			LineCode:                    Gs(R, i, 39),
			CreditPaymentStartYearTotal: Gd(R, i, 44),
			CreditPaymentStartYearDedlineReportingYear: Gd(R, i, 51),
			CreditPaymentEndYearTotal:                  Gd(R, i, 58),
			CreditPaymentEndYearTotalQuarter1:          Gd(R, i, 65),
			CreditPaymentEndYearTotalJanuary:           Gd(R, i, 72),
			CreditPaymentEndYearTotalQuarter2:          Gd(R, i, 79),
			CreditPaymentEndYearTotalQuarter3:          Gd(R, i, 86),
			CreditPaymentEndYearTotalQuarter4:          Gd(R, i, 93),
			CreditPaymentEndYearTotalNextYear:          Gd(R, i, 100),
			AmountDeferredTotal:                        Gd(R, i, 107),
			AmountDeferredSalary:                       Gd(R, i, 114),
			AmountDeferredClaims:                       Gd(R, i, 121),
			AmountDeferredNotReceived:                  Gd(R, i, 128),
			AmountDeferredOther:                        Gd(R, i, 135),
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
	R, err = f.GetRows("Лист15")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	R2, err = f.GetRows("Лист16")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	actualExpenses := ActualExpenses{
		Utilities: Utilities{
			Total:             Gd(R2, 11, 67),
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyMaintenanceServices: PropertyMaintenanceServices{
			Total:             Gd(R2, 11, 91),
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		PropertyTax: PropertyTax{
			Total:             "0",
			ReimbursedByUsers: "0",
			UnusedProperty:    "0",
		},
		Total: Gd(R2, 11, 58),
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

	for i := 26; i < 40; i++ {

		if len(R[i-1]) < 87 {
			continue
		}
		// fmt.Println(R[i-1], len(R[i-1]))
		// fmt.Println(R[i-1], Gs(R, i, 1), i)
		if tmp {
			actualExpenses = actualExpensesNull
		}
		e := EstateObject{
			Name:      Gs(R, i, 1),
			Address:   Gs(R, i, 18),
			CadNumber: Gs(R, i, 33),
			Oktmo: Oktmo{
				Code: Gs(R, i, 42),
				Name: "Name",
			},
			UniqueObjectCode: strconv.Itoa(i),

			Unit: Unit{
				Code:   Gs(R, i, 69),
				Symbol: Gs(R, i, 61),
			},
			TypeObject: TypeObject{
				Type:     "1000",
				LineCode: Gs(R, i, 75),
			},
			Used: Used{
				Total:               Gd(R, i, 80),
				MainPurposeTask:     Gd(R, i, 87),
				MainPurposeOverTask: Gd(R, i, 95),
				OtherPurpose:        Gd(R, i, 103),
			},
			ActualExpenses: actualExpenses,
		}
		if Gs(R, i, 56) != "" {
			e.BuildingYear = Gs(R, i, 56)
		}

		estateObject = append(estateObject, e)
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
	R, err = f.GetRows("Лист17")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	lendObject := []LendObject{}

	for i := 27; i < 33; i++ {
		if Gs(R, i, 1) == "" {
			break
		}
		// fmt.Println(Gs(R, i, 1))
		lendObject = append(lendObject, LendObject{
			Name:    Gs(R, i, 1),
			Address: Gs(R, i, 15),
			Oktmo: Oktmo{
				Code: Gs(R, i, 24),
				Name: "name",
			},
			CadNumber: Gs(R, i, 29),
			Unit: Unit{
				Code:   Gs(R, i, 42),
				Symbol: Gs(R, i, 35),
			},
			LendObjectUse: LendObjectUse{
				LineCode: Gs(R, i, 47),
				Total:    Gd(R, i, 2),
				Used: UsedL{
					Total:               Gd(R, i, 58),
					MainPurposeTask:     Gd(R, i, 64),
					MainPurposeOverTask: Gd(R, i, 70),
					UsedOther:           Gd(R, i, 76),
				},
				UsedByEasement: Gd(R, i, 82),
				NotUsed: NotUsed{
					Total:                    Gd(R, i, 88),
					TemporaryUsedRent:        Gd(R, i, 94),
					TemporaryUsedFree:        Gd(R, i, 100),
					TemporaryUsedWithoutRigt: Gd(R, i, 106),
					TemporaryUsedOther:       Gd(R, i, 112),
				},
				ExpensesForLand: ExpensesForLand{
					Total:       Gd(R, i, 118),
					CostsTotal:  Gd(R, i, 124),
					CostsRefund: Gd(R, i, 130),
					LandTax:     Gd(R, i, 136),
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
	R, err = f.GetRows("Лист21")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}
	R2, err = f.GetRows("Лист22")
	if err != nil {
		return fmt.Errorf("не удалось прочитать лист: %w", err)
	}

	r := []int{33, 35}
	r2 := []int{18, 20}

	valuableMovablePropertyObject := []ValuableMovablePropertyObject{}
	for i := range r {
		valuableMovablePropertyObject = append(valuableMovablePropertyObject, ValuableMovablePropertyObject{
			Name: "53151",
			TypeMovableObject: TypeMovableObject{
				Type:     "2000",
				LineCode: Gd(R, r[i], 32),
			},
			AvailabilityEndPeriod: AvailabilityEndPeriod{
				Total:                  Gd(R, r[i], 38),
				UsedByAgency:           Gd(R, r[i], 51),
				TransferredForUseTotal: "0",
				ForRent:                "0",
				ForFree:                "0",
				NeedRepair:             "0",
				WaitWriteOffTotal:      "0",
				WaitWriteOffReplace:    "0",
			},
			ActualTermOfUse: ActualTermOfUse{
				ActualTerm: "Over121",
				Count:      Gd(R2, r2[i], 38),
				Cost:       Gd(R2, r2[i], 47),
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

	R, err = f.GetRows("Листы25-26")
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
					EndDate:      Gd(R, 49, 38),
					MiddleOfYear: Gd(R, 49, 51),
				},
				OperationManagment: OperationManagment{
					EndDate:      Gd(R, 49, 64),
					MiddleOfYear: Gd(R, 49, 77),
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
	R, err = f.GetRows("Листы29-30")
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
						EndDate:      Gd(R, 57, 22),
						MiddleOfYear: Gd(R, 57, 27),
					},
					OperationManagment: OperationManagment{
						EndDate:      Gd(R, 57, 32),
						MiddleOfYear: Gd(R, 57, 37),
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

	R, err = f.GetRows("Листы31-32")
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
					PeriodTotal:             Gd(R, 41, 37),
					TransportTax:            Gd(R, 41, 134),
					FuelCosts:               Gd(R, 41, 46),
					WheelCosts:              Gd(R, 41, 54),
					OsagoCosts:              Gd(R, 41, 62),
					VolunteerInsuranceCosts: Gd(R, 41, 70),
					RepairsCosts:            Gd(R, 41, 78),
					MaintenanceCosts:        Gd(R, 41, 86),
					GaragesRent:             Gd(R, 41, 102),
					GaragesMaintenance:      Gd(R, 41, 94),
					Drivers:                 Gd(R, 41, 110),
					ServisesPersonnel:       Gd(R, 41, 118),
					AdministrativePersonnel: Gd(R, 41, 126),
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
