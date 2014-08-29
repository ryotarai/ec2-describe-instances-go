ec2-describe-instances-go
=========================

```
$ ec2-describe-instances-go -r ap-northeast-1 | jq .
```

Installation
------------
### Download built binary

https://github.com/ryotarai/ec2-describe-instances-go/releases

### or go get

```
$ go get github.com/ryotarai/ec2-describe-instances-go
```

v.s. aws-cli
------------

```
$ for i in `seq 0 3`; do time aws ec2 describe-instances >/dev/null; done
aws ec2 describe-instances > /dev/null  2.15s user 0.11s system 21% cpu 10.548 total
aws ec2 describe-instances > /dev/null  2.17s user 0.11s system 25% cpu 9.009 total
aws ec2 describe-instances > /dev/null  2.14s user 0.11s system 20% cpu 10.771 total
aws ec2 describe-instances > /dev/null  2.15s user 0.11s system 25% cpu 8.899 total

$ for i in `seq 0 4`; do time ec2-describe-instances-go -r ap-northeast-1 >/dev/null; done
ec2-describe-instances-go -r ap-northeast-1 > /dev/null  0.54s user 0.02s system 23% cpu 2.405 total
ec2-describe-instances-go -r ap-northeast-1 > /dev/null  0.55s user 0.02s system 32% cpu 1.707 total
ec2-describe-instances-go -r ap-northeast-1 > /dev/null  0.55s user 0.02s system 31% cpu 1.766 total
ec2-describe-instances-go -r ap-northeast-1 > /dev/null  0.55s user 0.02s system 31% cpu 1.769 total
ec2-describe-instances-go -r ap-northeast-1 > /dev/null  0.55s user 0.02s system 30% cpu 1.860 total
```


