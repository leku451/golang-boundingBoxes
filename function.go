// Semi-axes of WGS-84 geoidal reference
const (
	WGS84_a = 6378137.0 // Major semiaxis [m]
	WGS84_b = 6356752.3 // Minor semiaxis [m]
)

type MapPoint struct {
	Longitude	float64
	Latitude	float64
}

type BoundingBox struct {
	MinPoint 	MapPoint
	MaxPoint 	MapPoint
}

// Deg2rad converts degrees to radians
func Deg2rad(degrees float64) float64 {
	return math.Pi * degrees / 180.0
}

// Rad2deg converts radians to degrees
func Rad2deg(radians float64) float64 {
	return 180.0 * radians / math.Pi
}

func WGS84EarthRadius(lat float64) float64 {
	An := WGS84_a * WGS84_a * math.Cos(lat)
	Bn := WGS84_b * WGS84_b * math.Sin(lat)
	Ad := WGS84_a * math.Cos(lat)
	Bd := WGS84_b * math.Sin(lat)
	return math.Sqrt((An*An + Bn*Bn) / (Ad*Ad + Bd*Bd))
}

// GetBoundingBox takes two arguments, MapPoint is a set of lat/lng,
// 'halfSideInKm' is the half length of the bounding box you want in kilometers.
func GetBoundingBox (point MapPoint, halfSideInKm float64) BoundingBox {
	// Bounding box surrounding the point at given coordinates,
	// assuming local approximation of Earth surface as a sphere
	// of radius given by WGS84
	lat := Deg2rad(point.Latitude)
	lon := Deg2rad(point.Longitude)
	halfSide := 1000 * halfSideInKm

	// Radius of Earth at given latitude
	radius := WGS84EarthRadius(lat)
	// Radius of the parallel at given latitude
	pradius := radius * math.Cos(lat)

	latMin := lat - halfSide / radius
	latMax := lat + halfSide / radius
	lonMin := lon - halfSide / pradius
	lonMax := lon + halfSide / pradius

	return BoundingBox{
		MinPoint: MapPoint{Latitude: Rad2deg(latMin), Longitude: Rad2deg(lonMin)},
		MaxPoint: MapPoint{Latitude: Rad2deg(latMax), Longitude: Rad2deg(lonMax)},
	}
}

// https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}