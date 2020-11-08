// Code generated by SQLBoiler 4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Discount is an object representing the database table.
type Discount struct {
	ID           int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name         string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	DiscountType string    `boil:"discount_type" json:"discount_type" toml:"discount_type" yaml:"discount_type"`
	Discount     int64     `boil:"discount" json:"discount" toml:"discount" yaml:"discount"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *discountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DiscountColumns = struct {
	ID           string
	Name         string
	DiscountType string
	Discount     string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	Name:         "name",
	DiscountType: "discount_type",
	Discount:     "discount",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// Generated where

var DiscountWhere = struct {
	ID           whereHelperint64
	Name         whereHelperstring
	DiscountType whereHelperstring
	Discount     whereHelperint64
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
}{
	ID:           whereHelperint64{field: "\"discounts\".\"id\""},
	Name:         whereHelperstring{field: "\"discounts\".\"name\""},
	DiscountType: whereHelperstring{field: "\"discounts\".\"discount_type\""},
	Discount:     whereHelperint64{field: "\"discounts\".\"discount\""},
	CreatedAt:    whereHelpertime_Time{field: "\"discounts\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"discounts\".\"updated_at\""},
}

// DiscountRels is where relationship names are stored.
var DiscountRels = struct {
	Coupons       string
	DiscountRules string
}{
	Coupons:       "Coupons",
	DiscountRules: "DiscountRules",
}

// discountR is where relationships are stored.
type discountR struct {
	Coupons       CouponSlice       `boil:"Coupons" json:"Coupons" toml:"Coupons" yaml:"Coupons"`
	DiscountRules DiscountRuleSlice `boil:"DiscountRules" json:"DiscountRules" toml:"DiscountRules" yaml:"DiscountRules"`
}

// NewStruct creates a new relationship struct
func (*discountR) NewStruct() *discountR {
	return &discountR{}
}

// discountL is where Load methods for each relationship are stored.
type discountL struct{}

var (
	discountAllColumns            = []string{"id", "name", "discount_type", "discount", "created_at", "updated_at"}
	discountColumnsWithoutDefault = []string{"name", "discount_type", "discount"}
	discountColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	discountPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscountSlice is an alias for a slice of pointers to Discount.
	// This should generally be used opposed to []Discount.
	DiscountSlice []*Discount
	// DiscountHook is the signature for custom Discount hook methods
	DiscountHook func(context.Context, boil.ContextExecutor, *Discount) error

	discountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discountType                 = reflect.TypeOf(&Discount{})
	discountMapping              = queries.MakeStructMapping(discountType)
	discountPrimaryKeyMapping, _ = queries.BindMapping(discountType, discountMapping, discountPrimaryKeyColumns)
	discountInsertCacheMut       sync.RWMutex
	discountInsertCache          = make(map[string]insertCache)
	discountUpdateCacheMut       sync.RWMutex
	discountUpdateCache          = make(map[string]updateCache)
	discountUpsertCacheMut       sync.RWMutex
	discountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var discountBeforeInsertHooks []DiscountHook
var discountBeforeUpdateHooks []DiscountHook
var discountBeforeDeleteHooks []DiscountHook
var discountBeforeUpsertHooks []DiscountHook

var discountAfterInsertHooks []DiscountHook
var discountAfterSelectHooks []DiscountHook
var discountAfterUpdateHooks []DiscountHook
var discountAfterDeleteHooks []DiscountHook
var discountAfterUpsertHooks []DiscountHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Discount) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Discount) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Discount) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Discount) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Discount) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Discount) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Discount) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Discount) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Discount) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range discountAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDiscountHook registers your hook function for all future operations.
func AddDiscountHook(hookPoint boil.HookPoint, discountHook DiscountHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		discountBeforeInsertHooks = append(discountBeforeInsertHooks, discountHook)
	case boil.BeforeUpdateHook:
		discountBeforeUpdateHooks = append(discountBeforeUpdateHooks, discountHook)
	case boil.BeforeDeleteHook:
		discountBeforeDeleteHooks = append(discountBeforeDeleteHooks, discountHook)
	case boil.BeforeUpsertHook:
		discountBeforeUpsertHooks = append(discountBeforeUpsertHooks, discountHook)
	case boil.AfterInsertHook:
		discountAfterInsertHooks = append(discountAfterInsertHooks, discountHook)
	case boil.AfterSelectHook:
		discountAfterSelectHooks = append(discountAfterSelectHooks, discountHook)
	case boil.AfterUpdateHook:
		discountAfterUpdateHooks = append(discountAfterUpdateHooks, discountHook)
	case boil.AfterDeleteHook:
		discountAfterDeleteHooks = append(discountAfterDeleteHooks, discountHook)
	case boil.AfterUpsertHook:
		discountAfterUpsertHooks = append(discountAfterUpsertHooks, discountHook)
	}
}

// One returns a single discount record from the query.
func (q discountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Discount, error) {
	o := &Discount{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for discounts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Discount records from the query.
func (q discountQuery) All(ctx context.Context, exec boil.ContextExecutor) (DiscountSlice, error) {
	var o []*Discount

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Discount slice")
	}

	if len(discountAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Discount records in the query.
func (q discountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count discounts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q discountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if discounts exists")
	}

	return count > 0, nil
}

// Coupons retrieves all the coupon's Coupons with an executor.
func (o *Discount) Coupons(mods ...qm.QueryMod) couponQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"coupons\".\"discount_id\"=?", o.ID),
	)

	query := Coupons(queryMods...)
	queries.SetFrom(query.Query, "\"coupons\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"coupons\".*"})
	}

	return query
}

// DiscountRules retrieves all the discount_rule's DiscountRules with an executor.
func (o *Discount) DiscountRules(mods ...qm.QueryMod) discountRuleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"discount_rules\".\"discount_id\"=?", o.ID),
	)

	query := DiscountRules(queryMods...)
	queries.SetFrom(query.Query, "\"discount_rules\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"discount_rules\".*"})
	}

	return query
}

// LoadCoupons allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (discountL) LoadCoupons(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDiscount interface{}, mods queries.Applicator) error {
	var slice []*Discount
	var object *Discount

	if singular {
		object = maybeDiscount.(*Discount)
	} else {
		slice = *maybeDiscount.(*[]*Discount)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &discountR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &discountR{}
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

	query := NewQuery(
		qm.From(`coupons`),
		qm.WhereIn(`coupons.discount_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load coupons")
	}

	var resultSlice []*Coupon
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice coupons")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on coupons")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for coupons")
	}

	if len(couponAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Coupons = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &couponR{}
			}
			foreign.R.Discount = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.DiscountID {
				local.R.Coupons = append(local.R.Coupons, foreign)
				if foreign.R == nil {
					foreign.R = &couponR{}
				}
				foreign.R.Discount = local
				break
			}
		}
	}

	return nil
}

// LoadDiscountRules allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (discountL) LoadDiscountRules(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDiscount interface{}, mods queries.Applicator) error {
	var slice []*Discount
	var object *Discount

	if singular {
		object = maybeDiscount.(*Discount)
	} else {
		slice = *maybeDiscount.(*[]*Discount)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &discountR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &discountR{}
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

	query := NewQuery(
		qm.From(`discount_rules`),
		qm.WhereIn(`discount_rules.discount_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discount_rules")
	}

	var resultSlice []*DiscountRule
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discount_rules")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on discount_rules")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for discount_rules")
	}

	if len(discountRuleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.DiscountRules = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &discountRuleR{}
			}
			foreign.R.Discount = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.DiscountID {
				local.R.DiscountRules = append(local.R.DiscountRules, foreign)
				if foreign.R == nil {
					foreign.R = &discountRuleR{}
				}
				foreign.R.Discount = local
				break
			}
		}
	}

	return nil
}

// AddCoupons adds the given related objects to the existing relationships
// of the discount, optionally inserting them as new records.
// Appends related to o.R.Coupons.
// Sets related.R.Discount appropriately.
func (o *Discount) AddCoupons(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Coupon) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DiscountID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"coupons\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"discount_id"}),
				strmangle.WhereClause("\"", "\"", 2, couponPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DiscountID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discountR{
			Coupons: related,
		}
	} else {
		o.R.Coupons = append(o.R.Coupons, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &couponR{
				Discount: o,
			}
		} else {
			rel.R.Discount = o
		}
	}
	return nil
}

// AddDiscountRules adds the given related objects to the existing relationships
// of the discount, optionally inserting them as new records.
// Appends related to o.R.DiscountRules.
// Sets related.R.Discount appropriately.
func (o *Discount) AddDiscountRules(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DiscountRule) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DiscountID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discount_rules\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"discount_id"}),
				strmangle.WhereClause("\"", "\"", 2, discountRulePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DiscountID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discountR{
			DiscountRules: related,
		}
	} else {
		o.R.DiscountRules = append(o.R.DiscountRules, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discountRuleR{
				Discount: o,
			}
		} else {
			rel.R.Discount = o
		}
	}
	return nil
}

// Discounts retrieves all the records using an executor.
func Discounts(mods ...qm.QueryMod) discountQuery {
	mods = append(mods, qm.From("\"discounts\""))
	return discountQuery{NewQuery(mods...)}
}

// FindDiscount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscount(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Discount, error) {
	discountObj := &Discount{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discounts\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, discountObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from discounts")
	}

	return discountObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Discount) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no discounts provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(discountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	discountInsertCacheMut.RLock()
	cache, cached := discountInsertCache[key]
	discountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			discountAllColumns,
			discountColumnsWithDefault,
			discountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(discountType, discountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discountType, discountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"discounts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"discounts\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "orm: unable to insert into discounts")
	}

	if !cached {
		discountInsertCacheMut.Lock()
		discountInsertCache[key] = cache
		discountInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Discount.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Discount) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	discountUpdateCacheMut.RLock()
	cache, cached := discountUpdateCache[key]
	discountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			discountAllColumns,
			discountPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update discounts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discounts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discountType, discountMapping, append(wl, discountPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update discounts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for discounts")
	}

	if !cached {
		discountUpdateCacheMut.Lock()
		discountUpdateCache[key] = cache
		discountUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q discountQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for discounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for discounts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscountSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"discounts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, discountPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in discount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all discount")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Discount) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no discounts provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(discountColumnsWithDefault, o)

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

	discountUpsertCacheMut.RLock()
	cache, cached := discountUpsertCache[key]
	discountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			discountAllColumns,
			discountColumnsWithDefault,
			discountColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			discountAllColumns,
			discountPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert discounts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discountPrimaryKeyColumns))
			copy(conflict, discountPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"discounts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(discountType, discountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discountType, discountMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
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
		return errors.Wrap(err, "orm: unable to upsert discounts")
	}

	if !cached {
		discountUpsertCacheMut.Lock()
		discountUpsertCache[key] = cache
		discountUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Discount record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Discount) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Discount provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discountPrimaryKeyMapping)
	sql := "DELETE FROM \"discounts\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from discounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for discounts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q discountQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no discountQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from discounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for discounts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscountSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(discountBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"discounts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, discountPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from discount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for discounts")
	}

	if len(discountAfterDeleteHooks) != 0 {
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
func (o *Discount) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDiscount(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscountSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DiscountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"discounts\".* FROM \"discounts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, discountPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in DiscountSlice")
	}

	*o = slice

	return nil
}

// DiscountExists checks if the Discount row exists.
func DiscountExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"discounts\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if discounts exists")
	}

	return exists, nil
}