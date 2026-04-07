package cs

import "encoding/xml"

type ActionGrant struct {
	XMLName  xml.Name `xml:"actionGrant"`
	Xmlns    string   `xml:"xmlns,attr"`
	XmlnsNs2 string   `xml:"xmlns:ns2,attr"`

	Header Header
	Body   Body
}

type Body struct {
	XMLName  xml.Name `xml:"body"`
	Position Position
}

type Header struct {
	XMLName        xml.Name `xml:"header"`
	Xmlns          string   `xml:"xmlns,attr"`
	ID             string   `xml:"id"`
	CreateDateTime string   `xml:"createDateTime"`
}

type Position struct {
	XMLName xml.Name `xml:"position"`

	PositionID         string `xml:"ns2:positionId"`
	ChangeDate         string `xml:"ns2:changeDate"`
	VersionNumber      string `xml:"ns2:versionNumber"`
	FinancialYear      string `xml:"ns2:financialYear"`
	ConfirmDate        string `xml:"ns2:confirmDate"`
	PlanBudgetaryFunds PlanBudgetaryFunds
	OtherGrantFunds    []OtherGrantFunds
}

type PlanBudgetaryFunds struct {
	XMLName           xml.Name `xml:"ns2:planBudgetaryFunds"`
	CapitalRealAssets string   `xml:"ns2:capitalRealAssets"`
	Total             string   `xml:"ns2:total"`
}
type OtherGrantFunds struct {
	XMLName xml.Name `xml:"ns2:otherGrantFunds"`
	Name    string   `xml:"ns2:name"`
	Funds   string   `xml:"ns2:funds"`
	Code    string   `xml:"ns2:code"`
	Kosgu   Kosgu
}
type Kosgu struct {
	XMLName xml.Name `xml:"ns2:kosgu"`
	Code    string   `xml:"ns2:code"`
}
