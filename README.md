# jeongmoji_go

## :pushpin: aws configuration
```
- ec2: aws linux ubuntu22.04LTS
- port: 443, 80, 20
- Elastic IP
- EBS: 20Gb
- Instance Type: t2.medium
- thi***.pem
```
## :pushpin: aws set root user: ubuntu to root
```bash
$ sudo passwd root
enter your password
$ sudo vim /etc/ssh/sshd_config
PermitRootLogin yes
$ sudo cp /home/ubuntu/.ssh/authorized_keys /root/.ssh
$ sudo systemctl restart sshd
$ ssh -i 'your_local_pc_home/.ssh/your_private_key.pem' root&your_aws_ip
```

## :pushpin: ubuntu configuration
```bash
$ sudo apt update
$ sudo apt upgrade
```

## :pushpin: install go
TODO

## :pushpin: install go packages
web/was server: gin-gonic
```bash
$ go get github.com/gin-gonic/gin
```
configuration : godotenv
```bash
$ go get github.com/joho/godotenv
```
live reload: air
```bash
$ go install github.com/cosmtrek/air@latest
$ air init
$ go mod init
$ go mod tidy
$ vim .air.toml
[log]
  time = true

[misc]
  clean_on_exit = true

[screen]
  clear_on_rebuild = true
```

## :pushpin: install elastic search
TODO


## :cherry_blossom: go web/was server run
```bash
$ go run main.go
```
or
```bash
$ air -c .air.toml
```