package handlers

import (
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/utils"
	"os"
	"testing"
)

var testApp config.AppConfig
var testRepo *DBRepo

func TestMain(m *testing.M) {

	a := config.AppConfig{}

	testApp = a

	testRepo = NewDBTestHandlers(&testApp)
	NewHandlers(testRepo, &testApp)
	utils.NewUtils(&testApp)
	os.Exit(m.Run())
}
