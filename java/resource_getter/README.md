# Resource Getter

no logic, just resource.

## create war file

### command line

```
$ cd /path/to/resource_getter
$ mvn install
$ mvn package
```

### eclipse(neon)

* right click on pom.xml
* Run As > 6 Maven Install
* Run As > 2 Maven Build
  * Select "package" is tailed
  * click ok

## launch

### command line

```
$ cd /path/to/resource_getter
$ java -jar target/resource_getter-{VERSION}.war -p {PORT_NUMBER}
ex. java -jar target/resource_getter-0.0.1-SNAPSHOT.war -p 3000
```

#### default port is 8080 (no need to specify port number)

```
$ java -jar target/resource_getter-0.0.1-SNAPSHOT.war
```

## check

[http://localhost:8080/test_dir/squirrel.png](http://localhost:8080/test_dir/squirrel.png)

[http://localhost:8080/notfound](http://localhost:8080/notfound)

# FAQ

## deploy with managed tomcat, not with emmbedded tomcat

[see wiki](https://github.com/ymatsukawa/resource_getter/wiki#deploy-with-managed-tomcat-not-embedded-one)
