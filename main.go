package main

import (
	"fmt"
	"plugin"
	"github.com/michaelhenkel/intentresourceinterface"
	"net"
	"os"
	"bufio"
	"strings"
	"encoding/json"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var apiList = map[string]plugin.Symbol{}

func main(){
	listener()
}

func listener(){
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
        fmt.Printf("Serving %s\n", c.RemoteAddr().String())
        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }

                data := strings.TrimSpace(string(netData))
                if data == "STOP" {
                        break
                } else {
			apiHandler(data)
		}

                c.Write([]byte("hello\n"))
        }
        c.Close()
}

func apiHandler(rawData string){
	fmt.Println("Data: ", rawData)
	jsonData := []byte(rawData)
	var data interface{}
	err := json.Unmarshal(jsonData, &data)
        if err != nil {
                fmt.Println(err)
        }
	dataInterface := data.(map[string]interface{})
	pluginName := dataInterface["plugin"]
	fmt.Println("Plugin: ", pluginName)
	switch pluginName {
	case "loadIRD":
		irdName := dataInterface["name"]
		fmt.Println("Loading IRD: ", irdName)
		loadPlugin(irdName.(string))
	case "runIR":
		irName := dataInterface["name"]
		irAction:= dataInterface["action"]
		irData := dataInterface["data"]
		fmt.Println("Running IR: ", irName)
		fmt.Println("\taction: ", irAction)
		fmt.Println("\tWith data: ", irData)
		runIr(irName.(string),irAction.(string),irData.(map[string]interface{}))
	}
}

func runIr(pluginName string, action string, data map[string]interface{}){
	intentResource := apiList[pluginName]
	intentresourceinterface, err := intentResource.(func(map[string]interface{}) (intentresourceinterface.IntentResourceInterface, error))(data)
	fmt.Printf("GetIntentResource. result: %T %v %v\n", intentresourceinterface, intentresourceinterface, err)
	switch action {
	case "create":
		intentresourceinterface.Create()
	}
}

func loadPlugin(pluginName string) {
	pluginPath := fmt.Sprintf("%s/%s.so", pluginName, pluginName)
	p, err := plugin.Open(pluginPath)
        if err != nil {
                panic(err)
        }

        GetIntentResource, err := p.Lookup("GetIntentResource")
        if err != nil {
                panic(err)
        }

	apiList[pluginName] = GetIntentResource
}
