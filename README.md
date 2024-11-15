# TDA596Labs
The labs of the chalmers course TDA596

## How to build

To build the HTTP server simpliy navigate to ```Lab1_HTTP_server``` folder, standing in the repository root:

```sh
cd Lab1_HTTP_server
go build .
```
This will generate a binary called ```httpserver```, which could be run like below:
```sh
./httpserver -port=2222 -v
```
**Note:** the flags are optional. ```-v``` runs the server in verbose mode, providing helpful prints for debugging. ```-port=<portnumber>``` specifies which port the server is run on, default is 1234 if this flag is omitted.

To build the proxy simpliy navigate to ```Proxy``` folder, standing in the repository root:
```sh
cd Lab1_HTTP_server/Proxy
go build .
```

The proxy is run is run in a very similar fashion as the HTTP server and the same flags apply.

## Working requests
The server only implements GET and POST all other requests will illicit a ```510 Not Implemented```.

## Testing the server
```curl``` and ```ab``` (from apache2-utils ```sudo apt install apache2-utils```) can be used to test the server and proxy. ```ab``` good for benchmarking and to make sure that your server doesn't exceed capacity when receiving multiple requests concurrently.

### Using curl
```sh
curl -v -X GET <serverip>:<port>/<file>
```

Testing GET:
```sh
curl -v -X GET localhost:1234/test.txt
```

Testing POST (-v for verbosity, -H for headers, use --help for info about curl):
```sh
curl -v -X POST localhost:1234/this.jpg -H "Content-Type: image/jpg" --data-binary @tmp.jpg
```

Testing the proxy aswell:
```sh
curl -v -X GET localhost:1234/test.txt -x localhost:5555
```

### Using ab
Structure of the command (use ab -h to get info about the command):
```sh
ab [options] [http[s]://]hostname[:port]/path
```

Testing the server by POSTing test.jpg 20 times, 10 at the time:
```sh
ab -v 2 -n 20 -c 10 -p tmp.jpg -T image/jpg localhost:1234/test.jpg
```
Benchmarking the server by GETing test.txt 100 times 20 requests at the time
```sh
ab -v 2 -n 100 -c 20 localhost:1234/test.txt
```