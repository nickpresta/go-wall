package gowall_test

import (
	"github.com/NickPresta/go-wall/gowall"
	"reflect"
	"testing"
)

var fetchtest = []struct {
	in       string
	expected gowall.User
}{
	{"gowalltest", gowall.User{0, "Location", "Go Wall Test", "501ed1b92a13290012000003", "gowalltest", map[string]string{"github": "gowalltest"}, []gowall.Badge{}}},
}

func TestFetchUser(t *testing.T) {
	for i, testcase := range fetchtest {
		user, err := gowall.FetchUser(testcase.in)
		if err != nil {
			t.Errorf("%s", i, err)
		}
		if !reflect.DeepEqual(user, testcase.expected) {
			t.Errorf("%+v != %+v", i, testcase.expected, user)
		}
	}
}

func TestFetchUserDoesNotExist(t *testing.T) {
	_, err := gowall.FetchUser("thisuserdoesnotexistandshouldneverexisthopefully")
	if err == nil {
		t.Errorf("FetchUser() should have returned with an error.")
	}
}
