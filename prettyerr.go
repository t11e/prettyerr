// Package prettyerr generates readble errors with optional stacktraces and causes
// (extracted from https://github.com/pkg/errors compatible errors).
package prettyerr

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// Format wraps an error for pretty printing.
type Format struct {
	Err    error  // wrapped error
	Flags  Flag   // format flags
	Prefix string // optional prefix for each line output
}

// Flag controls how the error should be formatted, by default full stacktraces and causes are included.
type Flag int

const (
	FlagNoCauses = 1 << iota // do not follow causes
	FlagNoStacks             // do not print stacks

	FlagNoLineNumbers    // do not include line numbers from stacktraces
	FlagNoGoRoot         // do not look for and substitute $GOROOT in paths
	FlagNoGoPath         // do not look for and substitute $GOPATH in paths
	FlagNoTrailingGoRoot // do not include any trailing stackframes within $GOROOT

	FlagTesting = FlagNoLineNumbers | FlagNoTrailingGoRoot // do not include data that would break test runs
)

func (f Format) String() string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s%s\n", f.Prefix, f.Err)
	generate(&w, f.Err, f.Prefix, f.Flags)
	return w.String()
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
type causer interface {
	Cause() error
}

func generate(w io.Writer, err error, prefix string, flags Flag) {
	if flags&FlagNoStacks == 0 {
		if st, ok := err.(stackTracer); ok {
			goroot := runtime.GOROOT()
			gopath := currentGOPATH()
			stack := st.StackTrace()
			lines := make([]string, 0, len(stack))
			for _, frame := range stack {
				line := fmt.Sprintf("%+v", frame)
				line = strings.Replace(line, "\n\t", "\t", 1)
				lines = append(lines, line)
			}
			if flags&FlagNoTrailingGoRoot != 0 {
				i := len(lines) - 1
				for ; i >= 0; i = i - 1 {
					if !strings.Contains(lines[i], goroot) {
						break
					}
				}
				if i > 0 && i+1 < len(lines) {
					lines = lines[:i+1]
				}
			}
			for _, s := range lines {
				if flags&FlagNoGoRoot == 0 {
					s = strings.Replace(s, goroot, "$GOROOT", 1)
				}
				if flags&FlagNoGoPath == 0 {
					s = strings.Replace(s, gopath, "$GOPATH", 1)
				}
				if flags&FlagNoLineNumbers != 0 {
					s = regexLineNums.ReplaceAllString(s, "")
				}
				fmt.Fprintf(w, "%s    at %s\n", prefix, s)
			}
		}
	}
	if flags&FlagNoCauses == 0 {
		if cr, ok := err.(causer); ok {
			c := cr.Cause()
			fmt.Fprintf(w, "%sCaused by: %s\n", prefix, c)
			generate(w, c, prefix, flags)
		}
	}
}

var regexLineNums = regexp.MustCompile(`:\d+$`)

// currentGOPATH exposes the current GOPATH.
// See http://stackoverflow.com/questions/32649770/how-to-get-current-gopath-from-code
func currentGOPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		return gopath
	}
	return defaultGOPATH()
}

// defaultGOPATH exposes the default GOPATH, used as a fallback when $GOPATH is not set.
// See http://stackoverflow.com/questions/32649770/how-to-get-current-gopath-from-code
func defaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	if home := os.Getenv(env); home != "" {
		def := filepath.Join(home, "go")
		if filepath.Clean(def) == filepath.Clean(runtime.GOROOT()) {
			// Don't set the default GOPATH to GOROOT,
			// as that will trigger warnings from the go tool.
			return ""
		}
		return def
	}
	return ""
}
