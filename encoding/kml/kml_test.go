package kml

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/twpayne/go-geom"
)

func Test(t *testing.T) {
	for _, tc := range []struct {
		g    geom.T
		want string
	}{
		{
			g:    geom.NewPoint(geom.XY),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    geom.NewPoint(geom.XY).MustSetCoords([]float64{0, 0}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    geom.NewPoint(geom.XYZ).MustSetCoords([]float64{0, 0, 0}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    geom.NewPoint(geom.XYZ).MustSetCoords([]float64{0, 0, 1}),
			want: `<Point><coordinates>0,0,1</coordinates></Point>`,
		},
		{
			g:    geom.NewPoint(geom.XYM).MustSetCoords([]float64{0, 0, 1}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    geom.NewPoint(geom.XYZM).MustSetCoords([]float64{0, 0, 0, 1}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    geom.NewPoint(geom.XYZM).MustSetCoords([]float64{0, 0, 1, 1}),
			want: `<Point><coordinates>0,0,1</coordinates></Point>`,
		},
		{
			g: geom.NewLineString(geom.XY).MustSetCoords([][]float64{
				{0, 0}, {1, 1},
			}),
			want: `<LineString><coordinates>0,0 1,1</coordinates></LineString>`,
		},
		{
			g: geom.NewLineString(geom.XYZ).MustSetCoords([][]float64{
				{1, 2, 3}, {4, 5, 6},
			}),
			want: `<LineString><coordinates>1,2,3 4,5,6</coordinates></LineString>`,
		},
		{
			g: geom.NewLineString(geom.XYM).MustSetCoords([][]float64{
				{1, 2, 3}, {4, 5, 6},
			}),
			want: `<LineString><coordinates>1,2 4,5</coordinates></LineString>`,
		},
		{
			g: geom.NewLineString(geom.XYZM).MustSetCoords([][]float64{
				{1, 2, 3, 4}, {5, 6, 7, 8},
			}),
			want: `<LineString><coordinates>1,2,3 5,6,7</coordinates></LineString>`,
		},
	} {
		b := &bytes.Buffer{}
		e := xml.NewEncoder(b)
		if err := e.Encode(Encode(tc.g)); err != nil {
			t.Errorf("Encode(Encode(%#v)) == %v, want nil", tc.g, err)
			continue
		}
		if got := b.String(); got != tc.want {
			t.Errorf("Encode(Encode(%#v))\nwrote %v\n want %v", tc.g, got, tc.want)
		}
	}
}