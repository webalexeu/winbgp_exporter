package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	//"github.com/mattn/go-ole/oleutil"
)

func wmi_query() {

	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, _ := oleutil.CreateObject("WbemScripting.SWbemLocator")
	defer unknown.Release()

	wmi, _ := unknown.QueryInterface(ole.IID_IDispatch)
	defer wmi.Release()

	serviceRaw, _ := oleutil.CallMethod(wmi, "ConnectServer")
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	//resultRaw, _ := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_Process")
	//result := resultRaw.ToIDispatch()
	//defer result.Release()
	//
	//countVar, _ := oleutil.GetProperty(result, "Count")
	//count := int(countVar.Val)
	//
	//for i := 0; i < count; i++ {
	//	itemRaw, _ := oleutil.CallMethod(result, "ItemIndex", i)
	//	item := itemRaw.ToIDispatch()
	//	defer item.Release()
	//
	//	processName, _ := oleutil.GetProperty(item, "Name")
	//	fmt.Println(processName.ToString())
	//}

	resultRaw, _ := oleutil.CallMethod(service, "ExecQuery", "select State from win32_service where name='WinRM'")
	result := resultRaw.ToIDispatch()
	defer result.Release()

	itemRaw, _ := oleutil.CallMethod(result, "ItemIndex", 0)
	item := itemRaw.ToIDispatch()
	defer item.Release()

	serviceStatus, _ := oleutil.GetProperty(item, "State")
	fmt.Println(serviceStatus.ToString())

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
