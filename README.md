> typora上传图片到七牛云

### 配置

#### 配置文件路径
```
默认路径
/usr/local/etc/qiniu.toml
```

#### 配置文件内容
```
[qiniu]
access_key = ""
secret_key = ""
bucket = ""
bucket_dir = ""
cdn_url = ""

[image]
compress_switch = 1 # 是否压缩
max_kb = 100 # 超过多少kb才压缩
width = 400     # 图片缩放宽度
quality = 80    # jpg图片质量 1-100

[logger]
log_file = "/usr/local/var/log/typora.log"  # 日志路径
```