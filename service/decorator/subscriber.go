package decorator

import (
	"pm5-emulator/config"
	"pm5-emulator/sm"

	"github.com/bettercap/gatt"
)

/*
 * Service Subscriber
 */

// ServiceSubscriber is a service decorator that wraps around a gatt service,
// and adds subscribing functionalities to the service.
type ServiceSubscriber struct {
	service *gatt.Service // service that this decorator wraps
}

// NewServiceSubscriber creates a new service subscriber decorator
func NewServiceSubscriber(service *gatt.Service) *ServiceSubscriber {
	return &ServiceSubscriber{service: service}
}

/** Functionalities Added **/

// AddCharacterstics adds a characterstics with a subscriber decorator
func (s *ServiceSubscriber) AddCharacteristic(uuid gatt.UUID) ICharDecorator {
	c := s.service.AddCharacteristic(uuid)
	return NewCharSubscriber(c)
}

/** Functionalities Not Added **/

// UUID returns the uuid of service
func (s *ServiceSubscriber) UUID() gatt.UUID {
	return s.service.UUID()
}

// Name returns the name of service
func (s *ServiceSubscriber) Name() string {
	return s.service.Name()
}

// Handle returns the handle to the service
func (s *ServiceSubscriber) Handle() uint16 {
	return s.service.Handle()
}

// EndHandle returns endhandle of service
func (s *ServiceSubscriber) EndHandle() uint16 {
	return s.service.EndHandle()
}

// SetHandle sets a handle to service
func (s *ServiceSubscriber) SetHandle(h uint16) {
	s.service.SetHandle(h)
}

// SetEndHandle sets an endhandle to service
func (s *ServiceSubscriber) SetEndHandle(endh uint16) {
	s.service.SetEndHandle(endh)
}

// SetCharacterstics sets provided characteristics to service
func (s *ServiceSubscriber) SetCharacteristics(chars []*gatt.Characteristic) {
	s.service.SetCharacteristics(chars)
}

// Characterstics returns all the characteristics in a service
func (s *ServiceSubscriber) Characteristics() []*gatt.Characteristic {
	return s.service.Characteristics()
}

/*
 * Characteristics Subscriber
 */

// CharSubscriber is a characteristics decorator that wraps around a gatt
// characteristics and adds subscribing functionalities to the characteristics.
type CharSubscriber struct {
	ch     *gatt.Characteristic
	stm    *sm.StateMachine
	notify chan bool
}

// NewCharSubscriber creates a new characteristics subscriber decorator
func NewCharSubscriber(ch *gatt.Characteristic) *CharSubscriber {

	stm := sm.NewStateMachine()
	// Start the state machine from READY state
	stm.Reset()

	return &CharSubscriber{
		ch:  ch,
		stm: stm,
	}
}

/** Functionalities Added **/

// Add subscriber workflow:
// 1. We need to create some client subscriber service
// 2. Client send notify request to a characteristic to subscribe
// 3. Whenever a write request arrives at a characteristic, we check whether client has subscribed to notification or
//    not. If yes, we can notify the client with the results.

// HandleReadFunc registers a functionto be called when a READ request arrives
func (c *CharSubscriber) HandleReadFunc(f func(rsp gatt.ResponseWriter, req *gatt.ReadRequest)) {
	c.ch.HandleReadFunc(f)
}

// HandleWriteFunc registers a functionto be called when a WRITE request arrives,
// it updates the state machine to INUSE state if it is IDLE
func (c *CharSubscriber) HandleWriteFunc(f func(r gatt.Request, data []byte) (status byte)) {
	fnc := func(r gatt.Request, data []byte) (status byte) {
		s := f(r, data)
		c.stm.Update(config.CSAFE_GOINUSE_CMD)
		return s
	}
	c.ch.HandleWriteFunc(fnc)
}

// HandleNotifyFunc registers a functionto be called when a NOTIFY request arrives.
// The function is called in response to a WRITE request.
func (c *CharSubscriber) HandleNotifyFunc(f func(r gatt.Request, n gatt.Notifier)) {
	fnc := func(r gatt.Request, n gatt.Notifier) {
		// Subscribe to the characteristics
		c.stm.Reset()
		c.stm.Update(config.CSAFE_GOIDLE_CMD)

		// Run the provided function if the characteristic's state machine is in INUSE state
		for {
			select {
			case <-c.notify:
				break
			default:
				if currState := c.stm.GetState(); currState == c.stm.INUSE {
					f(r, n)
					c.stm.Update(config.CSAFE_GOFINISHED_CMD)
					c.stm.Update(config.CSAFE_GOIDLE_CMD)
				}
			}
		}
	}
	c.ch.HandleNotifyFunc(fnc)
}

// StopNotify stops the notification capabilities
func (c *CharSubscriber) StopNotify() {
	c.notify <- true
}

/** Functionalities Not Added **/

// UUID returns the uuid of the characteristics
func (c *CharSubscriber) UUID() gatt.UUID {
	return c.ch.UUID()
}

// Name returns the name of the characteristics
func (c *CharSubscriber) Name() string {
	return c.ch.Name()
}

// Handle returns the handle of the characteristics
func (c *CharSubscriber) Handle() uint16 {
	return c.ch.Handle()
}

// VHandle returns the vhandle of the characteristics
func (c *CharSubscriber) VHandle() uint16 {
	return c.ch.VHandle()
}

// EndHandle returns the endhandle of the characteristics
func (c *CharSubscriber) EndHandle() uint16 {
	return c.ch.EndHandle()
}

// Descriptor returns the descriptor of the characteristics
func (c *CharSubscriber) Descriptor() *gatt.Descriptor {
	return c.ch.Descriptor()
}

// SetHandle sets the handle of the characteristics
func (c *CharSubscriber) SetHandle(h uint16) {
	c.ch.SetHandle(h)
}

// SetVHandle sets the vhandle of the characteristics
func (c *CharSubscriber) SetVHandle(vh uint16) {
	c.ch.SetVHandle(vh)
}

// SetEndHandle sets the endhandle of the characteristics
func (c *CharSubscriber) SetEndHandle(endh uint16) {
	c.ch.SetEndHandle(endh)
}

// SetDescriptor sets the descriptor of the characteristics
func (c *CharSubscriber) SetDescriptor(cccd *gatt.Descriptor) {
	c.ch.SetDescriptor(cccd)
}

// SetDescriptors sets the descriptors of the characteristics
func (c *CharSubscriber) SetDescriptors(descs []*gatt.Descriptor) {
	c.ch.SetDescriptors(descs)
}

// Service returns the service of the characteristics
func (c *CharSubscriber) Service() *gatt.Service {
	return c.ch.Service()
}

// Properties returns the properties of the characteristics
func (c *CharSubscriber) Properties() gatt.Property {
	return c.ch.Properties()
}

// Descriptors returns the descriptors of the characteristics
func (c *CharSubscriber) Descriptors() []*gatt.Descriptor {
	return c.ch.Descriptors()
}

// AddDescriptor adds a descriptor to the characteristics
func (c *CharSubscriber) AddDescriptor(u gatt.UUID) *gatt.Descriptor {
	return c.ch.AddDescriptor(u)
}

// SetValue sets the value of the characteristics
func (c *CharSubscriber) SetValue(b []byte) {
	c.ch.SetValue(b)
}

// HandleRead registers a read handler to the characteristics
func (c *CharSubscriber) HandleRead(h gatt.ReadHandler) {
	c.ch.HandleRead(h)
}

// GetReadHandler returns read handler of the characteristics
func (c *CharSubscriber) GetReadHandler() gatt.ReadHandler {
	return c.ch.GetReadHandler()
}

// HandleWrite registers a write handler to the characteristics
func (c *CharSubscriber) HandleWrite(h gatt.WriteHandler) {
	c.ch.HandleWrite(h)
}

// GetWriteHandler returns write handler of the characteristics
func (c *CharSubscriber) GetWriteHandler() gatt.WriteHandler {
	return c.ch.GetWriteHandler()
}

// HandleNotify registers a notify handler to the characteristics
func (c *CharSubscriber) HandleNotify(h gatt.NotifyHandler) {
	c.ch.HandleNotify(h)
}

// GetNotifyHandler returns notify handler of the characteristics
func (c *CharSubscriber) GetNotifyHandler() gatt.NotifyHandler {
	return c.ch.GetNotifyHandler()
}
