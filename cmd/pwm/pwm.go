package main

import (
	"os"
	"time"

	// Frameworks
	"github.com/djthorpe/gopi"

	// Modules
	_ "github.com/djthorpe/gopi-hw/sys/pwm"
	_ "github.com/djthorpe/gopi/sys/logger"
)

////////////////////////////////////////////////////////////////////////////////

type Motor struct {
	plus, minus gopi.GPIOPin
	factor      float32
	pwm         gopi.PWM
}

func NewMotor(pwm gopi.PWM, plus, minus gopi.GPIOPin, invert bool) *Motor {
	this := new(Motor)
	this.plus = plus
	this.minus = minus
	this.pwm = pwm
	if invert {
		this.factor = -1.0
	} else {
		this.factor = 1.0
	}
	return this
}

func (this *Motor) SetSpeed(speed float32) {
	speed = speed * this.factor
	if speed < -1.0 {
		speed = -1.0
	}
	if speed > 1.0 {
		speed = 1.0
	}
	if speed < 0.0 {
		this.pwm.SetDutyCycle(-speed, this.minus)
		this.pwm.SetDutyCycle(0, this.plus)
	} else if speed > 0.0 {
		this.pwm.SetDutyCycle(speed, this.plus)
		this.pwm.SetDutyCycle(0, this.minus)
	} else {
		this.pwm.SetDutyCycle(0.0, this.plus, this.minus)
	}
}

////////////////////////////////////////////////////////////////////////////////

const (
	// Motor, via DRV8833PWP Dual H-Bridge
	// In order to use this with pi-blaster, you will also need to
	// add the line to the file /etc/default/pi-blaster
	// DAEMON_OPTS=-g 19,20,21,26
	M1P = gopi.GPIOPin(19)
	M1N = gopi.GPIOPin(20)
	M2P = gopi.GPIOPin(21)
	M2N = gopi.GPIOPin(26)
)

////////////////////////////////////////////////////////////////////////////////

func mainLoop(app *gopi.AppInstance, done chan<- struct{}) error {

	if app.PWM == nil {
		return app.Logger.Error("Missing PWM module instance")
	}

	// Set up motors
	right := NewMotor(app.PWM, M1P, M1N, false)
	left := NewMotor(app.PWM, M2P, M2N, true)

	left.SetSpeed(1.0)
	right.SetSpeed(-1.0)

	time.Sleep(1.0 * time.Second)

	left.SetSpeed(0.0)
	right.SetSpeed(0.0)

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
