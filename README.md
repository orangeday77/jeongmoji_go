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
```bash
$ apt remove --autoremove golang-go
$ apt install golang-go
```
```bash
$HOME/.profile or /etc/profile
export PATH=$PATH:/usr/local/go/bin
$ go version
```

## :pushpin: set go PATH
```bash
$ go env
$ vim ~/.profile
GOROOT=/usr/lib/go-1.18
GOPATH=/root/go/bin
PATH=$PATH:$GOPATH
$ source ~/.profile
```

## :pushpin: configuration github
```bash
$ cd ~/.ssh
$ cp your_private_key_for_github /root/.ssh/
$ touch config
Host github.com
    User your_name
    Hostname github.com
    PreferredAuthentications publickey
    IdentityFile /root/.ssh/id_ed25519_your_private_key
$ ssh -T git@github.com
Hi your_name! You've successfully authenticated, but GitHub does not provide shell access.
```

## :pushpin: clone github project
```bash
$ git clone your_project.git
```

## :pushpin: set go module
```bash
$ go mod init github.com/your_name/your_project # 새로운 모듈 생성, go.mod 파일을 초기화
$ go mod tidy # 사용하지 않는 의존성 삭제
$ go test
$ go list -m all # 현재 모듈의 의존성을 모두 출력
```

## :pushpin: install go air
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