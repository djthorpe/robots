package main

import (
	"fmt"
	"os"

	// Frameworks
	"github.com/djthorpe/gopi"

	// Modules
	_ "github.com/djthorpe/gopi-hw/sys/pwm"
	_ "github.com/djthorpe/gopi/sys/logger"
)

const (
	// Motor, via DRV8833PWP Dual H-Bridge
	M1B = gopi.GPIOPin(19)
	M1F = gopi.GPIOPin(20)
	M2B = gopi.GPIOPin(21)
	M2F = gopi.GPIOPin(26)
)

////////////////////////////////////////////////////////////////////////////////

func mainLoop(app *gopi.AppInstance, done chan<- struct{}) error {

	if app.PWM == nil {
		return app.Logger.Error("Missing PWM module instance")
	}

	// Stop motors
	if err := app.PWM.SetDutyCycle(0, M1B, M1F, M2B, M2F); err != nil {
		return err
	}

	fmt.Println(app.PWM)

	// Finished
	done <- gopi.DONE
	return nil
}

func main() {
	// Create the configuration, load the spi instance
	config := gopi.NewAppConfig("pwm")

	// Run the command line tool
	os.Exit(gopi.CommandLineTool(config, mainLoop))
}
