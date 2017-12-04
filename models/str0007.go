package models

type STR0007 struct {
	Code                string `xml:"CodMsg,omitempty"`
	IFNumber            string `xml:"NumCtrlIF,omitempty"`
	SourceISPB          string `xml:"ISPBIFDebtd,omitempty"`
	SourcePersonType    string `xml:"TpPessoaRemet,omitempty"`
	SourceDocument      string `xml:"CNPJ_CPFRemet,omitempty"`
	SourceName          string `xml:"NomRemet,omitempty"`
	TargetISPB          string `xml:"ISPBIFCredtd,omitempty"`
	TargetBranch        string `xml:"AgCredtd,omitempty"`
	AccountType         string `xml:"TpCtCredtd,omitempty"`
	TargetAccount       string `xml:"CtCredtd,omitempty"`
	TargetPersonType    string `xml:"TpPessoaCredtd,omitempty"`
	TargetDocument      string `xml:"CNPJ_CPFCliCredtd,omitempty"`
	TargetClientName    string `xml:"NomCliCredtd,omitempty"`
	TargetControlNumber string `xml:"NumContrtoOpCred,omitempty"`
	Amount              string `xml:"VlrLanc,omitempty"`
	ReasonIF            string `xml:"FinlddIF,omitempty"`
	IDTransferCode      string `xml:"CodIdentdTransf,omitempty"`
	History             string `xml:"Hist,omitempty"`
	ScheduleDate        string `xml:"DtAgendt,omitempty"`
	ScheduleHour        string `xml:"HrAgendt,omitempty"`
	PreferenceLevel     string `xml:"NivelPref,omitempty"`
	MovtoDate           string `xml:"DtMovto,omitempty"`
}
