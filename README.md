<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h2 align="center">Golang - Ibm db2</h2> <br />
</p>


<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
    </li>
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## About The Project
This system is using for testing Golang connection with Ibm db2 database, which ibm db2 database installed on docker.


<!-- BUILD WITH -->
#### Built With

This section we explain what is backend is develop with 
* [Golang](https://golang.org)
* [Connector](https://github.com/ibmdb/go_ibm_db)
* [Db2](https://www.ibm.com/docs/en/db2woc?topic=installing-mac-os-x)
* [Db2 - Docker](https://hub.docker.com/r/ibmcom/db2)


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.


<!-- PREREQUISITES -->
#### Prerequisites

* First you need to install [Golang](https://golang.org/doc/install)

* After that check the installation and Golang version it must be above than 1.11 because we need the [Golang Modules](https://blog.golang.org/using-go-modules)
  ```sh
  > go version
  go version go1.16.3 darwin/amd64
   ```
* Create docker volume
```sh
    docker volume create ibm-data
```

* Install db2 on docker
```sh
    docker run -d --restart=always --name=ibm-exp --privileged=true -p 50000:50000 -e LICENSE=accept -e DB2INST1_PASSWORD=password -e DBNAME=testdb -v ibm-data:/database ibmcom/db2
```

* Check the logs
```sh
    ❯ docker logs ibm-exp | grep "Setup"

            (*) Setup has completed.


        _________________________________________________________________________

                            _____   DB2 Service Tools   _____

                                    I      B      M

                                    db2updv115

        This tool is a service utility designed to update a DB2 Version 11.5
        database to the current fix pack level.

        _________________________________________________________________________

```

* Enter the docker console for the db2 instance
```sh
    ❯ docker exec -it ibm-exp bash -c "su - db2inst1"
    Last login: Sun Aug  8 15:12:47 UTC 2021 on pts/0
    [db2inst1@da30f52444ce ~]$
```

* Check 1 
```sh
    [db2inst1@da30f52444ce ~]$ db2 get dbm cfg|grep SVCENAME
    TCP/IP Service name                          (SVCENAME) = db2c_db2inst1
    SSL service name                         (SSL_SVCENAME) =
```

* Check 2 
```sh
    [db2inst1@da30f52444ce ~]$ cat /etc/services | grep db2c_db2inst1
    db2c_db2inst1      50000/tcp
    db2c_db2inst1_ssl  50001/tcp
    [db2inst1@da30f52444ce ~]$
```

* Check 3
```sh
    [db2inst1@da30f52444ce ~]$ db2 list db directory

    System Database Directory

    Number of entries in the directory = 1

    Database 1 entry:

    Database alias                       = TESTDB
    Database name                        = TESTDB
    Local database directory             = /database/data
    Database release level               = 15.00
    Comment                              =
    Directory entry type                 = Indirect
    Catalog database partition number    = 0
    Alternate server hostname            =
    Alternate server port number         =

    [db2inst1@da30f52444ce ~]$
```


* Check 4 - Connect to the database from docker console
```sh
    [db2inst1@da30f52444ce ~]$ db2 connect to testdb

            Database Connection Information

        Database server        = DB2/LINUXX8664 11.5.6.0
        SQL authorization ID   = DB2INST1
        Local database alias   = TESTDB

        [db2inst1@da30f52444ce ~]$
```

<!-- INSTALLATION -->
#### Installation

* Install Db2 Driver for my Mac OS
```sh
   ❯ go get -d github.com/ibmdb/go_ibm_db
```

* Download the Db2 Driver from [ibm](https://www.ibm.com/docs/en/db2woc?topic=installing-mac-os-x)

* Add the Db2 installation path at the bottom, and make sure the dspath is correct
```sh
    ❯ nano ~/.zshrc
    # ibm db2 driver
    export DB2HOME=/Applications/dsdriver
    export CGO_CFLAGS=-I$DB2HOME/include
    export CGO_LDFLAGS=-L$DB2HOME/lib
    export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:/Applications/dsdriver/lib
    
    ❯ source  ~/.zshrc
``` 

* Done, try run the program 
```sh
    ❯ go run main.go
    NIM    NAMA   
    -------------------------------------
    1  satu   
``` 


<!-- LICENSE -->
## License
`MIT`.


<!-- CONTACT -->
## Contact
Dedy Styawan - [dedy.styawan](https://twitter.com/dedystyawan) - dedy.styawan@gmail.com

Shandy Siswandi - [shandysiswandi](https://github.com/shandysiswandi) - shandysiswandi@gmail.com


<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Installing Db2 on your coffee break](https://ajstorm.medium.com/installing-db2-on-your-coffee-break-5be1d811b052)
* [DB2 Connection refused](https://www.codenong.com/cs105573158/)
* [Installing Db2 on macOS Using Docker](https://www.idug.org/blogs/john-maenpaa1/2020/11/04/installing-db2-on-macos-using-docker)
* [Installing Db2 on your coffee break](https://ajstorm.medium.com/installing-db2-on-your-coffee-break-5be1d811b052)
* [Installing the Db2 driver package on Mac OS X](https://www.ibm.com/docs/en/db2woc?topic=installing-mac-os-x)
