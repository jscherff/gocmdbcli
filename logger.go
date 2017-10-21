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
	`log`
	`io`
	`io/ioutil`
	`os`
	`strings`
	`github.com/RackSec/srslog`
)

func newLoggers() (sl, cl, el *log.Logger) {

	var sw, cw, ew []io.Writer

	var newfl = func(f string) (h *os.File, err error) {

		if h, err = os.OpenFile(f, FileFlags, FileMode); err != nil {
			log.Println(err)
		}

		return h, err
	}

	var newsl = func(prot, raddr, tag string, pri srslog.Priority) (s *srslog.Writer, err error) {

		if s, err = srslog.Dial(prot, raddr, pri, tag); err != nil {
			log.Println(err)
		}

		return s, err
	}

	// File logging

	if f, err := newfl(conf.Logging.System.LogFile); err == nil {
		sw = append(sw, f)
	}
	if f, err := newfl(conf.Logging.Change.LogFile); err == nil {
		cw = append(cw, f)
	}
	if f, err := newfl(conf.Logging.Error.LogFile); err == nil {
		ew = append(ew, f)
	}

	// Console logging

	if conf.Logging.System.Console {
		sw = append(sw, os.Stdout)
	}

	if conf.Logging.Change.Console {
		cw = append(cw, os.Stdout)
	}

	if conf.Logging.Error.Console {
		ew = append(ew, os.Stderr)
	}

	// Syslog logging

	prot, port, host := conf.Syslog.Protocol, conf.Syslog.Port, conf.Syslog.Host
	raddr := strings.Join([]string{host, port}, `:`)

	if conf.Logging.System.Syslog {
		if s, err := newsl(prot, raddr, `cmdbc`, PriInfo); err == nil {
			sw = append(sw, s)
		}
	}
	if conf.Logging.Change.Syslog {
		if s, err := newsl(prot, raddr, `cmdbc`, PriInfo); err == nil {
			cw = append(cw, s)
		}
	}
	if conf.Logging.Error.Syslog {
		if s, err := newsl(prot, raddr, `cmdbc`, PriErr); err == nil {
			ew = append(ew, s)
		}
	}

	// Default logging (discard)

	if len(sw) == 0 {
		sw = append(sw, ioutil.Discard)
	}
	if len(cw) == 0 {
		cw = append(cw, ioutil.Discard)
	}
	if len(ew) == 0 {
		ew = append(ew, ioutil.Discard)
	}

	sl = log.New(io.MultiWriter(sw...), `system `, log.LstdFlags|log.Lshortfile)
	cl = log.New(io.MultiWriter(cw...), `change `, log.LstdFlags|log.Lshortfile)
	el = log.New(io.MultiWriter(ew...), `error `, log.LstdFlags|log.Lshortfile)

	return sl, cl, el
}
