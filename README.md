# muarachmann.com

My personal protfolio site. This can be used as a start off point for anyone willing to learn go
and do same!

**NB** This application uses [go]() 1.12.6 as main programming language. Install it if you don't yet have it via the official documentation.

## Setup

Clone the repo

```bash
git clone "github.com/muarachmann/muarachmann.com"
```

Install dependencies locally:

Get into the source's code directory

```bash
cd muarachmann.com
```

We use the [dep](https://golang.github.io/dep/) to managing the various dependencies.
Make sure to download and install it if you don't have. Run the following:

```bash
dep init
```

Configure your app environment variables with the sample `.env-example`

```bash
cp .env-example .env
```

Run the application by typing the following command:

```bash
go run main.go
```

Navigate to the browser and type

```bash
localhost:3000
```

Enjoy!!!