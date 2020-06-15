
## PM5 Emulator

[![alt text](https://godoc.org/github.com/raralabs/pm5-emulator?status.svg)](https://godoc.org/github.com/raralabs/pm5-emulator)
[![Go Report Card](https://goreportcard.com/badge/github.com/raralabs/pm5-emulator)](https://goreportcard.com/report/github.com/raralabs/pm5-emulator)
[![Build Status](https://api.travis-ci.com/raralabs/pm5-emulator.svg?token=DW7fs8Y8ozBN3DSsN2Uo&branch=master)](https://travis-ci.com/github/raralabs/pm5-emulator)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](code_of_conduct.md)

Emulating PM5 Rower via GATT server in BLE Device


**Instructions to Run**

Before starting the app, turn BLE service down: 
```
sudo hciconfig hci0 down 
```
Run tests and Build emulator:
```cassandraql
make all 
```

Advertise your custom PM5 service (as emulator):
```cassandraql
sudo ./pm5-emulator 
```


**Common Errors**

*rf-kill errror* 
```cassandraql
 error while opening device 0: operation not possible due to RF-kill
```

This error is due to blocked bluetooth service. Try ```rfkill unblock bluetooth```


**PM5 State Diagram**

![SM](docs/resources/StateDiagram.png)

**CONTRIBUTING**

All kind of contributions are welcome. See [here](CONTRIBUTING.md) for more details!


**REFERENCES**

[PM5-Specs](https://www.concept2.co.uk/files/pdf/us/monitors/PM5_BluetoothSmartInterfaceDefinition.pdf)

[Intro to GATT](https://www.oreilly.com/library/view/getting-started-with/9781491900550/ch04.html)

[BLE Stack](https://www.mathworks.com/help/comm/examples/ble-l2cap-frame-generation-and-decoding.html)


