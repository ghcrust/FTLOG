package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNewCreditCard(t *testing.T) {
	t.Parallel()
	want := "0123456789"
	cc, err := creditcard.NewCreditCard(want)
	if err != nil {
		t.Fatalf("got error %v for valid input", err)
	}
	got := cc.Number()
	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestNewInvalidCreditCard(t *testing.T) {
	t.Parallel()
	_, err := creditcard.NewCreditCard("")
	if err == nil {
		t.Fatalf("want error for invalid input, got nil")
	}
}
