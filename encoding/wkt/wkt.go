// Package wkt implements Well Known Text encoding and decoding.
package wkt

import (
	"github.com/twpayne/go-geom"
)

const (
	tPoint              = "POINT "
	tMultiPoint         = "MULTIPOINT "
	tLineString         = "LINESTRING "
	tMultiLineString    = "MULTILINESTRING "
	tPolygon            = "POLYGON "
	tMultiPolygon       = "MULTIPOLYGON "
	tGeometryCollection = "GEOMETRYCOLLECTION "
	tZ                  = "Z "
	tM                  = "M "
	tZm                 = "ZM "
	tEmpty              = "EMPTY"
)

// Marshal translates a geometry to the corresponding WKT.
func Marshal(g geom.T) (string, error) {
	return encode(g)
}

// Unmarshal translates a WKT to the corresponding geometry.
func Unmarshal(wkt string) (geom.T, error) {
	return decode(wkt)
}
