// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/dkrasnovdev/siberiana-api/ent/art"
	"github.com/dkrasnovdev/siberiana-api/ent/artgenre"
	"github.com/dkrasnovdev/siberiana-api/ent/artifact"
	"github.com/dkrasnovdev/siberiana-api/ent/artstyle"
	"github.com/dkrasnovdev/siberiana-api/ent/auditlog"
	"github.com/dkrasnovdev/siberiana-api/ent/book"
	"github.com/dkrasnovdev/siberiana-api/ent/bookgenre"
	"github.com/dkrasnovdev/siberiana-api/ent/category"
	"github.com/dkrasnovdev/siberiana-api/ent/collection"
	"github.com/dkrasnovdev/siberiana-api/ent/country"
	"github.com/dkrasnovdev/siberiana-api/ent/culture"
	"github.com/dkrasnovdev/siberiana-api/ent/district"
	"github.com/dkrasnovdev/siberiana-api/ent/ethnos"
	"github.com/dkrasnovdev/siberiana-api/ent/favourite"
	"github.com/dkrasnovdev/siberiana-api/ent/interview"
	"github.com/dkrasnovdev/siberiana-api/ent/keyword"
	"github.com/dkrasnovdev/siberiana-api/ent/license"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/medium"
	"github.com/dkrasnovdev/siberiana-api/ent/model"
	"github.com/dkrasnovdev/siberiana-api/ent/monument"
	"github.com/dkrasnovdev/siberiana-api/ent/mound"
	"github.com/dkrasnovdev/siberiana-api/ent/organization"
	"github.com/dkrasnovdev/siberiana-api/ent/periodical"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/personal"
	"github.com/dkrasnovdev/siberiana-api/ent/petroglyph"
	"github.com/dkrasnovdev/siberiana-api/ent/project"
	"github.com/dkrasnovdev/siberiana-api/ent/protectedarea"
	"github.com/dkrasnovdev/siberiana-api/ent/protectedareacategory"
	"github.com/dkrasnovdev/siberiana-api/ent/protectedareapicture"
	"github.com/dkrasnovdev/siberiana-api/ent/proxy"
	"github.com/dkrasnovdev/siberiana-api/ent/publication"
	"github.com/dkrasnovdev/siberiana-api/ent/publisher"
	"github.com/dkrasnovdev/siberiana-api/ent/region"
	"github.com/dkrasnovdev/siberiana-api/ent/set"
	"github.com/dkrasnovdev/siberiana-api/ent/settlement"
	"github.com/dkrasnovdev/siberiana-api/ent/technique"
	"github.com/dkrasnovdev/siberiana-api/ent/visit"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var artImplementors = []string{"Art", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Art) IsNode() {}

var artgenreImplementors = []string{"ArtGenre", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ArtGenre) IsNode() {}

var artstyleImplementors = []string{"ArtStyle", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ArtStyle) IsNode() {}

var artifactImplementors = []string{"Artifact", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Artifact) IsNode() {}

var auditlogImplementors = []string{"AuditLog", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AuditLog) IsNode() {}

var bookImplementors = []string{"Book", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Book) IsNode() {}

var bookgenreImplementors = []string{"BookGenre", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*BookGenre) IsNode() {}

var categoryImplementors = []string{"Category", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Category) IsNode() {}

var collectionImplementors = []string{"Collection", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Collection) IsNode() {}

var countryImplementors = []string{"Country", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Country) IsNode() {}

var cultureImplementors = []string{"Culture", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Culture) IsNode() {}

var districtImplementors = []string{"District", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*District) IsNode() {}

var ethnosImplementors = []string{"Ethnos", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Ethnos) IsNode() {}

var favouriteImplementors = []string{"Favourite", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Favourite) IsNode() {}

var interviewImplementors = []string{"Interview", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Interview) IsNode() {}

var keywordImplementors = []string{"Keyword", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Keyword) IsNode() {}

var licenseImplementors = []string{"License", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*License) IsNode() {}

var locationImplementors = []string{"Location", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Location) IsNode() {}

var mediumImplementors = []string{"Medium", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Medium) IsNode() {}

var modelImplementors = []string{"Model", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Model) IsNode() {}

var monumentImplementors = []string{"Monument", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Monument) IsNode() {}

var moundImplementors = []string{"Mound", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Mound) IsNode() {}

var organizationImplementors = []string{"Organization", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Organization) IsNode() {}

var periodicalImplementors = []string{"Periodical", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Periodical) IsNode() {}

var personImplementors = []string{"Person", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Person) IsNode() {}

var personalImplementors = []string{"Personal", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Personal) IsNode() {}

var petroglyphImplementors = []string{"Petroglyph", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Petroglyph) IsNode() {}

var projectImplementors = []string{"Project", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Project) IsNode() {}

var protectedareaImplementors = []string{"ProtectedArea", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ProtectedArea) IsNode() {}

var protectedareacategoryImplementors = []string{"ProtectedAreaCategory", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ProtectedAreaCategory) IsNode() {}

var protectedareapictureImplementors = []string{"ProtectedAreaPicture", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ProtectedAreaPicture) IsNode() {}

var proxyImplementors = []string{"Proxy", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Proxy) IsNode() {}

var publicationImplementors = []string{"Publication", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Publication) IsNode() {}

var publisherImplementors = []string{"Publisher", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Publisher) IsNode() {}

var regionImplementors = []string{"Region", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Region) IsNode() {}

var setImplementors = []string{"Set", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Set) IsNode() {}

var settlementImplementors = []string{"Settlement", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Settlement) IsNode() {}

var techniqueImplementors = []string{"Technique", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Technique) IsNode() {}

var visitImplementors = []string{"Visit", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Visit) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case art.Table:
		query := c.Art.Query().
			Where(art.ID(id))
		query, err := query.CollectFields(ctx, artImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case artgenre.Table:
		query := c.ArtGenre.Query().
			Where(artgenre.ID(id))
		query, err := query.CollectFields(ctx, artgenreImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case artstyle.Table:
		query := c.ArtStyle.Query().
			Where(artstyle.ID(id))
		query, err := query.CollectFields(ctx, artstyleImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case artifact.Table:
		query := c.Artifact.Query().
			Where(artifact.ID(id))
		query, err := query.CollectFields(ctx, artifactImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case auditlog.Table:
		query := c.AuditLog.Query().
			Where(auditlog.ID(id))
		query, err := query.CollectFields(ctx, auditlogImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case book.Table:
		query := c.Book.Query().
			Where(book.ID(id))
		query, err := query.CollectFields(ctx, bookImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case bookgenre.Table:
		query := c.BookGenre.Query().
			Where(bookgenre.ID(id))
		query, err := query.CollectFields(ctx, bookgenreImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case category.Table:
		query := c.Category.Query().
			Where(category.ID(id))
		query, err := query.CollectFields(ctx, categoryImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case collection.Table:
		query := c.Collection.Query().
			Where(collection.ID(id))
		query, err := query.CollectFields(ctx, collectionImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case country.Table:
		query := c.Country.Query().
			Where(country.ID(id))
		query, err := query.CollectFields(ctx, countryImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case culture.Table:
		query := c.Culture.Query().
			Where(culture.ID(id))
		query, err := query.CollectFields(ctx, cultureImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case district.Table:
		query := c.District.Query().
			Where(district.ID(id))
		query, err := query.CollectFields(ctx, districtImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case ethnos.Table:
		query := c.Ethnos.Query().
			Where(ethnos.ID(id))
		query, err := query.CollectFields(ctx, ethnosImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case favourite.Table:
		query := c.Favourite.Query().
			Where(favourite.ID(id))
		query, err := query.CollectFields(ctx, favouriteImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case interview.Table:
		query := c.Interview.Query().
			Where(interview.ID(id))
		query, err := query.CollectFields(ctx, interviewImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case keyword.Table:
		query := c.Keyword.Query().
			Where(keyword.ID(id))
		query, err := query.CollectFields(ctx, keywordImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case license.Table:
		query := c.License.Query().
			Where(license.ID(id))
		query, err := query.CollectFields(ctx, licenseImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case location.Table:
		query := c.Location.Query().
			Where(location.ID(id))
		query, err := query.CollectFields(ctx, locationImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case medium.Table:
		query := c.Medium.Query().
			Where(medium.ID(id))
		query, err := query.CollectFields(ctx, mediumImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case model.Table:
		query := c.Model.Query().
			Where(model.ID(id))
		query, err := query.CollectFields(ctx, modelImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case monument.Table:
		query := c.Monument.Query().
			Where(monument.ID(id))
		query, err := query.CollectFields(ctx, monumentImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case mound.Table:
		query := c.Mound.Query().
			Where(mound.ID(id))
		query, err := query.CollectFields(ctx, moundImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case organization.Table:
		query := c.Organization.Query().
			Where(organization.ID(id))
		query, err := query.CollectFields(ctx, organizationImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case periodical.Table:
		query := c.Periodical.Query().
			Where(periodical.ID(id))
		query, err := query.CollectFields(ctx, periodicalImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case person.Table:
		query := c.Person.Query().
			Where(person.ID(id))
		query, err := query.CollectFields(ctx, personImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case personal.Table:
		query := c.Personal.Query().
			Where(personal.ID(id))
		query, err := query.CollectFields(ctx, personalImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case petroglyph.Table:
		query := c.Petroglyph.Query().
			Where(petroglyph.ID(id))
		query, err := query.CollectFields(ctx, petroglyphImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case project.Table:
		query := c.Project.Query().
			Where(project.ID(id))
		query, err := query.CollectFields(ctx, projectImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case protectedarea.Table:
		query := c.ProtectedArea.Query().
			Where(protectedarea.ID(id))
		query, err := query.CollectFields(ctx, protectedareaImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case protectedareacategory.Table:
		query := c.ProtectedAreaCategory.Query().
			Where(protectedareacategory.ID(id))
		query, err := query.CollectFields(ctx, protectedareacategoryImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case protectedareapicture.Table:
		query := c.ProtectedAreaPicture.Query().
			Where(protectedareapicture.ID(id))
		query, err := query.CollectFields(ctx, protectedareapictureImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case proxy.Table:
		query := c.Proxy.Query().
			Where(proxy.ID(id))
		query, err := query.CollectFields(ctx, proxyImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case publication.Table:
		query := c.Publication.Query().
			Where(publication.ID(id))
		query, err := query.CollectFields(ctx, publicationImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case publisher.Table:
		query := c.Publisher.Query().
			Where(publisher.ID(id))
		query, err := query.CollectFields(ctx, publisherImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case region.Table:
		query := c.Region.Query().
			Where(region.ID(id))
		query, err := query.CollectFields(ctx, regionImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case set.Table:
		query := c.Set.Query().
			Where(set.ID(id))
		query, err := query.CollectFields(ctx, setImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case settlement.Table:
		query := c.Settlement.Query().
			Where(settlement.ID(id))
		query, err := query.CollectFields(ctx, settlementImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case technique.Table:
		query := c.Technique.Query().
			Where(technique.ID(id))
		query, err := query.CollectFields(ctx, techniqueImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case visit.Table:
		query := c.Visit.Query().
			Where(visit.ID(id))
		query, err := query.CollectFields(ctx, visitImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case art.Table:
		query := c.Art.Query().
			Where(art.IDIn(ids...))
		query, err := query.CollectFields(ctx, artImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case artgenre.Table:
		query := c.ArtGenre.Query().
			Where(artgenre.IDIn(ids...))
		query, err := query.CollectFields(ctx, artgenreImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case artstyle.Table:
		query := c.ArtStyle.Query().
			Where(artstyle.IDIn(ids...))
		query, err := query.CollectFields(ctx, artstyleImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case artifact.Table:
		query := c.Artifact.Query().
			Where(artifact.IDIn(ids...))
		query, err := query.CollectFields(ctx, artifactImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case auditlog.Table:
		query := c.AuditLog.Query().
			Where(auditlog.IDIn(ids...))
		query, err := query.CollectFields(ctx, auditlogImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case book.Table:
		query := c.Book.Query().
			Where(book.IDIn(ids...))
		query, err := query.CollectFields(ctx, bookImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case bookgenre.Table:
		query := c.BookGenre.Query().
			Where(bookgenre.IDIn(ids...))
		query, err := query.CollectFields(ctx, bookgenreImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case category.Table:
		query := c.Category.Query().
			Where(category.IDIn(ids...))
		query, err := query.CollectFields(ctx, categoryImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case collection.Table:
		query := c.Collection.Query().
			Where(collection.IDIn(ids...))
		query, err := query.CollectFields(ctx, collectionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case country.Table:
		query := c.Country.Query().
			Where(country.IDIn(ids...))
		query, err := query.CollectFields(ctx, countryImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case culture.Table:
		query := c.Culture.Query().
			Where(culture.IDIn(ids...))
		query, err := query.CollectFields(ctx, cultureImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case district.Table:
		query := c.District.Query().
			Where(district.IDIn(ids...))
		query, err := query.CollectFields(ctx, districtImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case ethnos.Table:
		query := c.Ethnos.Query().
			Where(ethnos.IDIn(ids...))
		query, err := query.CollectFields(ctx, ethnosImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case favourite.Table:
		query := c.Favourite.Query().
			Where(favourite.IDIn(ids...))
		query, err := query.CollectFields(ctx, favouriteImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case interview.Table:
		query := c.Interview.Query().
			Where(interview.IDIn(ids...))
		query, err := query.CollectFields(ctx, interviewImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case keyword.Table:
		query := c.Keyword.Query().
			Where(keyword.IDIn(ids...))
		query, err := query.CollectFields(ctx, keywordImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case license.Table:
		query := c.License.Query().
			Where(license.IDIn(ids...))
		query, err := query.CollectFields(ctx, licenseImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case location.Table:
		query := c.Location.Query().
			Where(location.IDIn(ids...))
		query, err := query.CollectFields(ctx, locationImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case medium.Table:
		query := c.Medium.Query().
			Where(medium.IDIn(ids...))
		query, err := query.CollectFields(ctx, mediumImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case model.Table:
		query := c.Model.Query().
			Where(model.IDIn(ids...))
		query, err := query.CollectFields(ctx, modelImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case monument.Table:
		query := c.Monument.Query().
			Where(monument.IDIn(ids...))
		query, err := query.CollectFields(ctx, monumentImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case mound.Table:
		query := c.Mound.Query().
			Where(mound.IDIn(ids...))
		query, err := query.CollectFields(ctx, moundImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case organization.Table:
		query := c.Organization.Query().
			Where(organization.IDIn(ids...))
		query, err := query.CollectFields(ctx, organizationImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case periodical.Table:
		query := c.Periodical.Query().
			Where(periodical.IDIn(ids...))
		query, err := query.CollectFields(ctx, periodicalImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case person.Table:
		query := c.Person.Query().
			Where(person.IDIn(ids...))
		query, err := query.CollectFields(ctx, personImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case personal.Table:
		query := c.Personal.Query().
			Where(personal.IDIn(ids...))
		query, err := query.CollectFields(ctx, personalImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case petroglyph.Table:
		query := c.Petroglyph.Query().
			Where(petroglyph.IDIn(ids...))
		query, err := query.CollectFields(ctx, petroglyphImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case project.Table:
		query := c.Project.Query().
			Where(project.IDIn(ids...))
		query, err := query.CollectFields(ctx, projectImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case protectedarea.Table:
		query := c.ProtectedArea.Query().
			Where(protectedarea.IDIn(ids...))
		query, err := query.CollectFields(ctx, protectedareaImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case protectedareacategory.Table:
		query := c.ProtectedAreaCategory.Query().
			Where(protectedareacategory.IDIn(ids...))
		query, err := query.CollectFields(ctx, protectedareacategoryImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case protectedareapicture.Table:
		query := c.ProtectedAreaPicture.Query().
			Where(protectedareapicture.IDIn(ids...))
		query, err := query.CollectFields(ctx, protectedareapictureImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case proxy.Table:
		query := c.Proxy.Query().
			Where(proxy.IDIn(ids...))
		query, err := query.CollectFields(ctx, proxyImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case publication.Table:
		query := c.Publication.Query().
			Where(publication.IDIn(ids...))
		query, err := query.CollectFields(ctx, publicationImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case publisher.Table:
		query := c.Publisher.Query().
			Where(publisher.IDIn(ids...))
		query, err := query.CollectFields(ctx, publisherImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case region.Table:
		query := c.Region.Query().
			Where(region.IDIn(ids...))
		query, err := query.CollectFields(ctx, regionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case set.Table:
		query := c.Set.Query().
			Where(set.IDIn(ids...))
		query, err := query.CollectFields(ctx, setImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case settlement.Table:
		query := c.Settlement.Query().
			Where(settlement.IDIn(ids...))
		query, err := query.CollectFields(ctx, settlementImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case technique.Table:
		query := c.Technique.Query().
			Where(technique.IDIn(ids...))
		query, err := query.CollectFields(ctx, techniqueImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case visit.Table:
		query := c.Visit.Query().
			Where(visit.IDIn(ids...))
		query, err := query.CollectFields(ctx, visitImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
