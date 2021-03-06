package motors_test

import (
	"context"
	"testing"
	"time"

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

////////////////////////////////////////////////////////////////////////////////
// CREATE MOTORS

func TestMotors_000(t *testing.T) {
	// Create app
	config := gopi.NewAppConfig("robots/motors")
	if app, err := gopi.NewAppInstance(config); err != nil {
		t.Fatal(err)
	} else if motors := app.ModuleInstance("robots/motors").(robots.Motors); motors == nil {
		t.Fatal("Motors module not found")
	} else if _, err := motors.Add(gopi.GPIOPin(19), gopi.GPIOPin(20), false); err != nil {
		t.Error(err)
	} else {
		t.Log(motors)
	}
}

func TestMotors_001(t *testing.T) {
	// Create app
	config := gopi.NewAppConfig("robots/motors")
	if app, err := gopi.NewAppInstance(config); err != nil {
		t.Fatal(err)
	} else if motors := app.ModuleInstance("robots/motors").(robots.Motors); motors == nil {
		t.Fatal("Motors module not found")
	} else if left, err := motors.Add(gopi.GPIOPin(19), gopi.GPIOPin(20), false); err != nil {
		t.Error(err)
	} else if right, err := motors.Add(gopi.GPIOPin(21), gopi.GPIOPin(26), false); err != nil {
		t.Error(err)
	} else if err := motors.Run(context.Background(), 1.0, left, right); err != nil {
		t.Error(err)
	} else {
		// Run for one second then stop
		time.Sleep(time.Second)
		if err := motors.Stop(left, right); err != nil {
			t.Error(err)
		} else {
			t.Log(motors)
		}
	}
}

func TestMotors_002(t *testing.T) {
	// Create app
	config := gopi.NewAppConfig("robots/motors")
	if app, err := gopi.NewAppInstance(config); err != nil {
		t.Fatal(err)
	} else if motors := app.ModuleInstance("robots/motors").(robots.Motors); motors == nil {
		t.Fatal("Motors module not found")
	} else if left, err := motors.Add(gopi.GPIOPin(19), gopi.GPIOPin(20), false); err != nil {
		t.Error(err)
	} else if right, err := motors.Add(gopi.GPIOPin(21), gopi.GPIOPin(26), false); err != nil {
		t.Error(err)
	} else {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5.0)
		if err := motors.Run(ctx, 1.0, left, right); err != nil {
			t.Error(err)
		} else if err := motors.Run(ctx, 0.5, left, right); err != nil {
			t.Error(err)
		} else {
			// Run for one second then cancel
			time.Sleep(time.Second)
		}
	}
}
