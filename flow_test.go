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
	`crypto/sha256`
	`fmt`
	`io/ioutil`
	`os`
	`path/filepath`
	`reflect`
	`strings`
	`testing`
	`time`
	`github.com/google/gousb`
	`github.com/jscherff/gocmdb/usbci`
	`github.com/jscherff/gotest`
)

/* TODO: test each path through the application by setting flags and
   passing an object to the router for both magtek and generic devices.

	Actions:

	  -audit
	        Audit devices
	  -checkin
	        Check devices in
	  -legacy
	        Legacy operation
	  -report
	        Report actions
	  -reset
	        Reset device
	  -serial
	        Set serial number


	Audit Options:

	  -local
	        Audit against local state
	  -server
	        Audit against server state


	Report Options:

	  -console
	        Write reports to console
	  -folder <path>
	        Write reports to <path>
	  -format <format>
	        Report <format> {csv|nvp|xml|json}


	Serial Options:

	  -copy
	        Copy factory serial number
	  -erase
	        Erase current serial number
	  -fetch
	        Fetch serial number from server
	  -force
	        Force serial number change
	  -set "TESTING"
	        Set serial number to <string>
*/

func TestFlowAudit(t *testing.T) {

	var err error

	t.Run(`Flags: -audit -local`, func(t *testing.T) {

		resetFlags(t)
		*fActionAudit = true
		*fAuditLocal = true

		// Remove audit file artifacts from previous tests.

		af := fmt.Sprintf(`%s-%s-%s.json`, mag1.VID(), mag1.PID(), mag1.ID())
		os.RemoveAll(filepath.Join(conf.Paths.StateDir, af))

		// Send device to router.

		err = magtekRouter(mag1)
		gotest.Assert(t, err != nil, `first run should result in file-not-found error`)

		// Determine whether there are no changes recorded when auditing same device.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		gotest.Assert(t, len(mag1.Changes) == 0, `device change log should be empty`)

		// Determine whether device differences are recorded in device change log.

		err = magtekRouter(mag2)
		gotest.Ok(t, err)

		gotest.Assert(t, reflect.DeepEqual(mag2.Changes, magChanges),
			`device change log does not contain known device differences`)

		// Determine whether device differences are recorded in app change log.

		fb, err := ioutil.ReadFile(conf.Files.ChangeLog)
		gotest.Ok(t, err)

		fs := string(fb)
		gotest.Assert(t, strings.Contains(fs, ClogCh1) && strings.Contains(fs, ClogCh2),
			`application change log does not contain known device differences`)
	})

	t.Run(`Flags: -audit -server`, func(t *testing.T) {

		resetFlags(t)
		*fActionAudit = true
		*fAuditServer = true

		// Send device to router.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		// Determine whether there are no changes recorded when auditing same device.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		// Determine whether device differences are recorded in device change log.

		err = magtekRouter(mag2)
		gotest.Ok(t, err)

		gotest.Assert(t, reflect.DeepEqual(mag2.Changes, magChanges),
			`device change log does not contain known device differences`)

		// Determine whether device differences are recorded in app change log.

		fb, err := ioutil.ReadFile(conf.Files.ChangeLog)
		gotest.Ok(t, err)

		fs := string(fb)
		gotest.Assert(t, strings.Contains(fs, ClogCh1) && strings.Contains(fs, ClogCh2),
			`application change log does not contain known device differences`)
	})

	restoreState(t)
}

func TestFlowCheckin(t *testing.T) {

	var err error

	t.Run(`Flags: -checkin`, func(t *testing.T) {

		resetFlags(t)
		*fActionCheckin = true

		// Change a property.

		mag2.VendorName = `Check-in Test`

		// Send device to router.

		err = magtekRouter(mag2)
		gotest.Ok(t, err)

		// Checkout device and test if property change persisted.

		b, err := checkoutDevice(mag2)
		gotest.Ok(t, err)

		err = mag2.RestoreJSON(b)
		gotest.Ok(t, err)

		gotest.Assert(t, mag2.VendorName == `Check-in Test`, `device changes did not persist after checkin`)
	})

	restoreState(t)
}

func TestFlowLegacy(t *testing.T) {

	var err error

	t.Run(`Flags: -legacy`, func(t *testing.T) {

		resetFlags(t)
		*fActionLegacy = true

		// Send device to router.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		// Test whether signature of report file content is correct.

		b, err := ioutil.ReadFile(conf.Files.Legacy)
		gotest.Ok(t, err)

		gotest.Assert(t, mag1SigLegacy == sha256.Sum256(b), `unexpected hash signature of Legacy report`)
	})

	restoreState(t)
}

func TestFlowReport(t *testing.T) {

	var err error

	t.Run(`Flags: -report -folder -format csv`, func(t *testing.T) {

		resetFlags(t)
		*fActionReport = true
		*fReportFormat = `csv`

		// Send device to router.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		// Test whether signature of report file content is correct.

		fn := filepath.Join(conf.Paths.ReportDir, mag1.Filename() + `.` + *fReportFormat)
		b, err := ioutil.ReadFile(fn)
		gotest.Ok(t, err)

		gotest.Assert(t, sha256.Sum256(b) == mag1SigCSV, `unexpected hash signature of CSV report`)
	})

	t.Run(`Flags: -report -folder -format nvp `, func(t *testing.T) {

		resetFlags(t)
		*fActionReport = true
		*fReportFormat = `nvp`

		// Send device to router.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		// Test whether signature of report file content is correct.

		fn := filepath.Join(conf.Paths.ReportDir, mag1.Filename() + `.` + *fReportFormat)
		b, err := ioutil.ReadFile(fn)
		gotest.Ok(t, err)

		gotest.Assert(t, sha256.Sum256(b) == mag1SigNVP, `unexpected hash signature of NVP report`)
	})

	t.Run(`Flags: -report -folder -format xml`, func(t *testing.T) {

		resetFlags(t)
		*fActionReport = true
		*fReportFormat = `xml`

		// Send device to router.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		// Test whether signature of report file content is correct.

		fn := filepath.Join(conf.Paths.ReportDir, mag1.Filename() + `.` + *fReportFormat)
		b, err := ioutil.ReadFile(fn)
		gotest.Ok(t, err)

		gotest.Assert(t, sha256.Sum256(b) == mag1SigPXML, `unexpected hash signature of XML report`)
	})

	t.Run(`Flags: -report -folder -format json`, func(t *testing.T) {

		resetFlags(t)
		*fActionReport = true
		*fReportFormat = `json`

		// Send device to router.

		err = magtekRouter(mag1)
		gotest.Ok(t, err)

		// Test whether signature of report file content is correct.

		fn := filepath.Join(conf.Paths.ReportDir, mag1.Filename() + `.` + *fReportFormat)
		b, err := ioutil.ReadFile(fn)
		gotest.Ok(t, err)

		gotest.Assert(t, sha256.Sum256(b) == mag1SigPJSON, `unexpected hash signature of JSON report`)
	})
}

func TestFlowSerial(t *testing.T) {

	var (
		mdev *usbci.Magtek
		err error
	)

	ctx := gousb.NewContext()
	defer ctx.Close()

	if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
		t.Skip(`device not found`)
	}

	oldSn := mdev.DeviceSN
	newSn := `TESTING`

	err = mdev.SetDeviceSN(newSn)
	gotest.Ok(t, err)
	mdev.Close()

	t.Run(`Flags: -serial -copy (serial number exists)`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		if mdev.FactorySN == `` {
			t.Skip(`factory SN empty`)
		}

		if mdev.DeviceSN == `` {
			t.Skip(`device SN empty`)
		}

		resetFlags(t)
		*fActionSerial = true
		*fSerialCopy = true

		err = magtekRouter(mdev)
		gotest.Assert(t, err != nil, `attempt to set SN when one already exists should produce error`)
	})

	t.Run(`Flags: -serial -fetch (serial number exists)`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		if mdev.DeviceSN == `` {
			t.Skip(`device SN empty`)
		}

		resetFlags(t)
		*fActionSerial = true
		*fSerialFetch = true

		err = magtekRouter(mdev)
		gotest.Assert(t, err != nil, `attempt to set SN when one already exists should produce error`)
	})

	t.Run(`Flags: -serial -set <string> (serial number exists)`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		if mdev.DeviceSN == `` {
			t.Skip(`device SN empty`)
		}

		resetFlags(t)
		*fActionSerial = true
		*fSerialSet = newSn

		err = magtekRouter(mdev)
		gotest.Assert(t, err != nil, `attempt to set SN when one already exists should produce error`)
	})

	t.Run(`Flags: -serial -erase -copy`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		if mdev.FactorySN == `` {
			t.Skip(`factory SN empty`)
		}

		resetFlags(t)
		*fActionSerial = true
		*fSerialErase = true
		*fSerialCopy = true

		err = magtekRouter(mdev)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.DeviceSN == mdev.FactorySN[:7], `attempt to set device SN to factory SN failed`)
	})

	t.Run(`Flags: -serial -erase -fetch`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		resetFlags(t)
		*fActionSerial = true
		*fSerialErase = true
		*fSerialFetch = true

		err = magtekRouter(mdev)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.DeviceSN[:3] == `24F`, `attempt to set device SN from server failed`)
	})

	t.Run(`Flags: -serial -erase -set <string>`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		resetFlags(t)
		*fActionSerial = true
		*fSerialErase = true
		*fSerialSet = newSn

		err = magtekRouter(mdev)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.DeviceSN == newSn, `attempt to set device SN to string failed`)
	})

	t.Run(`Flags: -serial -force -copy (serial number exists)`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		if mdev.FactorySN == `` {
			t.Skip(`factory SN empty`)
		}

		resetFlags(t)
		*fActionSerial = true
		*fSerialForce = true
		*fSerialCopy = true

		err = magtekRouter(mdev)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.DeviceSN == mdev.FactorySN[:7], `attempt to set device SN to factory SN failed`)
	})

	t.Run(`Flags: -serial -force -fetch (serial number exists)`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		resetFlags(t)
		*fActionSerial = true
		*fSerialForce = true
		*fSerialFetch = true

		err = magtekRouter(mdev)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.DeviceSN[:3] == `24F`, `attempt to set device SN from server failed`)
	})

	t.Run(`Flags: -serial -force -set <string> (serial number exists)`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		resetFlags(t)
		*fActionSerial = true
		*fSerialForce = true
		*fSerialSet = newSn

		err = magtekRouter(mdev)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.DeviceSN == newSn, `attempt to set device SN to string failed`)
	})

	if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
		t.Skip(`device not found`)
	}

	err = mdev.SetDeviceSN(oldSn)
	gotest.Ok(t, err)
	mdev.Close()

	t.Run(`Flags: -reset`, func(t *testing.T) {

		mux.Lock()
		defer mux.Unlock()

		if mdev, err = getMagtekDevice(t, ctx); mdev == nil {
			t.Skip(`device not found`)
		}

		defer mdev.Close()

		resetFlags(t)
		*fActionReset = true

		err = magtekRouter(mdev)
		gotest.Ok(t, err)

		time.Sleep(5 * time.Second)
	})
}
