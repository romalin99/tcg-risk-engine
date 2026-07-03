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
package util

import (
	"testing"
)

func TestNumberic(t *testing.T) {
	t.Log(GetType(20))
	t.Log(GetType(30))
	t.Log(GetType("1111"))
	t.Log(GetType("11.11"))
	t.Log(GetType("aa11"))
	t.Log(GetType(true))
	t.Log(GetType("true"))
	t.Log(GetType(1111))
	t.Log(GetType(1111.1111))
}

func TestFloat(t *testing.T) {
	t.Log(ToFloat64("-1.588"))
	t.Log(ToFloat64("1.588"))
	t.Log(ToFloat64(1.588))
	t.Log(ToFloat64(1))
	t.Log(ToFloat64(-1))
	t.Log(ToFloat64(-1.0))
	t.Log(ToFloat64(true))
	t.Log(ToFloat64("1.35e"))
	t.Log(ToFloat64("1.35e5"))
}
