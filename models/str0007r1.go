package models

type STR0007R1 struct {
	Code       string `xml:"CodMsg,omitempty"`
	IFNumber   string `xml:"NumCtrlIF,omitempty"`
	SourceISPB string `xml:"ISPBIFDebtd,omitempty"`
	StrCode    string `xml:"NumCtrlSTR,omitempty"`
	StrStatus  string `xml:"SitLancSTR,omitempty"`
	MovtoDate  string `xml:"DtMovto,omitempty"`
}
