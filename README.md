# pseudo-reddit
## Pseudo Reddit Postgres DB implementation

# Go installation

1. In your terminal, run `sudo tar -C /usr/local -xzf (downloaded go version, e.g. go1.14.linux-amd64.tar.gz)`

2. Check the correctness by running `ls -ls/usr/local/go`

3. Edit profile `nano ~/.profile` and add `export PATH=$PATH:/usr/local/go/bin` line to the end of the file.
   Also export the home directory of the user `export PATH=$PATH:/home/your_directory/go/bin`(e.g. export PATH=$PATH:/home/goweb/go/bin )

# Installation

1. Install unique module `go mod int github.com/monkrus/pseudo-reddit`

2. Install migration tool `tar -xf migrate,linux-amd64.tar.gz` .Then follow by `sudo mv migrate.linux-amd64 /usr/local/bin/migrate`, 
   checking installation by typing `which migrate`, and `migrate --help`

3. Install docker https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04

4. After creating makefile, run DB `make postgres`, `make adminer`, and migrations `run migrate`.

5. sqlx package is used for DB queries `go get github.com/jmoiron/sqlx`

6. Download driver for en/decoding messages b/w Go and DB :`go get github.com/lib/pq`

7. Install chi router :`go get -u github.com/go-chi/chi`

8. Run server `go run cmd/goreddit/main.go`

