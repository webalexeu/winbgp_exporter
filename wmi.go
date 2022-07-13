package main

import (
	"fmt"
	"log"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	//"github.com/mattn/go-ole/oleutil"
)

func wmi_query() {
	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		log.Fatal(err)
	}

	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		log.Fatal(err)
	}
	defer unknown.Release()

	wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatal(err)
	}
	defer wmi.Release()

	serviceRaw, _ := oleutil.CallMethod(wmi, "ConnectServer", nil, "ROOT/Microsoft/Windows/RemoteAccess")
	//serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer", nil, "ROOT/cimv2")
	if err != nil {
		log.Fatal(err)
	}
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

	//resultRaw, _ := oleutil.CallMethod(service, "Get", "PS_BgpCustomRoute")

	//resultRaw, err := oleutil.CallMethod(service, "ExecQuery", "select state from win32_service where name='WinRM'")
	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", "select * from PS_BgpCustomRoute")
	if err != nil {
		log.Fatalf("wmi execution error: %v", err)
	}
	result := resultRaw.ToIDispatch()
	defer result.Release()

	itemRaw, err := oleutil.CallMethod(result, "ItemIndex", 0)
	if err != nil {
		log.Fatalf("wmi index error: %v", err)
	}
	item := itemRaw.ToIDispatch()
	defer item.Release()

	serviceStatus, err := oleutil.GetProperty(item, "Network")
	if err != nil {
		log.Fatalf("wmi property error: %v", err)
	}
	fmt.Println("WMI TEST %s", serviceStatus.ToString())
}
