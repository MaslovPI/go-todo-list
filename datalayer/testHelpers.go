package datalayer

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Didn't expect error, but got: %v", err)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got error %q want %q", got, want)
	}
}

func assertTasksEqual(t testing.TB, got, want Task) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want: %#v, but got: %#v", want, got)
	}
}

func assertNumberEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("Want: %d, but got: %d", want, got)
	}
}

func assertStringEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %q, but got %q", want, got)
	}
}

func assertMapVaultsEqual(t testing.TB, want, got MapTaskVault) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v, but got: %v", want, got)
	}
}

func assertParseError(t testing.TB, err error) {
	t.Helper()
	_, ok := err.(*time.ParseError)
	if !ok {
		t.Errorf("Expected ParseError, got: %v", err)
	}
}

func assertNumError(t testing.TB, err error) {
	t.Helper()
	_, ok := err.(*strconv.NumError)
	if !ok {
		t.Errorf("Expected NumError, got: %v", err)
	}
}
