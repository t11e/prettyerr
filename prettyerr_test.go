package prettyerr_test

import (
	goerrors "errors"
	"fmt"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/t11e/prettyerr"
)

func errorsNew(message string) error {
	return goerrors.New(message)
}

func pkgErrorsNew(message string) error {
	return errors.New(message)
}

func return_errorsNew(message string) error {
	return errorsNew(message)
}

func return_pkgErrorsNew(message string) error {
	return pkgErrorsNew(message)
}

func withStack_errorsNew(message string) error {
	return errors.WithStack(errorsNew(message))
}

func withStack_pkgErrorsNew(message string) error {
	return errors.WithStack(pkgErrorsNew(message))
}

func withMessage_errorsNew(message, message2 string) error {
	return errors.WithMessage(errorsNew(message), message2)
}

func withMessage_pkgErrorsNew(message, message2 string) error {
	return errors.WithMessage(pkgErrorsNew(message), message2)
}

func wrap_errorsNew(message, message2 string) error {
	return errors.Wrap(errorsNew(message), message2)
}

func wrap_pkgErrorsNew(message, message2 string) error {
	return errors.Wrap(pkgErrorsNew(message), message2)
}

func ExampleFormat() {
	err := errors.New("ran out of polish")
	err = errors.WithMessage(err, "problem polishing widget")
	fmt.Println(prettyerr.Format{
		Err:   err,
		Flags: prettyerr.FlagTesting,
	})
	// Output: problem polishing widget: ran out of polish
	// Caused by: ran out of polish
	//     at github.com/t11e/prettyerr_test.ExampleFormat	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
	//     at testing.runExample	$GOROOT/src/testing/example.go
	//     at testing.runExamples	$GOROOT/src/testing/example.go
	//     at testing.(*M).Run	$GOROOT/src/testing/testing.go
	//     at main.main	github.com/t11e/prettyerr/_test/_testmain.go
}

func TestFormat_String_NoStacks_NoCauses(t *testing.T) {
	for idx, test := range []struct {
		err      error
		expected string
	}{
		{
			err:      errorsNew("errorsNew"),
			expected: "errorsNew\n",
		},
		{
			err:      pkgErrorsNew("pkgErrorsNew"),
			expected: "pkgErrorsNew\n",
		},
		{
			err:      return_errorsNew("return_errorsNew"),
			expected: "return_errorsNew\n",
		},
		{
			err:      return_pkgErrorsNew("return_pkgErrorsNew"),
			expected: "return_pkgErrorsNew\n",
		},
		{
			err:      withStack_errorsNew("withStack_errorsNew"),
			expected: "withStack_errorsNew\n",
		},
		{
			err:      withStack_pkgErrorsNew("withStack_pkgErrorsNew"),
			expected: "withStack_pkgErrorsNew\n",
		},
		{
			err:      withMessage_errorsNew("withMessage_errorsNew", "extra message"),
			expected: "extra message: withMessage_errorsNew\n",
		},
		{
			err:      withMessage_pkgErrorsNew("withMessage_pkgErrorsNew", "extra message"),
			expected: "extra message: withMessage_pkgErrorsNew\n",
		},
		{
			err:      wrap_errorsNew("wrap_errorsNew", "extra message"),
			expected: "extra message: wrap_errorsNew\n",
		},
		{
			err:      wrap_pkgErrorsNew("wrap_pkgErrorsNew", "extra message"),
			expected: "extra message: wrap_pkgErrorsNew\n",
		},
	} {
		name := test.err.Error()
		t.Run(name, func(t *testing.T) {
			expected := test.expected
			actual := prettyerr.Format{
				Err:   test.err,
				Flags: prettyerr.FlagNoStacks | prettyerr.FlagNoCauses | prettyerr.FlagTesting,
			}.String()
			assert.Equal(t, expected, actual, "[%d] %s\nExpected:\n%sActual:\n%s\n", idx, name, expected, actual)
		})
	}
}

func TestFormat_String_Flags_NoStacks(t *testing.T) {
	for idx, test := range []struct {
		err      error
		expected string
	}{
		{
			err: errorsNew("errorsNew"),
			expected: `
errorsNew
`,
		},
		{
			err: pkgErrorsNew("pkgErrorsNew"),
			expected: `
pkgErrorsNew
`,
		},
		{
			err: return_errorsNew("return_errorsNew"),
			expected: `
return_errorsNew
`,
		},
		{
			err: return_pkgErrorsNew("return_pkgErrorsNew"),
			expected: `
return_pkgErrorsNew
`,
		},
		{
			err: withStack_errorsNew("withStack_errorsNew"),
			expected: `
withStack_errorsNew
Caused by: withStack_errorsNew
`,
		},
		{
			err: withStack_pkgErrorsNew("withStack_pkgErrorsNew"),
			expected: `
withStack_pkgErrorsNew
Caused by: withStack_pkgErrorsNew
`,
		},
		{
			err: withMessage_errorsNew("withMessage_errorsNew", "extra message"),
			expected: `
extra message: withMessage_errorsNew
Caused by: withMessage_errorsNew
`,
		},
		{
			err: withMessage_pkgErrorsNew("withMessage_pkgErrorsNew", "extra message"),
			expected: `
extra message: withMessage_pkgErrorsNew
Caused by: withMessage_pkgErrorsNew
`,
		},
		{
			err: wrap_errorsNew("wrap_errorsNew", "extra message"),
			expected: `
extra message: wrap_errorsNew
Caused by: extra message: wrap_errorsNew
Caused by: wrap_errorsNew
`,
		},
		{
			err: wrap_pkgErrorsNew("wrap_pkgErrorsNew", "extra message"),
			expected: `
extra message: wrap_pkgErrorsNew
Caused by: extra message: wrap_pkgErrorsNew
Caused by: wrap_pkgErrorsNew
`,
		},
	} {
		name := test.err.Error()
		t.Run(name, func(t *testing.T) {
			expected := strings.TrimPrefix(test.expected, "\n")
			actual := prettyerr.Format{
				Err:   test.err,
				Flags: prettyerr.FlagNoStacks,
			}.String()
			assert.Equal(t, expected, actual, "[%d] %s\nExpected:\n%sActual:\n%s\n", idx, name, expected, actual)
		})
	}
}

func TestFormat_String_NoCauses(t *testing.T) {
	for idx, test := range []struct {
		err      error
		expected string
	}{
		{
			err: errorsNew("errorsNew"),
			expected: `
errorsNew
`,
		},
		{
			err: pkgErrorsNew("pkgErrorsNew"),
			expected: `
pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_NoCauses	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: return_errorsNew("return_errorsNew"),
			expected: `
return_errorsNew
`,
		},
		{
			err: return_pkgErrorsNew("return_pkgErrorsNew"),
			expected: `
return_pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.return_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_NoCauses	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: withStack_errorsNew("withStack_errorsNew"),
			expected: `
withStack_errorsNew
    at github.com/t11e/prettyerr_test.withStack_errorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_NoCauses	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: withStack_pkgErrorsNew("withStack_pkgErrorsNew"),
			expected: `
withStack_pkgErrorsNew
    at github.com/t11e/prettyerr_test.withStack_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_NoCauses	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: withMessage_errorsNew("withMessage_errorsNew", "extra message"),
			expected: `
extra message: withMessage_errorsNew
`,
		},
		{
			err: withMessage_pkgErrorsNew("withMessage_pkgErrorsNew", "extra message"),
			expected: `
extra message: withMessage_pkgErrorsNew
`,
		},
		{
			err: wrap_errorsNew("wrap_errorsNew", "extra message"),
			expected: `
extra message: wrap_errorsNew
    at github.com/t11e/prettyerr_test.wrap_errorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_NoCauses	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: wrap_pkgErrorsNew("wrap_pkgErrorsNew", "extra message"),
			expected: `
extra message: wrap_pkgErrorsNew
    at github.com/t11e/prettyerr_test.wrap_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_NoCauses	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
	} {
		name := test.err.Error()
		t.Run(name, func(t *testing.T) {
			expected := strings.TrimPrefix(test.expected, "\n")
			actual := prettyerr.Format{
				Err:   test.err,
				Flags: prettyerr.FlagNoCauses | prettyerr.FlagTesting,
			}.String()
			assert.Equal(t, expected, actual, "[%d] %s\nExpected:\n%sActual:\n%s\n", idx, name, expected, actual)
		})
	}
}

func TestFormat_String(t *testing.T) {
	for idx, test := range []struct {
		err      error
		expected string
	}{
		{
			err: errorsNew("errorsNew"),
			expected: `
errorsNew
`,
		},
		{
			err: pkgErrorsNew("pkgErrorsNew"),
			expected: `
pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: return_errorsNew("return_errorsNew"),
			expected: `
return_errorsNew
`,
		},
		{
			err: return_pkgErrorsNew("return_pkgErrorsNew"),
			expected: `
return_pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.return_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: withStack_errorsNew("withStack_errorsNew"),
			expected: `
withStack_errorsNew
    at github.com/t11e/prettyerr_test.withStack_errorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
Caused by: withStack_errorsNew
`,
		},
		{
			err: withStack_pkgErrorsNew("withStack_pkgErrorsNew"),
			expected: `
withStack_pkgErrorsNew
    at github.com/t11e/prettyerr_test.withStack_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
Caused by: withStack_pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.withStack_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: withMessage_errorsNew("withMessage_errorsNew", "extra message"),
			expected: `
extra message: withMessage_errorsNew
Caused by: withMessage_errorsNew
`,
		},
		{
			err: withMessage_pkgErrorsNew("withMessage_pkgErrorsNew", "extra message"),
			expected: `
extra message: withMessage_pkgErrorsNew
Caused by: withMessage_pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.withMessage_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			err: wrap_errorsNew("wrap_errorsNew", "extra message"),
			expected: `
extra message: wrap_errorsNew
    at github.com/t11e/prettyerr_test.wrap_errorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
Caused by: extra message: wrap_errorsNew
Caused by: wrap_errorsNew
`,
		},
		{
			err: wrap_pkgErrorsNew("wrap_pkgErrorsNew", "extra message"),
			expected: `
extra message: wrap_pkgErrorsNew
    at github.com/t11e/prettyerr_test.wrap_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
Caused by: extra message: wrap_pkgErrorsNew
Caused by: wrap_pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.wrap_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
	} {
		name := test.err.Error()
		t.Run(name, func(t *testing.T) {
			expected := strings.TrimPrefix(test.expected, "\n")
			actual := prettyerr.Format{
				Err:   test.err,
				Flags: prettyerr.FlagTesting,
			}.String()
			assert.Equal(t, expected, actual, "[%d] %s\nExpected:\n%sActual:\n%s\n", idx, name, expected, actual)
		})
	}
}

func TestFormat_String_Prefix(t *testing.T) {
	for idx, test := range []struct {
		indent   string
		expected string
	}{
		{
			indent: "",
			expected: `
extra message: wrap_pkgErrorsNew
    at github.com/t11e/prettyerr_test.wrap_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_Prefix.func1	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
Caused by: extra message: wrap_pkgErrorsNew
Caused by: wrap_pkgErrorsNew
    at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.wrap_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
    at github.com/t11e/prettyerr_test.TestFormat_String_Prefix.func1	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
		{
			indent: "abcd ",
			expected: `
abcd extra message: wrap_pkgErrorsNew
abcd     at github.com/t11e/prettyerr_test.wrap_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
abcd     at github.com/t11e/prettyerr_test.TestFormat_String_Prefix.func1	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
abcd Caused by: extra message: wrap_pkgErrorsNew
abcd Caused by: wrap_pkgErrorsNew
abcd     at github.com/t11e/prettyerr_test.pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
abcd     at github.com/t11e/prettyerr_test.wrap_pkgErrorsNew	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
abcd     at github.com/t11e/prettyerr_test.TestFormat_String_Prefix.func1	$GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go
`,
		},
	} {
		name := fmt.Sprintf("%d_%s", len(test.indent), strings.Replace(test.indent, " ", ".", -1))
		t.Run(name, func(t *testing.T) {
			expected := strings.TrimPrefix(test.expected, "\n")
			actual := prettyerr.Format{
				Err:    wrap_pkgErrorsNew("wrap_pkgErrorsNew", "extra message"),
				Flags:  prettyerr.FlagTesting,
				Prefix: test.indent,
			}.String()
			assert.Equal(t, expected, actual, "[%d] %s\nExpected:\n%sActual:\n%s\n", idx, name, expected, actual)
		})
	}
}
