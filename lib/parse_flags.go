package lib

import (
	"flag"
	"os"
	"strconv"

	"github.com/golang/glog"
)

// FlagTokenKey is the parameter name for token
const FlagTokenKey = "token"

// FlagActiveKey is the parameter name for active
const FlagActiveKey = "active"

// FlagVerboseLogLevelKey is the parameter name for enabling verbose logging
const FlagVerboseLogLevelKey = "v"

var flagToken string
var flagActive string

func init() {
	flag.StringVar(&flagToken, FlagTokenKey, "", "token for SCIM API authentication of Slack. (REQUIRED)")
	flag.StringVar(&flagActive, FlagActiveKey, "", "active that controls whether to activate/deactivate users via SCIM API . (by default FALSE)")
}

// ParseFlags parses command line arguments/environment variables
func ParseFlags() (token string, active bool) {
	// Parse command line arguments
	flag.Parse()

	token = flagToken
	if flagActive == "true" {
		active = true
	} else {
		active = false
	}

	// Replace with environment variable if command line argument missing
	if token == "" {
		token = os.Getenv(FlagTokenKey)
	}

	// Validate required parameters
	if token == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// glog - Enforce all logs being sent to STDERR
	f := flag.Lookup("logtostderr")
	if f == nil {
		glog.Fatal("Lookup of flag `logtostderr` failed")
	}
	err := f.Value.Set("true")
	if err != nil {
		glog.Fatal(err)
	}

	// glog - Verbose logging
	vLevel := getVLevel()

	glog.Infoln("---------- Flags: ----------")
	glog.Infof("- token[%s], ", "********")
	if flagActive == "true" {
		glog.Infof("- active[true], ")
	} else {
		glog.Infof("- active[false], ")
	}
	glog.Infof("- verbose logging level[%d], ", vLevel)

	return
}

func getVLevel() (vLevel int) {
	f := flag.Lookup(FlagVerboseLogLevelKey).Value.String()
	vLevel, err := strconv.Atoi(f)
	if err != nil {
		glog.Fatal(err)
	}

	// If no command line flags are provided then the default value is 0.
	if vLevel == 0 {
		// Override flag value with environment value
		envValue := os.Getenv(FlagVerboseLogLevelKey)
		if envValue != "" {
			vLevel, err := strconv.Atoi(envValue)
			if err != nil {
				glog.Fatal(err)
			}
			// Write this into the flags so that glog sees it
			f := flag.Lookup("v")
			if f == nil {
				glog.Fatal("Lookup of flag `v` failed")
			}
			err = f.Value.Set(strconv.Itoa(vLevel))
			if err != nil {
				glog.Fatal(err)
			}
		}
	}

	return
}
