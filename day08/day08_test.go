package day08

import "testing"

func TestCompileString(t *testing.T) {
	if s := CompileString("\"\""); s != "" {
		t.Errorf("\"\"\"\" compiles to \"\", not \"%s\"", s)
	}
	if s := CompileString("\"abc\""); s != "abc" {
		t.Errorf("\"\"abc\"\" compiles to \"abc\", not \"%s\"", s)
	}
	if s := CompileString("\"aaa\"aaa\""); s != "" {
		t.Errorf("\"\"aaa\"aaa\"\" compiles to \"aaa\", not \"%s\"", s)
	}
	if s := CompileString("\"\\x27\""); s != "" {
		t.Errorf("\"\"\\x27\"\" compiles to \"'\", not \"%s\"", s)
	}
}
