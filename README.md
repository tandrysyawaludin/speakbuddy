# speakbuddybe-api-k8s

## env
- staging (non-live)
- production

## preparation

### docker setup
- install docker or download docker desktop https://www.docker.com/products/docker-desktop/
- install kubernetes or enable kubernetes in docker desktop https://docs.docker.com/desktop/features/kubernetes/

### mysql setup
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

### sftp setup
1. set docker network for sftpgo
```
> make SFTP_IP_RANGE=<sftp_ip_range> set-sftp-network
```

2. setup sftpgo
```
> make ENV=<env> SFTP_TARGET_IP=<sftp_target_ip> serve-sftp
```

3. visit http://localhost:8080/web/admin

4. create admin account

5. create user account

6. visit http://localhost:8080/web/client (logout from prev session or new incognito)

7. login with user account that already created before

8. create folder speakbuddy_storage

9. and you can watch if there is new file uploaded


## run app

1. setup docker, build and publish the image
```
> make ENV=<env> VERSION=<version> docker-speakbuddybeapi-setup
```

2. setup speakbuddybeapi
```
> make ENV=<env> VERSION=<version> speakbuddybeapi-setup
```

3. serve the speakbuddybeapi
```
> make serve-speakbuddybeapi
```

4. access the speakbuddybeapi with this host http://localhost:8081

## reset everything
- reset all
```
> make clean-all
```

- reset database only
```
> make clean-db-ns
```

- reset app only
```
> make clean-app-ns
```

- reset docker only
```
> make clean-docker
```

## other (if needed)
- ssh-keygen -R "[localhost]:2022"
- kubectl get storageclasses.storage.k8s.io
- kubectl get endpoints
- kubectl get namespaces
- kubectl get <keyword>
- kubectl get pods -n <namespace>
- kubectl describe pod <pod-name> -n <namespace>
- kubectl logs <pod-name> -n <namespace>
- kubectl delete ns <namespace>