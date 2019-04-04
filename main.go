package main

import (
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/csvmonitor/csv"
	"time"
)

var (
	app   = kingpin.New("darksky_exporter", "DarkSky Exporter for DarkSky Weather API").Author("Billy Wooten")
	csvloc = app.Flag("csv", "File path of the CSV file").Envar("CSVLOC").Required().String()
	slackhook = app.Flag("slackhook", "Webhook URL for Slacks").Envar("SLACKHOOK").Required().String()
)

func initEight() {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, t.Location())
	d := n.Sub(t)
	if d < 0 {
		n = n.Add(24 * time.Hour)
		d = n.Sub(t)
	}
	for {
		time.Sleep(d)
		d = 24 * time.Hour
		csvparse.Parse_csv(*csvloc, *slackhook)
	}
}


func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	initEight()
}

