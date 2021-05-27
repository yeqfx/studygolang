package geo

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func (c *Coordinate) SetLatitude(latitude float64) {
	c.Latitude = latitude
}

func (c *Coordinate) SetLongitude(longitude float64) {
	c.Longitude = longitude
}
