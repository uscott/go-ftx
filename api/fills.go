package api

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uscott/go-ftx/models"
)

const (
	apiGetFills = "/fills"
)

type Fills struct {
	client *Client
}
