package fhd

import "encoding/xml"

// FinancialActivityPlan2020 представляет корневой элемент XML
type FinancialActivityPlan2020 struct {
	XMLName  xml.Name `xml:"financialActivityPlan2020"`
	Xmlns    string   `xml:"xmlns,attr"`
	XmlnsNs2 string   `xml:"xmlns:ns2,attr"`
	XmlnsNs3 string   `xml:"xmlns:ns3,attr"`
	Header   Header   `xml:"ns2:header"`
	Body     Body     `xml:"body"`
}

// Header представляет ns2:header
type Header struct {
	ID             string `xml:"ns2:id"`
	CreateDateTime string `xml:"ns2:createDateTime"`
}

// Body представляет тело XML
type Body struct {
	Position Position `xml:"position"`
}

// Position представляет ns2:position
type Position struct {
	PositionID       string             `xml:"ns2:positionId"`
	ChangeDate       string             `xml:"ns2:changeDate"`
	Placer           Placer             `xml:"ns2:placer"`
	Initiator        Initiator          `xml:"ns2:initiator"`
	VersionNumber    string             `xml:"ns2:versionNumber"`
	FinancialYear    string             `xml:"ns3:financialYear"`
	GeneralData      GeneralData        `xml:"ns3:generalData"`
	PlanPaymentIndex []PlanPaymentIndex `xml:"ns3:planPaymentIndex"`
	PlanPaymentTRU   []PlanPaymentTRU   `xml:"ns3:planPaymentTRU"`
}

// Placer представляет ns2:placer
type Placer struct {
	RegNum   string `xml:"ns2:regNum"`
	FullName string `xml:"ns2:fullName"`
	INN      string `xml:"ns2:inn"`
	KPP      string `xml:"ns2:kpp"`
}

// Initiator представляет ns2:initiator
type Initiator struct {
	RegNum   string `xml:"ns2:regNum"`
	FullName string `xml:"ns2:fullName"`
	INN      string `xml:"ns2:inn"`
	KPP      string `xml:"ns2:kpp"`
}

// GeneralData представляет ns3:generalData
type GeneralData struct {
	Date                   string           `xml:"ns3:date"`
	DateApprovel           string           `xml:"ns3:dateApprovel"`
	FounderAuthority       FounderAuthority `xml:"ns3:founderAuthority"`
	OKEI                   OKEI             `xml:"ns3:okei"`
	ManagerName            string           `xml:"ns3:managerName"`
	ManagerPosition        string           `xml:"ns3:managerPosition"`
	ExecutorName           string           `xml:"ns3:executorName"`
	ExecutorPosition       string           `xml:"ns3:executorPosition"`
	Phone                  string           `xml:"ns3:phone"`
	SignDate               string           `xml:"ns3:signDate"`
	FounderManagerName     string           `xml:"ns3:founderManagerName"`
	FounderManagerPosition string           `xml:"ns3:founderManagerPosition"`
	FounderSignDate        string           `xml:"ns3:founderSignDate"`
}

// FounderAuthority представляет ns3:founderAuthority
type FounderAuthority struct {
	RegNum   string `xml:"ns2:regNum"`
	FullName string `xml:"ns2:fullName"`
	INN      string `xml:"ns2:inn"`
	KPP      string `xml:"ns2:kpp"`
}

// OKEI представляет ns3:okei
type OKEI struct {
	Code   string `xml:"ns2:code"`
	Symbol string `xml:"ns2:symbol"`
}

// PlanPaymentIndex представляет ns3:planPaymentIndex
type PlanPaymentIndex struct {
	Name     string `xml:"ns3:name"`
	LineCode string `xml:"ns3:lineCode"`
	KBK      string `xml:"ns3:kbk,omitempty"`
	Sum      Sum    `xml:"ns3:sum"`
}

// PlanPaymentTRU представляет ns3:planPaymentTRU
type PlanPaymentTRU struct {
	Name      string `xml:"ns3:name"`
	LineCode  string `xml:"ns3:lineCode"`
	YearStart string `xml:"ns3:yearStart,omitempty"`
	Sum       Sum    `xml:"ns3:sum"`
}

// Sum представляет ns3:sum
type Sum struct {
	FinancialYearSum string `xml:"ns3:financialYearSum"`
	PlanFirstYearSum string `xml:"ns3:planFirstYearSum"`
	PlanLastYearSum  string `xml:"ns3:planLastYearSum"`
}
