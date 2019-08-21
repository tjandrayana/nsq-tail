package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
	"github.com/tjandrayana/nsq-tail/jsonapi"
	"github.com/urfave/negroni"
	grace "gopkg.in/paytm/grace.v1"
)

func main() {
	client := GetRoute()
	InitRoute(client)

	n := negroni.New()
	n.UseHandler(client)
	log.Fatal(grace.Serve(fmt.Sprintf(":%s", "8005"), n))
}

type NSQModule struct {
	LookUpURL     string `json:"lookup_url"`
	LookUpPort    string `json:"lookup_port"`
	Topic         string `json:"topic"`
	TotalMessages int64  `json:"total_messages"`
}

func GetRoute() *httprouter.Router {
	return httprouter.New()

}
func InitRoute(r *httprouter.Router) {
	r.POST("/get-nsq-messages", DoExecuteQuery)
}

func DoExecuteQuery(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var reqData NSQModule
	api := jsonapi.New(res, "")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	ms := reqData.ReadMessages()
	res.Write(ms)
}

func (m *NSQModule) ReadMessages() []byte {
	args := []string{
		"-lookupd-http-address",
		m.LookUpURL + ":" + m.LookUpPort,
		"-topic",
		m.Topic,
		"-n",
		fmt.Sprintf("%d", m.TotalMessages),
	}

	msg, err := exec.Command("nsq_tail", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %s\n", string(msg))

	return msg
}
