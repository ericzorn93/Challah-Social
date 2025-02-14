// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"apps/services/accounts-graphql/internal/graph/models"
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/google/uuid"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Entity() EntityResolver
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Account struct {
		CreatedAt    func(childComplexity int) int
		EmailAddress func(childComplexity int) int
		ID           func(childComplexity int) int
		UpdatedAt    func(childComplexity int) int
	}

	Entity struct {
		FindAccountByID func(childComplexity int, id uuid.UUID) int
	}

	Mutation struct {
		DeleteAccount func(childComplexity int, commonID uuid.UUID) int
		Empty         func(childComplexity int) int
	}

	Query struct {
		Account            func(childComplexity int, input models.RetrieveAccountInput) int
		Empty              func(childComplexity int) int
		Viewer             func(childComplexity int, commonID uuid.UUID) int
		__resolve__service func(childComplexity int) int
		__resolve_entities func(childComplexity int, representations []map[string]any) int
	}

	Viewer struct {
		CreatedAt    func(childComplexity int) int
		EmailAddress func(childComplexity int) int
		Empty        func(childComplexity int) int
		ID           func(childComplexity int) int
		UpdatedAt    func(childComplexity int) int
	}

	_Service struct {
		SDL func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]any) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Account.createdAt":
		if e.complexity.Account.CreatedAt == nil {
			break
		}

		return e.complexity.Account.CreatedAt(childComplexity), true

	case "Account.emailAddress":
		if e.complexity.Account.EmailAddress == nil {
			break
		}

		return e.complexity.Account.EmailAddress(childComplexity), true

	case "Account.id":
		if e.complexity.Account.ID == nil {
			break
		}

		return e.complexity.Account.ID(childComplexity), true

	case "Account.updatedAt":
		if e.complexity.Account.UpdatedAt == nil {
			break
		}

		return e.complexity.Account.UpdatedAt(childComplexity), true

	case "Entity.findAccountByID":
		if e.complexity.Entity.FindAccountByID == nil {
			break
		}

		args, err := ec.field_Entity_findAccountByID_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Entity.FindAccountByID(childComplexity, args["id"].(uuid.UUID)), true

	case "Mutation.deleteAccount":
		if e.complexity.Mutation.DeleteAccount == nil {
			break
		}

		args, err := ec.field_Mutation_deleteAccount_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteAccount(childComplexity, args["commonID"].(uuid.UUID)), true

	case "Mutation.empty":
		if e.complexity.Mutation.Empty == nil {
			break
		}

		return e.complexity.Mutation.Empty(childComplexity), true

	case "Query.account":
		if e.complexity.Query.Account == nil {
			break
		}

		args, err := ec.field_Query_account_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Account(childComplexity, args["input"].(models.RetrieveAccountInput)), true

	case "Query.empty":
		if e.complexity.Query.Empty == nil {
			break
		}

		return e.complexity.Query.Empty(childComplexity), true

	case "Query.viewer":
		if e.complexity.Query.Viewer == nil {
			break
		}

		args, err := ec.field_Query_viewer_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Viewer(childComplexity, args["commonID"].(uuid.UUID)), true

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

	case "Query._entities":
		if e.complexity.Query.__resolve_entities == nil {
			break
		}

		args, err := ec.field_Query__entities_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.__resolve_entities(childComplexity, args["representations"].([]map[string]any)), true

	case "Viewer.createdAt":
		if e.complexity.Viewer.CreatedAt == nil {
			break
		}

		return e.complexity.Viewer.CreatedAt(childComplexity), true

	case "Viewer.emailAddress":
		if e.complexity.Viewer.EmailAddress == nil {
			break
		}

		return e.complexity.Viewer.EmailAddress(childComplexity), true

	case "Viewer.empty":
		if e.complexity.Viewer.Empty == nil {
			break
		}

		return e.complexity.Viewer.Empty(childComplexity), true

	case "Viewer.id":
		if e.complexity.Viewer.ID == nil {
			break
		}

		return e.complexity.Viewer.ID(childComplexity), true

	case "Viewer.updatedAt":
		if e.complexity.Viewer.UpdatedAt == nil {
			break
		}

		return e.complexity.Viewer.UpdatedAt(childComplexity), true

	case "_Service.sdl":
		if e.complexity._Service.SDL == nil {
			break
		}

		return e.complexity._Service.SDL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	opCtx := graphql.GetOperationContext(ctx)
	ec := executionContext{opCtx, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputRetrieveAccountInput,
	)
	first := true

	switch opCtx.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, opCtx.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, opCtx.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schemas/accounts.graphql", Input: `"""
Account interface defines key
shared account properties
"""
interface AccountInterface {
  """
  The unique identifier for the account
  """
  id: UUID!
  """
  The email address of the account
  """
  emailAddress: String!
  """
  The createdAt time of the account
  """
  createdAt: Time!
  """
  The updatedAt time of the account
  """
  updatedAt: Time!
}

"""
Account has default account interface with additional
account properties
"""
type Account implements AccountInterface @key(fields: "id") {
  """
  The unique identifier for the account
  """
  id: UUID!
  """
  The email address of the account
  """
  emailAddress: String!
  """
  The createdAt time of the account
  """
  createdAt: Time!
  """
  The updatedAt time of the account
  """
  updatedAt: Time!
}

extend type Viewer implements AccountInterface {
  """
  The unique identifier for the account
  """
  id: UUID!
  """
  The email address of the account
  """
  emailAddress: String!
  """
  The createdAt time of the account
  """
  createdAt: Time!
  """
  The updatedAt time of the account
  """
  updatedAt: Time!
}

"""
Uses either the commonID or email address to fetch user
"""
input RetrieveAccountInput {
  """
  commonID is the unique identifier for an account
  """
  commonID: UUID

  """
  emailAddress defines the unique email address for the account
  """
  emailAddress: String
}

extend type Query {
  """
  Obtains the account by commonID or email address
  """
  account(input: RetrieveAccountInput!): Account
}

extend type Mutation {
  """
  Delete account by commonID
  """
  deleteAccount(commonID: UUID!): Time!
}
`, BuiltIn: false},
	{Name: "../schemas/schema.graphql", Input: `# GraphQL schema example
#
# https://gqlgen.com/getting-started/

# Scalars
scalar Time
scalar UUID

# Directives
directive @goModel(
  model: String
  models: [String!]
  forceGenerate: Boolean
) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION

directive @goField(
  forceResolver: Boolean
  name: String
  omittable: Boolean
) on INPUT_FIELD_DEFINITION | FIELD_DEFINITION

directive @goTag(
  key: String!
  value: String
) on INPUT_FIELD_DEFINITION | FIELD_DEFINITION

"""
Viewer is the root query object for the user
"""
type Viewer {
  empty: Boolean!
}

type Query {
  empty: Boolean!

  """
  Viewer is the root query object for the user
  """
  viewer(
    """
    The unique identifier for the viewer
    """
    commonID: UUID!
  ): Viewer @goField(forceResolver: true)
}

type Mutation {
  empty: Boolean!
}
`, BuiltIn: false},
	{Name: "../../../federation/directives.graphql", Input: `
	directive @authenticated on FIELD_DEFINITION | OBJECT | INTERFACE | SCALAR | ENUM
	directive @composeDirective(name: String!) repeatable on SCHEMA
	directive @extends on OBJECT | INTERFACE
	directive @external on OBJECT | FIELD_DEFINITION
	directive @key(fields: FieldSet!, resolvable: Boolean = true) repeatable on OBJECT | INTERFACE
	directive @inaccessible on
	  | ARGUMENT_DEFINITION
	  | ENUM
	  | ENUM_VALUE
	  | FIELD_DEFINITION
	  | INPUT_FIELD_DEFINITION
	  | INPUT_OBJECT
	  | INTERFACE
	  | OBJECT
	  | SCALAR
	  | UNION
	directive @interfaceObject on OBJECT
	directive @link(import: [String!], url: String!) repeatable on SCHEMA
	directive @override(from: String!, label: String) on FIELD_DEFINITION
	directive @policy(policies: [[federation__Policy!]!]!) on
	  | FIELD_DEFINITION
	  | OBJECT
	  | INTERFACE
	  | SCALAR
	  | ENUM
	directive @provides(fields: FieldSet!) on FIELD_DEFINITION
	directive @requires(fields: FieldSet!) on FIELD_DEFINITION
	directive @requiresScopes(scopes: [[federation__Scope!]!]!) on
	  | FIELD_DEFINITION
	  | OBJECT
	  | INTERFACE
	  | SCALAR
	  | ENUM
	directive @shareable repeatable on FIELD_DEFINITION | OBJECT
	directive @tag(name: String!) repeatable on
	  | ARGUMENT_DEFINITION
	  | ENUM
	  | ENUM_VALUE
	  | FIELD_DEFINITION
	  | INPUT_FIELD_DEFINITION
	  | INPUT_OBJECT
	  | INTERFACE
	  | OBJECT
	  | SCALAR
	  | UNION
	scalar _Any
	scalar FieldSet
	scalar federation__Policy
	scalar federation__Scope
`, BuiltIn: true},
	{Name: "../../../federation/entity.graphql", Input: `
# a union of all types that use the @key directive
union _Entity = Account

# fake type to build resolver interfaces for users to implement
type Entity {
	findAccountByID(id: UUID!,): Account!
}

type _Service {
  sdl: String
}

extend type Query {
  _entities(representations: [_Any!]!): [_Entity]!
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
