# Mongobeat

Mongobeat is a lightweight agent that formats and ships infomation about a MongoDB instance.

This project is currently in an early development phase, and is not intended for production use!


## Data Shipped By Mongobeat

| Field | Mongo CLI Command |
| -- | -- |
| `dbStats` | db.stats() |

### dbStats

Mongobeat iterates through the list of running databases on the instance and calls db.stats().

The results are appended to the field `dbStats` as part of the Mongobeat event.

Sample `dbStats` field:

```
...

"dbStats": [
{
  "AvgObjSize": 0,
  "Collections": 4,
  "DataSize": 1168,
  "Db": "admin",
  "FileSize": 67108864,
  "IndexSize": 24528,
  "Indexes": 3,
  "NumExtents": 4,
  "Objects": 11,
  "Ok": 1,
  "StorageSize": 28672
},
{
  "AvgObjSize": 0,
  "Collections": 3,
  "DataSize": 36560,
  "Db": "local",
  "FileSize": 67108864,
  "IndexSize": 8176,
  "Indexes": 1,
  "NumExtents": 3,
  "Objects": 34,
  "Ok": 1,
  "StorageSize": 10498048
}
],

...
```

Ensure that this folder is at the following location:
`${GOPATH}/github.com/scottcrespo`

## Getting Started with Mongobeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Mongobeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Mongobeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/scottcrespo/mongobeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Mongobeat run the command below. This will generate a binary
in the same directory with the name mongobeat.

```
make
```


### Run

To run Mongobeat with debugging output enabled, run:

```
./mongobeat -c mongobeat.yml -e -d "*"
```


### Test

To test Mongobeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/mongobeat.template.json and etc/mongobeat.asciidoc

```
make update
```


### Cleanup

To clean  Mongobeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Mongobeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/scottcrespo
cd ${GOPATH}/github.com/scottcrespo
git clone https://github.com/scottcrespo/mongobeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
