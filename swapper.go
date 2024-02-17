package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

var (
	removeToken string
	tokens      string
	removeguild string
	vanityCode  string
	keepguild   string
	client      = &http.Client{}
)

func main() {
	fmt.Print("Urlyi Silicek Olan Hesabın Tokeni")
	fmt.Scanln(&removeToken)

	fmt.Print("Silinecek Olan Url ")
	fmt.Scanln(&vanityCode)

	fmt.Print("Urlsi Silinecek Olan sunucu ")
	fmt.Scanln(&removeguild)

	fmt.Print("Silinen Urlyi Alıcağımız Sunucu ")
	fmt.Scanln(&keepguild)

	fmt.Println("Swap başladı")
	runtime.GOMAXPROCS(runtime.NumCPU())

	removeURL(removeToken)

	var wg sync.WaitGroup
	wg.Add(len(tokens))

	for _, token := range tokens {
		go func(token string) {
			defer wg.Done()
			spamURL(token, vanityCode, keepguild)
		}(token)
	}

	wg.Wait()

	fmt.Println("belli degil aga")
}

func removeURL(token string) {
	jsonData := []byte(`{"code":null}`)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://canary.discord.com/api/v7/guilds/%s/vanity-url", removeguild), bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println(fmt.Sprintf("Urlyi Sildik| StatusCode: %d", resp.StatusCode))
	} else {
		fmt.Println(fmt.Sprintf("Url Silinemedi | StatusCode: %d", resp.StatusCode))
	}
}

func spamURL(token, vanityCode, keepguild string) {
	jsonData := []byte(fmt.Sprintf(`{"code":"%s"}`, vanityCode))
	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://canary.discord.com/api/v7/guilds/%s/vanity-url", keepguild), bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println(fmt.Sprintf("Url Swaplandı -hollandalı dev| StatusCode: %d", resp.StatusCode))
	} else {
		fmt.Println(fmt.Sprintf("Url Swaplanamadı| StatusCode: %d", resp.StatusCode))
	}
}
