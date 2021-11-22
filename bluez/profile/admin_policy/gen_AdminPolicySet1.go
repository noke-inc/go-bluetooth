// Code generated by go-bluetooth generator DO NOT EDIT.

package admin_policy

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/muka/go-bluetooth/bluez"
	"github.com/muka/go-bluetooth/props"
	"github.com/muka/go-bluetooth/util"
)

var AdminPolicySet1Interface = "org.bluez.AdminPolicySet1"

// NewAdminPolicySet1 create a new instance of AdminPolicySet1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}
func NewAdminPolicySet1(objectPath dbus.ObjectPath) (*AdminPolicySet1, error) {
	a := new(AdminPolicySet1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: AdminPolicySet1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(AdminPolicySet1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

/*
AdminPolicySet1 Admin Policy Set hierarchy

*/
type AdminPolicySet1 struct {
	client                 *bluez.Client
	propertiesSignal       chan *dbus.Signal
	objectManagerSignal    chan *dbus.Signal
	objectManager          *bluez.ObjectManager
	Properties             *AdminPolicySet1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// AdminPolicySet1Properties contains the exposed properties of an interface
type AdminPolicySet1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`
}

//Lock access to properties
func (p *AdminPolicySet1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *AdminPolicySet1Properties) Unlock() {
	p.lock.Unlock()
}

// Close the connection
func (a *AdminPolicySet1) Close() {
	a.unregisterPropertiesSignal()
	a.client.Disconnect()
}

// Path return AdminPolicySet1 object path
func (a *AdminPolicySet1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return AdminPolicySet1 dbus client
func (a *AdminPolicySet1) Client() *bluez.Client {
	return a.client
}

// Interface return AdminPolicySet1 interface
func (a *AdminPolicySet1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *AdminPolicySet1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}

// ToMap convert a AdminPolicySet1Properties to map
func (a *AdminPolicySet1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an AdminPolicySet1Properties
func (a *AdminPolicySet1Properties) FromMap(props map[string]interface{}) (*AdminPolicySet1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an AdminPolicySet1Properties
func (a *AdminPolicySet1Properties) FromDBusMap(props map[string]dbus.Variant) (*AdminPolicySet1Properties, error) {
	s := new(AdminPolicySet1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *AdminPolicySet1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *AdminPolicySet1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *AdminPolicySet1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *AdminPolicySet1) GetProperties() (*AdminPolicySet1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *AdminPolicySet1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *AdminPolicySet1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *AdminPolicySet1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *AdminPolicySet1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *AdminPolicySet1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *AdminPolicySet1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}

/*
SetServiceAllowList 			This method sets the service allowlist by specifying
			service UUIDs.
			When SetServiceAllowList is called, bluez will block
			incoming and outgoing connections to the service not in
			UUIDs for all of the clients.
			Any subsequent calls to this method will supersede any
			previously set allowlist values.  Calling this method
			with an empty array will allow any service UUIDs to be
			used.
			The default value is an empty array.
			Possible errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.Failed

*/
func (a *AdminPolicySet1) SetServiceAllowList(UUIDs []string) error {
	return a.client.Call("SetServiceAllowList", 0, UUIDs).Store()
}
