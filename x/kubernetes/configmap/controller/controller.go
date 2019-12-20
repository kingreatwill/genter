package controller

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	client    kubernetes.Interface
	indexer   cache.Indexer
	queue     workqueue.RateLimitingInterface
	informer  cache.Controller
	namespace string
}

func NewController(client kubernetes.Interface, namespace string) (*Controller, error) {
	c := Controller{
		client:    client,
		namespace: namespace,
	}
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	listWatcher := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "configMaps", namespace, fields.Everything())

	indexer, informer := cache.NewIndexerInformer(listWatcher, &v1.ConfigMap{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc:    c.Add,
		UpdateFunc: c.Update,
		DeleteFunc: c.Delete,
	}, cache.Indexers{})

	c.indexer = indexer
	c.informer = informer
	c.queue = queue
	return &c, nil
}

// Add function to add a new object to the queue in case of creating a resource
func (c *Controller) Add(obj interface{}) {
	switch object := obj.(type) {
	case *v1.ConfigMap:
		fmt.Println(object)
	case *v1.Secret:
		fmt.Println(object)
	default:
		fmt.Println(object)
	}

}

// Update function to add an old object and a new object to the queue in case of updating a resource
func (c *Controller) Update(old interface{}, new interface{}) {
	fmt.Println(old)
	fmt.Println(new)
}

// Delete function to add an object to the queue in case of deleting a resource
func (c *Controller) Delete(old interface{}) {
	// Todo: Any future delete event can be handled here
}

//Run function for controller which handles the queue
func (c *Controller) Run() {

}
