package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "07-04-2019 15:04:05" // 'murica
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	// log.Info("Some info. Earth is not flat.")
	// log.Warning("This is a warning")
	// log.Error("Not fatal. An error. Won't stop execution")
	// log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	// log.Panic("Do not panic")
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
