MyRpc interface { 
    CreateExpenseReport() (status Status)
    RemoveExpenseReport() (status Status)
    PrintExpenseReport() (list []string, status Status)
    SubmitExpenseReport() (status Status)
    AddItemExpenseReport(name, date, description string, amount int64) (itemNumber int32, status Status)
    RemoveItemExpensetReport(itemNumber int32) (status Status)
}

