package elon

import "fmt"

func (car *Car) Drive() {

	// Assumptions:
	// - the receiver, car Struct, is already created.

	// we have the juice to drive
	if car.battery >= car.batteryDrain {

		car.battery -= car.batteryDrain
		car.distance += car.speed
	}

}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {

	return car.battery/car.batteryDrain*car.speed >= trackDistance
}
