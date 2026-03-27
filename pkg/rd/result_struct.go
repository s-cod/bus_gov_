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
	Id             string `xml:"id"`
	CreateDateTime string `xml:"createDateTime"`
}

type Body struct {
	Position Position
}
type Position struct {
	XMLName           xml.Name `xml:"position"`
	PositionId        string   `xml:"ns2:positionId"`
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
	NumberEmployeesRemuneration NumberEmployeesRemuneration
	OpenedCreditAccounts        OpenedCreditAccounts
}

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

type NonFinancialAssetsChange struct {
	XMLName xml.Name `xml:"nonFinancialAssetsChange"`
}

type NumberEmployeesRemuneration struct {
	XMLName                           xml.Name `xml:"numberEmployeesRemuneration"`
	NumberEmployeesRemunerationGroups NumberEmployeesRemunerationGroups
}

type NumberEmployeesRemunerationGroups struct {
	XMLName                          xml.Name `xml:"numberEmployeesRemunerationGroups"`
	GroupStaff                       GroupStaff
	NumberEmployeesRemunerationGroup []NumberEmployeesRemunerationGroup
	Employees                        string `xml:"employees"`
	Salary                           string `xml:"salary"`
}

type GroupStaff struct {
	XMLName  xml.Name `xml:"groupStaff"`
	Name     string   `xml:"name"`
	LineCode string   `xml:"lineCode"`
}

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

type AssetsUse struct {
	XMLName xml.Name `xml:"assetsUse"`
	Xmlns   string   `xml:"xmlns,attr"`
}

// Сведения об особо ценном движимом имуществе
// (за исключение м транспортны х средств)
type ValuableMovableProperty struct {
	XMLName                        xml.Name `xml:"valuableMovableProperty"`
	ValuableMovablePropertyObjects ValuableMovablePropertyObjects
	TypeMovableObject              TypeMovableObject
	ValuableMovablePropertyObject  []ValuableMovablePropertyObject
	AvailabilityEndPeriod          AvailabilityEndPeriod
	ActualTermOfUse                ActualTermOfUse
	ResidualValue                  ResidualValue
}

type TypeMovableObject struct {
	XMLName  xml.Name `xml:"typeMovableObject"`
	Type     string   `xml:"type"`
	LineCode string   `xml:"lineCode"`
}

type ValuableMovablePropertyObject struct {
	XMLName               xml.Name `xml:"valuableMovablePropertyObject"`
	Name                  string
	TypeMovableObject     TypeMovableObject
	AvailabilityEndPeriod AvailabilityEndPeriod
	ActualTermOfUse       ActualTermOfUse
	ResidualValue         ResidualValue
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

type ValuableMovablePropertyObjects struct {
	XMLName xml.Name `xml:"valuableMovablePropertyObjects"`
}

type EffectiveActivity struct {
	XMLName xml.Name `xml:"effectiveActivity"`
	Xmlns   string   `xml:"xmlns,attr"`
}
