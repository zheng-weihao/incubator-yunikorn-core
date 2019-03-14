/*
Copyright 2019 The Unity Scheduler Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scheduler

import "fmt"

type SchedulingAllocation struct {
    SchedulingAsk *SchedulingAllocationAsk
    NumAllocation int32
    NodeId        string
}

func NewSchedulingAllocation(ask *SchedulingAllocationAsk, nodeId string) *SchedulingAllocation {
    return &SchedulingAllocation{SchedulingAsk: ask, NodeId: nodeId, NumAllocation: 1}
}

func (m *SchedulingAllocation) String() string {
    return fmt.Sprintf("{AllocatioKey=%s,NumAllocation=%d,Node=%s", m.SchedulingAsk.AskProto.AllocationKey, m.NumAllocation, m.NodeId)
}