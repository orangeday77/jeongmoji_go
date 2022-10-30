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
```bash
$ curl -fsSL https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo gpg --dearmor -o /usr/share/keyrings/elastic.gpg
$ echo "deb [signed-by=/usr/share/keyrings/elastic.gpg] https://artifacts.elastic.co/packages/7.x/apt stable main" | sudo tee -a /etc/apt/sources.list.d/elastic-7.x.list
$ sudo apt update
$ sudo apt install elasticsearch
```

## :pushpin: configuration elasticsearch
```bash
$ sudo vim /etc/elasticsearch/elasticsearch.yml
cluster.name: your_project_name-cluster
node.name: node-1
network.host: 0.0.0.0
http.port: 9200
cluster.initial_master_nodes: ["node-1"]
```

## :pushpin: install elasticsearch plugin nori
```bash
$ cd /usr/share/elasticsearch
$ bin/elasticsearch-plugin install analysis-nori
```

## :pushpin: elasticsearch CORS
```bash
http.cors.enabled : true
http.cors.allow-origin : "*"
http.cors.allow-methods : OPTIONS, HEAD, GET, POST
http.cors.allow-headers : X-Requested-With, X-Auth-Token, Content-Type, Content-Length, Authorization, Access-Control-Allow-Headers, Accept
```

## :pushpin: registration elasticsearch
```bash
$ systemctl start elasticsearch
$ systemctl enable elasticsearch
$ curl -X GET 'http://localhost:9200' # test
```
## :pushpin: registration elasticsearch python project
```bash
$ mkdir jeongmoji_python
$ cd jeongmoji_python
$ cp your_source jeong_moji.py
$ apt install python3-pip
$ pip3 install tqdm
$ pip3 install elasticsearch
$ touch jeong_moji.json 
```

## :pushpin: elasticsearch 인덱스 재생성 또는 갱신
```bash
$ python3 jeong_moji.py
```

## :pushpin: elasticsearch nori 점검
```
curl -XPOST "your_site/_search"  -H 'Content-Type: application/json' -d'
  {
     "query": {
        "query_string": {
            "query": "99_1_3_안드로이드.pdf"
          }
       }
  }'
```

## :cherry_blossom: go web/was server run
```bash
$ go run main.go
```
or
```bash
$ air -c .air.toml
```

## :cherry_blossom: registration go service
make your_go_binary_file
```bash
$ go build your_go_binary_file
```

```bash
$ cd 
$ vim your_service_name
[Unit]
Description=Golang Gin Gonic webapp Daemon
#After=network.target
StartLimitIntervalSec=0
[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/your_project_directory/
ExecStart=/your_project_directory/your_go_binary_file
TimeoutStopSec=300
[Install]
WantedBy=multi-user.target
```

registration
```bash
$ systemctl start your_service_name
$ systemctl enable your_service_name
```

restart
```bash
$ systemctl list-units

# elasticsearch
$ systemctl status elasticsearch.service
$ systemctl restart elasticsearch

# gonic web searver
$ systemctl status jeongmoji.service
$ systemctl restart jeongmoji.service
```