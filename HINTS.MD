# TO DOWNLOAD CARGO UNDER PROXY

- on .cargo/config:

```
[http]
proxy = "http://proxy.example.com:8000/"

[https]
proxy = "https://proxy.example.com:8000/"
```

- and on command line:

```
export HTTP_PROXY=http://proxy.example.com:8000/
export HTTPS_PROXY=https://proxy.example.com:8000/
export FTP_PROXY=ftp://proxy.example.com:8000/
```
