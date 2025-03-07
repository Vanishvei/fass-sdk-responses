package responses

// File       : responses.go
// Path       : responses
// Time       : CST 2023/4/10 14:44
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"encoding/json"
	"github.com/google/uuid"
)

func prettify(i interface{}) string {
	resp, _ := json.MarshalIndent(i, "", "   ")
	return string(resp)
}

type SuzakuResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId uuid.UUID   `json:"request_id"`
	Token     *uuid.UUID  `json:"token"`
	Total     *int        `json:"total"`
	PageNum   *int        `json:"page_num"`
	PageSize  *int        `json:"page_size"`
}

func (s *SuzakuResponse) String() string {
	if s.Token != nil {
		return prettify(s)
	}

	return prettify(struct {
		Code      int         `json:"code"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
		RequestId uuid.UUID   `json:"request_id"`
	}{
		Code:      s.Code,
		Message:   s.Message,
		Data:      s.Data,
		RequestId: s.RequestId,
	})
}

type RetrievePool struct {
	Name                string `json:"name"`
	CreateTime          string `json:"create_time"`
	EcRatio             string `json:"ec_ratio"`
	RedundancyType      string `json:"redundancy_type"`
	ProvisionedCapacity int64  `json:"provisioned_capacity"`
	TotalCapacity       int64  `json:"total_capacity"`
	ReplicaNum          int    `json:"replica_num"`
	SectorSize          int    `json:"sector_size"`
	Split               int    `json:"split"`
}

func (r *RetrievePool) String() string {
	return prettify(r)
}

type subsysVolume struct {
	Bps            int    `json:"bps"`
	BpsBurst       int    `json:"bps_burst"`
	BurstPeriod    int    `json:"burst_period"`
	Capacity       int64  `json:"capacity"`
	CloneStatus    string `json:"clone_status"`
	CreateTime     string `json:"create_time"`
	EcRatio        string `json:"ec_ratio"`
	Format         string `json:"format"`
	Iops           int    `json:"iops"`
	IopsBurst      int    `json:"iops_burst"`
	NsID           int    `json:"nsid"`
	PoolName       string `json:"pool_name"`
	RedundancyType string `json:"redundancy_type"`
	ReplicaNum     int    `json:"replica_num"`
	SectorSize     int    `json:"sector_size"`
	Sharding       string `json:"sharding"`
	UUID           string `json:"uuid"`
	VolumeName     string `json:"volume_name"`
}

func (s *subsysVolume) String() string {
	return prettify(s)
}

type subsysResponse struct {
	MaxNsID    int            `json:"max_nsid"`
	Iqn        string         `json:"iqn"`
	Nqn        string         `json:"nqn"`
	ISCSI      string         `json:"iSCSI"`
	NVMeoF     string         `json:"NVMeoF"`
	SCSIID     string         `json:"SCSI_ID"`
	CreateTime string         `json:"create_time"`
	SubsysName string         `json:"subsys_name"`
	VIPList    []string       `json:"vip_list"`
	VolList    []subsysVolume `json:"vol_list"`
}

func (s *subsysResponse) String() string {
	return prettify(s)
}

type CreateSubsys = subsysResponse

type ListSubsys struct {
	Data     []subsysResponse `json:"data"`
	PageNum  int              `json:"page_num"`
	PageSize int              `json:"page_size"`
	Total    int              `json:"total"`
	Token    string           `json:"token"`
}

func (ls *ListSubsys) String() string {
	return prettify(ls)
}

type RetrieveSubsys = subsysResponse

type RetrieveSubsysVLAN struct {
	Subsys   string   `json:"subsys"`
	VLANList []string `json:"vlan_list"`
}

type volumeSummaryResponse struct {
	ReplicaNum     int    `json:"replica_num"`
	SectorSize     int    `json:"sector_size"`
	BurstPeriod    int    `json:"burst_period"`
	IopsBurst      int    `json:"iops_burst"`
	BpsBurst       int    `json:"bps_burst"`
	Iops           int    `json:"iops"`
	Bps            int    `json:"bps"`
	Capacity       int64  `json:"capacity"`
	UUID           string `json:"uuid"`
	Format         string `json:"format"`
	EcRatio        string `json:"ec_ratio"`
	Sharding       string `json:"sharding"`
	PoolName       string `json:"pool_name"`
	VolumeName     string `json:"volume_name"`
	CreateTime     string `json:"create_time"`
	CloneStatus    string `json:"clone_status"`
	RedundancyType string `json:"redundancy_type"`
}

func (s *volumeSummaryResponse) String() string {
	return prettify(s)
}

type ListVolume struct {
	Data     []volumeSummaryResponse `json:"data"`
	PageNum  int                     `json:"page_num"`
	PageSize int                     `json:"page_size"`
	Total    int                     `json:"total"`
	Token    string                  `json:"token"`
}

func (l *ListVolume) String() string {
	return prettify(l)
}

type RetrieveVolume struct {
	Bps            int    `json:"bps"`
	Iops           int    `json:"iops"`
	BpsBurst       int    `json:"bps_burst"`
	IopsBurst      int    `json:"iops_burst"`
	ReplicaNum     int    `json:"replica_num"`
	SectorSize     int    `json:"sector_size"`
	BurstPeriod    int    `json:"burst_period"`
	Used           int64  `json:"used"`
	Capacity       int64  `json:"capacity"`
	UUID           string `json:"uuid"`
	Format         string `json:"format"`
	EcRatio        string `json:"ec_ratio"`
	Sharding       string `json:"sharding"`
	PoolName       string `json:"pool_name"`
	VolumeName     string `json:"volume_name"`
	CreateTime     string `json:"create_time"`
	CloneStatus    string `json:"clone_status"`
	RedundancyType string `json:"redundancy_type"`
}

func (r *RetrieveVolume) String() string {
	return prettify(r)
}

type RetrieveSubsysHostGroup struct {
	HostGroup string `json:"host_group"`
	Subsys    string `json:"subsys"`
}

func (r *RetrieveSubsysHostGroup) String() string {
	return prettify(r)
}

type RetrieveSubsysChap struct {
	Chap   string `json:"chap"`
	Subsys string `json:"subsys"`
}

func (r *RetrieveSubsysChap) String() string {
	return prettify(r)
}

type flattenVolumeResponse struct {
	Status string `json:"status"`
	TaskId string `json:"task_id"`
}

func (f *flattenVolumeResponse) IsDone() bool {
	if f.Status == "success" {
		return true
	}
	return false
}

func (f *flattenVolumeResponse) String() string {
	return prettify(f)
}

type FlattenVolume = flattenVolumeResponse

type FlattenVolumeProgress = flattenVolumeResponse

type snapshotDesc struct {
	Size         int64  `json:"size"`
	Used         int64  `json:"used"`
	Reference    int    `json:"reference"`
	CreateTime   string `json:"create_time"`
	VolumeName   string `json:"volume_name"`
	SnapshotName string `json:"snapshot_name"`
}

func (s *snapshotDesc) String() string {
	return prettify(s)
}

type CreateSnapshot = snapshotDesc

type RetrieveSnapshot = snapshotDesc

type snapshotSummary struct {
	Size         int64  `json:"size"`
	Reference    int    `json:"reference"`
	CreateTime   string `json:"create_time"`
	VolumeName   string `json:"volume_name"`
	SnapshotName string `json:"snapshot_name"`
}

func (s *snapshotSummary) String() string {
	return prettify(s)
}

type ListSnapshot = []snapshotSummary

type accountInfo struct {
	AccountName string `json:"account_name"`
	CreateTime  string `json:"create_time"`
}

func (a *accountInfo) String() string {
	return prettify(a)
}

type ListAccount = []accountInfo

type RetrieveAccount = accountInfo

type CreateAccount = accountInfo

type hostGroupInfo struct {
	GroupName string            `json:"group_name"`
	ISCSI     map[string]string `json:"iSCSI"`
	NVMeoF    map[string]string `json:"NVMeoF"`
}

func (g *hostGroupInfo) String() string {
	return prettify(g)
}

type ListHostGroup = []hostGroupInfo

type AddHostToHostGroup = hostGroupInfo

type RemoveHostFromHostGroup = hostGroupInfo
