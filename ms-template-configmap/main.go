package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	v13 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	configmap     *v12.ConfigMap
	configMapName = "ms-template"
	clientset     *kubernetes.Clientset
)

func main() {
	initialize()
	r := gin.Default()
	r.GET("/health", healthCheck)
	r.GET("/configmap", printConfigmap)
	r.Run("0.0.0.0:8080")
	log.SetOutput(os.Stdout)
}

func initialize() {
	// initialize client
	config, _ := rest.InClusterConfig()
	var configErr error
	clientset, configErr = kubernetes.NewForConfig(config)
	if configErr != nil {
		panic(configErr.Error())
	}

	configmaps := clientset.CoreV1().ConfigMaps("default")
	initializeConfigmap(&configmaps)
	//Let this run in the background
	go startConfigMapWatch(&configmaps)

}

func initializeConfigmap(configmaps *v13.ConfigMapInterface) {
	var err error
	cfgs := *configmaps
	configmap, err = cfgs.Get(context.TODO(), configMapName, v1.GetOptions{})
	if err != nil {
		fmt.Printf("Unable to retreive configmap %v, with errror:%v", configMapName, err)
		panic(err.Error())
	}
	fmt.Printf("\nCreated/Refreshed Configmap Object Value from k8s configmap %v", configMapName)
}

func startConfigMapWatch(configmaps *v13.ConfigMapInterface) {
	cfgs := *configmaps
	watcher, err := cfgs.Watch(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("Unable to Create a watcher on configmap %v, with errror:%v", configMapName, err)
		panic(err.Error())
	}

	for event := range watcher.ResultChan() {
		cfg := event.Object.(*v12.ConfigMap)
		switch event.Type {
		case watch.Modified:
			if strings.Contains(cfg.ObjectMeta.Name, configMapName) {
				fmt.Printf("\nConfigmap %s/%s modified", cfg.ObjectMeta.Namespace, cfg.ObjectMeta.Name)
				initializeConfigmap(configmaps)
			}
		}
	}
}

func printConfigmap(c *gin.Context) {
	c.JSON(http.StatusOK, configmap.Data)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
