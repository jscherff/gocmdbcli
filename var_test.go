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

import `github.com/jscherff/gocmdb/usbci`

var (
	gen1JSON = []byte(
	`{
		"host_name": "John-SurfacePro",
		"vendor_id": "0acd",
		"product_id": "2030",
		"vendor_name": "ID TECH",
		"product_name": "TM3 Magstripe USB-HID Keyboard Reader",
		"serial_num": "",
		"software_id": "",
		"product_ver": "",
		"bus_number": 1,
		"bus_address": 14,
		"port_number": 1,
		"buffer_size": 0,
		"max_pkt_size": 8,
		"usb_spec": "2.00",
		"usb_class": "per-interface",
		"usb_subclass": "per-interface",
		"usb_protocol": "0",
		"device_speed": "full",
		"device_ver": "1.00",
		"object_type": "*usbci.Generic",
		"device_sn": "",
		"factory_sn": "",
		"descriptor_sn": ""
	}`)

	gen2JSON = []byte(
	`{
		"host_name": "John-SurfacePro",
		"vendor_id": "0acd",
		"product_id": "2030",
		"vendor_name": "ID TECH",
		"product_name": "TM3 Magstripe USB-HID Keyboard Reader",
		"serial_num": "",
		"software_id": "",
		"product_ver": "",
		"bus_number": 1,
		"bus_address": 14,
		"port_number": 1,
		"buffer_size": 0,
		"max_pkt_size": 8,
		"usb_spec": "2.00",
		"usb_class": "per-interface",
		"usb_subclass": "per-interface",
		"usb_protocol": "0",
		"device_speed": "full",
		"device_ver": "1.00",
		"object_type": "*usbci.Generic",
		"device_sn": "",
		"factory_sn": "",
		"descriptor_sn": ""
	}`)

	mag1JSON = []byte(
	`{
		"host_name": "John-SurfacePro",
		"vendor_id": "0801",
		"product_id": "0001",
		"vendor_name": "Mag-Tek",
		"product_name": "USB Swipe Reader",
		"serial_num": "24F0014",
		"software_id": "21042818B01",
		"product_ver": "",
		"bus_number": 1,
		"bus_address": 13,
		"port_number": 1,
		"buffer_size": 24,
		"max_pkt_size": 8,
		"usb_spec": "1.10",
		"usb_class": "per-interface",
		"usb_subclass": "per-interface",
		"usb_protocol": "0",
		"device_speed": "full",
		"device_ver": "1.00",
		"object_type": "*usbci.Magtek",
		"device_sn": "24F0014",
		"factory_sn": "",
		"descriptor_sn": "24F0014"
	}`)

	mag2JSON = []byte(
	`{
		"host_name": "John-SurfacePro",
		"vendor_id": "0801",
		"product_id": "0001",
		"vendor_name": "Mag-Tek",
		"product_name": "USB Swipe Reader",
		"serial_num": "24F0014",
		"software_id": "21042818B02",
		"product_ver": "",
		"bus_number": 1,
		"bus_address": 13,
		"port_number": 1,
		"buffer_size": 24,
		"max_pkt_size": 8,
		"usb_spec": "2.00",
		"usb_class": "per-interface",
		"usb_subclass": "per-interface",
		"usb_protocol": "0",
		"device_speed": "full",
		"device_ver": "1.00",
		"object_type": "*usbci.Magtek",
		"device_sn": "24F0014",
		"factory_sn": "",
		"descriptor_sn": "24F0014"
	}`)

	mag1SigPrettyJSON = [32]byte{
		0x52,0x09,0x1d,0x84,0x2a,0x76,0xf5,0x1e,
		0x8c,0x99,0xb7,0x8b,0x37,0x22,0x88,0x22,
		0x50,0xdd,0x5a,0x75,0x0c,0x7d,0x1e,0xf5,
		0x3b,0x2f,0xb5,0x5a,0x05,0x12,0xf3,0x20,
	}

	mag1SigPrettyXML = [32]byte{
		0x0f,0x05,0x4e,0x13,0x51,0x5e,0x90,0x9d,
		0x3d,0x39,0xfb,0xb8,0x6a,0x14,0x20,0xcb,
		0x3a,0xd0,0xb6,0x79,0xa5,0x56,0xad,0xf7,
		0xce,0xff,0x31,0xdc,0x56,0x2a,0xbd,0x92,
	}

	mag1SigJSON = [32]byte{
		0xf8,0xec,0xeb,0x90,0xed,0x14,0x16,0xf8,
		0x82,0xbe,0x0a,0xd5,0x03,0x39,0x22,0x98,
		0x99,0x1c,0x81,0xe6,0x40,0x65,0x12,0x9c,
		0x41,0x92,0x18,0x0a,0xf8,0x26,0xcc,0x7c,
	}

	mag1SigXML = [32]byte{
		0x82,0xc7,0x14,0x84,0xee,0xb5,0x4a,0x91,
		0xfc,0x92,0xa6,0x8b,0xeb,0xf7,0xd4,0x66,
		0x93,0xad,0xc0,0x6b,0x89,0x3e,0x99,0x11,
		0x28,0xfc,0x7e,0x61,0xf3,0x4f,0x7c,0xed,
	}

	mag1SigCSV = [32]byte{
		0x98,0xd5,0xe9,0x1d,0x6f,0xa9,0xe8,0xfe,
		0x7c,0xd6,0xa8,0xa0,0x7e,0x88,0x48,0xd4,
		0xcf,0x8b,0x04,0x9c,0x05,0x3e,0x1b,0x58,
		0x41,0x3c,0xf8,0x3e,0x27,0x8a,0x98,0xea,
	}

	mag1SigNVP = [32]byte{
		0xd0,0xc4,0xea,0x8b,0x3c,0x80,0xae,0x79,
		0xe8,0x0e,0x17,0x1e,0xd3,0x55,0x09,0x88,
		0xbb,0x2b,0x11,0x84,0xac,0x3d,0xd9,0x42,
		0x50,0xc4,0x5d,0x5e,0x70,0xd3,0x65,0xe2,
	}

	mag2SigPrettyJSON = [32]byte{
		0xb3,0xc3,0x83,0x2f,0x77,0x4e,0xfb,0x15,
		0xc5,0x37,0x07,0xca,0x1c,0xee,0x23,0x2d,
		0x5b,0x9c,0x5d,0x7f,0xda,0x5d,0x7a,0x81,
		0x71,0x07,0x7e,0x08,0x0e,0xcc,0xdd,0x8a,
	}

	mag2SigPrettyXML = [32]byte{
		0x4e,0xea,0x5a,0xd6,0x73,0x93,0xba,0x1e,
		0x99,0x90,0xbd,0xcb,0x9e,0x1b,0x54,0xef,
		0xe3,0x8d,0x93,0xaf,0xe6,0x8a,0xb6,0x9d,
		0x70,0x17,0x03,0xec,0xa5,0xf0,0x70,0xb1,
	}

	mag2SigJSON = [32]byte{
		0x86,0x78,0xd2,0x1f,0xc9,0x00,0xe8,0xaf,
		0xff,0x8e,0xa1,0x4d,0x97,0x6c,0xc2,0x62,
		0x38,0x87,0xf2,0x5b,0x88,0x95,0xce,0x4a,
		0x4a,0xde,0x9f,0x33,0xee,0xc0,0xf5,0xa3,
	}

	mag2SigXML = [32]byte{
		0x4c,0x9c,0x48,0x3b,0x18,0x92,0x72,0x5d,
		0xc9,0x78,0x1d,0x0a,0x96,0x08,0x99,0x22,
		0xb8,0x8c,0x4e,0x5d,0x73,0x0c,0x76,0x60,
		0xd2,0x15,0xc5,0x13,0x7f,0xb4,0x69,0xe1,
	}

	mag2SigCSV = [32]byte{
		0xfd,0x0a,0xca,0x37,0x25,0x74,0xd9,0xfe,
		0x64,0xd6,0xd4,0xbb,0x28,0x8d,0x99,0x1a,
		0xee,0x7d,0x83,0x0c,0xf2,0xaa,0xac,0xc9,
		0xc5,0x21,0xf2,0x75,0xf0,0xa5,0x2f,0xcb,
	}

	mag2SigNVP = [32]byte{
		0x08,0x22,0x11,0x8f,0x73,0xf6,0xfc,0xa6,
		0xf3,0x8f,0x20,0xe2,0x2a,0x32,0xa8,0x5f,
		0xc1,0xd8,0x8f,0xa3,0x04,0x7d,0x32,0xd2,
		0x9b,0xeb,0xe5,0x34,0x4b,0xfb,0xc3,0x9d,
	}

	mag1SigLegacy = [32]byte{
		0xb3,0xb5,0x58,0x2b,0xb2,0xd9,0x88,0x4a,
		0x78,0xd5,0xf4,0x2d,0x98,0x0c,0x2b,0x81,
		0xfd,0xd1,0x43,0xb6,0xcc,0x58,0x14,0x39,
		0x23,0x30,0x50,0x2f,0xe3,0x59,0x88,0x5a,
	}

	mag2SigLegacy = [32]byte{
		0xb3,0xb5,0x58,0x2b,0xb2,0xd9,0x88,0x4a,
		0x78,0xd5,0xf4,0x2d,0x98,0x0c,0x2b,0x81,
		0xfd,0xd1,0x43,0xb6,0xcc,0x58,0x14,0x39,
		0x23,0x30,0x50,0x2f,0xe3,0x59,0x88,0x5a,
	}

	gen1SigPrettyJSON = [32]byte{
		0x95,0xb1,0x7b,0x53,0xa9,0xda,0x70,0x66,
		0x2a,0x81,0xe7,0x53,0xda,0x3e,0x2b,0x03,
		0x88,0x6f,0x8b,0xb1,0x66,0x06,0xb6,0x6c,
		0x39,0x3e,0x14,0xa4,0xec,0x39,0xfd,0x1b,
	}

	gen1SigPrettyXML = [32]byte{
		0xde,0xcd,0x0d,0x47,0xdb,0x7f,0xb9,0xc3,
		0x42,0xf4,0xcc,0x3c,0x0b,0x2a,0x59,0x00,
		0x76,0x9f,0x6b,0x28,0x39,0x7a,0x2e,0xa9,
		0xf4,0xfe,0xd0,0x7e,0xdc,0x79,0xb5,0x31,
	}

	gen1SigJSON = [32]byte{
		0xd6,0xef,0xf6,0x63,0xae,0x85,0x84,0xf3,
		0xff,0x8a,0xb0,0xa1,0x9a,0x62,0x9d,0x0d,
		0xa9,0xd7,0x82,0x63,0x4f,0x9e,0xf2,0xc2,
		0xa4,0xda,0x4b,0x7a,0x65,0xc1,0xd1,0xe1,
	}

	gen1SigXML = [32]byte{
		0x77,0x7d,0xcc,0x77,0x18,0x54,0xf5,0x1e,
		0x32,0xb3,0x1d,0xc6,0x28,0xa4,0x8f,0x7f,
		0x36,0x03,0x3c,0x66,0xd3,0x24,0x59,0x3d,
		0x42,0x3f,0xd5,0x6d,0x00,0x18,0xb9,0x52,
	}

	gen1SigCSV = [32]byte{
		0x7c,0xe0,0x04,0x96,0x62,0x2f,0xea,0x24,
		0xb1,0x80,0x7f,0x86,0x80,0xe4,0xce,0xae,
		0xad,0x7b,0x7b,0x5a,0xee,0x99,0xb8,0x15,
		0x25,0x27,0x18,0x11,0x04,0xaa,0x0a,0x97,
	}

	gen1SigNVP = [32]byte{
		0x0b,0xcf,0xed,0x21,0x18,0xf7,0x7a,0xb8,
		0x28,0x11,0xab,0x99,0xf1,0x34,0x41,0x4d,
		0x10,0xe5,0xa1,0x7e,0xa5,0x8d,0x29,0xd1,
		0x7c,0x5a,0xe9,0x29,0xd0,0x11,0x4d,0x31,
	}

	gen2PrettyJSON = [32]byte{
		0x95,0xb1,0x7b,0x53,0xa9,0xda,0x70,0x66,
		0x2a,0x81,0xe7,0x53,0xda,0x3e,0x2b,0x03,
		0x88,0x6f,0x8b,0xb1,0x66,0x06,0xb6,0x6c,
		0x39,0x3e,0x14,0xa4,0xec,0x39,0xfd,0x1b,
	}

	gen2SigPrettyXML = [32]byte{
		0xde,0xcd,0x0d,0x47,0xdb,0x7f,0xb9,0xc3,
		0x42,0xf4,0xcc,0x3c,0x0b,0x2a,0x59,0x00,
		0x76,0x9f,0x6b,0x28,0x39,0x7a,0x2e,0xa9,
		0xf4,0xfe,0xd0,0x7e,0xdc,0x79,0xb5,0x31,
	}

	gen2SigJSON = [32]byte{
		0xd6,0xef,0xf6,0x63,0xae,0x85,0x84,0xf3,
		0xff,0x8a,0xb0,0xa1,0x9a,0x62,0x9d,0x0d,
		0xa9,0xd7,0x82,0x63,0x4f,0x9e,0xf2,0xc2,
		0xa4,0xda,0x4b,0x7a,0x65,0xc1,0xd1,0xe1,
	}

	gen2SigXML = [32]byte{
		0x77,0x7d,0xcc,0x77,0x18,0x54,0xf5,0x1e,
		0x32,0xb3,0x1d,0xc6,0x28,0xa4,0x8f,0x7f,
		0x36,0x03,0x3c,0x66,0xd3,0x24,0x59,0x3d,
		0x42,0x3f,0xd5,0x6d,0x00,0x18,0xb9,0x52,
	}

	gen2SigCSV = [32]byte{
		0x7c,0xe0,0x04,0x96,0x62,0x2f,0xea,0x24,
		0xb1,0x80,0x7f,0x86,0x80,0xe4,0xce,0xae,
		0xad,0x7b,0x7b,0x5a,0xee,0x99,0xb8,0x15,
		0x25,0x27,0x18,0x11,0x04,0xaa,0x0a,0x97,
	}

	gen2SigNVP = [32]byte{
		0x0b,0xcf,0xed,0x21,0x18,0xf7,0x7a,0xb8,
		0x28,0x11,0xab,0x99,0xf1,0x34,0x41,0x4d,
		0x10,0xe5,0xa1,0x7e,0xa5,0x8d,0x29,0xd1,
		0x7c,0x5a,0xe9,0x29,0xd0,0x11,0x4d,0x31,
	}

	mag1, mag2 *usbci.Magtek
	gen1, gen2 *usbci.Generic

	magChanges = make([][]string, 2)
)
