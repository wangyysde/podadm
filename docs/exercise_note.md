

## prepareEnv
- wget https://github.com/kubernetes-sigs/kubebuilder/releases/latest/download/kubebuilder_linux_amd64
- mv kubebuilder_linux_amd64 kubebuilder
- chmod +x kubebuilder
- mv kubebuilder /usr/local/bin/
- kubebuilder version

- mkdir podadm
- kubebuilder init --domain sysadm.cn --owner "Wayne Wang<net_use@bzhy.com>" 
- kubebuilder create api --controller --force --resource --group podadm --kind PodAdm --version v1beta1

## REF: 
- https://blog.csdn.net/litaimin/article/details/132754180
- https://blog.csdn.net/CBGCampus/article/details/130614915
- https://blog.csdn.net/CBGCampus/article/details/130614915

