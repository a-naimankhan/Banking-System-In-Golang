package domain_test

import (
	"math"
	"testing"

	domain "BankingSystem/internal/domain"
)

const eps = 1e-6

func floatEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) < eps
}

func TestDeposit(t *testing.T) {
	cases := []struct {
		name        string
		initial     float32
		amount      float32
		wantErr     bool
		wantBalance float32
	}{
		{"positive amount", 0, 100, false, 100},
		{"increase existing", 50, 25, false, 75},
		{"zero amount", 10, 0, true, 10},
		{"negative amount", 10, -5, true, 10},
	}

	for _, tc := range cases {
		tc := tc // capture
		t.Run(tc.name, func(t *testing.T) {
			acc := domain.Account{Balance: tc.initial}
			before := acc.Balance
			err := acc.Deposit(tc.amount)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for case %q, got nil", tc.name)
				}
				if !floatEqual(acc.Balance, before) {
					t.Fatalf("balance changed on error: before=%v after=%v", before, acc.Balance)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !floatEqual(acc.Balance, tc.wantBalance) {
				t.Fatalf("balance mismatch: want=%v got=%v", tc.wantBalance, acc.Balance)
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	cases := []struct {
		name        string
		initial     float32
		amount      float32
		wantErr     bool
		wantBalance float32
	}{
		{"withdraw less than balance", 100, 40, false, 60},
		{"withdraw equal to balance", 75, 75, false, 0},
		{"withdraw more than balance", 30, 50, true, 30},
		{"zero amount", 20, 0, true, 20},
		{"negative amount", 20, -5, true, 20},
	}

	for _, tc := range cases {
		tc := tc // capture
		t.Run(tc.name, func(t *testing.T) {
			acc := domain.Account{Balance: tc.initial}
			before := acc.Balance
			err := acc.Withdraw(tc.amount)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for case %q, got nil", tc.name)
				}
				if !floatEqual(acc.Balance, before) {
					t.Fatalf("balance changed on error: before=%v after=%v", before, acc.Balance)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !floatEqual(acc.Balance, tc.wantBalance) {
				t.Fatalf("balance mismatch: want=%v got=%v", tc.wantBalance, acc.Balance)
			}
		})
	}
}
