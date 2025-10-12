# ged

minimal encrypt/decrypt tool using AES-256-GCM

```bash
# encrypt
ged enc ./path/to/file.txt -k foobar     # out binary ./path/to/file.txt.enc

# decrypt
ged dec ./path/to/file.txt.enc -k foobar # out ./path/to/file.txt.dec
# decrypt fail
ged dec ./path/to/file.txt.enc -k woof   # auth failed
```

alt tool is openssl. use at your own risk.
