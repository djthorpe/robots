package motors

import (
	"context"
	"fmt"
	"sync"

	// Frameworks
	"github.com/djthorpe/gopi"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type motor struct {
	plus, minus gopi.GPIOPin
	speed       float32
	factor      float32
	cancel      context.CancelFunc

	sync.Mutex
}

////////////////////////////////////////////////////////////////////////////////
// NEW

// Create new Motors object
func NewMotor(plus, minus gopi.GPIOPin, invert bool) *motor {
	this := new(motor)
	this.plus = plus
	this.minus = minus
	this.speed = 0

	// Set factor
	if invert {
		this.factor = -1.0
	} else {
		this.factor = 1.0
	}

	return this
}

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *motor) String() string {
	return fmt.Sprintf("<sys.robots.Motor>{ plus=%v minus=%v factor=%v speed=%v }", this.plus, this.minus, this.factor, this.speed)
}

////////////////////////////////////////////////////////////////////////////////
// IMPLEMENTATION

func (this *motor) Plus() gopi.GPIOPin {
	return this.plus
}

func (this *motor) Minus() gopi.GPIOPin {
	return this.minus
}

func (this *motor) Speed() float32 {
	return this.speed
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *motor) do_stop() {
	this.Lock()
	defer this.Unlock()
	if this.cancel != nil {
		this.cancel()
		this.cancel = nil
		this.speed = 0.0
	}
	fmt.Println("do_stop m=", this)
}

func (this *motor) do_start(speed float32, cancel context.CancelFunc) {
	this.Lock()
	defer this.Unlock()
	if this.cancel != nil {
		this.cancel()
		this.speed = 0.0
	}
	if cancel != nil {
		this.cancel = cancel
	}
	this.speed = speed * this.factor
	fmt.Println("do_start m=", this)
}
