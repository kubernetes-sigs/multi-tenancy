# Block cluster resources

**Profile Applicability:** 

Level 1

**Type:**

Behavioral Check

**Category:**

Control Plane Isolation

**Description:** 

Tenants should not be able to view, edit, create, or delete cluster (non-namespaced) resources such Node, ClusterRole, ClusterRoleBinding, etc. 

**Rationale:** 

Access controls should be configured for tenants so that a tenant cannot list, create, modify or delete cluster resources

**Audit:**

Run the following commands to retrieve the list of non namespaced resources:

  	kubectl --kubeconfig cluster-admin api-resources --namespaced=false

For all non namespaced resources,  issue the following command
	
	kubectl --kubeconfig tenant-a get <resource>

Each command must return 403 FORBIDDEN