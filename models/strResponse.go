package models

type StrResponse struct {
	ID              string `json:"id,omitempty"`
	Source          string `json:"sistemaOrigem,omitempty"`
	Event           string `json:"tipoEvento,omitempty"`
	DateTime        string `json:"dataHoraResposta,omitempty"`
	CompanyNumber   string `json:"numEmpresa,omitempty"`
	OperationID1    string `json:"operacaoOrigemId1,omitempty"`
	OperationID2    string `json:"operacaoOrigemId2,omitempty"`
	TransactionType string `json:"tipoTransacaoOrigem,omitempty"`
	TargetISPB      string `json:"situacaoNm,omitempty"`
	Status          string `json:"situacaoLancamento,omitempty"`
	Domain          string `json:"dominioSistema,omitempty"`
	Message         string `json:"message,omitempty"`
}
