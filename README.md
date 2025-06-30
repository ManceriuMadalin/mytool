# mytool

A multifunctional CLI tool written in Go that can:
- 🔐 Generate secure passwords (`genpass`)
- 🗂️ Perform incremental backups (`backup`)
- 🌐 Download images from URLs (`download`)

## Usage

```bash
./mytool genpass --length 20
./mytool backup --source ./test --dest ./backup
./mytool download --url https://example.com/image.jpg --dest ./downloads
