module exFiber

go 1.22.2

require (
	github.com/gofiber/fiber/v2 v2.52.4
	github.com/gofiber/storage/redis/v3 v3.1.1
	middleware/session v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/redis/go-redis/v9 v9.5.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.52.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	storage/memory v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	middleware/session => ./internal/session
	storage/memory => ./internal/storage/memory
)
