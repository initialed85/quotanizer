# quotanizer
Small Go app to manage a quota on folders by deleting the oldest things

## What does it do?

You give it some paths, some quotas (in GB), some extensions and a period (optional, default 60 seconds) and it spins around maintaining the quotas by deleting the oldest files (until the quota is met).

## How do I build it?

    ./build.sh
    
## How do I test it?

    ./test.sh
    
## How do I run it?

This is just an example (ensure you have run `./build.sh` first):

    docker run --rm -it \
        -v ~/Desktop:/mnt/Desktop \
        -v ~/Music:/mnt/Music \
        -v ~/Pictures:/mnt/Pictures \
        quotanizer \
            -path /mnt/Desktop \
            -quota 1 \
            -path /mnt/Music \
            -quota 1 \
            -path /mnt/Pictures \
            -quota 1 \
            -suffix .mkv \
            -suffix .mp4 \
            -period 5
