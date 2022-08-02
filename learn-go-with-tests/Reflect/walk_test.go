package Reflect

import (
	"reflect"
	"testing"
)

//func TestWalk(t *testing.T)  {
//	expected := "Chiris"
//	var got []string
//
//	x := struct {
//		Name string
//	}{expected}
//
//	walk(x, func(input string) {
//		got = append(got, input)
//	})
//	if got[0] != expected {
//		t.Errorf("got %s, want %s.", got[0], expected)
//	}
//	if len(got) != 1 {
//		t.Errorf("wrong number of function calls, got %d, want %d", len(got), 1)
//	}
//
//}
type Profile struct {
	Age int
	City string
}
type Person struct {
	Name string
	Profile Profile
}
func TestWalk(t *testing.T)  {
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with  one string field",
			struct {
				Name string
			}{"Chirs"},
			[]string{"Chirs"},
		},
		{
			"Struct with two string field",
			struct {
			Name string
			City string
			}{"Chirs", "London"},
			[]string {"Chirs", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age int
			}{"Chirs", 33},
			[]string{"Chirs"},
		},
		{
			"Nested fields",
			Person {
				"Chris",
				Profile {33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointer to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Beijing"},
			},
			[]string{"London", "Beijing"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls){
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")

	})
}
func assertContains(t *testing.T, haystack []string, needle string)  {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %s, but didn't", haystack, needle)
	}
}