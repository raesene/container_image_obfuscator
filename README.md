# Container Image Obfuscator

This is a dumb PoC to make a point, do not use this in production! It takes a tarball of a container image and extracts it at runtime using a small Golang program. This essentially means that any tools that scan the image won't see the actual contents of the image, they'll see a tarball :)

## re-compiling

If you want to change the `golang` code you need to statically compile it

`CGO_ENABLED=0 go build -ldflags="-s -w" main.go`

If you want to avoid it coming up with any vulnerabilities for the golang binary you can use `upx` to compress it.

## building the image

`docker build -t obfuscator .`

## running the image

`docker run obfuscator [command]`

so to run `ls` you would do `docker run obfuscator ls`

