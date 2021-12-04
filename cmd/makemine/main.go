package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/natemarks/makemine/model"
	"github.com/natemarks/makemine/version"
	"github.com/rs/zerolog"
)

// EnsureDataDir create the data directory if it doesn't eist
func EnsureDataDir() error {
	err := os.MkdirAll(model.DataDir, 0755)
	return err
}

// makemine processes each argument FromUrl and FromFile, using the first one that works
// if there are no arguments it prints usage and exits with an error code
// there is no logging unless the debug logging is set
func main() {
	var data model.MyData
	var err error

	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	// Exit immediately if we can't ensure the data directory
	err = EnsureDataDir()
	if err != nil {
		logger.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger.Debug().Msgf("Beginning to process data source arguments")
	if len(flag.Args()) == 0 {
		logger.Debug().Msg("no arguments provided. prompting for user data")
		data = model.MyDataFromInput()
	}
	for _, v := range flag.Args() {
		// try to get data from a file
		data, err = model.MyDataFromFilePath(v)
		if err != nil {
			logger.Debug().Msgf("Unable to get MyData from file: %s", v)
		} else {
			// if successful break out of the loop
			break
		}

		data, err = model.MyDataFromURL(v)
		if err != nil {
			logger.Debug().Msgf("Unable to get MyData from url: %s", v)
		} else {
			// if successful break out of the loop
			break
		}
	}
	// if we haven't succeeded in getting the data after processing the args
	if err != nil {
		data = model.MyDataFromInput()
	}

	// At this point we have data from some source
	// Write to json
	err = data.ToJSOM()
	if err != nil {
		logger.Fatal().Err(err)
	}
	// write to yaml
	err = data.ToYaml()
	if err != nil {
		logger.Fatal().Err(err)
	}
	// write to source script
	err = data.ToSourceScript()
	if err != nil {
		logger.Fatal().Err(err)
	}

	// finally, print the useradd command associated with the user data
	addUserCmd := fmt.Sprintf("sudo useradd -m -d /home/%[1]s -s /bin/bash -g sudo %[1]s && passwd %[1]s", data.LocalUser)
	fmt.Println(addUserCmd)
}
