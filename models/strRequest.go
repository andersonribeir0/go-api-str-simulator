package models

type StrRequest struct {
	Source          string `json:"sistemaOrigem,omitempty"`
	Event           string `json:"tipoEvento,omitempty"`
	DateTime        string `json:"dataHoraAgendamento,omitempty"`
	CompanyNumber   string `json:"numEmpresa,omitempty"`
	Operator        string `json:"operador,omitempty"`
	OperationID1    string `json:"operacaoOrigemId1,omitempty"`
	OperationID2    string `json:"operacaoOrigemId2,omitempty"`
	TransactionType string `json:"tipoTransacaoOrigem,omitempty"`
	TargetISPB      string `json:"ispbDestinatario,omitempty"`
	Description     string `json:"descricao,omitempty"`
	Domain          string `json:"dominioSistema,omitempty"`
	Message         string `json:"message,omitempty"`
	URL             string `json:"url,omitempty"`
}
