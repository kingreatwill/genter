package main

import (
	"fmt"
	"github.com/openjw/genter/x/kubernetes/configmap/controller"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/signal"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("https://192.168.1.135:6443", "C:/Users/35084/Desktop/config")

	//config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err1 := kubernetes.NewForConfig(config)
	if err1 != nil {
		panic(err1)
	}

	cm, e := clientset.CoreV1().ConfigMaps("default").Get("queen-golang-config", metav1.GetOptions{})
	fmt.Println(cm, e)

	listWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "configMaps", "default", fields.Everything())

	_, informer := cache.NewIndexerInformer(listWatcher, &v1.ConfigMap{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println(oldObj, newObj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println(obj)
		},
	}, cache.Indexers{})
	stop := make(chan struct{})
	defer close(stop)
	informer.Run(stop)

	c, err2 := controller.NewController(clientset, "default")
	if err2 != nil {
		panic(err2)
	}
	c.Run()

	q := make(chan os.Signal)
	signal.Notify(q, os.Interrupt)
	<-q
	fmt.Println("退出")
	//informers.NewSharedInformerFactory(clientset, time.Second).Core().V1().ConfigMaps().Lister().ConfigMaps("sdf")

}
