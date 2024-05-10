package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var distance, battery int
	if car.batteryDrain > car.battery {
		distance = car.distance
		battery = car.battery
	} else {
		distance = car.distance + car.speed
		battery = car.battery - car.batteryDrain
	}
	return Car{
		distance:     distance,
		battery:      battery,
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	carTotalDistance := car.speed * int(car.battery/car.batteryDrain)
	return carTotalDistance >= track.distance
}
