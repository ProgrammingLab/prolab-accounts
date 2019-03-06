// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Invitation is an object representing the database table.
type Invitation struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code      string    `boil:"code" json:"code" toml:"code" yaml:"code"`
	Email     string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	InviterID int64     `boil:"inviter_id" json:"inviter_id" toml:"inviter_id" yaml:"inviter_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *invitationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L invitationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InvitationColumns = struct {
	ID        string
	Code      string
	Email     string
	InviterID string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Code:      "code",
	Email:     "email",
	InviterID: "inviter_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var InvitationWhere = struct {
	ID        whereHelperint64
	Code      whereHelperstring
	Email     whereHelperstring
	InviterID whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: `id`},
	Code:      whereHelperstring{field: `code`},
	Email:     whereHelperstring{field: `email`},
	InviterID: whereHelperint64{field: `inviter_id`},
	CreatedAt: whereHelpertime_Time{field: `created_at`},
	UpdatedAt: whereHelpertime_Time{field: `updated_at`},
}

// InvitationRels is where relationship names are stored.
var InvitationRels = struct {
	Inviter string
}{
	Inviter: "Inviter",
}

// invitationR is where relationships are stored.
type invitationR struct {
	Inviter *User
}

// NewStruct creates a new relationship struct
func (*invitationR) NewStruct() *invitationR {
	return &invitationR{}
}

// invitationL is where Load methods for each relationship are stored.
type invitationL struct{}

var (
	invitationColumns               = []string{"id", "code", "email", "inviter_id", "created_at", "updated_at"}
	invitationColumnsWithoutDefault = []string{"code", "email", "inviter_id", "created_at", "updated_at"}
	invitationColumnsWithDefault    = []string{"id"}
	invitationPrimaryKeyColumns     = []string{"id"}
)

type (
	// InvitationSlice is an alias for a slice of pointers to Invitation.
	// This should generally be used opposed to []Invitation.
	InvitationSlice []*Invitation
	// InvitationHook is the signature for custom Invitation hook methods
	InvitationHook func(context.Context, boil.ContextExecutor, *Invitation) error

	invitationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	invitationType                 = reflect.TypeOf(&Invitation{})
	invitationMapping              = queries.MakeStructMapping(invitationType)
	invitationPrimaryKeyMapping, _ = queries.BindMapping(invitationType, invitationMapping, invitationPrimaryKeyColumns)
	invitationInsertCacheMut       sync.RWMutex
	invitationInsertCache          = make(map[string]insertCache)
	invitationUpdateCacheMut       sync.RWMutex
	invitationUpdateCache          = make(map[string]updateCache)
	invitationUpsertCacheMut       sync.RWMutex
	invitationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var invitationBeforeInsertHooks []InvitationHook
var invitationBeforeUpdateHooks []InvitationHook
var invitationBeforeDeleteHooks []InvitationHook
var invitationBeforeUpsertHooks []InvitationHook

var invitationAfterInsertHooks []InvitationHook
var invitationAfterSelectHooks []InvitationHook
var invitationAfterUpdateHooks []InvitationHook
var invitationAfterDeleteHooks []InvitationHook
var invitationAfterUpsertHooks []InvitationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Invitation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Invitation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Invitation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Invitation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Invitation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Invitation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Invitation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Invitation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Invitation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invitationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInvitationHook registers your hook function for all future operations.
func AddInvitationHook(hookPoint boil.HookPoint, invitationHook InvitationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		invitationBeforeInsertHooks = append(invitationBeforeInsertHooks, invitationHook)
	case boil.BeforeUpdateHook:
		invitationBeforeUpdateHooks = append(invitationBeforeUpdateHooks, invitationHook)
	case boil.BeforeDeleteHook:
		invitationBeforeDeleteHooks = append(invitationBeforeDeleteHooks, invitationHook)
	case boil.BeforeUpsertHook:
		invitationBeforeUpsertHooks = append(invitationBeforeUpsertHooks, invitationHook)
	case boil.AfterInsertHook:
		invitationAfterInsertHooks = append(invitationAfterInsertHooks, invitationHook)
	case boil.AfterSelectHook:
		invitationAfterSelectHooks = append(invitationAfterSelectHooks, invitationHook)
	case boil.AfterUpdateHook:
		invitationAfterUpdateHooks = append(invitationAfterUpdateHooks, invitationHook)
	case boil.AfterDeleteHook:
		invitationAfterDeleteHooks = append(invitationAfterDeleteHooks, invitationHook)
	case boil.AfterUpsertHook:
		invitationAfterUpsertHooks = append(invitationAfterUpsertHooks, invitationHook)
	}
}

// One returns a single invitation record from the query.
func (q invitationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Invitation, error) {
	o := &Invitation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: failed to execute a one query for invitations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Invitation records from the query.
func (q invitationQuery) All(ctx context.Context, exec boil.ContextExecutor) (InvitationSlice, error) {
	var o []*Invitation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "record: failed to assign all query results to Invitation slice")
	}

	if len(invitationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Invitation records in the query.
func (q invitationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to count invitations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q invitationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "record: failed to check if invitations exists")
	}

	return count > 0, nil
}

// Inviter pointed to by the foreign key.
func (o *Invitation) Inviter(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.InviterID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadInviter allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (invitationL) LoadInviter(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvitation interface{}, mods queries.Applicator) error {
	var slice []*Invitation
	var object *Invitation

	if singular {
		object = maybeInvitation.(*Invitation)
	} else {
		slice = *maybeInvitation.(*[]*Invitation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &invitationR{}
		}
		args = append(args, object.InviterID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &invitationR{}
			}

			for _, a := range args {
				if a == obj.InviterID {
					continue Outer
				}
			}

			args = append(args, obj.InviterID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(invitationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Inviter = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.InviterInvitations = append(foreign.R.InviterInvitations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.InviterID == foreign.ID {
				local.R.Inviter = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.InviterInvitations = append(foreign.R.InviterInvitations, local)
				break
			}
		}
	}

	return nil
}

// SetInviter of the invitation to the related item.
// Sets o.R.Inviter to related.
// Adds o to related.R.InviterInvitations.
func (o *Invitation) SetInviter(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"invitations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"inviter_id"}),
		strmangle.WhereClause("\"", "\"", 2, invitationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.InviterID = related.ID
	if o.R == nil {
		o.R = &invitationR{
			Inviter: related,
		}
	} else {
		o.R.Inviter = related
	}

	if related.R == nil {
		related.R = &userR{
			InviterInvitations: InvitationSlice{o},
		}
	} else {
		related.R.InviterInvitations = append(related.R.InviterInvitations, o)
	}

	return nil
}

// Invitations retrieves all the records using an executor.
func Invitations(mods ...qm.QueryMod) invitationQuery {
	mods = append(mods, qm.From("\"invitations\""))
	return invitationQuery{NewQuery(mods...)}
}

// FindInvitation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvitation(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Invitation, error) {
	invitationObj := &Invitation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"invitations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, invitationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: unable to select from invitations")
	}

	return invitationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Invitation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("record: no invitations provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invitationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	invitationInsertCacheMut.RLock()
	cache, cached := invitationInsertCache[key]
	invitationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			invitationColumns,
			invitationColumnsWithDefault,
			invitationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(invitationType, invitationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(invitationType, invitationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invitations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invitations\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "record: unable to insert into invitations")
	}

	if !cached {
		invitationInsertCacheMut.Lock()
		invitationInsertCache[key] = cache
		invitationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Invitation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Invitation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	invitationUpdateCacheMut.RLock()
	cache, cached := invitationUpdateCache[key]
	invitationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			invitationColumns,
			invitationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("record: unable to update invitations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"invitations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, invitationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(invitationType, invitationMapping, append(wl, invitationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update invitations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by update for invitations")
	}

	if !cached {
		invitationUpdateCacheMut.Lock()
		invitationUpdateCache[key] = cache
		invitationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q invitationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all for invitations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected for invitations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InvitationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("record: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"invitations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, invitationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all in invitation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected all in update all invitation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Invitation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("record: no invitations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invitationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	invitationUpsertCacheMut.RLock()
	cache, cached := invitationUpsertCache[key]
	invitationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			invitationColumns,
			invitationColumnsWithDefault,
			invitationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			invitationColumns,
			invitationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("record: unable to upsert invitations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(invitationPrimaryKeyColumns))
			copy(conflict, invitationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"invitations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(invitationType, invitationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(invitationType, invitationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "record: unable to upsert invitations")
	}

	if !cached {
		invitationUpsertCacheMut.Lock()
		invitationUpsertCache[key] = cache
		invitationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Invitation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Invitation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Invitation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), invitationPrimaryKeyMapping)
	sql := "DELETE FROM \"invitations\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete from invitations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by delete for invitations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q invitationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("record: no invitationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from invitations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for invitations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InvitationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Invitation slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(invitationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"invitations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invitationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from invitation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for invitations")
	}

	if len(invitationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Invitation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvitation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InvitationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InvitationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"invitations\".* FROM \"invitations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invitationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "record: unable to reload all in InvitationSlice")
	}

	*o = slice

	return nil
}

// InvitationExists checks if the Invitation row exists.
func InvitationExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"invitations\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "record: unable to check if invitations exists")
	}

	return exists, nil
}
