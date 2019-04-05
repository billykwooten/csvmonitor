package main

import (
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/csvmonitor/csv"
	"time"
	"github.com/azer/logger"
)

var (
	app   = kingpin.New("darksky_exporter", "DarkSky Exporter for DarkSky Weather API").Author("Billy Wooten")
	csvloc = app.Flag("csv", "File path of the CSV file").Envar("CSVLOC").Required().String()
	slackhook = app.Flag("slackhook", "Webhook URL for Slacks").Envar("SLACKHOOK").Required().String()
	csvhour = app.Flag("csvhour", "Set the hour in which the script will run").Envar("CSVHOUR").Default("8").Int()
	csvminute = app.Flag("csvminute", "Set the minute in which the script will run").Envar("CSVMINUTE").Default("0").Int()
	log = logger.New("CSVMonitor")

)

func initEight() {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), *csvhour, *csvminute, 0, 0, t.Location())
	d := n.Sub(t)
	if d < 0 {
		log.Info("Waiting 24 hours to run again")
		log.Info("Will run again at hour %d", *csvhour, ":", *csvminute)
		n = n.Add(24 * time.Hour)
		d = n.Sub(t)
	}
	for {
		time.Sleep(d)
		d = 24 * time.Hour
		log.Info("Executing Parse CSV Function")
		csvparse.Parse_csv(*csvloc, *slackhook)
	}
}


func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.Info("Starting CSVMonitor")
	initEight()
}

