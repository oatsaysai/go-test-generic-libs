# Go test generic libs

## Prerequisites

- Go 1.22 or later
- Mariadb 10.3.7 or later
- S3/Min.IO
- make

## Getting Started

1. Start deps docker

```sh
cd scripts-for-local
./start-deps-env.sh
```

2. Run service

- Back to root dir (cd ..)

```sh
make run
```
