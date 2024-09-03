package eoldate

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/formatter"
	"github.com/projectdiscovery/gologger/levels"
	"os"
	"runtime"
	"time"
)

// LogError ...
func LogError(err error) error {
	timestamp := time.Now().Format("01-02-2006")
	fname := fmt.Sprintf("eoldate-error-log-%s.json", timestamp)

	f, openFileErr := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if openFileErr != nil {
		return openFileErr
	}
	defer f.Close()
	teeFormatter := formatter.NewTee(formatter.NewCLI(false), f)
	gologger.DefaultLogger.SetFormatter(teeFormatter)
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		gologger.Warning().Msg("Failed to retrieve Caller information")
	}
	fn := runtime.FuncForPC(pc).Name()
	gologger.DefaultLogger.SetMaxLevel(levels.LevelError)
	gologger.Error().Msgf("Error in function %s, called from %s:%d:\n %v", fn, file, line, err)
	return err
}
