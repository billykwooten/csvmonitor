package csvparse

import (
	"os"
	"github.com/gocarina/gocsv"
	"strings"
	"strconv"
	"fmt"
	"time"
	"github.com/csvmonitor/slack"
	"github.com/azer/logger"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Date      string `csv:"date"`
	Name    string `csv:"client_name"`
	Payment     string `csv:"payment_amount"`
	NotUsed string `csv:"-"`
}

func Parse_csv(csvloc string, slackhook string) {
	var log = logger.New("Parse_csv")

	log.Info("Opening CSV file")
	clientsFile, err := os.OpenFile(csvloc, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Error("Fatal could not open CSV")
		log.Error("error: %s\n", err)
		panic(err)
	}
	defer clientsFile.Close()

	var clients []*Client

	log.Info("Unmarshaling CSV file")
	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		log.Error("Fatal could not unmarshal CSV")
		log.Error("error: %s\n", err)
		panic(err)
	}

	for _, client := range clients {
		var month int
		var day int
		var year int

		paymentdate := strings.Split(client.Date, "/")

		for i := range paymentdate {
			if i == 0 {
				month, err = strconv.Atoi(paymentdate[i])
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			}
			if i == 1 {
				day, err = strconv.Atoi(paymentdate[i])
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			}
			if i == 2 {
				year, err = strconv.Atoi(paymentdate[i])
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			}

		}

		var paymentconv int
		paymentconv, err = strconv.Atoi(client.Payment)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		monthstoadd := 0

		if paymentconv >= 5 {
			monthstoadd = paymentconv / 5
		}

		log.Info("Calculating expiration dates")
		start := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		expirationdate := start.AddDate(0, monthstoadd, 0)

		log.Info("Sending slack Webhook")
		slackwebhook.SlackWebHook(slackhook, client.Name, client.Date, client.Payment, expirationdate.String(), expirationdate.Before(time.Now()))

	}

}
