// Copyright (c) 2021 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/common/types/testdata"
)

func TestHostInfo(t *testing.T) {
	for _, item := range []*types.HostInfo{nil, {}, &testdata.HostInfo} {
		assert.Equal(t, item, ToHostInfo(FromHostInfo(item)))
	}
}
func TestMembershipInfo(t *testing.T) {
	for _, item := range []*types.MembershipInfo{nil, {}, &testdata.MembershipInfo} {
		assert.Equal(t, item, ToMembershipInfo(FromMembershipInfo(item)))
	}
}
func TestDomainCacheInfo(t *testing.T) {
	for _, item := range []*types.DomainCacheInfo{nil, {}, &testdata.DomainCacheInfo} {
		assert.Equal(t, item, ToDomainCacheInfo(FromDomainCacheInfo(item)))
	}
}
func TestRingInfo(t *testing.T) {
	for _, item := range []*types.RingInfo{nil, {}, &testdata.RingInfo} {
		assert.Equal(t, item, ToRingInfo(FromRingInfo(item)))
	}
}
func TestTransientDecisionInfo(t *testing.T) {
	for _, item := range []*types.TransientDecisionInfo{nil, {}, &testdata.TransientDecisionInfo} {
		assert.Equal(t, item, ToTransientDecisionInfo(FromTransientDecisionInfo(item)))
	}
}
func TestVersionHistories(t *testing.T) {
	for _, item := range []*types.VersionHistories{nil, {}, &testdata.VersionHistories} {
		assert.Equal(t, item, ToVersionHistories(FromVersionHistories(item)))
	}
}
func TestVersionHistory(t *testing.T) {
	for _, item := range []*types.VersionHistory{nil, {}, &testdata.VersionHistory} {
		assert.Equal(t, item, ToVersionHistory(FromVersionHistory(item)))
	}
}
func TestVersionHistoryItem(t *testing.T) {
	for _, item := range []*types.VersionHistoryItem{nil, {}, &testdata.VersionHistoryItem} {
		assert.Equal(t, item, ToVersionHistoryItem(FromVersionHistoryItem(item)))
	}
}
func TestHostInfoArray(t *testing.T) {
	for _, item := range [][]*types.HostInfo{nil, {}, testdata.HostInfoArray} {
		assert.Equal(t, item, ToHostInfoArray(FromHostInfoArray(item)))
	}
}
func TestVersionHistoryArray(t *testing.T) {
	for _, item := range [][]*types.VersionHistory{nil, {}, testdata.VersionHistoryArray} {
		assert.Equal(t, item, ToVersionHistoryArray(FromVersionHistoryArray(item)))
	}
}
func TestRingInfoArray(t *testing.T) {
	for _, item := range [][]*types.RingInfo{nil, {}, testdata.RingInfoArray} {
		assert.Equal(t, item, ToRingInfoArray(FromRingInfoArray(item)))
	}
}
func TestDomainTaskAttributes(t *testing.T) {
	for _, item := range []*types.DomainTaskAttributes{nil, &testdata.DomainTaskAttributes} {
		assert.Equal(t, item, ToDomainTaskAttributes(FromDomainTaskAttributes(item)))
	}
}
func TestFailoverMarkerAttributes(t *testing.T) {
	for _, item := range []*types.FailoverMarkerAttributes{nil, {}, &testdata.FailoverMarkerAttributes} {
		assert.Equal(t, item, ToFailoverMarkerAttributes(FromFailoverMarkerAttributes(item)))
	}
}
func TestFailoverMarkerToken(t *testing.T) {
	for _, item := range []*types.FailoverMarkerToken{nil, {}, &testdata.FailoverMarkerToken} {
		assert.Equal(t, item, ToFailoverMarkerToken(FromFailoverMarkerToken(item)))
	}
}
func TestHistoryTaskV2Attributes(t *testing.T) {
	for _, item := range []*types.HistoryTaskV2Attributes{nil, {}, &testdata.HistoryTaskV2Attributes} {
		assert.Equal(t, item, ToHistoryTaskV2Attributes(FromHistoryTaskV2Attributes(item)))
	}
}
func TestReplicationMessages(t *testing.T) {
	for _, item := range []*types.ReplicationMessages{nil, {}, &testdata.ReplicationMessages} {
		assert.Equal(t, item, ToReplicationMessages(FromReplicationMessages(item)))
	}
}
func TestReplicationTaskInfo(t *testing.T) {
	for _, item := range []*types.ReplicationTaskInfo{nil, {}, &testdata.ReplicationTaskInfo} {
		assert.Equal(t, item, ToReplicationTaskInfo(FromReplicationTaskInfo(item)))
	}
}
func TestReplicationToken(t *testing.T) {
	for _, item := range []*types.ReplicationToken{nil, {}, &testdata.ReplicationToken} {
		assert.Equal(t, item, ToReplicationToken(FromReplicationToken(item)))
	}
}
func TestSyncActivityTaskAttributes(t *testing.T) {
	for _, item := range []*types.SyncActivityTaskAttributes{nil, {}, &testdata.SyncActivityTaskAttributes} {
		assert.Equal(t, item, ToSyncActivityTaskAttributes(FromSyncActivityTaskAttributes(item)))
	}
}
func TestSyncShardStatus(t *testing.T) {
	for _, item := range []*types.SyncShardStatus{nil, {}, &testdata.SyncShardStatus} {
		assert.Equal(t, item, ToSyncShardStatus(FromSyncShardStatus(item)))
	}
}
func TestSyncShardStatusTaskAttributes(t *testing.T) {
	for _, item := range []*types.SyncShardStatusTaskAttributes{nil, {}, &testdata.SyncShardStatusTaskAttributes} {
		assert.Equal(t, item, ToSyncShardStatusTaskAttributes(FromSyncShardStatusTaskAttributes(item)))
	}
}
func TestReplicationTaskInfoArray(t *testing.T) {
	for _, item := range [][]*types.ReplicationTaskInfo{nil, {}, testdata.ReplicationTaskInfoArray} {
		assert.Equal(t, item, ToReplicationTaskInfoArray(FromReplicationTaskInfoArray(item)))
	}
}
func TestReplicationTaskArray(t *testing.T) {
	for _, item := range [][]*types.ReplicationTask{nil, {}, testdata.ReplicationTaskArray} {
		assert.Equal(t, item, ToReplicationTaskArray(FromReplicationTaskArray(item)))
	}
}
func TestReplicationTokenArray(t *testing.T) {
	for _, item := range [][]*types.ReplicationToken{nil, {}, testdata.ReplicationTokenArray} {
		assert.Equal(t, item, ToReplicationTokenArray(FromReplicationTokenArray(item)))
	}
}
func TestReplicationMessagesMap(t *testing.T) {
	for _, item := range []map[int32]*types.ReplicationMessages{nil, {}, testdata.ReplicationMessagesMap} {
		assert.Equal(t, item, ToReplicationMessagesMap(FromReplicationMessagesMap(item)))
	}
}
func TestReplicationTask(t *testing.T) {
	for _, item := range []*types.ReplicationTask{
		nil,
		{},
		&testdata.ReplicationTask_Domain,
		&testdata.ReplicationTask_Failover,
		&testdata.ReplicationTask_History,
		&testdata.ReplicationTask_SyncActivity,
		&testdata.ReplicationTask_SyncShard,
	} {
		assert.Equal(t, item, ToReplicationTask(FromReplicationTask(item)))
	}
}
func TestFailoverMarkerTokenArray(t *testing.T) {
	for _, item := range [][]*types.FailoverMarkerToken{nil, {}, testdata.FailoverMarkerTokenArray} {
		assert.Equal(t, item, ToFailoverMarkerTokenArray(FromFailoverMarkerTokenArray(item)))
	}
}
func TestVersionHistoryItemArray(t *testing.T) {
	for _, item := range [][]*types.VersionHistoryItem{nil, {}, testdata.VersionHistoryItemArray} {
		assert.Equal(t, item, ToVersionHistoryItemArray(FromVersionHistoryItemArray(item)))
	}
}
func TestEventIDVersionPair(t *testing.T) {
	assert.Nil(t, FromEventIDVersionPair(nil, nil))
	assert.Nil(t, ToEventID(nil))
	assert.Nil(t, ToEventVersion(nil))

	pair := FromEventIDVersionPair(common.Int64Ptr(testdata.EventID1), common.Int64Ptr(testdata.Version1))
	assert.Equal(t, testdata.EventID1, *ToEventID(pair))
	assert.Equal(t, testdata.Version1, *ToEventVersion(pair))
}

func TestCrossClusterTaskInfo(t *testing.T) {
	for _, item := range []*types.CrossClusterTaskInfo{nil, {}, &testdata.CrossClusterTaskInfo} {
		assert.Equal(t, item, ToCrossClusterTaskInfo(FromCrossClusterTaskInfo(item)))
	}
}

func TestCrossClusterTaskRequest(t *testing.T) {
	for _, item := range []*types.CrossClusterTaskRequest{
		nil,
		{},
		&testdata.CrossClusterTaskRequestStartChildExecution,
		&testdata.CrossClusterTaskRequestCancelExecution,
		&testdata.CrossClusterTaskRequestSignalExecution,
	} {
		assert.Equal(t, item, ToCrossClusterTaskRequest(FromCrossClusterTaskRequest(item)))
	}
}

func TestCrossClusterTaskResponse(t *testing.T) {
	for _, item := range []*types.CrossClusterTaskResponse{
		nil,
		{},
		&testdata.CrossClusterTaskResponseStartChildExecution,
		&testdata.CrossClusterTaskResponseCancelExecution,
		&testdata.CrossClusterTaskResponseSignalExecution,
	} {
		assert.Equal(t, item, ToCrossClusterTaskResponse(FromCrossClusterTaskResponse(item)))
	}
}

func TestCrossClusterTaskRequestArray(t *testing.T) {
	for _, item := range [][]*types.CrossClusterTaskRequest{nil, {}, testdata.CrossClusterTaskRequestArray} {
		assert.Equal(t, item, ToCrossClusterTaskRequestArray(FromCrossClusterTaskRequestArray(item)))
	}
}

func TestCrossClusterTaskResponseArray(t *testing.T) {
	for _, item := range [][]*types.CrossClusterTaskResponse{nil, {}, testdata.CrossClusterTaskResponseArray} {
		assert.Equal(t, item, ToCrossClusterTaskResponseArray(FromCrossClusterTaskResponseArray(item)))
	}
}

func TestCrossClusterTaskRequestMap(t *testing.T) {
	for _, item := range []map[int32][]*types.CrossClusterTaskRequest{nil, {}, testdata.CrossClusterTaskRequestMap} {
		assert.Equal(t, item, ToCrossClusterTaskRequestMap(FromCrossClusterTaskRequestMap(item)))
	}
	assert.Equal(
		t,
		map[int32][]*types.CrossClusterTaskRequest{
			0: {},
		},
		ToCrossClusterTaskRequestMap(FromCrossClusterTaskRequestMap(
			map[int32][]*types.CrossClusterTaskRequest{
				0: nil,
			},
		)),
	)
}

func TestGetTaskFailedCauseMap(t *testing.T) {
	for _, item := range []map[int32]types.GetTaskFailedCause{nil, {}, testdata.GetCrossClusterTaskFailedCauseMap} {
		assert.Equal(t, item, ToGetTaskFailedCauseMap(FromGetTaskFailedCauseMap(item)))
	}
}
