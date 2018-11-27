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

// A set of motors which can be controlled forwards, backwards or stopped
type Motors interface {
	// Add a motor into the list of motors, indicting the GPIO pins for + and -
	Add(plus, minus gopi.GPIOPin) (Motor, error)

	// Run one or more motors at speed, which should be between -1.0 and 1.0
	// The function will return immediately, and the context can be used to cancel
	// the operation (and return the motors to 0)
	Run(ctx context.Context, speed float32, motors ...Motor) error

	// Return a list of motors
	Motors() []Motor
}

// A direct current motor interface
type Motor interface {
	// Return pins used to drive the motor
	Plus() gopi.GPIOPin
	Minus() gopi.GPIOPin
}
