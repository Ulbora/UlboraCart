package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddCustomer(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 1

	sdb.MockAddCustomerSuccess = true
	sdb.MockCustomerID = 2
	resp := m.AddCustomer(&cus)
	if !resp.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCustomerFail(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 1

	//sdb.MockAddCustomerSuccess = true
	sdb.MockCustomerID = 2
	resp := m.AddCustomer(&cus)
	if resp.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCustomer(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 1

	sdb.MockUpdateCustomerSuccess = true
	res := m.UpdateCustomer(&cus)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCustomerFail(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 1

	//sdb.MockUpdateCustomerSuccess = true
	res := m.UpdateCustomer(&cus)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomer(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6

	sdb.MockCustomer = &cus
	fcus := m.GetCustomer("test", 2)
	if fcus.ID != cus.ID {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomerID(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6

	sdb.MockCustomer = &cus
	fcus := m.GetCustomerID(4)
	if fcus.ID != cus.ID {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomerList(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6

	var cuslst []sdbi.Customer
	cuslst = append(cuslst, cus)

	sdb.MockCustomerList = &cuslst
	fcuslst := m.GetCustomerList(3)
	if (*fcuslst)[0].ID != cus.ID {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCustomer(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	sdb.MockDeleteCustomerSuccess = true
	res := m.DeleteCustomer(5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCustomerFail(t *testing.T) {

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	//sdb.MockDeleteCustomerSuccess = true
	res := m.DeleteCustomer(5)
	if res.Success {
		t.Fail()
	}
}
