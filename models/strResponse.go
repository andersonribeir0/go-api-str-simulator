package models

type StrResponse struct {
	ID              string `json:"id,omitempty"`
	Source          string `json:"sistemaOrigem,omitempty"`
	Event           string `json:"tipoEvento,omitempty"`
	DateTime        string `json:"dataHoraResposta,omitempty"`
	CompanyNumber   int    `json:"numEmpresa,omitempty"`
	OperationID1    string `json:"operacaoOrigemId1,omitempty"`
	OperationID2    string `json:"operacaoOrigemId2,omitempty"`
	TransactionType string `json:"tipoTransacaoOrigem,omitempty"`
	Status          int    `json:"situacaoNm,omitempty"`
	StatusThrow     string `json:"situacaoLancamento,omitempty"`
	Domain          string `json:"dominioSistema,omitempty"`
	Message         string `json:"message,omitempty"`
}
