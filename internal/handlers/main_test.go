package handlers

import (
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/utils"
	"testing"
)

var testApp config.AppConfig
var testRepo *DBRepo

func TestMain(m *testing.M) {

	a := config.AppConfig{}

	testApp = a

	testRepo = NewTestHandlers(&testApp)
	NewHandlers(testRepo, &testApp)
	utils.NewUtils(&testApp)
}
