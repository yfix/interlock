package client

import (
	"encoding/json"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

// ServiceCreate creates a new Service.
<<<<<<< HEAD
func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, headers map[string][]string) (types.ServiceCreateResponse, error) {
=======
func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error) {
	var headers map[string][]string

	if options.EncodedRegistryAuth != "" {
		headers = map[string][]string{
			"X-Registry-Auth": []string{options.EncodedRegistryAuth},
		}
	}

>>>>>>> 12a5469... start on swarm services; move to glade
	var response types.ServiceCreateResponse
	resp, err := cli.post(ctx, "/services/create", nil, service, headers)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.body).Decode(&response)
	ensureReaderClosed(resp)
	return response, err
}