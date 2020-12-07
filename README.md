# NetShare
Share files by http and may be other protocols.


## HOW TO

* With config.  
You can run it using the config.
Config example placed in project dir. To do this, copy the config to your user's directory.
```
cp .netshare.yml ~/  
go run main.go
```

* With only command line.  
Running the program if the config is located in the user's directory.  
```
go run main.go
go run main.go -p 8181
go run main.go -s 127.0.0.1
go run main.go -p 8181 -s 127.0.0.1
go run main.go -p 8182 -s localhost -d data -t web
```

![screenshot of sample](https://github.com/yvv4git/netshare/blob/main/about.png)