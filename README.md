# certmonitor

A simple website certificate monitor tool

## How to use?

### For general user

You can download the pre-compiled binaries for the corresponding platform from [release page](https://github.com/mritd/certmonitor/releases).
Next create a configuration file named `certmonitor.yaml`, like this:

``` yaml
alarm:
- type: smtp
  targets:
  - mritd1234@gmail.com
- type: webhook
  targets:
  - https://google.com
- type: telegram
  targets:
  - "-124568340456"
monitor:
  websites:
  - name: bleem
    description: 博客主站点
    address: https://mritd.com
  - name: baidu
    description: 百度首页
    address: https://baidu.com
  cron: '@every 1h'
  beforetime: 168h0m0s
  timeout: 10s
smtp:
  username: mritd
  password: password
  from: mritd@mritd.me
  server: smtp.qq.com:465
telegram:
  api: https://api.telegram.org
  token: token_example
webhook:
  method: get
  timeout: 5s
```

Finally run it(Suppose the file you downloaded is named `certmonitor_linux_amd64`):

``` sh
chmod +x certmonitor_linux_amd64
./certmonitor_linux_amd64
```

### For docker user(Advanced)

build docker image

``` sh
make docker
```

create a config named `certmonitor.yaml`

``` yaml
alarm:
- type: smtp
  targets:
  - mritd1234@gmail.com
- type: webhook
  targets:
  - https://google.com
- type: telegram
  targets:
  - "-124568340456"
monitor:
  websites:
  - name: bleem
    description: 博客主站点
    address: https://mritd.com
  - name: baidu
    description: 百度首页
    address: https://baidu.com
  cron: '@every 1h'
  beforetime: 168h0m0s
  timeout: 10s
smtp:
  username: mritd
  password: password
  from: mritd@mritd.me
  server: smtp.qq.com:465
telegram:
  api: https://api.telegram.org
  token: token_example
webhook:
  method: get
  timeout: 5s
```

run a container

``` sh
docker run -dt --name cermonitor -v ./certmonitor.yaml:/certmonitor.yaml mritd/certmonitor:CURRENT_VERSION
```
