package devicebundle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/domgoodwin/go-api/app/bundles/prime"

	"github.com/gorilla/mux"

	"github.com/domgoodwin/go-api/app/bundles/bitbucket"
	"github.com/domgoodwin/go-api/app/bundles/common"
	"github.com/domgoodwin/go-api/app/bundles/db"
	"github.com/domgoodwin/go-api/app/bundles/r53"
	"github.com/domgoodwin/go-api/app/bundles/wait"
)

// DeviceController struct
type DeviceController struct {
	common.Controller
}

// Index func return all devices in database
func (c *DeviceController) Index(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		[]*Device{NewDevice("host", "192.168.0.1", "80.80.80.80")},
		http.StatusOK,
	)
}

// Create func return all devices in database
func (c *DeviceController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(json.NewDecoder(r.Body))
	res := db.PutItem("home-devices", NewDevice("hello", "world", "indeed"))
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}

// NextPrime gets next prime number
func (c *DeviceController) NextPrime(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	curs, ok := vals["cur"]
	cur, err := strconv.ParseInt(curs[0], 10, 64)
	fmt.Println("Getting next prime on from: " + strconv.Itoa(int(cur)))
	res := int64(0)
	if ok && err == nil {
		res = prime.GetNextPrime(cur)
	}
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}

// GetRandom gets next prime number
func (c *DeviceController) Wait(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	dur, ok := vals["time"]
	res := int64(0)
	if !ok || len(dur[0]) < 1 {
		fmt.Println("Waiting random time 0<X<30")
		res = wait.WaitDurationRandom()
	} else {
		duration, err := strconv.ParseInt(dur[0], 10, 64)
		if err == nil {
			fmt.Println("Waiting fixed time of : " + strconv.Itoa(int(duration)))
			res = wait.WaitDurationFixed(duration)
		}
	}
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}

// GetRecords returns records sets of supplied id
func (c *DeviceController) GetRecords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := r53.GetRecordSets(vars["id"])
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}

// UpdateRecordSet sets record values and type of supplied domain
func (c *DeviceController) UpdateRecordSet(w http.ResponseWriter, r *http.Request) {
	var rs r53.RecordSet
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	println(rs.HostedZoneId)
	res := r53.UpdateRecordSet(rs)
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}

// HandlePROpen sets record values and type of supplied domain
func (c *DeviceController) HandlePROpen(w http.ResponseWriter, r *http.Request) {
	var body interface{}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(body)
	res := bitbucket.HandlePROpen("aaa")
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}
