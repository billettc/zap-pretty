package zappretty

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/logrusorgru/aurora"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var errNonZapLine = errors.New("non-zap line")

var debug = log.New(ioutil.Discard, "", 0)
var debugEnabled = false
var severityToColor map[string]Color

func init() {
	if os.Getenv("ZAP_PRETTY_DEBUG") != "" {
		debug = log.New(os.Stderr, "[pretty-debug] ", 0)
		debugEnabled = true
	}

	severityToColor = make(map[string]Color)
	severityToColor["debug"] = BlueFg
	severityToColor["info"] = GreenFg
	severityToColor["warning"] = BrownFg
	severityToColor["error"] = RedFg
	severityToColor["dpanic"] = RedFg
	severityToColor["panic"] = RedFg
	severityToColor["fatal"] = RedFg
}

func PrintVersion() {
	fmt.Printf("zap-pretty %s (commit: %s, date: %v)\n", version, commit, date)
}
