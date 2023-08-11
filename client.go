package deepl

import (
	"errors"
	"fmt"
	"github.com/EYERCORD/deepl-sdk-go/types"
)

type Client struct {
	AuthenticationKey string
	APIPlanType       types.APIPlan
	EndpointBase      string
}

// NewClient returns new DeepL API client.
//
//	Authentication key: DEEPL_API_AUTHN_KEY
//	API plan: DEEPL_API_PLAN
//	          `free` or `pro` (ignore upper case or lower case)
func NewClient(AuthKey string, DeeplAPIPlan string) (*Client, error) {
	var c Client

	authnKey := AuthKey
	if authnKey == "" {
		return nil, errors.New("DEEPL_API_AUTHN_KEY is empty")
	}
	c.AuthenticationKey = authnKey

	planName := DeeplAPIPlan
	plan, err := types.ToAPIPlan(planName)
	if err != nil {
		return nil, errors.New("DEEPL_API_PLAN is empty or invalid. Set `free` or `pro`.")
	}
	c.APIPlanType = plan

	if plan == types.APIPlanFree {
		c.EndpointBase = fmt.Sprintf("https://%s/%s/", types.DomainFree, types.APIVersion)
	} else if plan == types.APIPlanPro {
		c.EndpointBase = fmt.Sprintf("https://%s/%s/", types.DomainPro, types.APIVersion)
	}

	return &c, nil
}
