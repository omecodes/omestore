// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/omecodes/omestore/ent/access"
	"github.com/omecodes/omestore/ent/predicate"
)

// AccessDelete is the builder for deleting a Access entity.
type AccessDelete struct {
	config
	hooks      []Hook
	mutation   *AccessMutation
	predicates []predicate.Access
}

// Where adds a new predicate to the delete builder.
func (ad *AccessDelete) Where(ps ...predicate.Access) *AccessDelete {
	ad.predicates = append(ad.predicates, ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AccessDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ad.hooks) == 0 {
		affected, err = ad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccessMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ad.mutation = mutation
			affected, err = ad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ad.hooks) - 1; i >= 0; i-- {
			mut = ad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AccessDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AccessDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: access.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: access.FieldID,
			},
		},
	}
	if ps := ad.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
}

// AccessDeleteOne is the builder for deleting a single Access entity.
type AccessDeleteOne struct {
	ad *AccessDelete
}

// Exec executes the deletion query.
func (ado *AccessDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{access.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AccessDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
