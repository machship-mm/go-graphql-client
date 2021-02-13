package graphql_test

import (
	"testing"
	"time"

	"go-graphql-client"
)

func TestNewScalars(t *testing.T) {

	if got := graphql.NewBool(false); got == nil {
		t.Error("NewBoolean returned nil")
	}

	if got := graphql.NewFloat64(0.0); got == nil {
		t.Error("NewFloat returned nil")
	}

	if got := graphql.NewInt64(0); got == nil {
		t.Error("NewInt returned nil")
	}

	if got := graphql.NewString(""); got == nil {
		t.Error("NewString returned nil")
	}

	if got := graphql.NewTime(time.Now()); got == nil {
		t.Error("NewTime returned nil")
	}
}
