package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemcachedSpec defines the desired state of Memcached
type MemcachedSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Size is the size of the memcached deployment, SIZE , SIZE, SIZE!
	Size int32 `json:"size"`
}

// MemcachedStatus defines the observed state of Memcached
type MemcachedStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

// MemcachedDnsTTL defines the ...
type MemcachedDnsTTL struct {
	// How many seconds a resolver is supposed to cache DNS records
	Ttl int32 `json:"ttl"`
}

// MemcachedSplitBrainThreshold defines the split brain threshold
type MemcachedSplitBrainThreshold struct {
	// Defines split brain threshold in minutes
	Threshold int32 `json:"threshold"`
}

// DNS interval
type DnsInterval struct {
	// External-dns sync interval to update etcd backend of coredns
	Interval int32 `json:"interval"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Memcached is the Schema for the memcacheds API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=memcacheds,scope=Namespaced
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec        `json:"spec,omitempty"`
	Status MemcachedStatus      `json:"status,omitempty"`

	SyncInterval DnsInterval	`json:"syncInterval,omitempty"`
	DnsTTL MemcachedDnsTTL	`json:"ttl,omitempty"`
	SplitBrainThreshold MemcachedSplitBrainThreshold	`json:"txtExpiration"`

}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}
