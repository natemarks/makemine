package main

import (
	"flag"
	"os"

	"github.com/natemarks/makemine/output"

	"github.com/natemarks/makemine/input"
	"github.com/natemarks/makemine/version"
	"github.com/rs/zerolog"
)

// Try to get data as a local file or url
func tryString(istr string, log *zerolog.Logger) (input.MakeMineInput, error) {
	var data = input.MakeMineInput{}
	var err error

	log.Debug().Msgf("Trying to us input argument: %s", istr)

	// If we get good data from local file, return
	data, err = input.FromFile(istr)
	if err == nil {
		return data, err
	}
	// If we also fail to get it as a url, return empty data and the error
	data, err = input.FromUrl(istr)

	return data, err
}

// makemine processes each argument FromUrl and FromFile, using the first one that works
// if there are no arguments it prints usage and exits with an error code
// there is no logging unless the debug logging is set
func main() {
	var data input.MakeMineInput
	var err error

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger.Debug().Msgf("version: %s", version.Version)

	if len(flag.Args()) == 0 {
		data = input.FromUser()
	}
	for _, v := range flag.Args() {
		data, err = tryString(v, &logger)
		if err == nil {
			break
		}
	}
	if err != nil {
		data = input.FromUser()
	}
	err = output.WriteData(data)
	if err != nil {
		logger.Fatal().Err(err)
		os.Exit(1)
	}
	os.Exit(0)
}
