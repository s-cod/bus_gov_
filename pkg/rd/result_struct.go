package rd

import "encoding/xml"

type ReportActivityResult struct {
	XMLName  xml.Name `xml:"reportActivityResult"`
	Xmlns    string   `xml:"xmlns,attr"`
	XmlnsNs2 string   `xml:"xmlns:ns2,attr"`
	Header   Header   `xml:"header"`
	Body     Body     `xml:"body"`
}

type Header struct {
	Xmlns          string `xml:"xmlns,attr"`
	ID             string `xml:"id"`
	CreateDateTime string `xml:"createDateTime"`
}

type Body struct {
	Position Position
}
type Position struct {
	XMLName           xml.Name `xml:"position"`
	PositionID        string   `xml:"ns2:positionId"`
	ChangeDate        string   `xml:"ns2:changeDate"`
	Placer            Placer
	Initiator         Initiator
	VersionNumber     string `xml:"ns2:versionNumber"`
	SignerDetails     SignerDetails
	ReportMonth       string `xml:"ns2:reportMonth"`
	ReportYearShort   string `xml:"ns2:reportYearShort"`
	AgencyTypeCode    string `xml:"ns2:agencyTypeCode"`
	FounderAgencyName string `xml:"ns2:founderAgencyName"`
	GlavaCode         string `xml:"ns2:glavaCode"`
	PpoName           string `xml:"ns2:ppoName"`
	OktmoCode         string `xml:"ns2:oktmoCode"`
	ReportYear        string `xml:"ns2:reportYear"`
	ConfirmDate       string `xml:"ns2:confirmDate"`
	Result            Result
	AssetsUse         AssetsUse
	EffectiveActivity EffectiveActivity
}

type Placer struct {
	XMLName  xml.Name `xml:"placer"`
	Xmlns    string   `xml:"xmlns,attr"`
	RegNum   string   `xml:"regNum"`
	FullName string   `xml:"fullName"`
	Inn      string   `xml:"inn"`
	Kpp      string   `xml:"kpp"`
}
type Initiator struct {
	XMLName  xml.Name `xml:"initiator"`
	Xmlns    string   `xml:"xmlns,attr"`
	RegNum   string   `xml:"regNum"`
	FullName string   `xml:"fullName"`
	Inn      string   `xml:"inn"`
	Kpp      string   `xml:"kpp"`
}

type SignerDetails struct {
	XMLName           xml.Name `xml:"signerDetails"`
	Xmlns             string   `xml:"xmlns,attr"`
	ManagerName       string   `xml:"managerName"`
	ManagerPosition   string   `xml:"managerPosition"`
	ExecutorName      string   `xml:"executorName"`
	ExecutorPosition  string   `xml:"executorPosition"`
	Phone             string   `xml:"phone"`
	SignDate          string   `xml:"signDate"`
	SignerDetailsType string   `xml:"signerDetailsType"`
}

type Result struct {
	XMLName                     xml.Name `xml:"result"`
	Xmlns                       string   `xml:"xmlns,attr"`
	ReceiptsAndPayments         ReceiptsAndPayments
	NonFinancialAssetsChange    NonFinancialAssetsChange
	CreditPayment               []CreditPayment
	NumberEmployeesRemuneration NumberEmployeesRemuneration
	OpenedCreditAccounts        OpenedCreditAccounts
}

// Сведения о поступлениях и выплатах учреждения
type ReceiptsAndPayments struct {
	XMLName  xml.Name `xml:"receiptsAndPayments"`
	Receipts []Receipts
	Payments []Payments
}

type Receipts struct {
	XMLName                xml.Name `xml:"receipts"`
	Name                   string   `xml:"name"`
	LineCode               string   `xml:"lineCode"`
	ReceiptsTotal          string   `xml:"receiptsTotal"`
	ReportingFinancialYear string   `xml:"reportingFinancialYear"`
	PrecedingFinancialYear string   `xml:"precedingFinancialYear"`
	Change                 string   `xml:"change"`
	TotalAmountShare       string   `xml:"totalAmountShare"`
}
type Payments struct {
	XMLName                               xml.Name `xml:"payments"`
	Name                                  string   `xml:"name"`
	LineCode                              string   `xml:"lineCode"`
	PaymentsTotal                         string   `xml:"paymentsTotal"`
	PaymentsTotalShare                    string   `xml:"paymentsTotalShare"`
	FinancialSupportSubsidiesStateTasks   string   `xml:"financialSupportSubsidiesStateTasks"`
	PaymentsTotalShare6                   string   `xml:"paymentsTotalShare6"`
	FinancialSupportSubsidiesOther        string   `xml:"financialSupportSubsidiesOther"`
	PaymentsTotalShare8                   string   `xml:"paymentsTotalShare8"`
	GrantSubsidiesStateBudget             string   `xml:"grantSubsidiesStateBudget"`
	PaymentsTotalShare10                  string   `xml:"paymentsTotalShare10"`
	GrantSubsidiesStateBudgetSubject      string   `xml:"grantSubsidiesStateBudgetSubject"`
	PaymentsTotalShare12                  string   `xml:"paymentsTotalShare12"`
	FinancialSupportOms                   string   `xml:"financialSupportOms"`
	PaymentsTotalShare14                  string   `xml:"paymentsTotalShare14"`
	FinancialSupportIncomeActivitiesTotal string   `xml:"financialSupportIncomeActivitiesTotal"`
	PaymentsTotalShare16                  string   `xml:"paymentsTotalShare16"`
	FundsIncome                           string   `xml:"fundsIncome"`
	PaymentsTotalShare18                  string   `xml:"paymentsTotalShare18"`
	FundsIncomeGratuitousReceipts         string   `xml:"fundsIncomeGratuitousReceipts"`
	PaymentsTotalShare20                  string   `xml:"paymentsTotalShare20"`
}

// Сведения об оказываемых услугах, выполняемых работах сверх установленного
// государственного (муниципального) задания, а также выпускаемой продукции»
type NonFinancialAssetsChange struct {
	XMLName xml.Name `xml:"nonFinancialAssetsChange"`
}

// Сведения о кредиторской задолженности и обязательствах учреждения
type CreditPayment struct {
	XMLName                                    xml.Name `xml:"creditPayment"`
	Name                                       string   `xml:"name"`
	LineCode                                   string   `xml:"lineCode"`
	CreditPaymentStartYearTotal                string   `xml:"creditPaymentStartYearTotal"`
	CreditPaymentStartYearDedlineReportingYear string   `xml:"creditPaymentStartYearDedlineReportingYear"`
	CreditPaymentEndYearTotal                  string   `xml:"creditPaymentEndYearTotal"`
	CreditPaymentEndYearTotalQuarter1          string   `xml:"creditPaymentEndYearTotalQuarter1"`
	CreditPaymentEndYearTotalJanuary           string   `xml:"creditPaymentEndYearTotalJanuary"`
	CreditPaymentEndYearTotalQuarter2          string   `xml:"creditPaymentEndYearTotalQuarter2"`
	CreditPaymentEndYearTotalQuarter3          string   `xml:"creditPaymentEndYearTotalQuarter3"`
	CreditPaymentEndYearTotalQuarter4          string   `xml:"creditPaymentEndYearTotalQuarter4"`
	CreditPaymentEndYearTotalNextYear          string   `xml:"creditPaymentEndYearTotalNextYear"`
	AmountDeferredTotal                        string   `xml:"amountDeferredTotal"`
	AmountDeferredSalary                       string   `xml:"amountDeferredSalary"`
	AmountDeferredClaims                       string   `xml:"amountDeferredClaims"`
	AmountDeferredNotReceived                  string   `xml:"amountDeferredNotReceived"`
	AmountDeferredOther                        string   `xml:"amountDeferredOther"`
}

//  Сведения о численности сотрудников и оплате
type NumberEmployeesRemuneration struct {
	XMLName                           xml.Name `xml:"numberEmployeesRemuneration"`
	NumberEmployeesRemunerationGroups []NumberEmployeesRemunerationGroups
}

// Сведения о численности и оплате труда по
type NumberEmployeesRemunerationGroups struct {
	XMLName    xml.Name `xml:"numberEmployeesRemunerationGroups"`
	GroupStaff GroupStaff
	// NumberEmployeesRemunerationGroup []NumberEmployeesRemunerationGroup
	Employees Employees
	Salary    Salary
}

type GroupStaff struct {
	XMLName  xml.Name `xml:"groupStaff"`
	Name     string   `xml:"name"`
	LineCode string   `xml:"lineCode"`
}

//ведения о численности и оплате труда по
// группам сотрудников
type NumberEmployeesRemunerationGroup struct {
	XMLName    xml.Name `xml:"numberEmployeesRemunerationGroup"`
	Name       string   `xml:"name"`
	GroupStaff GroupStaff
	Employees  Employees
	Salary     Salary
}

type Employees struct {
	XMLName           xml.Name `xml:"employees"`
	StaffingStartYear StaffingStartYear
	AverageOfYear     AverageOfYear
	Contracts         Contracts
	StaffingEndYear   StaffingEndYear
}

type StaffingStartYear struct {
	XMLName       xml.Name `xml:"staffingStartYear"`
	StaffingTotal string   `xml:"staffingTotal"`
	StaffingBase  string   `xml:"staffingBase"`
	Replaced      string   `xml:"replaced"`
	Vacancy       string   `xml:"vacancy"`
}

type AverageOfYear struct {
	XMLName         xml.Name `xml:"averageOfYear"`
	Total           string   `xml:"total"`
	Base            string   `xml:"base"`
	MainPlace       string   `xml:"mainPlace"`
	InsidePartTme   string   `xml:"insidePartTme"`
	OutsidePartTime string   `xml:"outsidePartTime"`
}

type Contracts struct {
	XMLName      xml.Name `xml:"contracts"`
	Employees    string   `xml:"employees"`
	NotEmployees string   `xml:"notEmployees"`
}

type StaffingEndYear struct {
	XMLName       xml.Name `xml:"staffingEndYear"`
	StaffingTotal string   `xml:"staffingTotal"`
	StaffingBase  string   `xml:"staffingBase"`
	Replaced      string   `xml:"replaced"`
	Vacancy       string   `xml:"vacancy"`
}

type Salary struct {
	XMLName        xml.Name `xml:"salary"`
	SalaryFund     SalaryFund
	Accrued        Accrued
	SalaryAnalytic SalaryAnalytic
}

type SalaryFund struct {
	XMLName         xml.Name `xml:"salaryFund"`
	Total           string   `xml:"total"`
	Base            string   `xml:"base"`
	BaseFullTime    string   `xml:"baseFullTime"`
	MainPartTme     string   `xml:"mainPartTme"`
	InsideCombining string   `xml:"insideCombining"`
	Outside         string   `xml:"outside"`
}

type Accrued struct {
	XMLName      xml.Name `xml:"accrued"`
	Employees    string   `xml:"employees"`
	NotEmployees string   `xml:"notEmployees"`
}

type SalaryAnalytic struct {
	XMLName              xml.Name `xml:"salaryAnalytic"`
	BaseWorkplace        BaseWorkplace
	InsideCombining      InsideCombining
	OutsideWorkplace     OutsideWorkplace
	EmployeesContract    EmployeesContract
	NotEmployeesContract NotEmployeesContract
}

type BaseWorkplace struct {
	XMLName                xml.Name `xml:"baseWorkplace"`
	SubsidiesContract      string   `xml:"subsidiesContract"`
	SubsidiesOther         string   `xml:"subsidiesOther"`
	SubsidiesGrantFederal  string   `xml:"subsidiesGrantFederal"`
	SubsidiesGrantRegional string   `xml:"subsidiesGrantRegional"`
	Oms                    string   `xml:"oms"`
	IncomeActivities       string   `xml:"incomeActivities"`
}

type InsideCombining struct {
	XMLName                xml.Name `xml:"insideCombining"`
	SubsidiesContract      string   `xml:"subsidiesContract"`
	SubsidiesOther         string   `xml:"subsidiesOther"`
	SubsidiesGrantFederal  string   `xml:"subsidiesGrantFederal"`
	SubsidiesGrantRegional string   `xml:"subsidiesGrantRegional"`
	Oms                    string   `xml:"oms"`
	IncomeActivities       string   `xml:"incomeActivities"`
}

type OutsideWorkplace struct {
	XMLName                xml.Name `xml:"outsideWorkplace"`
	SubsidiesContract      string   `xml:"subsidiesContract"`
	SubsidiesOther         string   `xml:"subsidiesOther"`
	SubsidiesGrantFederal  string   `xml:"subsidiesGrantFederal"`
	SubsidiesGrantRegional string   `xml:"subsidiesGrantRegional"`
	Oms                    string   `xml:"oms"`
	IncomeActivities       string   `xml:"incomeActivities"`
}

type EmployeesContract struct {
	XMLName                xml.Name `xml:"employeesContract"`
	SubsidiesContract      string   `xml:"subsidiesContract"`
	SubsidiesOther         string   `xml:"subsidiesOther"`
	SubsidiesGrantFederal  string   `xml:"subsidiesGrantFederal"`
	SubsidiesGrantRegional string   `xml:"subsidiesGrantRegional"`
	Oms                    string   `xml:"oms"`
	IncomeActivities       string   `xml:"incomeActivities"`
}

type NotEmployeesContract struct {
	XMLName                xml.Name `xml:"notEmployeesContract"`
	SubsidiesContract      string   `xml:"subsidiesContract"`
	SubsidiesOther         string   `xml:"subsidiesOther"`
	SubsidiesGrantFederal  string   `xml:"subsidiesGrantFederal"`
	SubsidiesGrantRegional string   `xml:"subsidiesGrantRegional"`
	Oms                    string   `xml:"oms"`
	IncomeActivities       string   `xml:"incomeActivities"`
}

type OpenedCreditAccounts struct {
	XMLName                xml.Name `xml:"openedCreditAccounts"`
	CreditAccountsRF       string   `xml:"CreditAccountsRF"`
	CreditAccountsCurrency string   `xml:"CreditAccountsCurrency"`
}

type CreditAccountsRF struct {
	XMLName xml.Name `xml:"CreditAccountsRF"`
}
type CreditAccountsCurrency struct {
	XMLName xml.Name `xml:"CreditAccountsRF"`
}

// сведения о имуществе
type AssetsUse struct {
	XMLName                 xml.Name `xml:"assetsUse"`
	Xmlns                   string   `xml:"xmlns,attr"`
	EstateExceptLand        EstateExceptLand
	LandPermanentUse        LandPermanentUse
	ValuableMovableProperty ValuableMovableProperty
	Vehicles                Vehicles
}

// Сведения о недвижимом имуществе, за исключение м земельных  участков,
// закрепленно м на праве оперативног о управления
type EstateExceptLand struct {
	XMLName                 xml.Name `xml:"estateExceptLand"`
	EstateExceptLandObjects EstateExceptLandObjects
}
type EstateExceptLandObjects struct {
	XMLName      xml.Name `xml:"estateExceptLandObjects"`
	TypeObject   TypeObject
	EstateObject []EstateObject
}
type TypeObject struct {
	XMLName  xml.Name `xml:"typeObject"`
	Type     string   `xml:"type"`
	LineCode string   `xml:"lineCode"`
}

type EstateObject struct {
	XMLName          xml.Name `xml:"estateObject"`
	Name             string   `xml:"name"`
	Address          string   `xml:"address"`
	CadNumber        string   `xml:"cadNumber"`
	Oktmo            Oktmo
	UniqueObjectCode string `xml:"uniqueObjectCode"`
	BuildingYear     string `xml:"buildingYear"`
	Unit             Unit
	TypeObject       TypeObject
	Used             Used
	ActualExpenses   ActualExpenses
}
type Oktmo struct {
	XMLName xml.Name `xml:"oktmo"`
	Code    string   `xml:"code"`
	Name    string   `xml:"name"`
}
type Unit struct {
	XMLName xml.Name `xml:"unit"`
	Code    string   `xml:"code"`
	Symbol  string   `xml:"symbol"`
}

type Used struct {
	XMLName             xml.Name `xml:"used"`
	Total               string   `xml:"total"`
	MainPurposeTask     string   `xml:"mainPurposeTask"`
	MainPurposeOverTask string   `xml:"mainPurposeOverTask"`
	OtherPurpose        string   `xml:"otherPurpose"`
}

type ActualExpenses struct {
	XMLName                     xml.Name `xml:"actualExpenses"`
	Utilities                   Utilities
	PropertyMaintenanceServices PropertyMaintenanceServices
	PropertyTax                 PropertyTax
	Total                       string `xml:"total"`
}
type Utilities struct {
	XMLName           xml.Name `xml:"utilities"`
	Total             string   `xml:"total"`
	ReimbursedByUsers string   `xml:"reimbursedByUsers"`
	UnusedProperty    string   `xml:"unusedProperty"`
}
type PropertyMaintenanceServices struct {
	XMLName           xml.Name `xml:"propertyMaintenanceServices"`
	Total             string   `xml:"total"`
	ReimbursedByUsers string   `xml:"reimbursedByUsers"`
	UnusedProperty    string   `xml:"unusedProperty"`
}
type PropertyTax struct {
	XMLName           xml.Name `xml:"propertyTax"`
	Total             string   `xml:"total"`
	ReimbursedByUsers string   `xml:"reimbursedByUsers"`
	UnusedProperty    string   `xml:"unusedProperty"`
}

// Сведения о земельных участках, предоставленных на праве постоянного
// (бессрочного ) пользования
type LandPermanentUse struct {
	XMLName       xml.Name `xml:"landPermanentUse"`
	LendObject    []LendObject
	LendObjectUse LendObjectUse
}
type LendObject struct {
	XMLName       xml.Name `xml:"lendObject"`
	Name          string   `xml:"name"`
	Address       string   `xml:"address"`
	Oktmo         Oktmo
	CadNumber     string `xml:"cadNumber"`
	Unit          Unit
	LendObjectUse LendObjectUse
}
type LendObjectUse struct {
	XMLName         xml.Name `xml:"lendObjectUse"`
	LineCode        string   `xml:"lineCode"`
	Total           string   `xml:"total"`
	Used            UsedL
	UsedByEasement  string `xml:"usedByEasement"`
	NotUsed         NotUsed
	ExpensesForLand ExpensesForLand
}

type UsedL struct {
	XMLName             xml.Name `xml:"used"`
	Total               string   `xml:"total"`
	MainPurposeTask     string   `xml:"mainPurposeTask"`
	MainPurposeOverTask string   `xml:"mainPurposeOverTask"`
	UsedOther           string   `xml:"usedOther"`
}

type NotUsed struct {
	XMLName                  xml.Name `xml:"notUsed"`
	Total                    string   `xml:"total"`
	TemporaryUsedRent        string   `xml:"temporaryUsedRent"`
	TemporaryUsedFree        string   `xml:"temporaryUsedFree"`
	TemporaryUsedWithoutRigt string   `xml:"temporaryUsedWithoutRigt"`
	TemporaryUsedOther       string   `xml:"temporaryUsedOther"`
}
type ExpensesForLand struct {
	XMLName     xml.Name `xml:"expensesForLand"`
	Total       string   `xml:"total"`
	CostsTotal  string   `xml:"costsTotal"`
	CostsRefund string   `xml:"costsRefund"`
	LandTax     string   `xml:"landTax"`
}

// Сведения об особо ценном движимом имуществе
// (за исключение м транспортны х средств)
type ValuableMovableProperty struct {
	XMLName                        xml.Name `xml:"valuableMovableProperty"`
	ValuableMovablePropertyObjects ValuableMovablePropertyObjects
}

type TypeMovableObject struct {
	XMLName  xml.Name `xml:"typeMovableObject"`
	Type     string   `xml:"type"`
	LineCode string   `xml:"lineCode"`
}

type ValuableMovablePropertyObject struct {
	XMLName               xml.Name `xml:"valuableMovablePropertyObject"`
	Name                  string   `xml:"name"`
	TypeMovableObject     TypeMovableObject
	AvailabilityEndPeriod AvailabilityEndPeriod
	ActualTermOfUse       ActualTermOfUse
	// ResidualValue         ResidualValue
}

type AvailabilityEndPeriod struct {
	XMLName                xml.Name `xml:"availabilityEndPeriod"`
	Total                  string   `xml:"total"`
	UsedByAgency           string   `xml:"usedByAgency"`
	TransferredForUseTotal string   `xml:"transferredForUseTotal"`
	ForRent                string   `xml:"forRent"`
	ForFree                string   `xml:"forFree"`
	NeedRepair             string   `xml:"needRepair"`
	WaitWriteOffTotal      string   `xml:"waitWriteOffTotal"`
	WaitWriteOffReplace    string   `xml:"waitWriteOffReplace"`
}
type ActualTermOfUse struct {
	XMLName    xml.Name `xml:"actualTermOfUse"`
	ActualTerm string   `xml:"actualTerm"`
	Count      string   `xml:"count"`
	Cost       string   `xml:"cost"`
}
type ResidualValue struct {
	XMLName   xml.Name `xml:"residualValue"`
	ValueTerm string   `xml:"ValueTerm"`
	Cost      string   `xml:"cost"`
}

// ==============================
// Транспортные средства
// ==============================

type Vehicles struct {
	XMLName                 xml.Name `xml:"vehicles"`
	UsedVehicles            UsedVehicles
	NotUsedVehicles         NotUsedVehicles
	DirectionOfUse          DirectionOfUse
	CostMaintenanceVehicles CostMaintenanceVehicles
}

type UsedVehicles struct {
	XMLName                xml.Name `xml:"usedVehicles"`
	UsedVehiclesTransports UsedVehiclesTransports
}

type UsedVehiclesTransports struct {
	XMLName         xml.Name `xml:"usedVehiclesTransports"`
	TransportsType  TransportsType
	VehiclesDetails VehiclesDetails
}

type TransportsType struct {
	XMLName    xml.Name `xml:"TransportsType"`
	TranspType string   `xml:"type"`
	LineCode   string   `xml:"lineCode"`
}
type VehiclesDetails struct {
	XMLName            xml.Name `xml:"vehiclesDetails"`
	Name               string   `xml:"name"`
	TransportsType     TransportsType
	Total              Total
	OperationManagment OperationManagment
}

type Total struct {
	XMLName      xml.Name `xml:"total"`
	EndDate      string   `xml:"endDate"`
	MiddleOfYear string   `xml:"middleOfYear"`
}

type OperationManagment struct {
	XMLName      xml.Name `xml:"operationManagment"`
	EndDate      string   `xml:"endDate"`
	MiddleOfYear string   `xml:"middleOfYear"`
}

type NotUsedVehicles struct {
	XMLName                   xml.Name `xml:"notUsedVehicles"`
	NotUsedVehiclesTransports NotUsedVehiclesTransports
}
type NotUsedVehiclesTransports struct {
	XMLName        xml.Name `xml:"notUsedVehiclesTransports"`
	TransportsType TransportsType
}

type DirectionOfUse struct {
	XMLName                  xml.Name `xml:"directionOfUse"`
	DirectionOfUseTransports DirectionOfUseTransports
}
type DirectionOfUseTransports struct {
	XMLName                 xml.Name `xml:"directionOfUseTransports"`
	TransportsType          TransportsType
	DirectionOfUseTransport DirectionOfUseTransport
}
type DirectionOfUseTransport struct {
	XMLName              xml.Name `xml:"directionOfUseTransport"`
	Name                 string   `xml:"name"`
	TransportsType       TransportsType
	DirectlyUsedVehicles DirectlyUsedVehicles
	// PurposeServicingPersonnel PurposeServicingPersonnel
}

type DirectlyUsedVehicles struct {
	XMLName            xml.Name `xml:"directlyUsedVehicles"`
	Total              Total
	OperationManagment OperationManagment
	UnderLease         UnderLease
	UnderGratuitous    UnderGratuitous
}

type UnderLease struct {
	XMLName      xml.Name `xml:"underLease"`
	EndDate      string   `xml:"endDate"`
	MiddleOfYear string   `xml:"middleOfYear"`
}
type UnderGratuitous struct {
	XMLName      xml.Name `xml:"underGratuitous"`
	EndDate      string   `xml:"endDate"`
	MiddleOfYear string   `xml:"middleOfYear"`
}

type CostMaintenanceVehicles struct {
	XMLName                           xml.Name `xml:"costMaintenanceVehicles"`
	CostMaintenanceVehiclesTransports CostMaintenanceVehiclesTransports
}
type CostMaintenanceVehiclesTransports struct {
	XMLName                          xml.Name `xml:"costMaintenanceVehiclesTransports"`
	TransportsType                   TransportsType
	CostMaintenanceVehiclesTransport CostMaintenanceVehiclesTransport
}

type CostMaintenanceVehiclesTransport struct {
	XMLName                     xml.Name `xml:"costMaintenanceVehiclesTransport"`
	Name                        string   `xml:"name"`
	TransportsType              TransportsType
	VehiclesExpensesMaintenance VehiclesExpensesMaintenance
}
type VehiclesExpensesMaintenance struct {
	XMLName                 xml.Name `xml:"vehiclesExpensesMaintenance"`
	PeriodTotal             string   `xml:"periodTotal"`
	TransportTax            string   `xml:"transportTax"`
	FuelCosts               string   `xml:"fuelCosts"`
	WheelCosts              string   `xml:"wheelCosts"`
	OsagoCosts              string   `xml:"osagoCosts"`
	VolunteerInsuranceCosts string   `xml:"volunteerInsuranceCosts"`
	RepairsCosts            string   `xml:"repairsCosts"`
	MaintenanceCosts        string   `xml:"maintenanceCosts"`
	GaragesRent             string   `xml:"garagesRent"`
	GaragesMaintenance      string   `xml:"garagesMaintenance"`
	Drivers                 string   `xml:"drivers"`
	ServisesPersonnel       string   `xml:"servisesPersonnel"`
	AdministrativePersonnel string   `xml:"administrativePersonnel"`
}

// ==============================
// ==============================
// ==============================
type ValuableMovablePropertyObjects struct {
	XMLName                       xml.Name `xml:"valuableMovablePropertyObjects"`
	TypeMovableObject             TypeMovableObject
	ValuableMovablePropertyObject []ValuableMovablePropertyObject
	// AvailabilityEndPeriod         AvailabilityEndPeriod
	// ActualTermOfUse               ActualTermOfUse
	// ResidualValue                 ResidualValue
}

type EffectiveActivity struct {
	XMLName xml.Name `xml:"effectiveActivity"`
	Xmlns   string   `xml:"xmlns,attr"`
}
