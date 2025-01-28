# speakbuddybe-api-k8s

## preparation
- install docker or download docker desktop
- install kubernetes or enable kubernetes in docker desktop

## env
- staging (non-live)
- production

## run
1. setup mysql as a db (if have not setup db yet)
```
> make ENV=<env> mysql-setup
```

2. open connection to mysql and create db (if have not setup db yet)
```
> make mysql-open
```

if meet this error below please wait a few minutes and try again

```
> error: Internal error occurred: unable to upgrade connection: container not found ("mysql")
```

3. create new db (if have not setup db yet)

```
> mysql -u root -p
> create database evergreen_speakbuddybe_db;
```

4. setup docker, build and publish the image
```
> make ENV=<env> VERSION=<version> docker-setup
```

5. setup speakbuddybeapi
```
> make ENV=<env> VERSION=<version> speakbuddybeapi-setup
```

6. serve the application
```
> make serve
```

7. access the application with this host http://localhost:8081

## other
- kubectl get storageclasses.storage.k8s.io
- kubectl get endpoints
- kubectl get namespaces
- kubectl get <keyword>
- kubectl get pods -n <namespace>
- kubectl describe pod <pod-name> -n <namespace>
- kubectl logs <pod-name> -n <namespace>