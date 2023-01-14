// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Carts", testCarts)
	t.Run("Orders", testOrders)
	t.Run("Phones", testPhones)
	t.Run("Users", testUsers)
	t.Run("Webapps", testWebapps)
}

func TestDelete(t *testing.T) {
	t.Run("Carts", testCartsDelete)
	t.Run("Orders", testOrdersDelete)
	t.Run("Phones", testPhonesDelete)
	t.Run("Users", testUsersDelete)
	t.Run("Webapps", testWebappsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Carts", testCartsQueryDeleteAll)
	t.Run("Orders", testOrdersQueryDeleteAll)
	t.Run("Phones", testPhonesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("Webapps", testWebappsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Carts", testCartsSliceDeleteAll)
	t.Run("Orders", testOrdersSliceDeleteAll)
	t.Run("Phones", testPhonesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("Webapps", testWebappsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Carts", testCartsExists)
	t.Run("Orders", testOrdersExists)
	t.Run("Phones", testPhonesExists)
	t.Run("Users", testUsersExists)
	t.Run("Webapps", testWebappsExists)
}

func TestFind(t *testing.T) {
	t.Run("Carts", testCartsFind)
	t.Run("Orders", testOrdersFind)
	t.Run("Phones", testPhonesFind)
	t.Run("Users", testUsersFind)
	t.Run("Webapps", testWebappsFind)
}

func TestBind(t *testing.T) {
	t.Run("Carts", testCartsBind)
	t.Run("Orders", testOrdersBind)
	t.Run("Phones", testPhonesBind)
	t.Run("Users", testUsersBind)
	t.Run("Webapps", testWebappsBind)
}

func TestOne(t *testing.T) {
	t.Run("Carts", testCartsOne)
	t.Run("Orders", testOrdersOne)
	t.Run("Phones", testPhonesOne)
	t.Run("Users", testUsersOne)
	t.Run("Webapps", testWebappsOne)
}

func TestAll(t *testing.T) {
	t.Run("Carts", testCartsAll)
	t.Run("Orders", testOrdersAll)
	t.Run("Phones", testPhonesAll)
	t.Run("Users", testUsersAll)
	t.Run("Webapps", testWebappsAll)
}

func TestCount(t *testing.T) {
	t.Run("Carts", testCartsCount)
	t.Run("Orders", testOrdersCount)
	t.Run("Phones", testPhonesCount)
	t.Run("Users", testUsersCount)
	t.Run("Webapps", testWebappsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Carts", testCartsHooks)
	t.Run("Orders", testOrdersHooks)
	t.Run("Phones", testPhonesHooks)
	t.Run("Users", testUsersHooks)
	t.Run("Webapps", testWebappsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Carts", testCartsInsert)
	t.Run("Carts", testCartsInsertWhitelist)
	t.Run("Orders", testOrdersInsert)
	t.Run("Orders", testOrdersInsertWhitelist)
	t.Run("Phones", testPhonesInsert)
	t.Run("Phones", testPhonesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("Webapps", testWebappsInsert)
	t.Run("Webapps", testWebappsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CartToUserUsingUser", testCartToOneUserUsingUser)
	t.Run("CartToPhoneUsingPhone", testCartToOnePhoneUsingPhone)
	t.Run("OrderToUserUsingUser", testOrderToOneUserUsingUser)
	t.Run("OrderToPhoneUsingPhone", testOrderToOnePhoneUsingPhone)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PhoneToCarts", testPhoneToManyCarts)
	t.Run("PhoneToOrders", testPhoneToManyOrders)
	t.Run("UserToCarts", testUserToManyCarts)
	t.Run("UserToOrders", testUserToManyOrders)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CartToUserUsingCarts", testCartToOneSetOpUserUsingUser)
	t.Run("CartToPhoneUsingCarts", testCartToOneSetOpPhoneUsingPhone)
	t.Run("OrderToUserUsingOrders", testOrderToOneSetOpUserUsingUser)
	t.Run("OrderToPhoneUsingOrders", testOrderToOneSetOpPhoneUsingPhone)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PhoneToCarts", testPhoneToManyAddOpCarts)
	t.Run("PhoneToOrders", testPhoneToManyAddOpOrders)
	t.Run("UserToCarts", testUserToManyAddOpCarts)
	t.Run("UserToOrders", testUserToManyAddOpOrders)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Carts", testCartsReload)
	t.Run("Orders", testOrdersReload)
	t.Run("Phones", testPhonesReload)
	t.Run("Users", testUsersReload)
	t.Run("Webapps", testWebappsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Carts", testCartsReloadAll)
	t.Run("Orders", testOrdersReloadAll)
	t.Run("Phones", testPhonesReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("Webapps", testWebappsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Carts", testCartsSelect)
	t.Run("Orders", testOrdersSelect)
	t.Run("Phones", testPhonesSelect)
	t.Run("Users", testUsersSelect)
	t.Run("Webapps", testWebappsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Carts", testCartsUpdate)
	t.Run("Orders", testOrdersUpdate)
	t.Run("Phones", testPhonesUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("Webapps", testWebappsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Carts", testCartsSliceUpdateAll)
	t.Run("Orders", testOrdersSliceUpdateAll)
	t.Run("Phones", testPhonesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("Webapps", testWebappsSliceUpdateAll)
}
