package motor_test

import (
	"testing"

	// Frameworks
	"github.com/djthorpe/gopi"

	// Modules
	_ "github.com/djthorpe/gopi-hw/sys/pwm"
	_ "github.com/djthorpe/gopi/sys/logger"
	_ "github.com/djthorpe/robots/sys/motor"
)

////////////////////////////////////////////////////////////////////////////////
// CREATE MODULES / APPS

func TestConfig_000(t *testing.T) {
	// Create config file
	config := gopi.NewAppConfig("robots/motor")
	t.Log(config)
}

func TestApp_000(t *testing.T) {
	// Create app
	config := gopi.NewAppConfig("robots/motor")
	if app, err := gopi.NewAppInstance(config); err != nil {
		t.Fatal(err)
	} else if motor := app.ModuleInstance("robots/motor"); motor == nil {
		t.Fatal("Motor module not found")
	} else {
		t.Log(motor)
	}
}
