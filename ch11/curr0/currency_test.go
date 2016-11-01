package curr0

import "testing"

var data = "../data.csv"

func TestLoad(t *testing.T) {
	table := Load(data)
	if len(table) != 278 {
		t.Log("Table size = ", len(table))
		t.Fatal("Expected row count not loaded")
	}
}

func TestFindByCode(t *testing.T) {
	table := Load(data)
	if len(Find(table, "XAU")) != 1 {
		t.Fatal("Find Failed to find 1 match by Code")
	}
}

func TestFindByNumber(t *testing.T) {
	table := Load(data)
	if len(Find(table, "949")) != 1 {
		t.Fatal("Find Failed to find 1 match by Number")
	}
}

func TestFindMultiple(t *testing.T) {
	table := Load(data)
	result := Find(table, "UNITED")
	if len(result) < 5 {
		t.Log("Found ", len(result))
		t.Fatal("Find failed to find several matches")
	}
}
