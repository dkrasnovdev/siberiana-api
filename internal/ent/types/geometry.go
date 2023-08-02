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

// Geometry represents a PostGIS Geometry.
// It stores a WKT (Well-Known Text) reference inside as a string.
type Geometry struct {
	wkt string
}

// NewGeometryFromWKT creates a new Geometry from the given WKT (Well-Known Text).
func NewGeometryFromWKT(wkt string) Geometry {
	return Geometry{wkt: wkt}
}

// Scan implements the Scanner interface, allowing the Geometry to be scanned from the database.
func (g *Geometry) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	// Decode the value (WKB encoded) into a geom.
	geom, err := ewkbhex.Decode(value.(string))
	if err != nil {
		return err
	}

	// Marshal the geom into WKT (Well-Known Text) format.
	wktOut, err := wkt.Marshal(geom)
	if err != nil {
		return err
	}

	g.wkt = wktOut
	return nil
}

// Value implements the driver Valuer interface, allowing the Geometry to be stored in the database.
func (g Geometry) Value() (driver.Value, error) {
	return g.wkt, nil
}

// WKT returns the Well-Known Text representation of the Geometry.
func (g Geometry) WKT() string {
	return g.wkt
}

// SetWKT sets the Well-Known Text for the Geometry.
func (g *Geometry) SetWKT(wkt string) {
	g.wkt = wkt
}

// String returns the string representation of the Geometry (same as WKT).
func (g Geometry) String() string {
	return g.WKT()
}

// FormatParam implements the sql.ParamFormatter interface to format the placeholder for a Geometry parameter.
func (g Geometry) FormatParam(placeholder string, info *sql.StmtInfo) string {
	return "ST_GeomFromText(" + placeholder + ", 4326)"
}

// GeometrySchemaType defines the schema-type of the Geometry type for different dialects.
func GeometrySchemaType() map[string]string {
	return map[string]string{
		dialect.Postgres: "geometry",
	}
}

// MarshalJSON marshals the Geometry to JSON.
func (g Geometry) MarshalJSON() ([]byte, error) {
	if g.wkt == "" {
		return nil, nil
	}

	// Wrap the WKT with quotes and return as JSON.
	return []byte(`"` + g.wkt + `"`), nil
}

// UnmarshalJSON unmarshals the Geometry from JSON.
func (g *Geometry) UnmarshalJSON(data []byte) error {
	// Remove the surrounding quotes from JSON data and store as WKT.
	g.wkt = strings.ReplaceAll(string(data), `"`, "")
	return nil
}

// UnmarshalGQL implements the interface graphql.Unmarshaler.
// UnmarshalGQL unmarshals the Geometry from GraphQL query variables.
func (g *Geometry) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("type Geometry must be a string")
	}

	// Unmarshal the GraphQL string value into the Geometry's WKT.
	return json.Unmarshal([]byte(`"`+value+`"`), g)
}

// MarshalGQL implements the interface graphql.Marshaler.
// MarshalGQL marshals the Geometry to GraphQL response.
func (g Geometry) MarshalGQL(w io.Writer) {
	// Encode the Geometry as JSON and write to the given io.Writer.
	_ = json.NewEncoder(w).Encode(g)
}
