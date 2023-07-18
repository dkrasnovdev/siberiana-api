package privacy

import (
	"context"

	"golang.org/x/exp/slices"
)

// UserViewer represents the user information obtained from the authentication process.
type UserViewer struct {
	Sub               string   `json:"sub"`                // Subject identifier
	Roles             []string `json:"roles"`              // User roles
	GivenName         string   `json:"given_name"`         // Given name
	PreferredUsername string   `json:"preferred_username"` // Preferred username
	FamilyName        string   `json:"family_name"`        // Family name
	Email             string   `json:"email"`              // Email address
}

// Viewer interface provides methods for accessing user information and roles.
type Viewer interface {
	GetUser() UserViewer
	GetSub() string
	GetPreferredUsername() string
	IsAdministrator() bool
	IsModerator() bool
	IsResearcher() bool
}

// GetUser returns the user information.
func (u UserViewer) GetUser() UserViewer {
	return u
}

// GetSub returns the sub of the user.
func (u UserViewer) GetSub() string {
	return u.Sub
}

// GetPreferredUsername returns the Preferred Username of the user.
func (u UserViewer) GetPreferredUsername() string {
	return u.PreferredUsername
}

// IsResearcher checks if the user has the "administartor" role.
func (u UserViewer) IsAdministrator() bool {
	return slices.Contains(u.Roles, "administartor")
}

// IsModerator checks if the user has the "moderator" role.
func (u UserViewer) IsModerator() bool {
	return slices.Contains(u.Roles, "moderator")
}

// IsResearcher checks if the user has the "researcher" role.
func (u UserViewer) IsResearcher() bool {
	return slices.Contains(u.Roles, "researcher")
}

// viewerCtxKey is a private key type used for context value.
type viewerCtxKey struct{}

// NewContext creates a new context with the provided Viewer value.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, viewerCtxKey{}, v)
}

// FromContext returns the Viewer value stored in the context.
// If the value is not found or the type is not compatible, it returns nil.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(viewerCtxKey{}).(Viewer)
	return v
}
