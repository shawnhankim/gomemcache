# memcached Docker

Memcached is a general-purpose distributed memory caching system. It is often used to speed up dynamic database-driven websites by caching data and objects in RAM to reduce the number of times an external data source (such as a database or API) must be read.

Memcached's APIs provide a very large hash table distributed across multiple machines. When the table is full, subsequent inserts cause older data to be purged in least recently used order. Applications using Memcached typically layer requests and additions into RAM before falling back on a slower backing store, such as a database.

## Docker Image
```bash
$ docker run --name my-memcache -p 11211:11211 -d memcached
```

## Setting Memory Usage
The following command would set the Memcached server to use 64 megabytes for storage. For infomation on configuring your memcached server, see the extensive [GitHub Wiki](https://github.com/memcached/memcached/wiki).

```bash
$ docker run --name my-memcache -d memcached memcached -m 64
```

## Reference
- [Docker Hub: memcached](https://hub.docker.com/_/memcached)
- [Wikipedia: Memcached](https://en.wikipedia.org/wiki/Memcached)
