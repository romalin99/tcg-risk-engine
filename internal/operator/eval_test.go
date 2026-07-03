// Copyright (c) 2023
//
// @author norman
// https://github.com/romalin99/tcg-risk-engine.git
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
package operator

import (
	"testing"

	"github.com/skyhackvip/risk_engine/internal/log"
)

func init() {
	log.InitLogger("console", "")
}

// TestEvaluate covers the bool-returning Evaluate, which mirrors the old
// govaluate contract (non-bool result -> "convert error").
func TestEvaluate(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		params  map[string]interface{}
		want    bool
		wantErr bool
	}{
		{"and true", "foo && bar", map[string]interface{}{"foo": true, "bar": true}, true, false},
		{"and false", "foo && bar", map[string]interface{}{"foo": true, "bar": false}, false, false},
		{"or", "foo || bar", map[string]interface{}{"foo": false, "bar": true}, true, false},
		{"not", "!foo", map[string]interface{}{"foo": false}, true, false},
		{"not paren", "!(foo && bar) || !a1", map[string]interface{}{"foo": true, "bar": false, "a1": true}, true, false},
		{"gt", "score > 60", map[string]interface{}{"score": 80}, true, false},
		{"lt", "score < 60", map[string]interface{}{"score": 80}, false, false},
		{"ge float", "amount >= 100.0", map[string]interface{}{"amount": 100.0}, true, false},
		{"compound range", "age >= 18 && age < 60", map[string]interface{}{"age": 30}, true, false},
		{"string eq", `name == "norman"`, map[string]interface{}{"name": "norman"}, true, false},
		{"in list", `city in ["beijing", "shanghai"]`, map[string]interface{}{"city": "beijing"}, true, false},
		{"not in list", `city in ["beijing", "shanghai"]`, map[string]interface{}{"city": "guangzhou"}, false, false},
		{"builtin max", "max(a, b) > 5", map[string]interface{}{"a": 3, "b": 8}, true, false},
		{"non-bool result", "a + b", map[string]interface{}{"a": 1, "b": 2}, false, true},
		{"syntax error", "a +", map[string]interface{}{"a": 1}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr, tt.params)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Evaluate(%q) error = %v, wantErr = %v", tt.expr, err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("Evaluate(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}

// TestEvaluateExpr covers the raw-result EvaluateExpr, exercising expr builtins
// that govaluate did not provide out of the box.
func TestEvaluateExpr(t *testing.T) {
	tests := []struct {
		name   string
		expr   string
		params map[string]interface{}
		want   interface{}
	}{
		{"max", "max(foo, bar)", map[string]interface{}{"foo": 1, "bar": 2}, 2},
		{"min", "min(foo, bar)", map[string]interface{}{"foo": 1, "bar": 2}, 1},
		{"add", "foo + bar", map[string]interface{}{"foo": 3, "bar": 4}, 7},
		{"len", "len(list)", map[string]interface{}{"list": []interface{}{1, 2, 3}}, 3},
		{"upper", "upper(name)", map[string]interface{}{"name": "norman"}, "NORMAN"},
		{"ternary", "score >= 60 ? 'pass' : 'fail'", map[string]interface{}{"score": 90}, "pass"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvaluateExpr(tt.expr, tt.params)
			if err != nil {
				t.Fatalf("EvaluateExpr(%q) unexpected error: %v", tt.expr, err)
			}
			if got != tt.want {
				t.Errorf("EvaluateExpr(%q) = %v (%T), want %v (%T)", tt.expr, got, got, tt.want, tt.want)
			}
		})
	}
}

func TestEvaluateExprError(t *testing.T) {
	if _, err := EvaluateExpr("foo +", map[string]interface{}{"foo": 1}); err == nil {
		t.Errorf("EvaluateExpr with syntax error should return error")
	}
}
