# speakbuddy
speakbuddy is app for upload and retrieve audio file. it only accept mp3 but will store file in the server as a wav file.

## business flow

## architechture diagram
![Untitled drawing](https://github.com/user-attachments/assets/19c92f76-5f04-4eb9-836f-989ed5e3436a)
- client (web app or app) connect to speakbuddybeapi via nginx ingress
- speakbuddybeapi store text data in mysql
- speakbuddybeapi store audio file in sftp

## entity relation diagram (db)
![image](https://github.com/user-attachments/assets/49ec4d9d-62c3-4cb0-801d-a0ccc8eaaac2)
- AudioFile: store data relation between user, phrase and audio file
- Phrase: store data phrase
- User: store data user

## sequence diagram
### upload audio file
![image](https://github.com/user-attachments/assets/d8621752-3603-4957-a6bd-37cc603f1cf1)
1. user open client (web app or app)
2. user fill input user, phrase and choose audio file in the form
3. user submit form
4. user get response (success or error)

> As a user,
> I want to upload an audio file along with my user ID and phrase ID,
> So that the system can associate the audio with my profile and a specific phrase.

### retreive audio file
![image](https://github.com/user-attachments/assets/04a3c5cb-099e-4069-84c2-985c5a43b654)
1. user open client (web app or app)
2. user fill input user and phrase in the form
3. user submit form
4. user get file as a response (success or error)

> As a user,
> I want to retrieve an audio file associated with my user ID and phrase ID,
> So that I can access and use the stored audio file in MP3 format.

## env
- staging (non-live)
- production

## preparation

### 1. docker setup
- install docker or download docker desktop https://www.docker.com/products/docker-desktop/
- install kubernetes or enable kubernetes in docker desktop https://docs.docker.com/desktop/features/kubernetes/

### 2. mysql setup
1. setup mysql as a db (if have not setup db yet)
```
> make ENV=<env> mysql-setup
```
you can set or check db base64 encoded password by check `data.rootpassword` this file `manifests/mysql/<env>/secret-mysql.yaml`
```
data:
  rootpassword: cGFzc3dvcmQ=
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

4. seed data (if needed)
```
> use evergreen_speakbuddybe_db;
> INSERT INTO phrases (created_at, updated_at, phrase) VALUES (now(), now(), "test");
> INSERT INTO users (created_at, updated_at, name) VALUES (now(), now(), "usertest");
```

### 3. sftp setup
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
```
usename: sftp_user
password: sftp_user
```

6. visit http://localhost:8080/web/client (logout from prev session or new incognito)

7. login with user account that already created before

8. create folder speakbuddy_storage

9. and you can watch if there is new file uploaded


## 4. speakbuddybeapi setup
1. change docker username with your own docker account in `Makefile`
```
DOCKERHUB_USERNAME = <dockerhub_username>
```

2. setup docker, build and publish the image
```
> make ENV=<env> VERSION=<version> docker-speakbuddybeapi-setup
```

3. setup speakbuddybeapi
```
> make ENV=<env> VERSION=<version> speakbuddybeapi-setup
```

4. serve the speakbuddybeapi
```
> make serve-speakbuddybeapi
```

5. access the speakbuddybeapi with this host http://localhost:8081

### 5. nginx ingress setup
1. install nginx ingress controller, if you don’t have it already, install the  nginx ingress controller using Helm
```
> helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
> helm repo update
> helm install nginx-ingress ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace
```

2. apply the ingress configuration
```
> kubectl apply -f ingress.yaml
```

3. if you're testing locally, update /etc/hosts to point speakbuddy.local to your cluster's IP
```
> echo "127.0.0.1 speakbuddy.local" | sudo tee -a /etc/hosts
```

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

## test result
device

<img width="207" alt="Screenshot 2025-02-02 at 16 01 41" src="https://github.com/user-attachments/assets/fcec6e3f-2c84-46e8-a7bf-13d7ac679358" />

result

<img width="998" alt="Screenshot 2025-02-02 at 16 00 35" src="https://github.com/user-attachments/assets/fe0863c5-dd75-46e7-ba95-f4419b7b3e86" />
<img width="992" alt="Screenshot 2025-02-02 at 16 02 45" src="https://github.com/user-attachments/assets/b07e1c8c-84d9-4fd2-9bc0-f45447033861" />


