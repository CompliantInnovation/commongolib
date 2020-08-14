package commongolib_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"github.com/CompliantInnovation/commongolib"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

type TestCase struct {
	unencoded int
	encoded   int
}

func TestAE2_Encode(t *testing.T) {

	cases := []TestCase{
		{127404, 12740469},
		{359109, 35910969},
		{641994, 64199466},
		{249709, 24970964},
		{809350, 80935067},
		{128225, 12822562},
		{136141, 13614167},
		{42848, 4284858},
		{564947, 56494768},
		{188567, 18856768},
		{213130, 21313061},
		{716528, 71652862},
		{328912, 32891267},
		{842389, 84238967},
		{388055, 38805562},
		{902718, 90271869},
		{801521, 80152168},
		{955560, 95556063},
		{30909, 3090953},
		{492360, 49236066},
	}

	for _, testCase := range cases {
		equals(t, testCase.encoded, commongolib.AE2Encode(testCase.unencoded))
	}
}

func TestAE2_Decode(t *testing.T) {

	cases := []TestCase{
		{127404, 12740469},
		{359109, 35910969},
		{641994, 64199466},
		{249709, 24970964},
		{809350, 80935067},
		{128225, 12822562},
		{136141, 13614167},
		{42848, 4284858},
		{564947, 56494768},
		{188567, 18856768},
		{213130, 21313061},
		{716528, 71652862},
		{328912, 32891267},
		{842389, 84238967},
		{388055, 38805562},
		{902718, 90271869},
		{801521, 80152168},
		{955560, 95556063},
		{30909, 3090953},
		{492360, 49236066},
	}

	for _, testCase := range cases {
		result, ok := commongolib.AE2Decode(testCase.encoded)
		if !ok {
			t.Fatalf("AE2Decode() failed: %d => %d", testCase.encoded, testCase.unencoded)
		}
		equals(t, testCase.unencoded, result)
	}

}
