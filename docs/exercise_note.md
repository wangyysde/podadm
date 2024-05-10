

## prepareEnv
- wget https://github.com/kubernetes-sigs/kubebuilder/releases/latest/download/kubebuilder_linux_amd64
- mv kubebuilder_linux_amd64 kubebuilder
- chmod +x kubebuilder
- mv kubebuilder /usr/local/bin/
- kubebuilder version

- mkdir podadm
- kubebuilder init --domain sysadm.cn --owner "Wayne Wang<net_use@bzhy.com>" 
- kubebuilder create api --controller --force --resource --group podadm --kind PodAdm --version v1beta1
- 修改api\v1beta1\podadm_types.go文件中的PodAdmSpec Struct,添加类似如下字段：
```
// PodAdmSpec defines the desired state of PodAdm
type PodAdmSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// NameSpace is the namespace of the configMap or Secret
	NameSpace string `json:"nameSpace,omitempty"`

	// ObjectName name of the configMap or Secret
	ObjectName string `json:"objectName,omitempty"`

	// Kind of configMap or Secret
	Kind string `json:"kind,omitempty"`
}

// PodAdmStatus defines the observed state of PodAdm
type PodAdmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// RefObjects list the objects which using the configMap or Secret
	RefObjects []RefObject `json:"refObjects"`
}

```
- make manifests
- make install
- scp config/crd/bases/podadm.sysadm.cn_podadms.yaml root@9.9.33.174:/root/
```上述文件为CRD定义的Yaml文件
    
```
- scp config/samples/podadm_v1beta1_podadm.yaml root@9.9.33.174:/root/
```
  上述文件为创建podadm 实例的Yaml文件例子
```
- kubectl apply -f config/samples/podadm_v1beta1_podadm.yaml




## REF: 
- https://blog.csdn.net/litaimin/article/details/132754180
- https://blog.csdn.net/CBGCampus/article/details/130614915
- https://blog.csdn.net/CBGCampus/article/details/130614915

