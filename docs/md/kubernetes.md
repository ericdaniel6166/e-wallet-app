##### Kubernetes
- Kubernetes
  - An open source container orchestration engine
  - For automating deployment, scaling, and management of containerized applications
- Kubernetes' components
  - Worker Node (a set of worker machines (aka nodes) that run containerized applications )
    - Kubelet agent (help all containers run properly inside Kubernetes pods)
    - Container runtimes: such as Docker, containerd, CRI-O
    - Kube-proxy: maintain network rules, allow communication with pods
  - Control plane (control plane runs on master nodes), manage the worker nodes and pods of the clusters
    - API server (the front end of the control plane), exposes the Kubernetes API 
    to interact with all the other components of the cluster
    - Persistence store etcd, acts as Kubernetes' backing store for all cluster data
    - Scheduler, watches for newly created Pods with no assigned nodes, 
    and select the nodes for them to run on
    - Controller manager
      - Node controller, is responsible for noticing and responding when nodes go down  
      - Job controller, watches for jobs, or one-off tasks, then creates Pods to run them
      - Endpoint controller, populates the endpoints object, or joins services and pods
      - Service account & token controller, create default account and API access tokens for new namespaces
    - Cloud controller manager, links cluster into cloud provider's API,
    separate out the components that interact with the cloud platform 
    with the components that only interact with cluster 
      - Node controller, checks cloud provider to determine if a non responding node has been deleted or not
      - Route controller, sets up routes in the underlying cloud infrastructure
      - Service controller, create, updates, and deletes cloud load balancers
- kubectl 
  - The Kubernetes command line tool
  - Use to deploy applications, inspect and manage cluster resources, and view logs
- K9s
  - a terminal based UI to interact with Kubernetes cluster