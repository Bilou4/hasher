# Hasher

Hasher is a command line tool that brings together the most common hash algorithms in order to have everything in one binary.


Hasher is faster than others command line tool computing hashes.

The `hyperfine` tool allows us to compare Hasher md5 with md5sum as an example.

```bash
# We create a large file to be computed
fallocate -l 5G largeFile

# Proceed to execution
hyperfine --warmup 3 'md5sum largeFile' './build/hasher-linux-amd64 hash md5 largeFile'
Benchmark 1: md5sum largeFile
  Time (mean ± σ):     13.133 s ±  0.370 s    [User: 10.824 s, System: 2.229 s]
  Range (min … max):   12.760 s … 13.864 s    10 runs

Benchmark 2: ./build/hasher-linux-amd64 hash md5 largeFile
  Time (mean ± σ):     12.870 s ±  0.711 s    [User: 10.498 s, System: 2.153 s]
  Range (min … max):   12.317 s … 14.703 s    10 runs

Summary
  './build/hasher-linux-amd64 hash md5 largeFile' ran
    1.02 ± 0.06 times faster than 'md5sum largeFile'
```

## TODO

- [ ] Adler-32	"hash/adler32"
- [ ] CRC-32	"hash/crc32"
- [ ] CRC-64	"hash/crc64"
- [ ] FNV   	"hash/fnv"