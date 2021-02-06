#!/usr/bin/env bash

version=$1
package_name="bippy"
if [[ -z "$version" ]]; then
  echo "usage: $0 <version>"
  exit 1
fi

mkdir release

platforms=("windows/amd64" "linux/amd64" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$version'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o ${output_name} ../
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    shasum -a 256 ${output_name} > ${output_name}.checksum
    mv ${output_name}* release
done
