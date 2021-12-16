// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for zookeeperreceiver metrics.
type MetricsSettings struct {
	ZookeeperApproximateDateSize   MetricSettings `mapstructure:"zookeeper.approximate_date_size"`
	ZookeeperConnectionsAlive      MetricSettings `mapstructure:"zookeeper.connections_alive"`
	ZookeeperEphemeralNodes        MetricSettings `mapstructure:"zookeeper.ephemeral_nodes"`
	ZookeeperFollowers             MetricSettings `mapstructure:"zookeeper.followers"`
	ZookeeperFsyncThresholdExceeds MetricSettings `mapstructure:"zookeeper.fsync_threshold_exceeds"`
	ZookeeperLatencyAvg            MetricSettings `mapstructure:"zookeeper.latency.avg"`
	ZookeeperLatencyMax            MetricSettings `mapstructure:"zookeeper.latency.max"`
	ZookeeperLatencyMin            MetricSettings `mapstructure:"zookeeper.latency.min"`
	ZookeeperMaxFileDescriptors    MetricSettings `mapstructure:"zookeeper.max_file_descriptors"`
	ZookeeperOpenFileDescriptors   MetricSettings `mapstructure:"zookeeper.open_file_descriptors"`
	ZookeeperOutstandingRequests   MetricSettings `mapstructure:"zookeeper.outstanding_requests"`
	ZookeeperPacketsReceived       MetricSettings `mapstructure:"zookeeper.packets.received"`
	ZookeeperPacketsSent           MetricSettings `mapstructure:"zookeeper.packets.sent"`
	ZookeeperPendingSyncs          MetricSettings `mapstructure:"zookeeper.pending_syncs"`
	ZookeeperSyncedFollowers       MetricSettings `mapstructure:"zookeeper.synced_followers"`
	ZookeeperWatches               MetricSettings `mapstructure:"zookeeper.watches"`
	ZookeeperZnodes                MetricSettings `mapstructure:"zookeeper.znodes"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		ZookeeperApproximateDateSize: MetricSettings{
			Enabled: true,
		},
		ZookeeperConnectionsAlive: MetricSettings{
			Enabled: true,
		},
		ZookeeperEphemeralNodes: MetricSettings{
			Enabled: true,
		},
		ZookeeperFollowers: MetricSettings{
			Enabled: true,
		},
		ZookeeperFsyncThresholdExceeds: MetricSettings{
			Enabled: true,
		},
		ZookeeperLatencyAvg: MetricSettings{
			Enabled: true,
		},
		ZookeeperLatencyMax: MetricSettings{
			Enabled: true,
		},
		ZookeeperLatencyMin: MetricSettings{
			Enabled: true,
		},
		ZookeeperMaxFileDescriptors: MetricSettings{
			Enabled: true,
		},
		ZookeeperOpenFileDescriptors: MetricSettings{
			Enabled: true,
		},
		ZookeeperOutstandingRequests: MetricSettings{
			Enabled: true,
		},
		ZookeeperPacketsReceived: MetricSettings{
			Enabled: true,
		},
		ZookeeperPacketsSent: MetricSettings{
			Enabled: true,
		},
		ZookeeperPendingSyncs: MetricSettings{
			Enabled: true,
		},
		ZookeeperSyncedFollowers: MetricSettings{
			Enabled: true,
		},
		ZookeeperWatches: MetricSettings{
			Enabled: true,
		},
		ZookeeperZnodes: MetricSettings{
			Enabled: true,
		},
	}
}

// metric holds data for generated metric and keeps track of data points slice capacity.
type metric struct {
	data     pdata.Metric // data buffer for generated metric.
	capacity int          // max observed number of data points added to the metric.
}

func (m *metric) updateCapacity(dpLen int) {
	if dpLen > m.capacity {
		m.capacity = dpLen
	}
}

func newMetric() metric {
	return metric{data: pdata.NewMetric()}
}

type metrics struct {
	zookeeperApproximateDateSize   metric
	zookeeperConnectionsAlive      metric
	zookeeperEphemeralNodes        metric
	zookeeperFollowers             metric
	zookeeperFsyncThresholdExceeds metric
	zookeeperLatencyAvg            metric
	zookeeperLatencyMax            metric
	zookeeperLatencyMin            metric
	zookeeperMaxFileDescriptors    metric
	zookeeperOpenFileDescriptors   metric
	zookeeperOutstandingRequests   metric
	zookeeperPacketsReceived       metric
	zookeeperPacketsSent           metric
	zookeeperPendingSyncs          metric
	zookeeperSyncedFollowers       metric
	zookeeperWatches               metric
	zookeeperZnodes                metric
}

func newMetrics(config MetricsSettings) metrics {
	ms := metrics{}
	if config.ZookeeperApproximateDateSize.Enabled {
		ms.zookeeperApproximateDateSize = newMetric()
	}
	if config.ZookeeperConnectionsAlive.Enabled {
		ms.zookeeperConnectionsAlive = newMetric()
	}
	if config.ZookeeperEphemeralNodes.Enabled {
		ms.zookeeperEphemeralNodes = newMetric()
	}
	if config.ZookeeperFollowers.Enabled {
		ms.zookeeperFollowers = newMetric()
	}
	if config.ZookeeperFsyncThresholdExceeds.Enabled {
		ms.zookeeperFsyncThresholdExceeds = newMetric()
	}
	if config.ZookeeperLatencyAvg.Enabled {
		ms.zookeeperLatencyAvg = newMetric()
	}
	if config.ZookeeperLatencyMax.Enabled {
		ms.zookeeperLatencyMax = newMetric()
	}
	if config.ZookeeperLatencyMin.Enabled {
		ms.zookeeperLatencyMin = newMetric()
	}
	if config.ZookeeperMaxFileDescriptors.Enabled {
		ms.zookeeperMaxFileDescriptors = newMetric()
	}
	if config.ZookeeperOpenFileDescriptors.Enabled {
		ms.zookeeperOpenFileDescriptors = newMetric()
	}
	if config.ZookeeperOutstandingRequests.Enabled {
		ms.zookeeperOutstandingRequests = newMetric()
	}
	if config.ZookeeperPacketsReceived.Enabled {
		ms.zookeeperPacketsReceived = newMetric()
	}
	if config.ZookeeperPacketsSent.Enabled {
		ms.zookeeperPacketsSent = newMetric()
	}
	if config.ZookeeperPendingSyncs.Enabled {
		ms.zookeeperPendingSyncs = newMetric()
	}
	if config.ZookeeperSyncedFollowers.Enabled {
		ms.zookeeperSyncedFollowers = newMetric()
	}
	if config.ZookeeperWatches.Enabled {
		ms.zookeeperWatches = newMetric()
	}
	if config.ZookeeperZnodes.Enabled {
		ms.zookeeperZnodes = newMetric()
	}
	return ms
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user configuration.
type MetricsBuilder struct {
	config    MetricsSettings
	startTime pdata.Timestamp
	metrics   metrics
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(config MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:    config,
		startTime: pdata.NewTimestampFromTime(time.Now()),
		metrics:   newMetrics(config),
	}

	for _, op := range options {
		op(mb)
	}

	mb.initMetrics()
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user configuration, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(metrics pdata.MetricSlice) {
	if mb.config.ZookeeperApproximateDateSize.Enabled && mb.metrics.zookeeperApproximateDateSize.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperApproximateDateSize.updateCapacity(mb.metrics.zookeeperApproximateDateSize.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperApproximateDateSize.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperConnectionsAlive.Enabled && mb.metrics.zookeeperConnectionsAlive.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperConnectionsAlive.updateCapacity(mb.metrics.zookeeperConnectionsAlive.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperConnectionsAlive.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperEphemeralNodes.Enabled && mb.metrics.zookeeperEphemeralNodes.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperEphemeralNodes.updateCapacity(mb.metrics.zookeeperEphemeralNodes.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperEphemeralNodes.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperFollowers.Enabled && mb.metrics.zookeeperFollowers.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperFollowers.updateCapacity(mb.metrics.zookeeperFollowers.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperFollowers.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperFsyncThresholdExceeds.Enabled && mb.metrics.zookeeperFsyncThresholdExceeds.data.Sum().DataPoints().Len() > 0 {
		mb.metrics.zookeeperFsyncThresholdExceeds.updateCapacity(mb.metrics.zookeeperFsyncThresholdExceeds.data.Sum().DataPoints().Len())
		mb.metrics.zookeeperFsyncThresholdExceeds.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperLatencyAvg.Enabled && mb.metrics.zookeeperLatencyAvg.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperLatencyAvg.updateCapacity(mb.metrics.zookeeperLatencyAvg.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperLatencyAvg.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperLatencyMax.Enabled && mb.metrics.zookeeperLatencyMax.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperLatencyMax.updateCapacity(mb.metrics.zookeeperLatencyMax.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperLatencyMax.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperLatencyMin.Enabled && mb.metrics.zookeeperLatencyMin.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperLatencyMin.updateCapacity(mb.metrics.zookeeperLatencyMin.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperLatencyMin.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperMaxFileDescriptors.Enabled && mb.metrics.zookeeperMaxFileDescriptors.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperMaxFileDescriptors.updateCapacity(mb.metrics.zookeeperMaxFileDescriptors.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperMaxFileDescriptors.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperOpenFileDescriptors.Enabled && mb.metrics.zookeeperOpenFileDescriptors.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperOpenFileDescriptors.updateCapacity(mb.metrics.zookeeperOpenFileDescriptors.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperOpenFileDescriptors.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperOutstandingRequests.Enabled && mb.metrics.zookeeperOutstandingRequests.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperOutstandingRequests.updateCapacity(mb.metrics.zookeeperOutstandingRequests.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperOutstandingRequests.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperPacketsReceived.Enabled && mb.metrics.zookeeperPacketsReceived.data.Sum().DataPoints().Len() > 0 {
		mb.metrics.zookeeperPacketsReceived.updateCapacity(mb.metrics.zookeeperPacketsReceived.data.Sum().DataPoints().Len())
		mb.metrics.zookeeperPacketsReceived.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperPacketsSent.Enabled && mb.metrics.zookeeperPacketsSent.data.Sum().DataPoints().Len() > 0 {
		mb.metrics.zookeeperPacketsSent.updateCapacity(mb.metrics.zookeeperPacketsSent.data.Sum().DataPoints().Len())
		mb.metrics.zookeeperPacketsSent.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperPendingSyncs.Enabled && mb.metrics.zookeeperPendingSyncs.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperPendingSyncs.updateCapacity(mb.metrics.zookeeperPendingSyncs.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperPendingSyncs.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperSyncedFollowers.Enabled && mb.metrics.zookeeperSyncedFollowers.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperSyncedFollowers.updateCapacity(mb.metrics.zookeeperSyncedFollowers.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperSyncedFollowers.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperWatches.Enabled && mb.metrics.zookeeperWatches.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperWatches.updateCapacity(mb.metrics.zookeeperWatches.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperWatches.data.MoveTo(metrics.AppendEmpty())
	}
	if mb.config.ZookeeperZnodes.Enabled && mb.metrics.zookeeperZnodes.data.Gauge().DataPoints().Len() > 0 {
		mb.metrics.zookeeperZnodes.updateCapacity(mb.metrics.zookeeperZnodes.data.Gauge().DataPoints().Len())
		mb.metrics.zookeeperZnodes.data.MoveTo(metrics.AppendEmpty())
	}

	// Reset metric data points collection.
	mb.initMetrics()
}

// initZookeeperApproximateDateSizeMetric builds new zookeeper.approximate_date_size metric.
func (mb *MetricsBuilder) initZookeeperApproximateDateSizeMetric() {
	metric := mb.metrics.zookeeperApproximateDateSize
	metric.data.SetName("zookeeper.approximate_date_size")
	metric.data.SetDescription("Size of data in bytes that a ZooKeeper server has in its data tree.")
	metric.data.SetUnit("By")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperConnectionsAliveMetric builds new zookeeper.connections_alive metric.
func (mb *MetricsBuilder) initZookeeperConnectionsAliveMetric() {
	metric := mb.metrics.zookeeperConnectionsAlive
	metric.data.SetName("zookeeper.connections_alive")
	metric.data.SetDescription("Number of active clients connected to a ZooKeeper server.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperEphemeralNodesMetric builds new zookeeper.ephemeral_nodes metric.
func (mb *MetricsBuilder) initZookeeperEphemeralNodesMetric() {
	metric := mb.metrics.zookeeperEphemeralNodes
	metric.data.SetName("zookeeper.ephemeral_nodes")
	metric.data.SetDescription("Number of ephemeral nodes that a ZooKeeper server has in its data tree.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperFollowersMetric builds new zookeeper.followers metric.
func (mb *MetricsBuilder) initZookeeperFollowersMetric() {
	metric := mb.metrics.zookeeperFollowers
	metric.data.SetName("zookeeper.followers")
	metric.data.SetDescription("The number of followers in sync with the leader. Only exposed by the leader.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperFsyncThresholdExceedsMetric builds new zookeeper.fsync_threshold_exceeds metric.
func (mb *MetricsBuilder) initZookeeperFsyncThresholdExceedsMetric() {
	metric := mb.metrics.zookeeperFsyncThresholdExceeds
	metric.data.SetName("zookeeper.fsync_threshold_exceeds")
	metric.data.SetDescription("Number of times fsync duration has exceeded warning threshold.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeSum)
	metric.data.Sum().SetIsMonotonic(true)
	metric.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

// initZookeeperLatencyAvgMetric builds new zookeeper.latency.avg metric.
func (mb *MetricsBuilder) initZookeeperLatencyAvgMetric() {
	metric := mb.metrics.zookeeperLatencyAvg
	metric.data.SetName("zookeeper.latency.avg")
	metric.data.SetDescription("Average time in milliseconds for requests to be processed.")
	metric.data.SetUnit("ms")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperLatencyMaxMetric builds new zookeeper.latency.max metric.
func (mb *MetricsBuilder) initZookeeperLatencyMaxMetric() {
	metric := mb.metrics.zookeeperLatencyMax
	metric.data.SetName("zookeeper.latency.max")
	metric.data.SetDescription("Maximum time in milliseconds for requests to be processed.")
	metric.data.SetUnit("ms")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperLatencyMinMetric builds new zookeeper.latency.min metric.
func (mb *MetricsBuilder) initZookeeperLatencyMinMetric() {
	metric := mb.metrics.zookeeperLatencyMin
	metric.data.SetName("zookeeper.latency.min")
	metric.data.SetDescription("Minimum time in milliseconds for requests to be processed.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperMaxFileDescriptorsMetric builds new zookeeper.max_file_descriptors metric.
func (mb *MetricsBuilder) initZookeeperMaxFileDescriptorsMetric() {
	metric := mb.metrics.zookeeperMaxFileDescriptors
	metric.data.SetName("zookeeper.max_file_descriptors")
	metric.data.SetDescription("Maximum number of file descriptors that a ZooKeeper server can open.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperOpenFileDescriptorsMetric builds new zookeeper.open_file_descriptors metric.
func (mb *MetricsBuilder) initZookeeperOpenFileDescriptorsMetric() {
	metric := mb.metrics.zookeeperOpenFileDescriptors
	metric.data.SetName("zookeeper.open_file_descriptors")
	metric.data.SetDescription("Number of file descriptors that a ZooKeeper server has open.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperOutstandingRequestsMetric builds new zookeeper.outstanding_requests metric.
func (mb *MetricsBuilder) initZookeeperOutstandingRequestsMetric() {
	metric := mb.metrics.zookeeperOutstandingRequests
	metric.data.SetName("zookeeper.outstanding_requests")
	metric.data.SetDescription("Number of currently executing requests.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperPacketsReceivedMetric builds new zookeeper.packets.received metric.
func (mb *MetricsBuilder) initZookeeperPacketsReceivedMetric() {
	metric := mb.metrics.zookeeperPacketsReceived
	metric.data.SetName("zookeeper.packets.received")
	metric.data.SetDescription("Number of ZooKeeper packets received by a server.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeSum)
	metric.data.Sum().SetIsMonotonic(true)
	metric.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

// initZookeeperPacketsSentMetric builds new zookeeper.packets.sent metric.
func (mb *MetricsBuilder) initZookeeperPacketsSentMetric() {
	metric := mb.metrics.zookeeperPacketsSent
	metric.data.SetName("zookeeper.packets.sent")
	metric.data.SetDescription("Number of ZooKeeper packets sent by a server.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeSum)
	metric.data.Sum().SetIsMonotonic(true)
	metric.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

// initZookeeperPendingSyncsMetric builds new zookeeper.pending_syncs metric.
func (mb *MetricsBuilder) initZookeeperPendingSyncsMetric() {
	metric := mb.metrics.zookeeperPendingSyncs
	metric.data.SetName("zookeeper.pending_syncs")
	metric.data.SetDescription("The number of pending syncs from the followers. Only exposed by the leader.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperSyncedFollowersMetric builds new zookeeper.synced_followers metric.
func (mb *MetricsBuilder) initZookeeperSyncedFollowersMetric() {
	metric := mb.metrics.zookeeperSyncedFollowers
	metric.data.SetName("zookeeper.synced_followers")
	metric.data.SetDescription("The number of followers in sync with the leader. Only exposed by the leader.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperWatchesMetric builds new zookeeper.watches metric.
func (mb *MetricsBuilder) initZookeeperWatchesMetric() {
	metric := mb.metrics.zookeeperWatches
	metric.data.SetName("zookeeper.watches")
	metric.data.SetDescription("Number of watches placed on Z-Nodes on a ZooKeeper server.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initZookeeperZnodesMetric builds new zookeeper.znodes metric.
func (mb *MetricsBuilder) initZookeeperZnodesMetric() {
	metric := mb.metrics.zookeeperZnodes
	metric.data.SetName("zookeeper.znodes")
	metric.data.SetDescription("Number of z-nodes that a ZooKeeper server has in its data tree.")
	metric.data.SetUnit("1")
	metric.data.SetDataType(pdata.MetricDataTypeGauge)
}

// initMetrics initializes metrics.
func (mb *MetricsBuilder) initMetrics() {
	if mb.config.ZookeeperApproximateDateSize.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperApproximateDateSizeMetric()
	}
	if mb.config.ZookeeperConnectionsAlive.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperConnectionsAliveMetric()
	}
	if mb.config.ZookeeperEphemeralNodes.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperEphemeralNodesMetric()
	}
	if mb.config.ZookeeperFollowers.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperFollowersMetric()
	}
	if mb.config.ZookeeperFsyncThresholdExceeds.Enabled {
		// TODO: Use metric.data.Sum().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperFsyncThresholdExceedsMetric()
	}
	if mb.config.ZookeeperLatencyAvg.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperLatencyAvgMetric()
	}
	if mb.config.ZookeeperLatencyMax.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperLatencyMaxMetric()
	}
	if mb.config.ZookeeperLatencyMin.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperLatencyMinMetric()
	}
	if mb.config.ZookeeperMaxFileDescriptors.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperMaxFileDescriptorsMetric()
	}
	if mb.config.ZookeeperOpenFileDescriptors.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperOpenFileDescriptorsMetric()
	}
	if mb.config.ZookeeperOutstandingRequests.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperOutstandingRequestsMetric()
	}
	if mb.config.ZookeeperPacketsReceived.Enabled {
		// TODO: Use metric.data.Sum().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperPacketsReceivedMetric()
	}
	if mb.config.ZookeeperPacketsSent.Enabled {
		// TODO: Use metric.data.Sum().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperPacketsSentMetric()
	}
	if mb.config.ZookeeperPendingSyncs.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperPendingSyncsMetric()
	}
	if mb.config.ZookeeperSyncedFollowers.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperSyncedFollowersMetric()
	}
	if mb.config.ZookeeperWatches.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperWatchesMetric()
	}
	if mb.config.ZookeeperZnodes.Enabled {
		// TODO: Use metric.data.Gauge().DataPoints().Clear() instead of rebuilding
		// the metrics once the Clear method is available.
		mb.initZookeeperZnodesMetric()
	}
}

// RecordZookeeperApproximateDateSizeDataPoint adds a data point to zookeeper.approximate_date_size metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperApproximateDateSizeDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperApproximateDateSize.Enabled {
		return
	}

	dp := mb.metrics.zookeeperApproximateDateSize.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperConnectionsAliveDataPoint adds a data point to zookeeper.connections_alive metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperConnectionsAliveDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperConnectionsAlive.Enabled {
		return
	}

	dp := mb.metrics.zookeeperConnectionsAlive.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperEphemeralNodesDataPoint adds a data point to zookeeper.ephemeral_nodes metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperEphemeralNodesDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperEphemeralNodes.Enabled {
		return
	}

	dp := mb.metrics.zookeeperEphemeralNodes.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperFollowersDataPoint adds a data point to zookeeper.followers metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperFollowersDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperFollowers.Enabled {
		return
	}

	dp := mb.metrics.zookeeperFollowers.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperFsyncThresholdExceedsDataPoint adds a data point to zookeeper.fsync_threshold_exceeds metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperFsyncThresholdExceedsDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperFsyncThresholdExceeds.Enabled {
		return
	}

	dp := mb.metrics.zookeeperFsyncThresholdExceeds.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperLatencyAvgDataPoint adds a data point to zookeeper.latency.avg metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperLatencyAvgDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperLatencyAvg.Enabled {
		return
	}

	dp := mb.metrics.zookeeperLatencyAvg.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperLatencyMaxDataPoint adds a data point to zookeeper.latency.max metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperLatencyMaxDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperLatencyMax.Enabled {
		return
	}

	dp := mb.metrics.zookeeperLatencyMax.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperLatencyMinDataPoint adds a data point to zookeeper.latency.min metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperLatencyMinDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperLatencyMin.Enabled {
		return
	}

	dp := mb.metrics.zookeeperLatencyMin.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperMaxFileDescriptorsDataPoint adds a data point to zookeeper.max_file_descriptors metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperMaxFileDescriptorsDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperMaxFileDescriptors.Enabled {
		return
	}

	dp := mb.metrics.zookeeperMaxFileDescriptors.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperOpenFileDescriptorsDataPoint adds a data point to zookeeper.open_file_descriptors metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperOpenFileDescriptorsDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperOpenFileDescriptors.Enabled {
		return
	}

	dp := mb.metrics.zookeeperOpenFileDescriptors.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperOutstandingRequestsDataPoint adds a data point to zookeeper.outstanding_requests metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperOutstandingRequestsDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperOutstandingRequests.Enabled {
		return
	}

	dp := mb.metrics.zookeeperOutstandingRequests.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperPacketsReceivedDataPoint adds a data point to zookeeper.packets.received metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperPacketsReceivedDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperPacketsReceived.Enabled {
		return
	}

	dp := mb.metrics.zookeeperPacketsReceived.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperPacketsSentDataPoint adds a data point to zookeeper.packets.sent metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperPacketsSentDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperPacketsSent.Enabled {
		return
	}

	dp := mb.metrics.zookeeperPacketsSent.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperPendingSyncsDataPoint adds a data point to zookeeper.pending_syncs metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperPendingSyncsDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperPendingSyncs.Enabled {
		return
	}

	dp := mb.metrics.zookeeperPendingSyncs.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperSyncedFollowersDataPoint adds a data point to zookeeper.synced_followers metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperSyncedFollowersDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperSyncedFollowers.Enabled {
		return
	}

	dp := mb.metrics.zookeeperSyncedFollowers.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperWatchesDataPoint adds a data point to zookeeper.watches metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperWatchesDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperWatches.Enabled {
		return
	}

	dp := mb.metrics.zookeeperWatches.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// RecordZookeeperZnodesDataPoint adds a data point to zookeeper.znodes metric.
// Any attribute of AttributeValueTypeEmpty type will be skipped.
func (mb *MetricsBuilder) RecordZookeeperZnodesDataPoint(ts pdata.Timestamp, val int64) {
	if !mb.config.ZookeeperZnodes.Enabled {
		return
	}

	dp := mb.metrics.zookeeperZnodes.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(mb.startTime)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// ServerState (State of the Zookeeper server (leader, standalone or follower).)
	ServerState string
	// ZkVersion (Zookeeper version of the instance.)
	ZkVersion string
}{
	"server.state",
	"zk.version",
}

// A is an alias for Attributes.
var A = Attributes
