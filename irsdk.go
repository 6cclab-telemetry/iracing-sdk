package irsdk

import (
	"fmt"
	"io/ioutil"

	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/hidez8891/shm"
	log "github.com/sirupsen/logrus"
)

// IrSdk is the main SDK object clients must use
type IrSdk struct {
	r                   reader
	h                   *header
	session             Session
	s                   []string
	tVars               *TelemetryVars
	lastUpdateTimestamp time.Time
}

func (sdk *IrSdk) WaitForData(timeout time.Duration) bool {
	sdk.readHeader()
	if !sessionStatusConnected(sdk.h.status) {
		return false
	}

	sRaw := readSessionData(sdk.r, sdk.h)
	err := yaml.Unmarshal([]byte(sRaw), &sdk.session)
	if err != nil {
		log.Debug("Failed to parse session data YAML: %v", err)
	}
	sdk.s = strings.Split(sRaw, "\n")
	sdk.tVars = readVariableHeaders(sdk.r, sdk.h)
	return readVariableValues(sdk)
}

func (sdk *IrSdk) GetSession() Session {
	return sdk.session
}

func (sdk *IrSdk) GetLastVersion() int {
	return sdk.tVars.lastVersion
}

func (sdk *IrSdk) GetAllVars() map[string]variable {
	return sdk.tVars.vars
}

func (sdk *IrSdk) GetVar(name string) (variable, error) {
	if v, ok := sdk.tVars.vars[name]; ok {
		return v, nil
	}
	return variable{}, fmt.Errorf("Telemetry variable %q not found", name)
}

func (sdk *IrSdk) GetSessionData(path string) (string, error) {
	return getSessionDataPath(sdk.s, path)
}

// Close clean up sdk resources
func (sdk *IrSdk) Close() {
	sdk.r.Close()
}

// Init creates a SDK instance to operate with
func Init(r reader) IrSdk {
	if r == nil {
		var err error
		r, err = shm.Open(fileMapName, fileMapSize)
		if err != nil {
			log.Fatalf("Failed to open SDK shared memory: %v", err)
		}
	}

	return sdk
}

func (sdk *IrSdk) readHeader() {
	h := readHeader(sdk.r)
	sdk.h = &h
}

func sessionStatusConnected(status int) bool {
	return (status & stConnected) > 0
}

// ExportIbtTo exports current memory data to a file
func (sdk *IrSdk) ExportIbtTo(fileName string) {
	rbuf := make([]byte, fileMapSize)
	_, err := sdk.r.ReadAt(rbuf, 0)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(fileName, rbuf, 0644)
}

// ExportSessionTo exports current session yaml data to a file
func (sdk *IrSdk) ExportSessionTo(fileName string) {
	y := strings.Join(sdk.s, "\n")
	ioutil.WriteFile(fileName, []byte(y), 0644)
}
