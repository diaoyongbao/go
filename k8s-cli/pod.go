package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//获取所有运行的POD
func GetPodList(clientset *kubernetes.Clientset) {
	pods,err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(pods.APIVersion,pods.Kind)
	//fmt.Println(pods.Items)
	for _,pod := range(pods.Items){
		fmt.Println(pod.Name)
		//fmt.Println(pod.Status)
	}
}

//获取单个POD的运行信息
//func GetPod(clientset *kubernetes.Clientset){
//
//}
//创建一个POD
func CreatePod(clientset *kubernetes.Clientset){
	//创建一个pod对象,并完善内容
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-test",
			Namespace: "default",
			},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "nginx",
					Image: "172.18.61.24/base/nginx:latest",
					ImagePullPolicy: "Always",
					Ports: []v1.ContainerPort{
						{ContainerPort: 80},
					},
				},
			},
		},
	}

	result,err := clientset.CoreV1().Pods("default").Create(pod)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Printf("create a pod, podName is %q",result)
}
//删除一个POD
func DeletePod(clientset *kubernetes.Clientset){
	err := clientset.CoreV1().Pods("default").Delete("nginx-test",&metav1.DeleteOptions{})
	if err!=nil{
		panic(err.Error())
	}else {
		fmt.Println("pod nginx-test 已删除" )
	}
}

//apply的实现，将yaml信息装填进annotions中，在下次apply时会与annotions中的信息进行对比，然后进行patch
func ApplyPod(clientset *kubernetes.Clientset){
	//clientset.CoreV1().Pods("default").
}

//更新一个POD
func UpdatePod(clientset *kubernetes.Clientset){

}