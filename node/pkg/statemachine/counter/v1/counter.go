// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	counterv1 "github.com/atomix/multi-raft-storage/api/atomix/multiraft/counter/v1"
	"github.com/atomix/multi-raft-storage/node/pkg/snapshot"
	"github.com/atomix/multi-raft-storage/node/pkg/statemachine"
	"github.com/atomix/runtime/sdk/pkg/errors"
	"github.com/gogo/protobuf/proto"
)

const Service = "atomix.multiraft.counter.v1.Counter"

func Register(registry *statemachine.PrimitiveTypeRegistry) {
	statemachine.RegisterPrimitiveType[*counterv1.CounterInput, *counterv1.CounterOutput](registry)(CounterType)
}

var CounterType = statemachine.NewPrimitiveType[*counterv1.CounterInput, *counterv1.CounterOutput](Service, counterCodec, newCounterStateMachine)

var counterCodec = statemachine.NewCodec[*counterv1.CounterInput, *counterv1.CounterOutput](
	func(bytes []byte) (*counterv1.CounterInput, error) {
		input := &counterv1.CounterInput{}
		if err := proto.Unmarshal(bytes, input); err != nil {
			return nil, err
		}
		return input, nil
	},
	func(output *counterv1.CounterOutput) ([]byte, error) {
		return proto.Marshal(output)
	})

func newCounterStateMachine(ctx statemachine.PrimitiveContext[*counterv1.CounterInput, *counterv1.CounterOutput]) statemachine.Primitive[*counterv1.CounterInput, *counterv1.CounterOutput] {
	sm := &CounterStateMachine{
		PrimitiveContext: ctx,
	}
	sm.init()
	return sm
}

type CounterStateMachine struct {
	statemachine.PrimitiveContext[*counterv1.CounterInput, *counterv1.CounterOutput]
	value     int64
	increment statemachine.Updater[*counterv1.CounterInput, *counterv1.CounterOutput, *counterv1.IncrementInput, *counterv1.IncrementOutput]
	decrement statemachine.Updater[*counterv1.CounterInput, *counterv1.CounterOutput, *counterv1.DecrementInput, *counterv1.DecrementOutput]
	get       statemachine.Reader[*counterv1.CounterInput, *counterv1.CounterOutput, *counterv1.GetInput, *counterv1.GetOutput]
}

func (s *CounterStateMachine) init() {
	s.increment = statemachine.NewUpdater[*counterv1.CounterInput, *counterv1.CounterOutput, *counterv1.IncrementInput, *counterv1.IncrementOutput](s).
		Name("Increment").
		Decoder(func(input *counterv1.CounterInput) (*counterv1.IncrementInput, bool) {
			if set, ok := input.Input.(*counterv1.CounterInput_Increment); ok {
				return set.Increment, true
			}
			return nil, false
		}).
		Encoder(func(output *counterv1.IncrementOutput) *counterv1.CounterOutput {
			return &counterv1.CounterOutput{
				Output: &counterv1.CounterOutput_Increment{
					Increment: output,
				},
			}
		}).
		Build(s.doIncrement)
	s.decrement = statemachine.NewUpdater[*counterv1.CounterInput, *counterv1.CounterOutput, *counterv1.DecrementInput, *counterv1.DecrementOutput](s).
		Name("Decrement").
		Decoder(func(input *counterv1.CounterInput) (*counterv1.DecrementInput, bool) {
			if set, ok := input.Input.(*counterv1.CounterInput_Decrement); ok {
				return set.Decrement, true
			}
			return nil, false
		}).
		Encoder(func(output *counterv1.DecrementOutput) *counterv1.CounterOutput {
			return &counterv1.CounterOutput{
				Output: &counterv1.CounterOutput_Decrement{
					Decrement: output,
				},
			}
		}).
		Build(s.doDecrement)
	s.get = statemachine.NewReader[*counterv1.CounterInput, *counterv1.CounterOutput, *counterv1.GetInput, *counterv1.GetOutput](s).
		Name("Get").
		Decoder(func(input *counterv1.CounterInput) (*counterv1.GetInput, bool) {
			if set, ok := input.Input.(*counterv1.CounterInput_Get); ok {
				return set.Get, true
			}
			return nil, false
		}).
		Encoder(func(output *counterv1.GetOutput) *counterv1.CounterOutput {
			return &counterv1.CounterOutput{
				Output: &counterv1.CounterOutput_Get{
					Get: output,
				},
			}
		}).
		Build(s.doGet)
}

func (s *CounterStateMachine) Snapshot(writer *snapshot.Writer) error {
	return writer.WriteVarInt64(s.value)
}

func (s *CounterStateMachine) Recover(reader *snapshot.Reader) error {
	i, err := reader.ReadVarInt64()
	if err != nil {
		return err
	}
	s.value = i
	return nil
}

func (s *CounterStateMachine) Update(proposal statemachine.Proposal[*counterv1.CounterInput, *counterv1.CounterOutput]) {
	switch proposal.Input().Input.(type) {
	case *counterv1.CounterInput_Increment:
		s.increment.Update(proposal)
	case *counterv1.CounterInput_Decrement:
		s.decrement.Update(proposal)
	default:
		proposal.Error(errors.NewNotSupported("proposal not supported"))
	}
}

func (s *CounterStateMachine) doIncrement(proposal statemachine.Proposal[*counterv1.IncrementInput, *counterv1.IncrementOutput]) {
	defer proposal.Close()
	s.value += proposal.Input().Delta
	proposal.Output(&counterv1.IncrementOutput{
		Value: s.value,
	})
}

func (s *CounterStateMachine) doDecrement(proposal statemachine.Proposal[*counterv1.DecrementInput, *counterv1.DecrementOutput]) {
	defer proposal.Close()
	s.value -= proposal.Input().Delta
	proposal.Output(&counterv1.DecrementOutput{
		Value: s.value,
	})
}

func (s *CounterStateMachine) Read(query statemachine.Query[*counterv1.CounterInput, *counterv1.CounterOutput]) {
	switch query.Input().Input.(type) {
	case *counterv1.CounterInput_Get:
		s.get.Read(query)
	default:
		query.Error(errors.NewNotSupported("query not supported"))
	}
}

func (s *CounterStateMachine) doGet(query statemachine.Query[*counterv1.GetInput, *counterv1.GetOutput]) {
	defer query.Close()
	query.Output(&counterv1.GetOutput{
		Value: s.value,
	})
}
