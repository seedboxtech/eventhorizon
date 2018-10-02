// Copyright (c) 2018 - The Event Horizon authors.
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

package local

import (
	"context"
	"testing"
	"time"

	"github.com/looplab/eventhorizon/eventhandler/projector"
	"github.com/stretchr/testify/assert"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/mocks"

	"github.com/looplab/eventhorizon/eventbus"
)

func TestEventBus(t *testing.T) {
	group := NewGroup()
	if group == nil {
		t.Fatal("there should be a group")
	}

	bus1 := NewEventBus(group)
	if bus1 == nil {
		t.Fatal("there should be a bus")
	}

	bus2 := NewEventBus(group)
	if bus2 == nil {
		t.Fatal("there should be a bus")
	}

	eventbus.AcceptanceTest(t, bus1, bus2, time.Second)

	bus1.Close()
	bus2.Close()
	bus1.Wait()
	bus2.Wait()
}

func TestCloseAndWait(t *testing.T) {
	bus := NewEventBus(nil)

	id, _ := eh.ParseUUID("c1138e5f-f6fb-4dd0-8e79-255c6c8d3756")
	event := eh.NewEventForAggregate(mocks.EventType, &mocks.EventData{Content: "event1"}, time.Now(), mocks.AggregateType, id, 1)

	repo := &mocks.Repo{}
	slowProjector := NewSlowProjector(repo, time.Second)
	projectorEventHandler := projector.NewEventHandler(slowProjector, repo)
	bus.AddHandler(eh.MatchAny(), projectorEventHandler)

	bus.PublishEvent(context.Background(), event)

	// Event can't be done processing yet
	assert.False(t, repo.SaveCalled)

	bus.Close()
	bus.Wait()

	// Event must be processed after the wait
	assert.True(t, repo.SaveCalled)
}

// SlowProjector is a projector that takes time to handle events
type SlowProjector struct {
	repo  eh.ReadWriteRepo
	delay time.Duration
}

// NewSlowProjector creates a new SlowProjector.
func NewSlowProjector(repo eh.ReadWriteRepo, delay time.Duration) *SlowProjector {
	return &SlowProjector{repo: repo, delay: delay}
}

// ProjectorType method of the eventhorizon.Projector interface.
func (p *SlowProjector) ProjectorType() projector.Type {
	return projector.Type("SlowProjector")
}

// Project method of the eventhorizon.Projector interface.
func (p *SlowProjector) Project(ctx context.Context, event eh.Event, entity eh.Entity) (eh.Entity, error) {
	time.Sleep(p.delay)

	// Calling save to be able to assert that something happened after the sleep
	p.repo.Save(ctx, entity)
	return entity, nil
}
