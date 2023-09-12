package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-vgo/robotgo"

	"github.com/gocolly/colly"
)

// commande pour ouvrir google chrome : open -a 'google chrome' http://www.google.com
// SMSRequestBody ...
type SMSRequestBody struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

// url with ur user conenction like => https://admissions.42.fr/my-user-id/introductions_users
const URL = "http://www.42/blalblbalba.com"
const NEXMO_API_KEY = "TOP SECRET KEY"
const NEXMP_API_SECRET = "TOP SECRET KEY"
const PHONE_NUMBER = "111111111111111111"

func Alert42() {
	const sentence = "Il n'y a pas de check-in disponible pour le moment, nous t'informerons dès qu'il y en aura un de disponible."

	// checker
	checker := false
	// scrapping
	c := colly.NewCollector(
		colly.AllowedDomains(URL, "https://admissions.42.fr/"),
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnHTML("h5.font-weight-bold", func(e *colly.HTMLElement) {
		if e.Text == sentence {
			checker = false
		} else {
			checker = true
		}
		fmt.Println(e.Text)
	})
	if checker {

		params := url.Values{}
		params.Add("from", `ME`)
		params.Add("text", `Check-in 42 vite !!!!!!!!!!`)
		params.Add("to", `1111`)
		params.Add("api_key", `11`)
		params.Add("api_secret", `111111`)
		body := strings.NewReader(params.Encode())

		req, err := http.NewRequest("POST", "https://rest.nexmo.com/sms/json", body)
		if err != nil {
			// handle err
			println("Error creating request: ", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
			println("Error: ", err)
		}
		defer resp.Body.Close()
	} else {
		log.Println("Pas de changement signalé pour le moment :(")
	}

	c.Visit(URL)
}
func Connexion() {

}
func main() {

	robotgo.Run("open -a 'safari'" + URL)

	Alert42()

}
