package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisFailover represents a Redis failover
type RedisFailover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFailoverSpec   `json:"spec"`
	Status            RedisFailoverStatus `json:"status,omitempty"`
}

// RedisFailoverSpec represents a Redis failover spec
type RedisFailoverSpec struct {
	// Redis defines its failover settings
	Redis RedisSettings `json:"redis,omitempty"`

	// Sentinel defines its failover settings
	Sentinel SentinelSettings `json:"sentinel,omitempty"`

	// HardAntiAffinity defines if the PodAntiAffinity on the deployments and
	// statefulsets has to be hard (it's soft by default)
	HardAntiAffinity bool `json:"hardAntiAffinity,omitempty"`

	// NodeAffinity defines the rules for scheduling the Redis and Sentinel
	// nodes
	NodeAffinity *corev1.NodeAffinity `json:"nodeAffinity,omitempty"`

	// Define Redis persistence settings and the volume to store the backup data on
	RedisFailoverPersistence RedisFailoverPersistence `json:"storage,omitempty"`
}

// RedisSettings defines the specification of the redis cluster
type RedisSettings struct {
	Replicas  int32                  `json:"replicas,omitempty"`
	Resources RedisFailoverResources `json:"resources,omitempty"`
	Exporter  bool                   `json:"exporter,omitempty"`
	Version   string                 `json:"version,omitempty"`
}

// SentinelSettings defines the specification of the sentinel cluster
type SentinelSettings struct {
	Replicas  int32                  `json:"replicas,omitempty"`
	Resources RedisFailoverResources `json:"resources,omitempty"`
}

// RedisFailoverResources sets the limits and requests for a container
type RedisFailoverResources struct {
	Requests CPUAndMem `json:"requests,omitempty"`
	Limits   CPUAndMem `json:"limits,omitempty"`
}

// CPUAndMem defines how many cpu and ram the container will request/limit
type CPUAndMem struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// RedisFailoverStatus has the status of the cluster
type RedisFailoverStatus struct {
	Phase      Phase       `json:"phase"`
	Conditions []Condition `json:"conditions"`
	Master     string      `json:"master"`
}

// Phase of the RF status
type Phase string

// Condition saves the state information of the redisfailover
type Condition struct {
	Type           ConditionType `json:"type"`
	Reason         string        `json:"reason"`
	TransitionTime string        `json:"transitionTime"`
}

// ConditionType defines the condition that the RF can have
type ConditionType string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisFailoverList represents a Redis failover list
type RedisFailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RedisFailover `json:"items"`
}

// Define the persistence for Redis
type RedisFailoverPersistence struct {
	// Enable persistence and mount a disk?
	Enabled bool `json:"enabled"`

	// A PVC to mount and store the RDB / AOF files on
	// This will be mounted on _every_ data node as running a (fast) restarting master
	// without persistence is dangerous
	PersistentVolumeClaim corev1.PersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`

	// VolumeMountDir  -  will be the mount point for the PVC _AND_ the `dir` setting in redis.conf
	PersistenceDir string `json:"persistencedir,omitempty"`

	// Set RDB (Snapshot) persistence
	Rdb RedisFailoverRdbPersistence `json:"rdb,omitempty"`

	// Set AOF (Append-only File) persistence
	Aof RedisFailoverAofPersistence `json:"aof,omitempty"`
}

type RedisFailoverRdbPersistence struct {
	Enabled                 bool                   `json:"enabled"`
	RdbFilename             string                 `json:"rdbfilename,omitempty"`
	Sizes                   []RedisFailoverRdbSize `json:"sizes,omitempty"`
	StopWritesOnBgSaveError bool                   `json:"stopwritesonbgsaveerror,omitempty"`
	RdbChecksum             bool                   `json:"rdbchecksum,omitempty"`
}

type RedisFailoverRdbSize struct {
	Seconds int32 `json:"seconds"`
	Keys    int32 `json:"keys"`
}

type RedisFailoverAofPersistence struct {
	Enabled           bool                           `json:"enabled"`
	AppendFilename    string                         `json:"appendfilename"`
	AppendFsync       string                         `json:"appendfsync,omitempty"`
	AofRewrite        RedisFailoverAofRewriteSetting `json:"aofrewrite,omitempty"`
	AutoloadTruncated bool                           `json:"autoloadtruncated,omitempty"`
}

type RedisFailoverAofRewriteSetting struct {
	Percentage int32 `json:"percentage"`
	MinSize    int32 `json:"minsize"`
}
