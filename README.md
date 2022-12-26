# PM5 Emulator

[![alt text](https://godoc.org/github.com/raralabs/pm5-emulator?status.svg)](https://godoc.org/github.com/raralabs/pm5-emulator)
[![Go Report Card](https://goreportcard.com/badge/github.com/raralabs/pm5-emulator)](https://goreportcard.com/report/github.com/raralabs/pm5-emulator)
[![Build Status](https://api.travis-ci.com/raralabs/pm5-emulator.svg?token=DW7fs8Y8ozBN3DSsN2Uo&branch=master)](https://travis-ci.com/github/raralabs/pm5-emulator)
[![Coverage Status](https://coveralls.io/repos/github/raralabs/pm5-emulator/badge.svg?branch=master)](https://coveralls.io/github/raralabs/pm5-emulator?branch=master)

[![Matrix](https://img.shields.io/matrix/pm5-emulator:matrix.org?label=chat%20on%20matrix)](https://matrix.to/#/#pm5-emulator:matrix.org)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](code_of_conduct.md)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE.md)

PM5 Rower is an indoor rower from Concept2.
For detailed information: [https://www.concept2.com/service/indoor-rowers]

This project aims to emulate a PM5 Rower via GATT server in BLE Device, so that
any third-party applications can use the emulator to test, debug and stuffs.

## PM5 Services and Their UUIDs

| Service       | UUIDs     |
|:-------       | :----     |
| GAP           | 0x1800    |
| GATT          | 0x1801    |
| Device Info   | 0x0010    |
| Control       | 0x0020    |
| Rowing        | 0x0030    |

## Instructions to Run

Before starting the app, turn BLE service down:

```bash
sudo hciconfig hci0 down
```

Run tests and Build emulator:

```bash
make all
```

Advertise your custom PM5 service (as emulator):

```bash
sudo ./pm5-emulator
```

## Common Errors

***rf-kill errror***

```bash
 error while opening device 0: operation not possible due to RF-kill
```

This error is due to blocked bluetooth service. Try:

```bash
rfkill unblock bluetooth
```

## PM5 State Diagram

![SM](docs/resources/StateDiagram.png)

## MAINTAINERS

[Anish Bhusal](https://www.github.com/anisbhsl)

[Nischal Nepal](https://www.github.com/n-is)

## CONTRIBUTING

All kind of contributions are welcome. See [here](CONTRIBUTING.md) for more details!

## CONTACT-US

[RARA LABS](https://raralabs.com/)
