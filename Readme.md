Masquerade
==========

Masquerade is library and command line utility.
The code under pkg can be used to tokenize and deal with several formats and fonts.
There is several executables that you can pipe bash style to fill your tokenization requirements.

Architecture
------------
Command line utilities use msgpack as binary format exchange format to avoid expensive json or csv transformation. If you need more performance, please, use it as a library and compile only one binary executable.

Installation
------------
Just make, it will put their binaries into $GOPATH/bin.
```bash
make
```
Ensure that $GOPATH/bin it's in your path:
```bash
export PATH=$(go env GOPATH)/bin:$(go env GOROOT)/bin:$PATH
```

Formats
=======

CSV
---
The simpliest format.

There is two executables: maskcsvin and maskcsvout.

### From CSV -> MsgPack
```bash
echo hello,World | maskcsvin > binary.out
```
Will return binary format of the csv.

### From MsgPack -> CSV
```bash
cat binary.out | maskcsvout
```
Will return "hello","World".

A complete usage may be.
```bash
echo hello,World | maskcsvin | maskcsvout
```
Will return "hello","World". Notice that our process add quotes, thats because our binary don't know how looks the original csv, so try to build the most "correct" one.

### Custom separator
You can use another separator like '|' or '@'. Just provide separator param.
```bash
echo hello@World | maskcsvin -separator '@' | maskcsvout -separator '|'
```
Will return "hello"|"World".

Sources
=======

RabbitMQ
--------
To read from Rabbit use:
```bash
maskrabbitin -dial amqp://guest:guest@localhost:5672/ -channel test
```

This command will consume the queue and output the content into Standart Output.

To write on Rabbit use:
```bash
cat data | maskrabbitout -dial amqp://guest:guest@localhost:5672/ -channel test
```

This comand will send lines from data file into RabbitMQ.

You can copy a queue using this commands together:
```bash
maskrabbitin -dial amqp://guest:guest@localhost:5672/ -channel topicA | maskrabbitout -dial amqp://guest:guest@localhost:5672/ -channel topicB
```

HDFS
----
HDFS has stdio support, just use is as follows.

To read:
```bash
hdfs dfs -cat data.csv
```

To write:
```bash
cat data.csv | hdfs dfs -put - data.csv
```

S3 + GCS (Google Cloud Storage) + Minio
---------------------------------------
This services can be accesed with stdio support thru [minio-cli](https://github.com/minio/mc#add-a-cloud-storage-service).

Roadmap
-------
- TODO: check https://github.com/idealista/format-preserving-encryption-java
