package day08

import (
	"fmt"
	"testing"
)

func TestCompileString(t *testing.T) {
	if s, _ := CompileString("\"\""); s != "" {
		t.Errorf("\"\"\"\" compiles to \"\", not \"%s\"", s)
	}
	if s, _ := CompileString("\"abc\""); s != "abc" {
		t.Errorf("\"\"abc\"\" compiles to \"abc\", not \"%s\"", s)
	}
	if s, _ := CompileString("\"aaa\\\"aaa\""); s != "aaa\"aaa" {
		t.Errorf("\"\"aaa\"aaa\"\" compiles to \"aaa\"aaa\", not \"%s\"", s)
	}
	if s, _ := CompileString("\"\\x27\""); s != "'" {
		t.Errorf("\"\"\\x27\"\" compiles to \"'\", not \"%s\"", s)
	}
}

func TestDeCompileString(t *testing.T) {
	s := DecompileString("\"\"")
	fmt.Println(s)
	if s != "\\\"\\\"" {
		t.Errorf("\"\"\"\" compiles to \\\"\\\", not \"%s\"", s)
	}

	s = DecompileString("\"abc\"")
	fmt.Println(s)
	if s != "\\\"abc\\\"" {
		t.Errorf("\"\"abc\"\" compiles to \\\"abc\\\", not \"%s\"", s)
	}

	s = DecompileString("\"aaa\\\"aaa\"")
	fmt.Println(s)
	if s != "\\\"aaa\\\\\\\"aaa\\\"" {
		t.Errorf("\"\"aaa\"aaa\"\" compiles to \\\"aaa\\\\\\\"aaa\\\", not \"%s\"", s)
	}

	s = DecompileString("\"\\x27\"")
	fmt.Println(s)
	if s != "\\\"\\\\x27\\\"" {
		t.Errorf("\"\"\\x27\"\" compiles to \\\"\\\\x27\\\", not \"%s\"", s)
	}
}
