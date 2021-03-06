/*
  Robots
  (c) Copyright David Thorpe 2018
  All Rights Reserved

  For Licensing and Usage information, please see LICENSE.md
  Documentation https://github.com/djthorpe/robots/


*/

package robots

import (
	"context"

	// Frameworks
	"github.com/djthorpe/gopi"
)

// Motors is a set of motors which can be controlled forwards, backwards or stopped
type Motors interface {
	// Add a motor into the list of motors, indicting the GPIO pins for - and +
	// and whether the speed should be inverted (where the wires are mixed up)
	Add(minus, plus gopi.GPIOPin, invert bool) (Motor, error)

	// Run one or more motors at speed, which should be between -1.0 and 1.0
	// The function will return immediately, and the context can be used to cancel
	// the operation (and return the motors to 0)
	Run(ctx context.Context, speed float32, motors ...Motor) error

	// Set speed to zero for all running motors, or for specific motors
	Stop(motors ...Motor) error
}

// Motor is direct current motor interface
type Motor interface {
	// Return pins used to drive the motor
	Plus() gopi.GPIOPin
	Minus() gopi.GPIOPin

	// Return current speed
	Speed() float32
}
