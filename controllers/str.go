package controllers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/str/models"
)

//StrController Trata requisições do caju simulando comportamento do sistema matera.
type StrController struct{}

//NewStrController Cria uma instância da controller str
func NewStrController() *StrController {
	return &StrController{}
}

//Str Trata requisições vindas simulando o sistema matera.
func (str StrController) Str(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var strRequest models.StrRequest
	var strResponse models.StrResponse
	var strRequestMessage models.STR0007
	var strResponseMessage models.STR0007R1

	w.Header().Add("content-type", "application/json")
	//decodeRequest
	err := json.NewDecoder(r.Body).Decode(&strRequest)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error decoding str request. Error: %v\n", err)
	}
	err = xml.Unmarshal([]byte(strRequest.Message), &strRequestMessage)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error decoding xml: %v\n", err)
	}

	//encodeResponse
	strResponseMessage.StrCode = strconv.Itoa(rand.Intn(999999))
	strResponseMessage.Code = "STR0007R1"
	strResponseMessage.IFNumber = strRequestMessage.IFNumber
	strResponseMessage.MovtoDate = time.Now().Format("2006-01-02")
	strResponseMessage.SourceISPB = strRequestMessage.SourceISPB
	strResponseMessage.StrStatus = "1"
	b, err := xml.Marshal(strResponseMessage)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v\n", err)
	}
	strResponse.Message = string(b)
	strResponse.CompanyNumber = strRequest.CompanyNumber
	strResponse.DateTime = "2017-12-13 15:00:00"
	strResponse.OperationID1 = strRequest.OperationID1
	strResponse.OperationID2 = strRequest.OperationID2
	strResponse.Status = "7"
	strResponse.ID = strconv.Itoa(rand.Intn(99999))
	strResponse.Domain = "SPB01"

	json.NewEncoder(w).Encode(strResponse)
	b, err = json.Marshal(strResponse)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v\n", err)
	}

	//Postback
	if strRequest.URL != "" {
		_, err = http.Post(strRequest.URL, "application/json", bytes.NewBuffer(b))
		if err != nil {
			fmt.Fprintf(w, "Postback failed: %s\nErr:%v\n", string(b), err)
		}
	}
}
