// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"

	"istio.io/mixer"
	"istio.io/mixer/adapters"
	"istio.io/mixer/adapters/denyChecker"
	"istio.io/mixer/adapters/factMapper"
	"istio.io/mixer/adapters/genericListChecker"
	"istio.io/mixer/adapters/ipListChecker"
)

// all the known fact converter adapter types
var factConverters = []adapters.Builder{
	factMapper.NewBuilder(),
}

// all the known fact updater adapter types
var factUpdaters = []adapters.Builder{}

// all the known list checker adapter types
var listCheckers = []adapters.Builder{
	ipListChecker.NewBuilder(),
	genericListChecker.NewBuilder(),
	denyChecker.NewBuilder(),
}

// AdapterManager keeps track of activated Adapter objects for different types of adapters
type AdapterManager struct {
	// FactUpdaters is the set of fact updater adapters
	FactUpdaters map[string]adapters.Builder

	// FactConverters is the set of fact converter adapters
	FactConverters map[string]adapters.Builder

	// ListCheckers is the set of list checker adapters
	ListCheckers map[string]adapters.Builder
}

// GetListCheckerAdapter returns a matching adapter for the given dispatchKey. If there is no existing adapter,
// it instantiates one, based on the provided adapter config.
func (mgr *AdapterManager) GetListCheckerAdapter(dispatchKey mixer.DispatchKey, config *adapters.AdapterConfig) (adapters.ListChecker, error) {
	// TODO: instantiation & caching of the adapters.
	return nil, errors.New("NYI")
}

// GetLoggerAdapter returns a matching adapter for the given dispatchKey. If there is no existing adapter,
// it instantiates one, based on the provided adapter config.
func (mgr *AdapterManager) GetLoggerAdapter(dispatchKey mixer.DispatchKey, config *adapters.AdapterConfig) (adapters.Logger, error) {
	// TODO: instantiation & caching of the adapters.
	return nil, errors.New("NYI")
}

// TODO: this implementation needs to be driven from external config instead of being hardcoded
func prepBuilders(l []adapters.Builder) (map[string]adapters.Builder, error) {
	m := make(map[string]adapters.Builder, len(l))

	for _, builder := range l {
		if _, exists := m[builder.Name()]; exists {
			return nil, errors.New("can't have multiple builders called " + builder.Name())
		}

		builder.Configure(builder.DefaultBuilderConfig())

		m[builder.Name()] = builder
	}

	return m, nil
}

// NewAdapterManager returns a new AdapterManager instance
func NewAdapterManager() (*AdapterManager, error) {
	var mgr = AdapterManager{}
	var err error

	if mgr.FactUpdaters, err = prepBuilders(factUpdaters); err != nil {
		return nil, err
	}

	if mgr.FactConverters, err = prepBuilders(factConverters); err != nil {
		return nil, err
	}

	if mgr.ListCheckers, err = prepBuilders(listCheckers); err != nil {
		return nil, err
	}

	return &mgr, nil
}
