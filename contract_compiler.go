package main

import (
	"strings"
)

type Compiler struct {
	Source string
	ByteCode []OpCode
}

func NewCompiler(source string) *Compiler {
	return &Compiler{
		Source: source,
	}
}

func (c *Compiler) Parse() {
	lines := strings.Split(c.Source, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch trimmed {
		case "add":
			c.ByteCode = append(c.ByteCode, OpAdd)
		case "sub":
			c.ByteCode = append(c.ByteCode, OpSub)
		case "mul":
			c.ByteCode = append(c.ByteCode, OpMul)
		case "store":
			c.ByteCode = append(c.ByteCode, OpStore)
		case "load":
			c.ByteCode = append(c.ByteCode, OpLoad)
		}
	}
}

func (c *Compiler) Compile() []OpCode {
	c.Parse()
	return c.ByteCode
}
