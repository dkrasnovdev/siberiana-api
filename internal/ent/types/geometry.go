package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
	"github.com/twpayne/go-geom/encoding/wkt"
)

// Geometry represents a Postgis Geometry.
// It stores a WKT (Well-Known text) reference inside as a string.
type Geometry struct {
	wkt string
}

// NewGeometryFromWKT creates a new Geometry from the WKT provided.
func NewGeometryFromWKT(wkt string) Geometry {
	return Geometry{wkt: wkt}
}

// Scan implements the Scanner interface.
func (g *Geometry) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	geom, err := ewkbhex.Decode(value.(string))
	if err != nil {
		return err
	}

	wktOut, err := wkt.Marshal(geom)
	if err != nil {
		return err
	}

	g.wkt = wktOut
	return nil
}

// Value implements the driver Valuer interface.
func (g Geometry) Value() (driver.Value, error) {
	return g.wkt, nil
}

// WKT returns the Well-Known text from the geometry.
func (g Geometry) WKT() string {
	return g.wkt
}

// SetWKT sets the WKT for the geometry.
func (g *Geometry) SetWKT(wkt string) {
	g.wkt = wkt
}

// String returns the string representation of the Geometry.
func (g Geometry) String() string {
	return g.WKT()
}

// FormatParam implements the sql.ParamFormatter interface to tell the SQL
// builder that the placeholder for a Geometry parameter needs to be formatted.
func (g Geometry) FormatParam(placeholder string, info *sql.StmtInfo) string {
	return "ST_GeomFromText(" + placeholder + ", 4326)"
}

// GeometrySchemaType defines the schema-type of the Geometry type.
func GeometrySchemaType() map[string]string {
	return map[string]string{
		dialect.Postgres: "geometry",
	}
}

// MarshalJSON marshal the Geometry to JSON.
func (g Geometry) MarshalJSON() ([]byte, error) {
	if g.wkt == "" {
		return nil, nil
	}

	return []byte(`"` + g.wkt + `"`), nil
}

// UnmarshalJSON unmarshal the Geometry from JSON.
func (g *Geometry) UnmarshalJSON(data []byte) error {
	g.wkt = strings.ReplaceAll(string(data), `"`, "")
	return nil
}

// UnmarshalGQL implements the interface graphql.Unmarshaler.
func (g *Geometry) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("type Geometry must be a string")
	}

	return json.Unmarshal([]byte(value), g)
}

// MarshalGQL implements the interface graphql.Marshaler.
func (g Geometry) MarshalGQL(w io.Writer) {
	_ = json.NewEncoder(w).Encode(g)
}
