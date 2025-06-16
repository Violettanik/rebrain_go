package convertor

import (
	"testing"
	"time"
)

type SimpleStruct struct {
	Name string `keyname:"name"`
	Age  int    `keyname:"age"`
}

type NestedStruct struct {
	ID      int         `keyname:"id"`
	Simple  SimpleStruct `keyname:"simple"`
	Created time.Time   `keyname:"created"`
}

type PointerStruct struct {
	ID      int           `keyname:"id"`
	Simple  *SimpleStruct `keyname:"simple"`
	Created time.Time     `keyname:"created"`
}

type UntaggedStruct struct {
	Name string
	Age  int
}

func TestStructToMap_Simple(t *testing.T) {
	s := SimpleStruct{
		Name: "John",
		Age:  30,
	}

	m := StructToMap(s)

	if m["name"] != "John" {
		t.Errorf("Expected name to be John, got %v", m["name"])
	}

	if m["age"] != 30 {
		t.Errorf("Expected age to be 30, got %v", m["age"])
	}
}

func TestMapToStruct_Simple(t *testing.T) {
	m := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	var s SimpleStruct
	err := MapToStruct(m, &s)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if s.Name != "John" {
		t.Errorf("Expected Name to be John, got %v", s.Name)
	}

	if s.Age != 30 {
		t.Errorf("Expected Age to be 30, got %v", s.Age)
	}
}

func TestStructToMap_Nested(t *testing.T) {
	n := NestedStruct{
		ID: 1,
		Simple: SimpleStruct{
			Name: "John",
			Age:  30,
		},
		Created: time.Now(),
	}

	m := StructToMap(n)

	if m["id"] != 1 {
		t.Errorf("Expected id to be 1, got %v", m["id"])
	}

	simple, ok := m["simple"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected simple to be a map")
	}

	if simple["name"] != "John" {
		t.Errorf("Expected nested name to be John, got %v", simple["name"])
	}
}

func TestMapToStruct_Nested(t *testing.T) {
	m := map[string]interface{}{
		"id": 1,
		"simple": map[string]interface{}{
			"name": "John",
			"age":  30,
		},
		"created": time.Now(),
	}

	var n NestedStruct
	err := MapToStruct(m, &n)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if n.ID != 1 {
		t.Errorf("Expected ID to be 1, got %v", n.ID)
	}

	if n.Simple.Name != "John" {
		t.Errorf("Expected nested Name to be John, got %v", n.Simple.Name)
	}
}

func TestStructToMap_Pointer(t *testing.T) {
	p := PointerStruct{
		ID: 1,
		Simple: &SimpleStruct{
			Name: "John",
			Age:  30,
		},
		Created: time.Now(),
	}

	m := StructToMap(p)

	simple, ok := m["simple"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected simple to be a map")
	}

	if simple["name"] != "John" {
		t.Errorf("Expected nested name to be John, got %v", simple["name"])
	}
}

func TestMapToStruct_Pointer(t *testing.T) {
	m := map[string]interface{}{
		"id": 1,
		"simple": map[string]interface{}{
			"name": "John",
			"age":  30,
		},
		"created": time.Now(),
	}

	var p PointerStruct
	err := MapToStruct(m, &p)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if p.Simple == nil {
		t.Fatalf("Expected Simple to be not nil")
	}

	if p.Simple.Name != "John" {
		t.Errorf("Expected nested Name to be John, got %v", p.Simple.Name)
	}
}

func TestStructToMap_Untagged(t *testing.T) {
	s := UntaggedStruct{
		Name: "John",
		Age:  30,
	}

	m := StructToMap(s)

	if m["Name"] != "John" {
		t.Errorf("Expected Name to be John, got %v", m["Name"])
	}

	if m["Age"] != 30 {
		t.Errorf("Expected Age to be 30, got %v", m["Age"])
	}
}

func TestMapToStruct_Untagged(t *testing.T) {
	m := map[string]interface{}{
		"Name": "John",
		"Age":  30,
	}

	var s UntaggedStruct
	err := MapToStruct(m, &s)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if s.Name != "John" {
		t.Errorf("Expected Name to be John, got %v", s.Name)
	}

	if s.Age != 30 {
		t.Errorf("Expected Age to be 30, got %v", s.Age)
	}
}

func TestRoundTrip(t *testing.T) {
	original := NestedStruct{
		ID: 1,
		Simple: SimpleStruct{
			Name: "John",
			Age:  30,
		},
		Created: time.Now().Truncate(time.Second), // truncate to avoid nanos issues
	}

	m := StructToMap(original)

	var restored NestedStruct
	err := MapToStruct(m, &restored)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if original.ID != restored.ID {
		t.Errorf("ID mismatch: original %v, restored %v", original.ID, restored.ID)
	}

	if original.Simple.Name != restored.Simple.Name {
		t.Errorf("Name mismatch: original %v, restored %v", original.Simple.Name, restored.Simple.Name)
	}

	if !original.Created.Equal(restored.Created) {
		t.Errorf("Created mismatch: original %v, restored %v", original.Created, restored.Created)
	}
}
