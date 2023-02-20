# bin-pub-crypt

- バイナリーファイルを公開鍵暗号方式でやり取りするためのツール。

## How to use

Read [example.sh](./example.sh)

```shell
go build -buildmode pie
bin-pub-crypt keygen [keyfilename]
bin-pub-crypt encrypt [publickeyfile] [inputfile] [outputfile]
bin-pub-crypt decrypt [privatekeyfile] [inputfile] [outputfile]
```
