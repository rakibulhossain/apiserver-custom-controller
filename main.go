package main

import (
	"flag"
	"time"
	clientset "github.com/Rakibul-Hossain/apiserver-custom-controller/pkg/generated/clientset/versioned"
	informers "github.com/Rakibul-Hossain/apiserver-custom-controller/pkg/generated/informers/externalversions"
	"k8s.io/klog/v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	kubeinformers "k8s.io/client-go/informers"
)

var (
	masterURL string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	stopCh := make(chan struct{})

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL,kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s",err.Error())
	}

	kubeclient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s",err.Error())
	}

	customclient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building customclient: %s",err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeclient, time.Second*30)
	customInformerFactory := informers.NewSharedInformerFactory(customclient, time.Second*30)

	controller := NewController(kubeclient,customclient,
		kubeInformerFactory.Apps().V1().Deployments(),
		customInformerFactory.Customcontroller().V1().Customs())

	kubeInformerFactory.Start(stopCh)
	customInformerFactory.Start(stopCh)

	if err = controller.Run(2,stopCh); err != nil {
		klog.Fatalf("Error running controller: %s",err.Error())
	}

}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}