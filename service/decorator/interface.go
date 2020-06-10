package decorator

import "github.com/bettercap/gatt"

// IServiceDecorator wraps a gatt service.
type IServiceDecorator interface {
	// AddCharacteristic adds a characteristic to a service.
	// AddCharacteristic panics if the service already contains another
	// characteristic with the same UUID.
	AddCharacteristic(u gatt.UUID) ICharDecorator

	// UUID returns the UUID of the service.
	UUID() gatt.UUID

	// Name returns the specification name of the service according to its UUID.
	// If the UUID is not assignee, Name returns an empty string.
	Name() string

	// Handle returns the Handle of the service.
	Handle() uint16

	// EndHandle returns the End Handle of the service.
	EndHandle() uint16

	// SetHandle sets the Handle of the service.
	SetHandle(h uint16)

	// SetEndHandle sets the End Handle of the service.
	SetEndHandle(endh uint16)

	// SetCharacteristics sets the Characteristics of the service.
	SetCharacteristics(chars []*gatt.Characteristic)

	// Characteristic returns the contained characteristic of this service.
	Characteristics() []*gatt.Characteristic
}


// ICharDecorator wraps a gatt characteristics.
type ICharDecorator interface {

	// UUID returns the UUID of the characteristic.
	UUID() gatt.UUID

	// Name returns the specification name of the characteristic according to its UUID.
	// If the UUID is not assignee, Name returns an empty string.
	Name() string

	// Handle returns the Handle of the characteristic.
	Handle() uint16

	// VHandle returns the Value Handle of the characteristic.
	VHandle() uint16

	// EndHandle returns the End Handle of the characteristic.
	EndHandle() uint16

	// Descriptor returns the Descriptor of the characteristic.
	Descriptor() *gatt.Descriptor

	// SetHandle sets the Handle of the characteristic.
	SetHandle(h uint16)

	// SetVHandle sets the Value Handle of the characteristic.
	SetVHandle(vh uint16)

	// SetEndHandle sets the End Handle of the characteristic.
	SetEndHandle(endh uint16)

	// SetDescriptor sets the Descriptor of the characteristic.
	SetDescriptor(cccd *gatt.Descriptor)

	// SetDescriptors sets the list of Descriptor of the characteristic.
	SetDescriptors(descs []*gatt.Descriptor)

	// Service returns the containing service of this characteristic.
	Service() *gatt.Service

	// Properties returns the properties of this characteristic.
	Properties() gatt.Property

	// Descriptors returns the contained descriptors of this characteristic.
	Descriptors() []*gatt.Descriptor

	// AddDescriptor adds a descriptor to a characteristic.
	// AddDescriptor panics if the characteristic already contains another
	// descriptor with the same UUID.
	AddDescriptor(u gatt.UUID) *gatt.Descriptor

	// SetValue makes the characteristic support read requests, and returns a
	// static value. SetValue must be called before the containing service is
	// added to a server.
	// SetValue panics if the characteristic has been configured with a ReadHandler.
	SetValue(b []byte)

	// HandleRead makes the characteristic support read requests, and routes read
	// requests to h. HandleRead must be called before the containing service is
	// added to a server.
	// HandleRead panics if the characteristic has been configured with a static value.
	HandleRead(h gatt.ReadHandler)

	// HandleReadFunc calls HandleRead(ReadHandlerFunc(f)).
	HandleReadFunc(f func(rsp gatt.ResponseWriter, req *gatt.ReadRequest))

	GetReadHandler() gatt.ReadHandler

	// HandleWrite makes the characteristic support write and write-no-response
	// requests, and routes write requests to h.
	// The WriteHandler does not differentiate between write and write-no-response
	// requests; it is handled automatically.
	// HandleWrite must be called before the containing service is added to a server.
	HandleWrite(h gatt.WriteHandler)

	// HandleWriteFunc calls HandleWrite(WriteHandlerFunc(f)).
	HandleWriteFunc(f func(r gatt.Request, data []byte) (status byte))

	GetWriteHandler() gatt.WriteHandler

	// HandleNotify makes the characteristic support notify requests, and routes
	// notification requests to h. HandleNotify must be called before the
	// containing service is added to a server.
	HandleNotify(h gatt.NotifyHandler)

	// HandleNotifyFunc calls HandleNotify(NotifyHandlerFunc(f)).
	HandleNotifyFunc(f func(r gatt.Request, n gatt.Notifier))

	GetNotifyHandler() gatt.NotifyHandler

	// StopNotify stops the notification channel if it's running
	StopNotify()
}