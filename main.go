package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	cur1 := args[1]
	cur2 := args[2]

	rate, err := getExchangeRate(cur1, cur2)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("The exchange rate of %v to %v is %v", cur1, cur2, rate)
}

func getExchangeRate(cur1, cur2 string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%s/%s.json", cur1, cur2))
	if err != nil {
		return 0, err
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return 0, err
	}

	x := make(map[string]any)
	err = json.Unmarshal([]byte(buf.String()), &x)
	if err != nil {
		return 0, err
	}

	v, ok := x[cur2].(float64)
	if !ok {
		return 0, nil
	}
	return v, err
}
