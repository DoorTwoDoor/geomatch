# Geomatch

### About
Geomatch is a Go application which provides API endpoints to retrieve a list of movers that are near a customer, update an available mover's current location and store location updates for active moves.

### Getting Started
1. Install [Go](https://www.golang.org/dl/) for your platform.
2. Install [Redis](https://www.redis.io/topics/quickstart/) for your platform.
3. Install [Google Cloud SDK](https://cloud.google.com/sdk/downloads/) for your platform.
4. Install the App Engine extension for Go.
```Bash
gcloud components install app-engine-go
```
5. Install the original [App Engine SDK for Go](https://cloud.google.com/appengine/docs/standard/go/download/) for your platform.
6. [Clone](https://help.github.com/articles/cloning-a-repository/) a copy of the repository into the current directory.
```Bash
git clone https://www.github.com/doortwodoor/geomatch
```
7. Navigate to the newly clone reponsitory.
```Bash
cd geomatch
```
8. Install all dependencies.
```Bash
make deps
```
9. Start the local development server.
```Bash
make serve
```

### License
Geomatch is licensed under the [Apache-2 License](https://www.github.com/DoorTwoDoor/geomatch/blob/master/LICENSE).
