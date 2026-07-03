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
	"errors"

	"github.com/expr-lang/expr"
	"github.com/romalin99/tcg-risk-engine/internal/log"
)

// EvaluateExpr compiles exprStr against params and returns the raw result.
// It is backed by expr-lang/expr (https://github.com/expr-lang/expr) and
// replaces the previous govaluate based implementation. Any expression that
// expr supports is accepted, including builtins such as max/min/len and the
// in / contains operators, e.g. `score > 60 && city in ["bj", "sh"]`.
func EvaluateExpr(exprStr string, params map[string]interface{}) (interface{}, error) {
	program, err := expr.Compile(exprStr, expr.Env(params))
	if err != nil {
		log.Infof("compile expr error: %s, %v", exprStr, err)
		return nil, err
	}
	output, err := expr.Run(program, params)
	if err != nil {
		log.Infof("run expr error: %s, %v", exprStr, err)
		return nil, err
	}
	return output, nil
}

// Evaluate runs a boolean expression and returns its bool result.
// It preserves the original govaluate-based contract: expressions whose
// result is not a bool return a "convert error".
func Evaluate(exprStr string, params map[string]interface{}) (bool, error) {
	output, err := EvaluateExpr(exprStr, params)
	if err != nil {
		return false, err
	}
	if result, ok := output.(bool); ok {
		return result, nil
	}
	return false, errors.New("convert error")
}
