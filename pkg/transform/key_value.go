package transform

import (
	"fmt"

	"entgo.io/ent"
)

// KeyValue formats a map of keys and values as a slice of strings.
func KeyValue(m map[string]ent.Value) []string {
	var kv []string
	for k, v := range m {
		kv = append(kv, fmt.Sprintf("%v: %v", k, v))
	}
	return kv
}
