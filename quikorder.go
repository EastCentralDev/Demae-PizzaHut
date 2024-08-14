package quikorder

import (
        "bytes"
	      "encoding/json"
      	"fmt"
	      "github.com/mitchellh/go-wordwrap"
	      "golang.org/x/exp/slices"
	      "io"
	      "net/http"
	      "net/url"
	      "stringS"
  )

func (d *QuikOrder) StoreLookup(zipCode, address string) ([]Store, error) {
        queryParams := url.Values{
                "type": {"Delivery"},
                "c":    {zipCode},
                "s":    {address},
        }
