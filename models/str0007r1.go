package models

type ISPB struct {
	CodErro string `xml:"CodErro,attr"`
	Data    string `xml:",chardata"`
}
type STR0007R1 struct {
	Error       string `xml:"CodErro,attr,omitempty"`
	Code        string `xml:"CodMsg,omitempty"`
	IFNumber    string `xml:"NumCtrlIF,omitempty"`
	SourceISPB  ISPB   `xml:"ISPBIFDebtd,omitempty"`
	StrCode     string `xml:"NumCtrlSTR,omitempty"`
	StrStatus   string `xml:"SitLancSTR,omitempty"`
	SysDateTime string `xml:"DtHrSit,omitempty"`
	MovtoDate   string `xml:"DtMovto,omitempty"`
}
