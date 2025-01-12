# GoRTB ( OpenRTB 2.5 )

[![Go Reference](https://pkg.go.dev/badge/github.com/MehrazRumman/gortb.svg)](https://pkg.go.dev/github.com/MehrazRumman/gortb)
[![Go Test](https://github.com/MehrazRumman/gortb/actions/workflows/go.yml/badge.svg)](https://github.com/MehrazRumman/gortb/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/MehrazRumman/gortb)](https://goreportcard.com/report/github.com/MehrazRumman/gortb)





GoRTB is a Go implementation of the OpenRTB (Real-Time Bidding) 2.5 specification. This library provides a robust and type-safe way to work with OpenRTB requests and responses in your Go applications.

## Go Version Support

GoRTB supports Go versions 1.13 and later.

## Features

- Full implementation of OpenRTB 2.5 specification
- Type-safe bid request and response handling
- Easy-to-use API for both bidders and exchanges
- Comprehensive validation of RTB objects
- High performance JSON serialization/deserialization
- Extensible architecture for custom implementations

## Installation

To install GoRTB, use `go get`:

```bash
go get github.com/MehrazRumman/gortb
```

## Quick Start

Here's a simple example of how to create a bid request:

```go
package main

import (
    "fmt"
    "github.com/MehrazRumman/gortb"
)

func main() {
    // Create a new bid request
    bidRequest := &gortb.BidRequest{
        ID: "bid-req-1",
        Imp: []gortb.Impression{{
            ID: "imp-1",
            Banner: &gortb.Banner{
                W: 300,
                H: 250,
            },
        }},
    }

    // Convert to JSON
    json, err := bidRequest.Marshal()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Bid Request: %s\n", string(json))
}
```

## Documentation

For detailed documentation and examples, please visit our [Go Package Documentation](https://pkg.go.dev/github.com/MehrazRumman/gortb).

### Main Components

- `BidRequest`: Represents an OpenRTB bid request
- `BidResponse`: Represents an OpenRTB bid response
- `Impression`: Describes an ad impression opportunity
- `Bid`: Contains bid information
- Various objects for different ad types (Banner, Video, Native, etc.)



## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please file an issue on the [GitHub issue tracker](https://github.com/MehrazRumman/gortb/issues).

## Acknowledgments

- OpenRTB 2.5 Specification
- The Go community for their valuable feedback and contributions

## Version

Current version: v1.0.0 
