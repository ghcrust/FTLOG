package bank_test

import (
	"testing"
	"bank"
)


func TestWithdrawal(t *testing.T) {
	t.Parallel()
	want := 80
	got, err := bank.Withdrawal(100, 20)
	if err != nil {
		t.Fatalf("want no error for valid input, got %v", err)
	}
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestInvalidWithdrawal(t *testing.T) {
	_, err := bank.Withdrawal(20,100)
	if err == nil {
		t.Fatalf("want error for invalid input, got nil")
	}
}
