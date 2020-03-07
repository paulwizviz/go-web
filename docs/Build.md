# Build process

Out-of-the-box you will find scripts to help you build native and container based apps based 
on the included skeleton ReactJS and Go Rest API codes. If you plan to extend from the skeletons, 
you just need to extend from these scripts. If you plan to replace the skeletons, please use this
build script as inspiration.

NOTE: All scripts are BASH based and located in the folder `scripts`

## Build locally

1. Open terminal and navigate to the root of your scaffold

### Native macOS, Linux and Windows

2. Run the command `./scripts/native/build.sh package`

3. You will find your build arfects in the folder `./build/package/<native platform>`

### Docker container based app

2. Edit the script `./scripts/container/build.sh` and modify the variable `$IMAGE_NAME` to one that suits your project

3. Run the command `./scripts/container/build.sh package <image tag>`, where `<image tag>` is any unique string value e.g. version number

4. You will find your build arfects in the folder `./build/package/<native platform>`

## Clean project

* Native apps run the command `./scripts/native/build.sh clean`

* Container run the command `./scripts/container/build.sh clean <image tag>` 