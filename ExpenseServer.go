package main

import (
	"ethos/altEthos"
	"ethos/myRpc"
	"ethos/syscall"
	"log"
)

type ExpenseItem struct {
    Name        string
    Date        string
    Description string
    Amount      int64
}

type Status struct {
	Code    int
	Message string
}

type ExpenseServer struct {
	fileSystem map[string]*ExpenseAccount
}


type ExpenseAccount = []ExpenseReport
type ExpenseReport []ExpenseItem
var expenseReport ExpenseReport
func init() {
	myRpc.SetUpMyRpcCreateExpenseReport(createExpense)
	myRpc.SetUpMyRpcDeleteExpenseReport(deleteExpense)
	myRpc.SetUpMyRpcSubmitExpenseReport(submitExpense)
	myRpc.SetUpMyRpcAddExpenseItem(addExpenseItem)
	myRpc.SetUpMyRpcRemoveExpenseItem(removeExpenseItem)
	myRpc.SetUpMyRpcPrintExpenseReport(printExpenseReport)
}

func createExpense () (status Status){
	expenseReport = make([]ExpenseItem, 0)
	log.Println(("Created Expense Report"))
	return Status{0, "Success"}
}

func deleteExpense () (status Status){
	log.Println(("Deleted Expense Report"))
	expenseReport = make([]ExpenseItem, 0)
	return Status{0, "Success"}
}

func submitExpense () (status Status){
	log.Println(("Submit Expense Report"))
	return Status{0, "Success"}
}

func addExpenseItem (name, date, description string, amount int64) (itemNumber int32, status Status){
	expenseReport = append(expenseReport, ExpenseItem{name, date, description, amount})
	log.Println(("Added Expense Item"))
	return int32(len(expenseReport)), Status{0, "Success"}
}

func removeExpenseItem(itemNumber int32) Status {
	if int(itemNumber) < 0 || int(itemNumber) >= len(expenseReport) {
		return Status{Code: 400, Message: "Invalid item number"}
	}

	expenseReport = append(expenseReport[:itemNumber], expenseReport[itemNumber+1:]...)
	log.Println("Removed Expense Item")
	return Status{Code: 0, Message: "Success"}
}


func printExpenseReport () (list []string, status Status) {
	log.Println(("Printed Expense Report"))
	list = make([]string, len(expenseReport))
	for i, item := range expenseReport {
		list[i] = item.Name + " " + item.Date + " " + item.Description + " " + string(item.Amount)
	}
	return list, Status{0, "Success"}
}	


func main() {
	altEthos.LogToDirectory("test/ExpenseServer")

	log.Println("before call")
	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}

}