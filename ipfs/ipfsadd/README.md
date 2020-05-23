# IPFSADD

Simple script to add a single file or a whole directory in a remote IPFS node

## Usage

### Command line

```
./ipfsadd -host 11.11.22.22:5001 -path /path/to/dir
```

### Docker

```
docker run -ti --rm -v /path/to/dir:/dir -e IPFS_API_HOST=11.11.22.22:5001 -e IPFS_ADD_PATH=/dir mqll/ipfsadd
```
