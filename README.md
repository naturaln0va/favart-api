## FavArt API
> A backend for a FileProviderExtension iOS app.
> This was created to power a sample project app for [RayWenderlich](https://www.raywenderlich.com).

### Run locally

In a shell in the root of this project, simply run `go run main.go`

### Endpoints

- [x] **GET** `/media?path={directory}` - Returns all the assets in a directory
- [ ] **POST** `/media` - Upload a new asset to a directory
- [x] **GET** `/file?path={path}&id={file_id}` - Returns the asset for the given id
- [ ] **DELETE** `/file?path={path}&id={file}` - Remove an asset at a path by a given id
- [ ] **PUT** `/file?path={path}&id={file}` - Updates an asset at a path by a given id
