package day08

import "testing"

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
	if s := DecompileString("\"\""); s != "\\\"\\\"" {
		t.Errorf("\"\"\"\" compiles to \\\"\\\", not \"%s\"", s)
	}
	if s := DecompileString("\"abc\""); s != "\\\"abc\\\"" {
		t.Errorf("\"\"abc\"\" compiles to \\\"abc\\\", not \"%s\"", s)
	}
	if s := DecompileString("\"aaa\\\"aaa\""); s != "\\\"aaa\\\\\\\"aaa\\\"" {
		t.Errorf("\"\"aaa\"aaa\"\" compiles to \\\"aaa\\\\\\\"aaa\\\", not \"%s\"", s)
	}
	if s := DecompileString("\"\\x27\""); s != "\\\"\\\\x27\\\"" {
		t.Errorf("\"\"\\x27\"\" compiles to \\\"\\\\x27\\\", not \"%s\"", s)
	}
}
