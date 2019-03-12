package main

import(
	"net/http"
	"io"
	"log"
	"os"
	"encoding/json"
	"bytes"
	"fmt"
)

func main(){
	type Task struct {
		Desc string `json:"desc"`
		Done bool   `json:"done"`
	}

	t := &Task{
			Desc: "จ่ายค่าบัตรเครดิต",
			Done: true,
	}
	b, err := json.Marshal(t)
	req, err := http.NewRequest(
		"POST",
		"https://httpbin.org/post",
		bytes.NewReader(b),
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil{
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var rt struct{
		Task Task `json:"json"`
	}
	respBuf := new(bytes.Buffer)

	io.Copy(io.MultiWriter(os.Stdout, respBuf), resp.Body)

	dec := json.NewDecoder(respBuf)
	dec.Decode(&rt)
	fmt.Printf("%+v\n", rt.Task)
}