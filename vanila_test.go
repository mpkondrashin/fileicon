package fileicon

import (
	"os"
	"testing"
)

func TestVanila(t *testing.T) {
	{
		ext := "a"
		data := VanilaIcon(ext)
		err := os.WriteFile(ext+".svg", []byte(data), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		ext := "ab"
		data := VanilaIcon(ext)
		err := os.WriteFile(ext+".svg", []byte(data), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		ext := "abc"
		data := VanilaIcon(ext)
		err := os.WriteFile(ext+".svg", []byte(data), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		ext := "abcd"
		data := VanilaIcon(ext)
		err := os.WriteFile(ext+".svg", []byte(data), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		ext := "abcde"
		data := VanilaIcon(ext)
		err := os.WriteFile(ext+".svg", []byte(data), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		ext := "abcdef"
		data := VanilaIcon(ext)
		err := os.WriteFile(ext+".svg", []byte(data), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		ext := "abcdefg"
		data := VanilaIcon(ext)
		err := os.WriteFile(ext+".svg", []byte(data), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
}
