[WIP]

There are many tools being developed around Kubernetes. Many times they communicate with Kubernetes API server.
While building such tools in golang, you need to write the code for unit tests.
We use mocking for external systems while writing unit tests. 
This article covers how one can write the unit tests for the code which interacts with K8s API server.
In golang, many libraries provide testing library as well which can help in writing unit-tests. 
https://golang.org/pkg/net/http/httptest/ is one such example.
Kubernete API client[k8s.io/client-go/kubernetes] also provides one such utility k8s.io/client-go/kubernetes/fake

Let's take one example:
Scenario: Listing all the pods in namespace 'default'.

pods.go
```
import(
  "k8s.io/client-go/kubernetes"
  coreV1 "k8s.io/api/core/v1"
  metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var k8sClient kubernetes.Interface

func listPods(namespace string) []string {
	//assuming k8sClient has been created before calling this method
	podList, err := k8sClient.CoreV1().Pods(namespace).List(metaV1.ListOptions{})
	if err!=nil{
		// log(err)
	}
  
	var res []string
	for _, k8sPod := range podList.Items {
		res = append(res, k8sPod.Name)
	}
	return res
}
```

pods_test.go

```

// helper function
func newPod(ns, name string) *coreV1.Pod {
	return &coreV1.Pod{
		ObjectMeta: metaV1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}
}

func TestListPods(t *testing.T) {
	k8sClient := fake.NewSimpleClientset(
		newPod("default", "pod-1"),
		newPod("default", "pod-2"),
		newPod("myNS", "mypod"),
	)
	
  expected := []string{
		"pod-1",
		"pod-2",
	}
  
	pods := listPods("default")
	assert.EqualValues(t, expected, pods)
  
	expected = []string{"mypod"}
  
	pods = listPods("myNS")
	assert.EqualValues(t, expected, pods)
}
```
