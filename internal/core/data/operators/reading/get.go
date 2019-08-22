/*
 * ******************************************************************************
 *  Copyright 2019 Dell Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software distributed under the License
 *  is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 *  or implied. See the License for the specific language governing permissions and limitations under
 *  the License.
 *  ******************************************************************************
 */

package reading

import (
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/edgexfoundry/edgex-go/internal/pkg/config"
)

// GetReadingsExecutor retrieves one or more readings.
type GetReadingsExecutor interface {
	Execute() ([]contract.Reading, error)
}

// getReadingsByValueDescriptorName encapsulates the data needed to obtain readings by a value descriptor name.
type getReadingsByValueDescriptorName struct {
	name   string
	limit  int
	loader Loader
	logger logger.LoggingClient
	config config.ServiceInfo
}

// Execute retrieves readings by value descriptor name.
func (g getReadingsByValueDescriptorName) Execute() ([]contract.Reading, error) {
	const op = "getReadingsByValueDescriptorName.Execute"
	r, err := g.loader.ReadingsByValueDescriptor(g.name, g.limit)

	if err != nil {
		return nil, contract.NewCommonEdgexError(op, contract.KindDatabaseError, err)
	}

	if len(r) > g.config.MaxResultCount {
		return nil, contract.NewCommonEdgexError(op, contract.KindLimitExceeded, fmt.Errorf("result count %v exceeds max result count of: %v", len(r), g.config.MaxResultCount))
	}

	return r, nil
}

// NewGetReadingsNameExecutor creates a GetReadingsExecutor which will retrieve readings by a value descriptor name.
func NewGetReadingsNameExecutor(name string, limit int, loader Loader, logger logger.LoggingClient, config config.ServiceInfo) GetReadingsExecutor {
	return getReadingsByValueDescriptorName{
		name:   name,
		limit:  limit,
		loader: loader,
		logger: logger,
		config: config,
	}
}
