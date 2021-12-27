# Easy Compression

An easy-to-use CLI-based compression tool.

## Usage

```
NAME:
   EasyCompression - A CLI-based tool for (de)compression

USAGE:
   EasyCompression [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

AUTHOR:
   Michael Tei

COMMANDS:
   compress, c    Compresses the input
   decompress, d  Decompresses the input
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output value, -o value  Sets the output name
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

## Examples

- Compress `path_to_file` to a file named `compressed`

```shell
EasyCompression(.exe) c path_to_file
```

- Compress `path_to_file` to a file named `output_name`

```shell
EasyCompression(.exe) -o output_name c path_to_file
```

- Decompress `path_to_file` to a file named `decompressed`

```shell
EasyCompression(.exe) d path_to_file
```

- Decompress `path_to_file` to a file named `output_name`

```shell
EasyCompression(.exe) -o output_name d path_to_file
```

## Dev

### Install

```shell
go install
```

### Build

```shell
go build
```

## References

- [Zstandard](https://facebook.github.io/zstd/)
- [LZ4](https://lz4.github.io/lz4/)
- [cli](https://github.com/urfave/cli)