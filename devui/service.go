package devui

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"os"
)

type service struct {
	log    *zap.SugaredLogger
	logger *zap.Logger
	devuis map[string]bool
}

func (s *service) GenerateDevUI() (devui string) {
	n := 8 //hex
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	devui = hex.EncodeToString(bytes)
	return
}

func (s *service) ValidateDevUI(devui string) (valid bool) {
	if devui == "" || len(devui) < 16 {
		return
	}

	uniquePart := devui[11:16]

	valid = s.devuis[uniquePart]

	if valid {
		s.log.Debug("devui exists", devui)
	} else {
		s.devuis[uniquePart] = true
	}

	return
}

func (s *service) Backup() {
	filename := "usedDevs.json"

	if fs, err := os.Create(filename); err == nil {
		devUIS := funk.Values(s.devuis).([]string)
		jsonb, err := json.Marshal(devUIS)
		if err == nil {
			fs.Write(jsonb)
		} else {
			s.log.Error(err)
		}
	}

}

func (s *service) Restore() {
	filename := "usedDevs.json"

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var devuis []string

	if err := json.Unmarshal(data, &devuis); err == nil {
		for _, i := range devuis {
			s.devuis[i] = true
		}
	} else {
		s.log.Fatalln("restore failed due to ", err)
	}
}
