# speakbuddybe-api-k8s

## preparation
- install docker or download docker desktop
- install kubernetes or enable kubernetes in docker desktop

## env
- staging (non-live)
- production

## run
1. make ENV=xxx VERSION=xxx docker-setup
2. make ENV=xxx VERSION=xxx db-setup (if have not setup db yet)
3. make ENV=xxx VERSION=xxx db-open (if have not setup db yet)
4. make ENV=xxx VERSION=xxx speakbuddybeapi-setup
5. make serve

# step
## docker build
- docker build -t speakbuddybeapik8s:0.1.0 .

## clean kubectl
- kubectl delete namespace db-ns
- kubectl delete namespace speakbuddybe-ns

## mysql
- kubectl create ns db-ns --dry-run -oyaml > manifests/mysql/ns.yaml
- kubectl create -f manifests/mysql/ns.yaml
- kubectl create -f manifests/mysql/ns-mysql.yaml
- kubectl create secret generic mysql-password -n db-ns
- kubectl create -f manifests/mysql/pvc-mysql.yaml
- kubectl create -f manifests/mysql/secret-mysql.yaml
- kubectl create -f manifests/mysql/service-mysql.yaml
- kubectl create -f manifests/mysql/deployment-mysql.yaml

## speakbuddybe
- kubectl create ns speakbuddybe-ns --dry-run -oyaml > manifests/speakbuddybeapi/ns.yaml
- kubectl apply -f manifests/speakbuddybeapi/ns.yaml
- kubectl apply -f manifests/speakbuddybeapi/ns-speakbuddybeapi.yaml
- kubectl apply -f manifests/speakbuddybeapi/secret-speakbuddybeapi.yaml 
- kubectl apply -f manifests/speakbuddybeapi/service-speakbuddybeapi.yaml 
- kubectl apply -f manifests/speakbuddybeapi/deployment-speakbuddybeapi.yaml 
- kubectl create configmap speakbuddybe-cm -n speakbuddybe-ns --from-literal serverport=8081 --from-literal dbuser=root --from-literal dbname=evergreen_speakbuddybe_db --dry-run -oyaml > manifests/speakbuddybeapi/configmap-speakbuddybeapi.yaml

## seeding
- kubectl exec -it -n db-ns mysql-6b69f446f7-jqr7f -- bash
- mysql -u root -p
- green

## serve
- kubectl port-forward -n speakbuddybe-ns svc/speakbuddybeapi 8081

## other
- kubectl delete configmap speakbuddybe-cm -n speakbuddybe-ns
- kubectl logs mysql-6b69f446f7-jqr7f -n db-ns
- kubectl get deployments -n db-ns
- kubectl get namespaces
- kubectl get pvc -n db-ns
- kubectl get secret mysql-password -n db-ns
- kubectl get pv
- kubectl get pods -n db-ns
- kubectl get svckubectl get endpoints
- kubectl get endpoints
- kubectl get ns | grep db-ns
- kubectl get storageclasses.storage.k8s.io