package main

import (
	"fmt"
	"github.com/openjw/genter/x/kubernetes/configmap/controller"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/signal"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("https://192.168.1.135:6443", "C:/Users/35084/Desktop/config")
	{
		cfg := &rest.Config{
			Host:        "https://192.168.1.135:6443",
			BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6InVjeWJsSDBPRTFpMUM1V2sta1YzOE5xc1M4cHJmVzByblRjMkQ3ck1sa1UifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkYXNoYm9hcmQtYWRtaW4tdG9rZW4tODZuOG4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGFzaGJvYXJkLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiN2E5NjYwMjktOGFiZC00YjFkLTkwNTItMmQyYmU3YWY5N2Y0Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmVybmV0ZXMtZGFzaGJvYXJkOmRhc2hib2FyZC1hZG1pbiJ9.bnjv5h86T6-GhiFP4tHZB4JCvfMU6uQIUnBJ5uNiPk5ZY0m7eWq7wHcvG7qJBcQVGg8y1I7WceKxU3XyliOb80BcLEWC_RAgfabZ-JcR3j16vvhyO0ju45Amg2AV2zqD75o8q8B4oVull6CjhajAFL7eV8BPmcw-dqPPW6NiiMjtt2UhoOs8ZIyBA9-e_d4eSurdK_qQ0vEnebr6AUv5SB0GJordkuFrqC_bdLCgSfgrYPK6VxwEy3adlK7OQ_Ke2tNxfgHW9MJFFYKy6CBu3k_hLtsghellu211nLV3uaaFuYjAv-gIsNbXIWy8vRy3rYsZWrFRn088C0bZli5yzg",
			TLSClientConfig: rest.TLSClientConfig{
				Insecure: true,
			},
		}
		fmt.Println(cfg)
		clientset, err1 := kubernetes.NewForConfig(cfg)
		if err1 != nil {
			panic(err1)
		}
		names, _ := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		fmt.Println(names)
	}
	//config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err1 := kubernetes.NewForConfig(config)
	if err1 != nil {
		panic(err1)
	}

	names, _ := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	fmt.Println(names)
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
