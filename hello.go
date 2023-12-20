package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func generateUUID() string {
	uuid := uuid.New()
	return fmt.Sprintf(`{"partnerUserId":"%s"}`, uuid)
}

func main() {
	client := &http.Client{}
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Loop indefinitely
	for {
		// Request data
		var data = strings.NewReader(generateUUID())
		req, err := http.NewRequest("POST", "https://api.discord.gx.games/v1/direct-fulfillment", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("authority", "api.discord.gx.games")
		req.Header.Set("accept", "*/*")
		req.Header.Set("accept-language", "en-US,en;q=0.9")
		req.Header.Set("content-type", "application/json")
		req.Header.Set("origin", "https://www.opera.com")
		req.Header.Set("referer", "https://www.opera.com/")
		req.Header.Set("sec-ch-ua", `"Opera GX";v="105", "Chromium";v="119", "Not?A_Brand";v="24"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		req.Header.Set("sec-fetch-dest", "empty")
		req.Header.Set("sec-fetch-mode", "cors")
		req.Header.Set("sec-fetch-site", "cross-site")
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 OPR/105.0.0.0")
		// Send request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Read response body
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Extract token from response
		var result map[string]string
		if err := json.Unmarshal(bodyText, &result); err != nil {
			log.Fatal(err)
		}
		token, exists := result["token"]
		if !exists {
			log.Fatal("Token not found in response")
		}

		// Generate URL
		url := fmt.Sprintf("https://discord.com/billing/partner-promotions/1180231712274387115/%s\n", token)

		if _, err := file.WriteString(url); err != nil {
			log.Fatal(err)
		}

		// Delay for 3 seconds
		time.Sleep(3 * time.Second)
	}
}
