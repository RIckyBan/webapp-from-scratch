// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWebapps(t *testing.T) {
	t.Parallel()

	query := Webapps()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWebappsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebappsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Webapps().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebappsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WebappSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebappsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WebappExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Webapp exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WebappExists to return true, but got false.")
	}
}

func testWebappsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	webappFound, err := FindWebapp(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if webappFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWebappsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Webapps().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWebappsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Webapps().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWebappsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	webappOne := &Webapp{}
	webappTwo := &Webapp{}
	if err = randomize.Struct(seed, webappOne, webappDBTypes, false, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}
	if err = randomize.Struct(seed, webappTwo, webappDBTypes, false, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = webappOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = webappTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Webapps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWebappsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	webappOne := &Webapp{}
	webappTwo := &Webapp{}
	if err = randomize.Struct(seed, webappOne, webappDBTypes, false, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}
	if err = randomize.Struct(seed, webappTwo, webappDBTypes, false, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = webappOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = webappTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func webappBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func webappAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Webapp) error {
	*o = Webapp{}
	return nil
}

func testWebappsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Webapp{}
	o := &Webapp{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, webappDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Webapp object: %s", err)
	}

	AddWebappHook(boil.BeforeInsertHook, webappBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	webappBeforeInsertHooks = []WebappHook{}

	AddWebappHook(boil.AfterInsertHook, webappAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	webappAfterInsertHooks = []WebappHook{}

	AddWebappHook(boil.AfterSelectHook, webappAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	webappAfterSelectHooks = []WebappHook{}

	AddWebappHook(boil.BeforeUpdateHook, webappBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	webappBeforeUpdateHooks = []WebappHook{}

	AddWebappHook(boil.AfterUpdateHook, webappAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	webappAfterUpdateHooks = []WebappHook{}

	AddWebappHook(boil.BeforeDeleteHook, webappBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	webappBeforeDeleteHooks = []WebappHook{}

	AddWebappHook(boil.AfterDeleteHook, webappAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	webappAfterDeleteHooks = []WebappHook{}

	AddWebappHook(boil.BeforeUpsertHook, webappBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	webappBeforeUpsertHooks = []WebappHook{}

	AddWebappHook(boil.AfterUpsertHook, webappAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	webappAfterUpsertHooks = []WebappHook{}
}

func testWebappsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWebappsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(webappColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWebappsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWebappsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WebappSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWebappsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Webapps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	webappDBTypes = map[string]string{`ID`: `varchar`, `AppliedAt`: `datetime`}
	_             = bytes.MinRead
)

func testWebappsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(webappPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(webappAllColumns) == len(webappPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, webappDBTypes, true, webappPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWebappsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(webappAllColumns) == len(webappPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Webapp{}
	if err = randomize.Struct(seed, o, webappDBTypes, true, webappColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, webappDBTypes, true, webappPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(webappAllColumns, webappPrimaryKeyColumns) {
		fields = webappAllColumns
	} else {
		fields = strmangle.SetComplement(
			webappAllColumns,
			webappPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WebappSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testWebappsUpsert(t *testing.T) {
	t.Parallel()

	if len(webappAllColumns) == len(webappPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLWebappUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Webapp{}
	if err = randomize.Struct(seed, &o, webappDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Webapp: %s", err)
	}

	count, err := Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, webappDBTypes, false, webappPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Webapp struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Webapp: %s", err)
	}

	count, err = Webapps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
