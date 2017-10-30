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
	`fmt`
	`github.com/jscherff/cmdb/ci/peripheral/usb`
)

// serial processes options and configures the the serial number.
func serial(dev usb.Serializer) (err error) {

	var s string

	if *fSerialErase {

		sl.Printf(`device %s-%s erasing serial %q`,
			dev.VID(), dev.PID(), dev.SN(),
		)

		if err = dev.EraseDeviceSN(); err != nil {
			return err
		}
	}

	if !*fSerialForce && dev.SN() != `` {

		return fmt.Errorf(`device %s-%s serial already set to %q`,
			dev.VID(), dev.PID(), dev.SN(),
		)

	}

	switch {

	case *fSerialFetch:

		if s, err = usbCiNewSnV1(dev); err != nil {
			break
		}

		sl.Printf(`device %s-%s setting serial to %q`,
			dev.VID(), dev.PID(), s,
		)

		if err = dev.SetDeviceSN(s); err != nil {
			break
		}

		sl.Printf(`device %s-%s-%s checking in with server`,
			dev.VID(), dev.PID(), dev.SN(),
		)

		err = usbCiCheckinV1(dev)

	case *fSerialDefault:

		sl.Printf(`device %s-%s setting serial to default`,
			dev.VID(), dev.PID(),
		)

		err = dev.SetDefaultSN()

	case *fSerialSet != ``:

		sl.Printf(`device %s-%s setting serial to %q`,
			dev.VID(), dev.PID(), *fSerialSet,
		)

		err = dev.SetDeviceSN(*fSerialSet)
	}

	return err
}