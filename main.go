package main

import (
	//"context"
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	//"os"
	"strconv"
	//"time"
)

/*type DeviceMethod struct {
	Method string `xml:"method,type"`
}
type DeviceID struct {
	ID string `json:"id"`
}
type ObisCode struct {
	ObisCode string `json:"obiscode"`
}
type DeviceType struct {
    DeviceType string `json:"devicetype"`
}*/
type Measurement struct {
	ID       string   `json:"id"`
	ObisCode string   `json:"obiscode"`
	DeviceType     string   `json:"devicetype"`
}

var measurements []Measurement

//CREATE
func createMeasurement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var measurement Measurement
	_ = json.NewDecoder(r.Body).Decode(&measurement)
	measurement.ID = strconv.Itoa(rand.Intn(10000000))
	measurements = append(measurements, measurement)
	json.NewEncoder(w).Encode(measurement)
}

//READ
/*func getDeviceMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/xml")
	xml.NewEncoder(w).Encode(measurements)
}*/

func readMeasurements(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurements)
}

func readMeasurement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, device := range measurements {
		if device.ID == params["id"]{
			json.NewEncoder(w).Encode(device)
			return
		}
	}
	json.NewEncoder(w).Encode(&Measurement{})
}
	
//UPDATE
func updateMeasurement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/xml")
	params := mux.Vars(r)
	for index, device := range measurements {
		if device.ID == params["id"] {
			measurements = append(measurements[:index], measurements[index+1:]...)
			var measurement Measurement
			_ = json.NewDecoder(r.Body).Decode(&measurement)
			measurement.ID = params["id"]
			measurements = append(measurements, measurement)
			json.NewEncoder(w).Encode(measurement)
			return
		}
	}
}

//DELETE
func deleteMeasurement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	params := mux.Vars(r)
	for index, device := range measurements {
		if device.ID == params["id"] {
			measurements = append(measurements[:index], measurements[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(measurements)
}

/*func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
}*/

func main() {
	
	//Context generator
	/*start := time.Now()
	gen := func(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()*/

/*	file, err := os.Create("/home/sbuddappagari/sources/src/test/result.csv")
	log.Fatal("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
*/
	//Context deadline
/*	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	
*/	
	/*for n := range gen(ctx) {
		//err := writer.Write(n)
		//log.Fatal("Cannot write to file", err)
		
		fmt.Println(n)
		if n == 10 {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("context generator took %s", elapsed)*/

	//REST api
	r := mux.NewRouter()

	measurements = append(measurements, Measurement{ID: "780934", ObisCode: "19284918", Type: "Electric"})
	measurements = append(measurements, Measurement{ID: "123488", ObisCode: "20957022", Type: "Thermal"})

	r.HandleFunc("/measurements", readMeasurements).Methods("GET")
	r.HandleFunc("/measurements", createMeasurement).Methods("POST")
	r.HandleFunc("/measurements/{id}", readMeasurement).Methods("GET")
	r.HandleFunc("/measurements/{id}", updateMeasurement).Methods("PUT")
	r.HandleFunc("/measurements/{id}", deleteMeasurement).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

	//Check status of weblink

	/*links := []string{
		"https://techdoc.cuculus.net",
		"https://jira.cuculus.net",
	}
	
	c := make(chan string)
	
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}*/

}
