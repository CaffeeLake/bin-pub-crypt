#!/usr/bin/env sh

set -e

go build -buildmode pie

echo "小田原熱海間に、軽便鉄道敷設の工事が始まったのは、良平の八つの年だった。良平は毎日村外れへ、その工事を見物に行った。工事を――といったところが、唯トロッコで土を運搬する――それが面白さに見に行ったのである。" > ./files/test-raw.txt

bin-pub-crypt keygen ./keys/test

bin-pub-crypt encrypt ./keys/test.pub ./files/test-raw.txt ./files/test-enc.txt
bin-pub-crypt decrypt ./keys/test.key ./files/test-enc.txt ./files/test-dec.txt

exit
