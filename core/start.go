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
package core

import (
	"github.com/romalin99/tcg-risk-engine/internal/log"
)

type StartNode struct {
	Info NodeInfo
}

func NewStartNode(name string) *StartNode {
	return &StartNode{
		Info: NodeInfo{Name: name, Kind: TypeStart.String()},
	}
}

func (node StartNode) GetName() string {
	return node.Info.Name
}

func (node StartNode) GetType() NodeType {
	return GetNodeType(node.Info.Kind)
}

func (node StartNode) GetInfo() NodeInfo {
	return node.Info
}

func (node StartNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (node StartNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (node StartNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	log.Info("======[trace] Start======")
	info := node.GetInfo()
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: node.GetType(), Tag: info.Tag, Label: info.Label}
	return nodeResult, nil
}
