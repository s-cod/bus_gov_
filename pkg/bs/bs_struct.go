package bs

import "encoding/xml"

type BudgetaryCircumstances struct {
	XMLName  xml.Name `xml:"budgetaryCircumstances"`
	Xmlns    string   `xml:"xmlns,attr"`
	XmlnsNs2 string   `xml:"xmlns:ns2,attr"`
	Header   Header   `xml:"ns2:header"`
	Body     Body     `xml:"body"`
}

type Header struct {
	ID             string `xml:"ns2:id"`
	CreateDateTime string `xml:"ns2:createDateTime"`
}

type Body struct {
	Position Position `xml:"position"`
}

type Position struct {
	PositionID    string                  `xml:"ns2:positionId"`
	ChangeDate    string                  `xml:"ns2:changeDate"`
	VersionNumber string                  `xml:"ns2:versionNumber"`
	FinancialYear string                  `xml:"ns2:financialYear"`
	ConfirmDate   string                  `xml:"ns2:confirmDate"`
	Section       string                  `xml:"ns2:section"`
	Circumstances []BudgetaryCircumstance `xml:"ns2:budgetaryCircumstance"`
}

type BudgetaryCircumstance struct {
	KbkBudget    KbkBudget `xml:"ns2:kbkBudget"`
	Circumstance string    `xml:"ns2:circumstance"`
}

type KbkBudget struct {
	Code   string `xml:"ns2:code"`
	Name   string `xml:"ns2:name"`
	Budget Budget `xml:"ns2:budget"`
}

type Budget struct {
	Code string `xml:"ns2:code"`
	Name string `xml:"ns2:name"`
}
