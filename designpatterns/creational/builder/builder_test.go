package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	md := ManufacturingDirector{}

	carbuilder := &CarBuilder{}
	md.SetBuilder(carbuilder)
	md.Construct()
	fmt.Printf("%+v", carbuilder.GetVehicle())
	bikebuilder := &MotorbikeBuilder{}
	md.SetBuilder(bikebuilder)
	md.Construct()
	fmt.Printf("%+v", bikebuilder.GetVehicle())
}
