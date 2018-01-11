package res

import (
	"log"
	"time"
)

// ResourceManager is used to manage the different resources on
// a CSP. It can be used to get e.g. all instances for all accounts
// in AWS.
type ResourceManager interface {
	// InstancesPerAccount returns a mapping from account/project
	// to its associated instances
	InstancesPerAccount() map[string][]Instance
	// ImagesPerAccount returns a mapping from account/project
	// to its associated images
	ImagesPerAccount() map[string][]Image
	// VolumesPerAccount returns a mapping from account/project
	// to its associated volumes
	VolumesPerAccount() map[string][]Volume
	// SnapshotsPerAccount returns a mapping from account/project
	// to its associated snaphots
	SnapshotsPerAccount() map[string][]Snapshot
	// AllResourcesPerAccount will return a mapping from account/project
	// to all of the resources associated with that account/project
	AllResourcesPerAccount() map[string]*ResourceCollection
}

// Resource represents a generic resource in any CSP. It should be
// concretizised further.
type Resource interface {
	ID() string
	Tags() map[string]string
	Location() string
	Public() bool
	CreationTime() time.Time
}

// Instance composes the Resource interface, and descibes an instance
// in any CSP.
type Instance interface {
	Resource
	InstanceType() string
}

// Image composes the Resource interface, and descibe an image in
// any CSP. Such as an AMI in AWS.
type Image interface {
	Resource
	Name() string
}

// Volume composes the Resource interface, and describe a volume in
// any CSP.
type Volume interface {
	Resource
	SizeGB() int64
	Attached() bool
	Encrypted() bool
	VolumeType() string
}

// Snapshot composes the Resource interface, and describe a snapshot
// in any CSP.
type Snapshot interface {
	Resource
	Encrypted() bool
	SizeGB() int64
}

// ResourceCollection encapsulates collections of multiple resources
type ResourceCollection struct {
	Instances []Instance
	Images    []Image
	Volumes   []Volume
	Snapshots []Snapshot
}

type csp int

const (
	// AWS is AWS
	AWS csp = iota
	// GCP is Google Cloud Platform
	GCP
)

// NewManager will build a new resource manager for the specified CSP
func NewManager(c csp, accounts ...string) ResourceManager {
	switch c {
	case AWS:
		log.Println("Initializing AWS Resource Manager")
		manager := &awsResourceManager{
			accounts: accounts,
		}
		return manager
	case GCP:
		log.Fatalln("Unfortunately, GCP is currently not supported")
	default:
		log.Fatalln("Invalid CSP specified")
	}
	return nil
}