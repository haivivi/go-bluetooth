// Code generated DO NOT EDIT

package mesh



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus/v5"
)

var Network1Interface = "org.bluez.mesh.Network1"


// NewNetwork1 create a new instance of Network1
//
// Args:

func NewNetwork1() (*Network1, error) {
	a := new(Network1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez.mesh",
			Iface: Network1Interface,
			Path:  dbus.ObjectPath("/org/bluez/mesh"),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Network1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Network1 Mesh Network Hierarchy

*/
type Network1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Network1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Network1Properties contains the exposed properties of an interface
type Network1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *Network1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Network1Properties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *Network1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Network1 object path
func (a *Network1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Network1 dbus client
func (a *Network1) Client() *bluez.Client {
	return a.client
}

// Interface return Network1 interface
func (a *Network1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Network1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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


// ToMap convert a Network1Properties to map
func (a *Network1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Network1Properties
func (a *Network1Properties) FromMap(props map[string]interface{}) (*Network1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Network1Properties
func (a *Network1Properties) FromDBusMap(props map[string]dbus.Variant) (*Network1Properties, error) {
	s := new(Network1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Network1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Network1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Network1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Network1) GetProperties() (*Network1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Network1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Network1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Network1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *Network1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Network1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Network1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
Join 		This is the first method that an application has to call to
		become a provisioned node on a mesh network. The call will
		initiate broadcasting of Unprovisioned Device Beacon.
		The app_root parameter is a D-Bus object root path of
		the application that implements org.bluez.mesh.Application1
		interface. The application represents a node where child mesh
		elements have their own objects that implement
		org.bluez.mesh.Element1 interface. The application hierarchy
		also contains a provision agent object that implements
		org.bluez.mesh.ProvisionAgent1 interface. The standard
		DBus.ObjectManager interface must be available on the
		app_root path.
		The uuid parameter is a 16-byte array that contains Device UUID.
		This UUID must be unique (at least from the daemon perspective),
		therefore attempting to call this function using already
		registered UUID results in an error.
		PossibleErrors:
			org.bluez.mesh.Error.InvalidArguments
			org.bluez.mesh.Error.AlreadyExists,

*/
func (a *Network1) Join(app_root dbus.ObjectPath, uuid []byte) error {
	
	return a.client.Call("Join", 0, app_root, uuid).Store()
	
}

/*
Cancel 
*/
func (a *Network1) Cancel() error {
	
	return a.client.Call("Cancel", 0, ).Store()
	
}

/*
Attach 		This is the first method that an application must call to get
		access to mesh node functionalities.
		The app_root parameter is a D-Bus object root path of
		the application that implements org.bluez.mesh.Application1
		interface. The application represents a node where child mesh
		elements have their own objects that implement
		org.bluez.mesh.Element1 interface. The standard
		DBus.ObjectManager interface must be available on the
		app_root path.
		The token parameter is a 64-bit number that has been assigned to
		the application when it first got provisioned/joined mesh
		network, i.e. upon receiving JoinComplete() method. The daemon
		uses the token to verify whether the application is authorized
		to assume the mesh node identity.
		In case of success, the method call returns mesh node object
		(see Mesh Node Hierarchy section) and current configuration
		settings. The return value of configuration parameter is an
		array, where each entry is a structure that contains element
		configuration. The element configuration structure is organized
		as follows:
		byte
			Element index, identifies the element to which this
			configuration entry pertains.
		array{struct}
			Models array where each entry is a structure with the
			following members:
			uint16
				Either a SIG Model Identifier or, if Vendor key
				is present in model configuration dictionary, a
				16-bit vendor-assigned Model Identifier
			dict
				A dictionary that contains model configuration
				with the following keys defined:
				array{uint16} Bindings
					Indices of application keys bound to the
					model
				uint32 PublicationPeriod
					Model publication period in milliseconds
				uint16 Vendor
					A 16-bit Company ID as defined by the
					Bluetooth SIG
				array{variant} Subscriptions
					Addresses the model is subscribed to.
					Each address is provided either as
					uint16 for group addresses, or
					as array{byte} for virtual labels.
		PossibleErrors:
			org.bluez.mesh.Error.InvalidArguments
			org.bluez.mesh.Error.NotFound,
			org.bluez.mesh.Error.AlreadyExists,
			org.bluez.mesh.Error.Failed

*/
func (a *Network1) Attach(app_root dbus.ObjectPath, token uint64) (dbus.ObjectPath, []map[byte][]map[uint16]map[string]interface{}, error) {
	
	var val0 dbus.ObjectPath
  var val1 []map[byte][]map[uint16]map[string]interface{}
	err := a.client.Call("Attach", 0, app_root, token).Store(&val0, &val1)
	return val0, val1, err	
}

/*
Leave 		This removes the configuration information about the mesh node
		identified by the 64-bit token parameter. The token parameter
		has been obtained as a result of successful Join() method call.
		PossibleErrors:
			org.bluez.mesh.Error.InvalidArguments

*/
func (a *Network1) Leave(token uint64) error {
	
	return a.client.Call("Leave", 0, token).Store()
	
}

/*
CreateNetwork 		This is the first method that an application calls to become
		a Provisioner node, and a Configuration Client on a newly
		created Mesh Network.
		The app_root parameter is a D-Bus object root path of the
		application that implements org.bluez.mesh.Application1
		interface, and a org.bluez.mesh.Provisioner1 interface. The
		application represents a node where child mesh elements have
		their own objects that implement org.bluez.mesh.Element1
		interface. The application hierarchy also contains a provision
		agent object that implements org.bluez.mesh.ProvisionAgent1
		interface. The standard DBus.ObjectManager interface must be
		available on the app_root path.
		The uuid parameter is a 16-byte array that contains Device UUID.
		This UUID must be unique (at least from the daemon perspective),
		therefore attempting to call this function using already
		registered UUID results in an error.
		The returned token must be preserved by the application in
		order to authenticate itself to the mesh daemon and attach to
		the network as a mesh node by calling Attach() method or
		permanently remove the identity of the mesh node by calling
		Leave() method.
		The other information the bluetooth-meshd daemon will preserve
		about the initial node, is to give it the initial primary
		unicast address (0x0001), and create and assign a net_key as the
		primary network net_index (0x000).
		PossibleErrors:
			org.bluez.mesh.Error.InvalidArguments
			org.bluez.mesh.Error.AlreadyExists,

*/
func (a *Network1) CreateNetwork(app_root dbus.ObjectPath, uuid []byte) (uint64, error) {
	
	var val0 uint64
	err := a.client.Call("CreateNetwork", 0, app_root, uuid).Store(&val0)
	return val0, err	
}

/*
Import 		This method creates a local mesh node based on node
		configuration that has been generated outside bluetooth-meshd.
		The app_root parameter is a D-Bus object root path of the
		application that implements org.bluez.mesh.Application1
		interface.
		The uuid parameter is a 16-byte array that contains Device UUID.
		This UUID must be unique (at least from the daemon perspective),
		therefore attempting to call this function using already
		registered UUID results in an error.
		The dev_key parameter is the 16-byte value of the dev key of
		the imported mesh node.
		Remaining parameters correspond to provisioning data:
		The net_key and net_index parameters describe the network (or a
		subnet, if net_index is not 0) the imported mesh node belongs
		to.
		The flags parameter is a dictionary containing provisioning
		flags. Supported values are:
			boolean IVUpdate
				When true, indicates that the network is in the
				middle of IV Index Update procedure.
			boolean KeyRefresh
				When true, indicates that the specified net key
				is in the middle of a key refresh procedure.
		The iv_index parameter is the current IV Index value used by
		the network. This value is known by the provisioner.
		The unicast parameter is the primary unicast address of the
		imported node.
		The returned token must be preserved by the application in
		order to authenticate itself to the mesh daemon and attach to
		the network as a mesh node by calling Attach() method or
		permanently remove the identity of the mesh node by calling
		Leave() method.
		PossibleErrors:
			org.bluez.mesh.Error.InvalidArguments,
			org.bluez.mesh.Error.AlreadyExists,
			org.bluez.mesh.Error.NotSupported,
			org.bluez.mesh.Error.Failed

*/
func (a *Network1) Import(app_root dbus.ObjectPath, uuid []byte, dev_key []byte, net_key []byte, net_index uint16, flags map[string]interface{}, iv_index uint32, unicast uint16) (uint64, error) {
	
	var val0 uint64
	err := a.client.Call("Import", 0, app_root, uuid, dev_key, net_key, net_index, flags, iv_index, unicast).Store(&val0)
	return val0, err	
}

