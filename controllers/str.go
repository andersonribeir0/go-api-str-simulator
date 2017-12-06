package controllers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
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

func init() {
	//log setup
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
}

//Str Trata requisições vindas simulando o sistema matera.
func (str StrController) Str(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var strRequest models.StrRequest
	var strResponse models.StrResponse
	var strRequestMessage models.STR0007
	var strResponseMessage models.STR0007R1

	//Response header
	w.Header().Add("content-type", "application/json")
	//decodeRequest
	err := json.NewDecoder(r.Body).Decode(&strRequest)
	log.Println("Incoming data: ", strRequest)
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
	strResponseMessage.SourceISPB.CodErro = "EGEN0017"
	strResponseMessage.SourceISPB.Data = strRequestMessage.SourceISPB
	strResponseMessage.StrStatus = "7"
	strResponseMessage.SysDateTime = time.Now().String()
	b, err := xml.Marshal(strResponseMessage)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v\n", err)
	}
	strResponse.Message = string(b)
	strResponse.CompanyNumber = strRequest.CompanyNumber
	strResponse.DateTime = time.Now().Format("20060102150405")
	strResponse.OperationID1 = strRequest.OperationID1
	strResponse.OperationID2 = strRequest.OperationID2
	strResponse.Status = 11
	strResponse.ID = strconv.Itoa(rand.Intn(99999))
	strResponse.Domain = "SPB01"
	strResponse.Event = strRequest.Event
	strResponse.StatusThrow = "5"
	strResponse.Source = strRequest.Source

	log.Println("XmlData strResponse: ", string(b))
	if err := json.NewEncoder(w).Encode(strResponse); err != nil {
		fmt.Println(reflect.TypeOf(err).String())
		fmt.Println(err)
	}

	//b, err = json.Marshal(strResponse)
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err = enc.Encode(&strResponse)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v\n", err)
	}

	//Postback
	if strRequest.URL != "" {
		var resp *http.Response
		log.Println("Sending postback to URL: ", strRequest.URL)
		log.Println("Sending data: ", buf.String())
		resp, err = http.Post(strRequest.URL, "application/json", bytes.NewBuffer(buf.Bytes()))
		if err != nil {
			log.Println(err)
		}
		log.Println("POSTBACK Status:", resp.Status)
	}
}
