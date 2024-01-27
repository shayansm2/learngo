package main

type BaseAccount struct {
	yearlyInterest int
	balance        int
}

func (account *BaseAccount) MonthlyInterest() int {
	return account.balance * account.yearlyInterest / 100 / 12
}

func (account *BaseAccount) Transfer(receiver Account, amount int) string {
	switch receiver.(type) {
	case *SavingsAccount:
	case *InvestmentAccount:
	case *CheckingAccount:
	default:
		return "Invalid receiver account"
	}

	if response := account.Withdraw(amount); response != "Success" {
		return response
	}

	if response := receiver.Deposit(amount); response != "Success" {
		return response
	}

	return "Success"
}

func (account *BaseAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	account.balance += amount
	return "Success"
}

func (account *BaseAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	if amount > account.balance {
		return "Account balance is not enough"
	}

	account.balance -= amount
	return "Success"
}

func (account *BaseAccount) CheckBalance() int {
	return account.balance
}

type SavingsAccount struct {
	BaseAccount
}

type CheckingAccount struct {
	BaseAccount
}

type InvestmentAccount struct {
	BaseAccount
}

func NewSavingsAccount() *SavingsAccount {
	result := &SavingsAccount{}
	result.yearlyInterest = 5
	return result
}

func NewCheckingAccount() *CheckingAccount {
	return &CheckingAccount{BaseAccount{yearlyInterest: 1}}
}

func NewInvestmentAccount() *InvestmentAccount {
	return &InvestmentAccount{BaseAccount{yearlyInterest: 2}}
}
