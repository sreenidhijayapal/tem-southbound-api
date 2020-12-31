package main

import (
	"fmt"
	"math/rand"
	"time"
	//"os"
)

/*var pool = "$%&"

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomStringByte(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(1001, 9999))
	}
	return string(bytes)
}

func randomStringPool(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}*/

func main() {
	var devices_num int
	var measurements_num int
	var measurement_terrupt int
	var days_num int
	var obis_num int
	
	devices := []string{
		"Electric",
		"Thermal",
		"Gas",
		"Water",
	}

	measurements := []string{
		"Berlin",
		"Paris",
		"Barcelona",
		"Lisbon",
	}

	obiss := []int{
		rand.Intn(99),
		rand.Intn(99),
		rand.Intn(255),
	}
	
	fmt.Print("Number of devices: ")
	_, err1 := fmt.Scanf("%d", &devices_num)
	if err1 != nil {
		fmt.Println(err1)
	}
	
	fmt.Print("Number of measurements: ")
	_, err2 := fmt.Scanf("%d", &measurements_num)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Print("Measurement intervals(in minutes):")
	_, err3 := fmt.Scanf("%d", &measurement_terrupt)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Print("Number of days:")
	_, err4 := fmt.Scanf("%d", &days_num)
	if err4 != nil {
		fmt.Println(err4)
	}

	fmt.Print("Number of OBIS codes:")
	_, err5 := fmt.Scanf("%d", &obis_num)
	if err5 != nil {
		fmt.Println(err5)
	}
	
	for i := 0; i < devices_num; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Println("Device number", rand.Intn(33))
		for i := 0; i < measurements_num*devices_num*(1440*days_num/measurement_terrupt); i++ {
			m := rand.Int() % len(devices)
			n := rand.Int() % len(measurements)
			o := rand.Int() % len(obiss)
			fmt.Println(devices[m], measurements[n], obiss[o], rand.Intn(99), "kWh")
		}
		/*fmt.Println(randomInt(10001, 90009))
		fmt.Println(randomStringByte(11))
		fmt.Println(randomStringPool(22))*/
		fmt.Println("\n")
	}

}
