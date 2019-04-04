package csvparse

import (
	"os"
	"github.com/gocarina/gocsv"
	"strings"
	"strconv"
	"fmt"
	"time"
	"github.com/csvmonitor/slack"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Date      string `csv:"date"`
	Name    string `csv:"client_name"`
	Payment     string `csv:"payment_amount"`
	NotUsed string `csv:"-"`
}

func Parse_csv(csvloc string, slackhook string) {

	clientsFile, err := os.OpenFile(csvloc, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	var clients []*Client

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
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

		if paymentconv > 5 {
			monthstoadd = paymentconv / 5
		}

		start := time.Date(year, time.Month(month+monthstoadd), day, 0, 0, 0, 0, time.UTC)
		expirationdate := start.AddDate(0, 1, 0)

		slackwebhook.SlackWebHook(slackhook, client.Name, client.Date, client.Payment, expirationdate.String(), expirationdate.Before(time.Now()))

	}

}