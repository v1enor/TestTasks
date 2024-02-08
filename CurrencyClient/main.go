package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const URL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

type CryptoData struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"current_price"`
}

func getCryptoData() ([]CryptoData, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []CryptoData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите количество криптовалют или название: ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	userInput := getUserInput()
	var num int
	var err error
	var cryptoName string

	data, err := getCryptoData()
	if err != nil {
		log.Fatal(err)
	}

	num, err = strconv.Atoi(userInput)
	if err != nil {
		found := false
		for _, crypto := range data {
			if strings.EqualFold(crypto.Name, userInput) {
				cryptoName = crypto.Name
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Криптовалюта не найдена")
			return
		}
	}

	for {
		data, err = getCryptoData()
		if err != nil {
			log.Fatal(err)
		}
		if err == nil {
			for i := 0; i < num && i < len(data); i++ {
				fmt.Println(data[i].Name, data[i].Price)
			}
		} else {
			for _, crypto := range data {
				if strings.EqualFold(crypto.Name, cryptoName) {
					fmt.Println(crypto.Name, crypto.Price)
					break
				}
			}
		}

		fmt.Println(time.Now().Format("15:04:05"))
		time.Sleep(10 * time.Minute)
	}
}