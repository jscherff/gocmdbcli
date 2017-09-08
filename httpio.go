// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	 http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jscherff/gocmdb"
)

// serialRequest obtains a serial number from the gocmdbd server.
func serialRequest(o gocmdb.Registerable) (s string, err error) {

	var j []byte
	url := fmt.Sprintf("%s/%s", config.ServerURL, config.SerialPath)

	if j, err = o.JSON(); err == nil {
		j, err = httpRequest(url, j)
	}
	if err == nil {
		err = json.Unmarshal(j, &s)
	}
	if err != nil {
		err = gocmdb.ErrorDecorator(err)
	}

	return s, err
}

// checkinRequest checks a device in with the gocmdbd server.
func checkinRequest(o gocmdb.Registerable) (err error) {

	var j []byte
	url := fmt.Sprintf("%s/%s", config.ServerURL, config.CheckinPath)

	if j, err = o.JSON(); err == nil {
		_, err = httpRequest(url, j)
	}
	if err != nil {
		err = gocmdb.ErrorDecorator(err)
	}

	return err
}

// auditRequest performs an audit and sends the results to the gocmdbd server.
func auditRequest(o gocmdb.Auditable) (err error) {

	var j []byte
	url := fmt.Sprintf("%s/%s", config.ServerURL, config.AuditPath)

	if len(o.ID()) == 0 {
		return gocmdb.ErrorDecorator(errors.New("no unique ID"))
	}
	if _, err = os.Stat(config.AuditDir); os.IsNotExist(err) {
		return gocmdb.ErrorDecorator(err)
	}

	f := filepath.Join(config.AuditDir, o.ID() + ".json")

	// If the audit file doesn't exist, create a change record indicating
	// a change from no serial number to a serial number, then create the
	// audit file. Otherwise, audit against the previous audit file.

	if _, err = os.Stat(f); os.IsNotExist(err) {
		o.AddChange("SerialNum", "", o.ID())
		err = o.Save(f)
	} else {
		err = o.AuditFile(f)
	}

	if err == nil {
		j, err = o.JSON()
	}
	if err == nil {
		_, err = httpRequest(url, j)
	}
	if err != nil {
		err = gocmdb.ErrorDecorator(err)
	}

	return err
}

// httpRequest sends JSON requests to the gocmdbd server for other functions.
// Error decoration will be handled by caller functions.
func httpRequest(url string, jreq []byte ) (jresp []byte, err error) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jreq))

	if err != nil {
		return jresp, gocmdb.ErrorDecorator(err)
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF8")
	req.Header.Add("Accept", "application/json; charset=UTF8")
	req.Header.Add("X-Custom-Header", "gocmdb")
	resp, err := client.Do(req)

	if err != nil {
		return jresp, gocmdb.ErrorDecorator(err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusAccepted:
	default:
		err = fmt.Errorf("http response status %s", resp.Status)
	}

	if err == nil {
		jresp, err = ioutil.ReadAll(resp.Body)
	}
	if err != nil {
		return jresp, gocmdb.ErrorDecorator(err)
	}

	return jresp, err
}
