package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Renato"},
			ExpectedCalls: []string{"Renato"},
		},
		{
			Name: "struct with two strings fields",
			Input: struct {
				Name string
				City string
			}{"Renato", "Caieiras"},
			ExpectedCalls: []string{"Renato", "Caieiras"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Renato", 27},
			ExpectedCalls: []string{"Renato"},
		},
		{
			Name:          "struct with nested fields",
			Input:         Person{"Renato", Profile{28, "Caieiras"}},
			ExpectedCalls: []string{"Renato", "Caieiras"},
		},
		{
			Name:          "pointers to things",
			Input:         &Person{"Renato", Profile{28, "Caieiras"}},
			ExpectedCalls: []string{"Renato", "Caieiras"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{33, "London"},
				{34, "Paris"},
			},
			ExpectedCalls: []string{"London", "Paris"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Paris"},
			},
			ExpectedCalls: []string{"London", "Paris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		a := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(a, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		a := make(chan Profile)

		go func() {
			a <- Profile{33, "London"}
			a <- Profile{34, "Berlin"}
			close(a)
		}()

		var got []string
		want := []string{"London", "Berlin"}

		walk(a, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		a := func() (Profile, Profile) {
			return Profile{27, "Caieiras"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Caieiras", "Katowice"}

		walk(a, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
