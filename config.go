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
	`net/http`
	`net/http/cookiejar`
	`path/filepath`
	`os`
	`time`
	`golang.org/x/net/publicsuffix`
)

const (
	FileAppend = os.O_APPEND|os.O_CREATE|os.O_WRONLY
	FileMode = 0640
	DirMode = 0750
)

var (
	// Program name and version.

	program string = filepath.Base(os.Args[0])
	version string = `undefined`

	// HTTP Transport and Client.

	httpTransport *http.Transport
	httpClient *http.Client

	// Configuration aliases.

	sl, cl, el *Logger
)

// Config holds the application configuration settings. The struct tags
// must match the field names in the JSON configuration file.
type Config struct {

	Client struct {

		HostName string				// Hostname or IP address of client
		Timeout time.Duration			// Time limit for entire request
		IdleConnTimeout time.Duration		// Time limit for idle connections
		ResponseHeaderTimeout time.Duration	// Time limit for response headers
		MaxResponseHeaderBytes int64		// Size limit for response headers
	}

	Server struct {

		Protocol string				// Protocol for server connections
		HostName string				// Hostname or IP address of server
		Port string				// TCP port on which server listens

		Auth struct {
			Username string			// Username for client utility
			Password string			// Password for client utility 
		}

		Endpoints map[string]string		// REST server API endpoints
	}

	Paths struct {
		ReportDir string
	}

	Syslog *Syslog
	Loggers *Loggers

	Include struct {
		VendorID map[string]bool
		ProductID map[string]map[string]bool
		Default bool
	}

	DebugLevel int
}

// newConfig retrieves the settings in the JSON configuration file and
// populates the fields in the runtime configuration. It also creates
// directories if they do not already exist.
func newConfig(cf string) (*Config, error) {

	this := &Config{}

	if dn := filepath.Dir(cf); dn == `` {
		cf = filepath.Join(filepath.Dir(os.Args[0]), cf)
	}

	// Load the configuration.

	if err := loadConfig(this, cf); err != nil {
		return nil, err
	}

	// Configure HTTP client.

	if hn, err := os.Hostname(); err != nil {
		return nil, err
	} else {
		this.Client.HostName = hn
	}

	httpTransport = &http.Transport{
		IdleConnTimeout: this.Client.IdleConnTimeout * time.Second,
		ResponseHeaderTimeout: this.Client.ResponseHeaderTimeout * time.Second,
		MaxResponseHeaderBytes: this.Client.MaxResponseHeaderBytes,
	}

	httpClient = &http.Client{
		Timeout: this.Client.Timeout * time.Second,
		Transport: httpTransport,
	}

	// Prepend protocol, host, and port to endpoints.

	for key, path := range this.Server.Endpoints {

		baseUrl := fmt.Sprintf(`%s://%s:%s`,
			this.Server.Protocol,
			this.Server.HostName,
			this.Server.Port,
		)

		this.Server.Endpoints[key] = baseUrl + path
	}

	// Create and initialize the Syslog object.

	if err := this.Syslog.Init(); err != nil {
		return nil, err
	}

	// Create and initialize the Loggers object.

	if err := this.Loggers.Init(this.Syslog); err != nil {
		return nil, err
	}

	// Ensure required loggers are present and create aliases.

        var ok bool

        if sl, ok = this.Loggers.Logger[`system`]; !ok {
                return nil, fmt.Errorf(`missing "system" log config`)
        }
        if cl, ok = this.Loggers.Logger[`change`]; !ok {
                return nil, fmt.Errorf(`missing "change" log config`)
        }
        if el, ok = this.Loggers.Logger[`error`]; !ok {
                return nil, fmt.Errorf(`missing "error" log config`)
        }

	// Create report directory.

	if dn, err := makePath(this.Paths.ReportDir); err != nil {
		return nil, err
	} else {
		this.Paths.ReportDir = dn
	}

	// Create http client cookie jar.

	if jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List}); err != nil {
		return nil, err
	} else {
		httpClient.Jar = jar
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
		path = filepath.Join(filepath.Dir(os.Args[0]), path)
	}

	return path, os.MkdirAll(path, DirMode)
}

// displayVersion displays the program version.
func displayVersion() {
        fmt.Fprintf(os.Stderr, "%s version %s\n", program, version)
}
