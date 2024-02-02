package main

import (
	"fmt"
	"os"
	"time"

	irsdk "github.com/iracing-telemetry-group/iracing-sdk"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
}

func main() {
	irsdk := irsdk.Init(nil)
	defer irsdk.Close()
	irsdk.WaitForData(1 * time.Second)
	session := irsdk.GetSession().DriverInfo.DriverUserID
	fmt.Println(session)
	// for {
	// 	log.Infof("======== New iteration =======")
	// 	if irsdk.WaitForData(1 * time.Second) {

	// 		// log.Infof("got new data")
	// 		// log.Infof("Version: %d", irsdk.GetLastVersion())
	// 		// // log.Infof("Session: %v", irsdk.GetSession())
	// 		// // log.Infof("Driver Info: %v", pretty.Sprint(irsdk.GetSession().DriverInfo))
	// 		// log.Infof("Estimated lap time: %v", pretty.Sprint(irsdk.GetSession().DriverInfo.DriverCarEstLapTime))
	// 		// carIdx := irsdk.GetSession().DriverInfo.DriverCarIdx
	// 		// estTimeVar, err := irsdk.GetVar("CarIdxEstTime")
	// 		// if err != nil {
	// 		// 	log.Fatalf("failed to get car idx est time: %e", err)
	// 		// }
	// 		// estTimes, ok := estTimeVar.Value.([]float32)
	// 		// if !ok {
	// 		// 	log.Fatalf("estimated time is not an array")
	// 		// }
	// 		// carEstTime := estTimes[carIdx]
	// 		// estDeltas := make([]float32, len(estTimes))
	// 		// for i := 0; i < len(estTimes); i++ {
	// 		// 	estDeltas[i] = estTimes[i] - carEstTime
	// 		// }
	// 		// log.Infof("Deltas: %v", pretty.Sprint(estDeltas))
	// 		// log.Infof("Drivers: %v", pretty.Sprint(irsdk.GetSession().DriverInfo.Drivers))
	// 		// log.Infof("All telemetry Vars: %v", pretty.Sprint(irsdk.GetAllVars()))
	// 	}
	// 	time.Sleep(500 * time.Millisecond)
	// }
}
