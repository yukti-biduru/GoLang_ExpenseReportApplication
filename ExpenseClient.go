package main

import (
	"ethos/altEthos"
	"ethos/myRpc"
	"ethos/syscall"
	"log"
)


func init () {
	// nothing needed here 
}

func main () {

	altEthos.LogToDirectory("test/ExpenseClient")	
	log.Println("before call")

	fd,status :=altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}	

	// create calls to the server functions
	call1 := myRpc.MyRpcCreateExpenseReport{}
	status = altEthos.ClientCall(fd, &call1)
	if status != syscall.StatusOk {
		log.Println("Client Failed")
		altEthos.Exit(status)
	}

	call3 := myRpc.MyRpcAddExpenseItem("Yukti", "01-01-2024", "Description", 1000) 
	status = altEthos.ClientCall(fd, &call3)
	if status != syscall.StatusOk {
		log.Println("Client Failed")
		altEthos.Exit(status)
	}
	
	call4 := myRpc.MyRpcRemoveExpenseItem(0) 
	status = altEthos.ClientCall(fd, &call4)
	if status != syscall.StatusOk {
		log.Println("Client Failed")
		altEthos.Exit(status)
	}

	call5 := myRpc.MyRpcPrintExpenseReport{}
	status = altEthos.ClientCall(fd, &call5)
	if status != syscall.StatusOk {
		log.Println("Client Failed")
		altEthos.Exit(status)
	}	

	call6 := myRpc.MyRpcSubmitExpenseReport{}
	status = altEthos.ClientCall(fd, &call6)
	if status != syscall.StatusOk {
		log.Println("Client Failed")
		altEthos.Exit(status)
	}

	call2 := myRpc.MyRpcDeleteExpenseReport{}
	status = altEthos.ClientCall(fd, &call2)
	if status != syscall.StatusOk {
		log.Println("Client Failed")
		altEthos.Exit(status)
	}

	log.Println("sysClient: Done")


}