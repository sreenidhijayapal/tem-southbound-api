package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"strconv"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == ""{
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	switch os := true; os {

	case ( m == "GET" && u == "/communication_topology") :
		communicationTopology(conn)
	case ( m == "GET" && u == "/connector_device_state"):
		connectorDeviceState(conn)
	case ( m == "GET" && u == "/connector_state"):
		connectorState(conn)
	case ( m == "GET" && u == "/device_parameter"):
		deviceParameter(conn)
	case ( m == "GET" && u == "/events"):
		events(conn)
	case ( m == "GET" && u == "/metered_data"):
		meteredData(conn)
	case ( m == "POST" && u == "/metered_data"):
		meteredDataValue(conn)
	case ( m == "GET" && u == "/tasks"):
		tasks(conn)
	}
}

func communicationTopology(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">><head><meta charset="UTF-8"><title></title></head><body>
        <strong>COMMUNICATION TOPOLOGY</strong><br>
        <a href="/communication_topology">communication_topology</a><br>
        <a href="/connector_device_state">connector_device_state</a><br>
        <a href="/connector_state">connector_state</a><br>
        <a href="/device_parameter">device_parameter</a><br>
        <a href="/events">events</a><br>
        <a href="/metered_data">metered_data</a><br>
        <a href="/tasks">tasks</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func connectorDeviceState(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">><head><meta charset="UTF-8"><title></title></head><body>
        <strong>CONNECTOR DEVICE STATE</strong><br>
        <a href="/communication_topology">communication_topology</a><br>
        <a href="/connector_device_state">connector_device_state</a><br>
        <a href="/connector_state">connector_state</a><br>
        <a href="/device_parameter">device_parameter</a><br>
        <a href="/events">events</a><br>
        <a href="/metered_data">metered_data</a><br>
        <a href="/tasks">tasks</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func connectorState(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">><head><meta charset="UTF-8"><title></title></head><body>
        <strong>CONNECTOR STATE</strong><br>
        <a href="/communication_topology">communication_topology</a><br>
        <a href="/connector_device_state">connector_device_state</a><br>
        <a href="/connector_state">connector_state</a><br>
        <a href="/device_parameter">device_parameter</a><br>
        <a href="/events">events</a><br>
        <a href="/metered_data">metered_data</a><br>
        <a href="/tasks">tasks</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func deviceParameter(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">><head><meta charset="UTF-8"><title></title></head><body>
        <strong>DEVICE PARAMETER</strong><br>
        <a href="/communication_topology">communication_topology</a><br>
        <a href="/connector_device_state">connector_device_state</a><br>
        <a href="/connector_state">connector_state</a><br>
        <a href="/device_parameter">device_parameter</a><br>
        <a href="/events">events</a><br>
        <a href="/metered_data">metered_data</a><br>
        <a href="/tasks">tasks</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func events(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">><head><meta charset="UTF-8"><title></title></head><body>
        <strong>EVENTS</strong><br>
        <a href="/communication_topology">communication_topology</a><br>
        <a href="/connector_device_state">connector_device_state</a><br>
        <a href="/connector_state">connector_state</a><br>
        <a href="/device_parameter">device_parameter</a><br>
        <a href="/events">events</a><br>
        <a href="/metered_data">metered_data</a><br>
        <a href="/tasks">tasks</a><br>
        <form method="POST" action="/events">
        <input type="submit" value="events">
        </form>
        </body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func meteredData(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">><head><meta charset="UTF-8"><title></title></head><body>
        <strong>METERED DATA</strong><br>
        <a href="/communication_topology">communication_topology</a><br>
        <a href="/connector_device_state">connector_device_state</a><br>
        <a href="/connector_state">connector_state</a><br>
        <a href="/device_parameter">device_parameter</a><br>
        <a href="/events">events</a><br>
        <a href="/metered_data">metered_data</a><br>
        <button onclick="run1()">Metered data 1</button>
        <button onclick="run2()">Metered data 2</button>
        <p id="demo1"></p>
        <p id="demo2"></p>  
        <script>
        function run1(){        
        var d = new Date();
        var n = d.valueOf();
        document.getElementById("demo1").innerHTML = (n);
        }
        var min = 11;
        var max = 100001;

        function getRandomInt(min,max) {
        return Math.floor(Math.random()*(max-min+1))+min;
        }

        var interval=setInterval(run2,2000);

        function run2(){
        min=getRandomInt(min+1,max);
        document.getElementById("demo2").innerHTML = (min);
        }
        </script>
        <a href="/tasks">tasks</a><br>
        </body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func meteredDataValue(conn net.Conn) {
    timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
    fmt.Fprintf(conn, timestamp)
}

func tasks(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">><head><meta charset="UTF-8"><title></title></head><body>
        <strong>TASKS</strong><br>
        <a href="/communication_topology">communication_topology</a><br>
        <a href="/connector_device_state">connector_device_state</a><br>
        <a href="/connector_state">connector_state</a><br>
        <a href="/device_parameter">device_parameter</a><br>
        <a href="/events">events</a><br>
        <a href="/metered_data">meter_data</a><br>
        <a href="/tasks">tasks</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
