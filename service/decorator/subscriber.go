package decorator

import (
	"pm5-emulator/config"
	"pm5-emulator/sm"

	"github.com/bettercap/gatt"
)

/*
 * Service Subscriber
 */
type ServiceSubscriber struct {
	service *gatt.Service // service that this decorator wraps
}

func NewServiceSubscriber(service *gatt.Service) *ServiceSubscriber {
	return &ServiceSubscriber{service: service}
}

/** Functionalities Added **/

func (s *ServiceSubscriber) AddCharacteristic(uuid gatt.UUID) ICharDecorator {
	c := s.service.AddCharacteristic(uuid)
	return NewCharSubscriber(c)
}

/** Functionalities Not Added **/

func (s *ServiceSubscriber) UUID() gatt.UUID {
	return s.service.UUID()
}

func (s *ServiceSubscriber) Name() string {
	return s.service.Name()
}

func (s *ServiceSubscriber) Handle() uint16 {
	return s.service.Handle()
}

func (s *ServiceSubscriber) EndHandle() uint16 {
	return s.service.EndHandle()
}

func (s *ServiceSubscriber) SetHandle(h uint16) {
	s.service.SetHandle(h)
}

func (s *ServiceSubscriber) SetEndHandle(endh uint16) {
	s.service.SetEndHandle(endh)
}

func (s *ServiceSubscriber) SetCharacteristics(chars []*gatt.Characteristic) {
	s.service.SetCharacteristics(chars)
}

func (s *ServiceSubscriber) Characteristics() []*gatt.Characteristic {
	return s.service.Characteristics()
}

/*
 * Characteristics Subscriber
 */

type CharSubscriber struct {
	ch     *gatt.Characteristic
	stm    *sm.StateMachine
	notify chan bool
}

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

func (c *CharSubscriber) HandleReadFunc(f func(rsp gatt.ResponseWriter, req *gatt.ReadRequest)) {
	c.ch.HandleReadFunc(f)
}

func (c *CharSubscriber) HandleWriteFunc(f func(r gatt.Request, data []byte) (status byte)) {
	fnc := func(r gatt.Request, data []byte) (status byte) {
		s := f(r, data)
		c.stm.Update(config.CSAFE_GOINUSE_CMD)
		return s
	}
	c.ch.HandleWriteFunc(fnc)
}

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

func (c *CharSubscriber) StopNotify() {
	c.notify <- true
}

/** Functionalities Not Added **/

func (c *CharSubscriber) UUID() gatt.UUID {
	return c.ch.UUID()
}

func (c *CharSubscriber) Name() string {
	return c.ch.Name()
}

func (c *CharSubscriber) Handle() uint16 {
	return c.ch.Handle()
}

func (c *CharSubscriber) VHandle() uint16 {
	return c.ch.VHandle()
}

func (c *CharSubscriber) EndHandle() uint16 {
	return c.ch.EndHandle()
}

func (c *CharSubscriber) Descriptor() *gatt.Descriptor {
	return c.ch.Descriptor()
}

func (c *CharSubscriber) SetHandle(h uint16) {
	c.ch.SetHandle(h)
}

func (c *CharSubscriber) SetVHandle(vh uint16) {
	c.ch.SetVHandle(vh)
}

func (c *CharSubscriber) SetEndHandle(endh uint16) {
	c.ch.SetEndHandle(endh)
}

func (c *CharSubscriber) SetDescriptor(cccd *gatt.Descriptor) {
	c.ch.SetDescriptor(cccd)
}

func (c *CharSubscriber) SetDescriptors(descs []*gatt.Descriptor) {
	c.ch.SetDescriptors(descs)
}

func (c *CharSubscriber) Service() *gatt.Service {
	return c.ch.Service()
}

func (c *CharSubscriber) Properties() gatt.Property {
	return c.ch.Properties()
}

func (c *CharSubscriber) Descriptors() []*gatt.Descriptor {
	return c.ch.Descriptors()
}

func (c *CharSubscriber) AddDescriptor(u gatt.UUID) *gatt.Descriptor {
	return c.ch.AddDescriptor(u)
}

func (c *CharSubscriber) SetValue(b []byte) {
	c.ch.SetValue(b)
}

func (c *CharSubscriber) HandleRead(h gatt.ReadHandler) {
	c.ch.HandleRead(h)
}

func (c *CharSubscriber) GetReadHandler() gatt.ReadHandler {
	return c.ch.GetReadHandler()
}

func (c *CharSubscriber) HandleWrite(h gatt.WriteHandler) {
	c.ch.HandleWrite(h)
}

func (c *CharSubscriber) GetWriteHandler() gatt.WriteHandler {
	return c.ch.GetWriteHandler()
}

func (c *CharSubscriber) HandleNotify(h gatt.NotifyHandler) {
	c.ch.HandleNotify(h)
}

func (c *CharSubscriber) GetNotifyHandler() gatt.NotifyHandler {
	return c.ch.GetNotifyHandler()
}
