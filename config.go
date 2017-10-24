// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	`encoding/json`
	`fmt`
	`path/filepath`
	`os`
)

const (
	FileAppend = os.O_APPEND|os.O_CREATE|os.O_WRONLY
	FileMode = 0640
	DirMode = 0750
)

var (
	progName = filepath.Base(os.Args[0])
	progPath = filepath.Dir(os.Args[0])
	version = `undefined`
)

// Config holds the application configuration settings. The struct tags
// must match the field names in the JSON configuration file.
type Config struct {

	Paths struct {
		ReportDir string
	}

	API struct {
		Server string
		Endpoint map[string]string
	}

	Include struct {
		VendorID map[string]bool
		ProductID map[string]map[string]bool
		Default bool
	}

	Syslog *Syslog
	Loggers *Loggers

	DebugLevel int
}

// newConfig retrieves the settings in the JSON configuration file and
// populates the fields in the runtime configuration. It also creates
// directories if they do not already exist.
func newConfig(cf string) (*Config, error) {

	this := &Config{}

	if !filepath.IsAbs(cf) {
		cf = filepath.Join(progPath, cf)
	}

	if err := loadConfig(this, cf); err != nil {
		return nil, err
	}

	if err := this.Syslog.Init(); err != nil {
		return nil, err
	}

	if err := this.Loggers.Init(this.Syslog); err != nil {
		return nil, err
	}

	if dn, err := makePath(this.Paths.ReportDir); err != nil {
		return nil, err
	} else {
		this.Paths.ReportDir = dn
	}

	return this, nil
}

// loadConfig loads a JSON configuration file into an object.
func loadConfig(t interface{}, cf string) error {

	if fh, err := os.Open(cf); err != nil {
		return err
	} else {
		defer fh.Close()
		jd := json.NewDecoder(fh)
		err = jd.Decode(&t)
		return err
	}
}

// makePath creates a directory and all intermediate path components.
// It prepends the program path if the given path is relative and 
// returns the resulting absolute path.
func makePath(path string) (string, error) {

	path = filepath.Clean(path)

	if !filepath.IsAbs(path) {
		path = filepath.Join(progPath, path)
	}

	return path, os.MkdirAll(path, DirMode)
}

// displayVersion displays the progName version.
func displayVersion() {
        fmt.Fprintf(os.Stderr, "%s version %s\n", progName, version)
}
