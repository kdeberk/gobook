package bank

type withdrawRequest struct {
	amount int
	okC    chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdrawRequest)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	okC := make(chan bool)
	withdraws <- withdrawRequest{amount: amount, okC: okC}
	return <-okC
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			if balance <= withdraw.amount {
				balance -= withdraw.amount
				withdraw.okC <- true
			} else {
				withdraw.okC <- false
			}
			close(withdraw.okC)
		}
	}
}

func init() {
	go teller()
}
