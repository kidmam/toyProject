package builder

// Product to build
type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

// Each builder should implement this interface
type VehicleBuilder interface {
	SetWheels() VehicleBuilder
	SetSeats() VehicleBuilder
	SetStructure() VehicleBuilder
	GetVehicle() Vehicle
}

// A builder of type Car
type CarBuilder struct {
	v Vehicle
}

func (c *CarBuilder) SetWheels() VehicleBuilder {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() VehicleBuilder {
	c.v.Seats = 4
	return c
}

func (c *CarBuilder) SetStructure() VehicleBuilder {
	c.v.Structure = "car"
	return c
}

func (c *CarBuilder) GetVehicle() Vehicle {
	return c.v
}

// A builder of type Motorbike
type MotorbikeBuilder struct {
	v Vehicle
}

func (b *MotorbikeBuilder) SetWheels() VehicleBuilder {
	b.v.Wheels = 2
	return b
}

func (b *MotorbikeBuilder) SetSeats() VehicleBuilder {
	b.v.Seats = 2
	return b
}

func (b *MotorbikeBuilder) SetStructure() VehicleBuilder {
	b.v.Structure = "motorbike"
	return b
}

func (b *MotorbikeBuilder) GetVehicle() Vehicle {
	return b.v
}

// Director
type ManufacturingDirector struct {
	builder VehicleBuilder
}

func (m *ManufacturingDirector) SetBuilder(b VehicleBuilder) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}
