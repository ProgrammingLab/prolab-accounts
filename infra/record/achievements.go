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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Achievement is an object representing the database table.
type Achievement struct {
	ID            int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title         string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Award         string      `boil:"award" json:"award" toml:"award" yaml:"award"`
	URL           string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	Description   string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	ImageFilename null.String `boil:"image_filename" json:"image_filename,omitempty" toml:"image_filename" yaml:"image_filename,omitempty"`
	HappenedAt    time.Time   `boil:"happened_at" json:"happened_at" toml:"happened_at" yaml:"happened_at"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *achievementR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L achievementL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AchievementColumns = struct {
	ID            string
	Title         string
	Award         string
	URL           string
	Description   string
	ImageFilename string
	HappenedAt    string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Title:         "title",
	Award:         "award",
	URL:           "url",
	Description:   "description",
	ImageFilename: "image_filename",
	HappenedAt:    "happened_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AchievementWhere = struct {
	ID            whereHelperint64
	Title         whereHelperstring
	Award         whereHelperstring
	URL           whereHelperstring
	Description   whereHelperstring
	ImageFilename whereHelpernull_String
	HappenedAt    whereHelpertime_Time
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperint64{field: `id`},
	Title:         whereHelperstring{field: `title`},
	Award:         whereHelperstring{field: `award`},
	URL:           whereHelperstring{field: `url`},
	Description:   whereHelperstring{field: `description`},
	ImageFilename: whereHelpernull_String{field: `image_filename`},
	HappenedAt:    whereHelpertime_Time{field: `happened_at`},
	CreatedAt:     whereHelpertime_Time{field: `created_at`},
	UpdatedAt:     whereHelpertime_Time{field: `updated_at`},
}

// AchievementRels is where relationship names are stored.
var AchievementRels = struct {
	AchievementUsers string
}{
	AchievementUsers: "AchievementUsers",
}

// achievementR is where relationships are stored.
type achievementR struct {
	AchievementUsers AchievementUserSlice
}

// NewStruct creates a new relationship struct
func (*achievementR) NewStruct() *achievementR {
	return &achievementR{}
}

// achievementL is where Load methods for each relationship are stored.
type achievementL struct{}

var (
	achievementColumns               = []string{"id", "title", "award", "url", "description", "image_filename", "happened_at", "created_at", "updated_at"}
	achievementColumnsWithoutDefault = []string{"title", "award", "url", "description", "image_filename", "happened_at", "created_at", "updated_at"}
	achievementColumnsWithDefault    = []string{"id"}
	achievementPrimaryKeyColumns     = []string{"id"}
)

type (
	// AchievementSlice is an alias for a slice of pointers to Achievement.
	// This should generally be used opposed to []Achievement.
	AchievementSlice []*Achievement
	// AchievementHook is the signature for custom Achievement hook methods
	AchievementHook func(context.Context, boil.ContextExecutor, *Achievement) error

	achievementQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	achievementType                 = reflect.TypeOf(&Achievement{})
	achievementMapping              = queries.MakeStructMapping(achievementType)
	achievementPrimaryKeyMapping, _ = queries.BindMapping(achievementType, achievementMapping, achievementPrimaryKeyColumns)
	achievementInsertCacheMut       sync.RWMutex
	achievementInsertCache          = make(map[string]insertCache)
	achievementUpdateCacheMut       sync.RWMutex
	achievementUpdateCache          = make(map[string]updateCache)
	achievementUpsertCacheMut       sync.RWMutex
	achievementUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var achievementBeforeInsertHooks []AchievementHook
var achievementBeforeUpdateHooks []AchievementHook
var achievementBeforeDeleteHooks []AchievementHook
var achievementBeforeUpsertHooks []AchievementHook

var achievementAfterInsertHooks []AchievementHook
var achievementAfterSelectHooks []AchievementHook
var achievementAfterUpdateHooks []AchievementHook
var achievementAfterDeleteHooks []AchievementHook
var achievementAfterUpsertHooks []AchievementHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Achievement) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Achievement) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Achievement) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Achievement) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Achievement) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Achievement) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Achievement) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Achievement) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Achievement) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range achievementAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAchievementHook registers your hook function for all future operations.
func AddAchievementHook(hookPoint boil.HookPoint, achievementHook AchievementHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		achievementBeforeInsertHooks = append(achievementBeforeInsertHooks, achievementHook)
	case boil.BeforeUpdateHook:
		achievementBeforeUpdateHooks = append(achievementBeforeUpdateHooks, achievementHook)
	case boil.BeforeDeleteHook:
		achievementBeforeDeleteHooks = append(achievementBeforeDeleteHooks, achievementHook)
	case boil.BeforeUpsertHook:
		achievementBeforeUpsertHooks = append(achievementBeforeUpsertHooks, achievementHook)
	case boil.AfterInsertHook:
		achievementAfterInsertHooks = append(achievementAfterInsertHooks, achievementHook)
	case boil.AfterSelectHook:
		achievementAfterSelectHooks = append(achievementAfterSelectHooks, achievementHook)
	case boil.AfterUpdateHook:
		achievementAfterUpdateHooks = append(achievementAfterUpdateHooks, achievementHook)
	case boil.AfterDeleteHook:
		achievementAfterDeleteHooks = append(achievementAfterDeleteHooks, achievementHook)
	case boil.AfterUpsertHook:
		achievementAfterUpsertHooks = append(achievementAfterUpsertHooks, achievementHook)
	}
}

// One returns a single achievement record from the query.
func (q achievementQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Achievement, error) {
	o := &Achievement{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: failed to execute a one query for achievements")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Achievement records from the query.
func (q achievementQuery) All(ctx context.Context, exec boil.ContextExecutor) (AchievementSlice, error) {
	var o []*Achievement

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "record: failed to assign all query results to Achievement slice")
	}

	if len(achievementAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Achievement records in the query.
func (q achievementQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to count achievements rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q achievementQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "record: failed to check if achievements exists")
	}

	return count > 0, nil
}

// AchievementUsers retrieves all the achievement_user's AchievementUsers with an executor.
func (o *Achievement) AchievementUsers(mods ...qm.QueryMod) achievementUserQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"achievement_users\".\"achievement_id\"=?", o.ID),
	)

	query := AchievementUsers(queryMods...)
	queries.SetFrom(query.Query, "\"achievement_users\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"achievement_users\".*"})
	}

	return query
}

// LoadAchievementUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (achievementL) LoadAchievementUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAchievement interface{}, mods queries.Applicator) error {
	var slice []*Achievement
	var object *Achievement

	if singular {
		object = maybeAchievement.(*Achievement)
	} else {
		slice = *maybeAchievement.(*[]*Achievement)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &achievementR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &achievementR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`achievement_users`), qm.WhereIn(`achievement_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load achievement_users")
	}

	var resultSlice []*AchievementUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice achievement_users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on achievement_users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for achievement_users")
	}

	if len(achievementUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AchievementUsers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &achievementUserR{}
			}
			foreign.R.Achievement = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.AchievementID {
				local.R.AchievementUsers = append(local.R.AchievementUsers, foreign)
				if foreign.R == nil {
					foreign.R = &achievementUserR{}
				}
				foreign.R.Achievement = local
				break
			}
		}
	}

	return nil
}

// AddAchievementUsers adds the given related objects to the existing relationships
// of the achievement, optionally inserting them as new records.
// Appends related to o.R.AchievementUsers.
// Sets related.R.Achievement appropriately.
func (o *Achievement) AddAchievementUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AchievementUser) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.AchievementID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"achievement_users\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"achievement_id"}),
				strmangle.WhereClause("\"", "\"", 2, achievementUserPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.AchievementID = o.ID
		}
	}

	if o.R == nil {
		o.R = &achievementR{
			AchievementUsers: related,
		}
	} else {
		o.R.AchievementUsers = append(o.R.AchievementUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &achievementUserR{
				Achievement: o,
			}
		} else {
			rel.R.Achievement = o
		}
	}
	return nil
}

// Achievements retrieves all the records using an executor.
func Achievements(mods ...qm.QueryMod) achievementQuery {
	mods = append(mods, qm.From("\"achievements\""))
	return achievementQuery{NewQuery(mods...)}
}

// FindAchievement retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAchievement(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Achievement, error) {
	achievementObj := &Achievement{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"achievements\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, achievementObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: unable to select from achievements")
	}

	return achievementObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Achievement) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("record: no achievements provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(achievementColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	achievementInsertCacheMut.RLock()
	cache, cached := achievementInsertCache[key]
	achievementInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			achievementColumns,
			achievementColumnsWithDefault,
			achievementColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(achievementType, achievementMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(achievementType, achievementMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"achievements\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"achievements\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "record: unable to insert into achievements")
	}

	if !cached {
		achievementInsertCacheMut.Lock()
		achievementInsertCache[key] = cache
		achievementInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Achievement.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Achievement) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	achievementUpdateCacheMut.RLock()
	cache, cached := achievementUpdateCache[key]
	achievementUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			achievementColumns,
			achievementPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("record: unable to update achievements, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"achievements\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, achievementPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(achievementType, achievementMapping, append(wl, achievementPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "record: unable to update achievements row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by update for achievements")
	}

	if !cached {
		achievementUpdateCacheMut.Lock()
		achievementUpdateCache[key] = cache
		achievementUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q achievementQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all for achievements")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected for achievements")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AchievementSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), achievementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"achievements\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, achievementPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all in achievement slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected all in update all achievement")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Achievement) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("record: no achievements provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(achievementColumnsWithDefault, o)

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

	achievementUpsertCacheMut.RLock()
	cache, cached := achievementUpsertCache[key]
	achievementUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			achievementColumns,
			achievementColumnsWithDefault,
			achievementColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			achievementColumns,
			achievementPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("record: unable to upsert achievements, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(achievementPrimaryKeyColumns))
			copy(conflict, achievementPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"achievements\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(achievementType, achievementMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(achievementType, achievementMapping, ret)
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
		return errors.Wrap(err, "record: unable to upsert achievements")
	}

	if !cached {
		achievementUpsertCacheMut.Lock()
		achievementUpsertCache[key] = cache
		achievementUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Achievement record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Achievement) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Achievement provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), achievementPrimaryKeyMapping)
	sql := "DELETE FROM \"achievements\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete from achievements")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by delete for achievements")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q achievementQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("record: no achievementQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from achievements")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for achievements")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AchievementSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Achievement slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(achievementBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), achievementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"achievements\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, achievementPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from achievement slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for achievements")
	}

	if len(achievementAfterDeleteHooks) != 0 {
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
func (o *Achievement) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAchievement(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AchievementSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AchievementSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), achievementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"achievements\".* FROM \"achievements\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, achievementPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "record: unable to reload all in AchievementSlice")
	}

	*o = slice

	return nil
}

// AchievementExists checks if the Achievement row exists.
func AchievementExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"achievements\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "record: unable to check if achievements exists")
	}

	return exists, nil
}