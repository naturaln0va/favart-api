## FavArt API
> A backend for a FileProviderExtension iOS app.
> This was created to power a sample project app for [RayWenderlich](https://www.raywenderlich.com).

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/naturaln0va/favart-api/tree/master)

### Run locally

In a shell in the root of this project, simply run `go run main.go`

### Endpoints

- [x] **GET** `/media?path={directory}` - Returns all the assets in a given directory
- [x] **POST** `/media` - Create directories to resolve a path
- [ ] **PUT** `/media` - Rename a directory at a given path
- [ ] **DELETE** `/media` - Delete a directory to resolve a path
- [x] **GET** `/file?path={path}&id={file_id}` - Returns the asset for the given id
- [x] **POST** `/file` - Upload a new asset to a directory
- [ ] **PUT** `/file?path={path}&id={file}` - Updates an asset at a path by a given id
- [ ] **DELETE** `/file?path={path}&id={file}` - Remove an asset at a path by a given id
