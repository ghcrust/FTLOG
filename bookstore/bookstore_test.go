package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"sort"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Book 1",
		Author: "Author 1",
		Copies: 8,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Author 2",
		Author: "Author 2",
		Copies: 8,
	}
	want := 7
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestBuyIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Book",
		Author: "Author",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Fatal("want error for invalid input, got nil")
	}
}

func TestCatalog(t *testing.T) {
	catalog := bookstore.Catalog{1: {Title: "For the love of go"}, 2: {Title: "The power of go: tools"}}
	want := []bookstore.Book{{Title: "For the love of go"}, {Title: "The power of go: tools"}}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	want := bookstore.Book{ID: 1, Title: "For the love of go"}
	catalog := bookstore.Catalog{1: want, 2: {ID: 2, Title: "The power of go: tools"}}
	got, _ := catalog.GetBook(1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIdReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(999)
	if err == nil {
		t.Fatal("Expected error for invalid input, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{Title: "Priced book", PriceCents: 10000, DiscountPercent: 50}
	want := 5000
	got := book.NetPriceCents()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
