package motors_test

import (
	"testing"

	// Frameworks
	"github.com/djthorpe/gopi"
	"github.com/djthorpe/robots"

	// Modules
	_ "github.com/djthorpe/gopi-hw/sys/pwm"
	_ "github.com/djthorpe/gopi/sys/logger"
	_ "github.com/djthorpe/robots/sys/motors"
)

////////////////////////////////////////////////////////////////////////////////
// CREATE MODULES / APPS

func TestConfig_000(t *testing.T) {
	// Create config file
	config := gopi.NewAppConfig("robots/motors")
	t.Log(config)
}

func TestApp_000(t *testing.T) {
	// Create app
	config := gopi.NewAppConfig("robots/motors")
	if app, err := gopi.NewAppInstance(config); err != nil {
		t.Fatal(err)
	} else if motors_ := app.ModuleInstance("robots/motors"); motors_ == nil {
		t.Fatal("Motors module not found")
	} else if motors := motors_.(robots.Motors); motors == nil {
		t.Fatal("Motors module not found")
	} else {
		t.Log(motors)
	}
}
