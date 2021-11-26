package main

import(
  "fmt"
  "net/http"
  //"net"
  "log"
  "bufio"
  "os"
  "strings"
  "math/rand"
  "encoding/json"
)

var factList []string
func Handler(writer http.ResponseWriter, request *http.Request){
  writer.Write([]byte("<h1>Ola</h1>"))
}
func RFHandler(writer http.ResponseWriter, request *http.Request){
  index := rand.Intn(len(factList))
  val := factList[index]
  data := make(map[string]string)
  data["fact"] = val
  fmt.Print(val)
  jsondata, err := json.MarshalIndent(data, "", "  ")
  if err != nil{
    log.Fatalf("Error happened in JSON marshal. Err: %s", err)
  }
  writer.Write(jsondata)
}

func main(){
  /*port := os.Getenv("PORT")
  if port == "" {
		log.Fatal("$PORT must be set")
	}*/port:= "80"
  dat, err := os.ReadFile("facts.txt")
  if err!= nil{
    panic(err)
  }
  fmt.Print(string(dat))
  scan := bufio.NewScanner(strings.NewReader(string(dat[:])))
  for scan.Scan() {
    factList = append(factList, scan.Text())
  }
  fmt.Println(factList)
  http.HandleFunc("/", Handler)
  http.HandleFunc("/random", RFHandler)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
