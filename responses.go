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

type SuzakuResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId uuid.UUID   `json:"request_id"`
}

type poolResponse struct {
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

type ListPoolResponse []poolResponse

func prettify(i interface{}) string {
	resp, _ := json.MarshalIndent(i, "", "   ")
	return string(resp)
}

func (l ListPoolResponse) String() string {
	return prettify(l)
}

type CreatePoolResponse poolResponse

func (c CreatePoolResponse) String() string {
	return prettify(c)
}

type RetrievePoolResponse poolResponse

func (r RetrievePoolResponse) String() string {
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

type subsysResponse struct {
	NVMeoF     string         `json:"NVMeoF"`
	SCSIID     string         `json:"SCSI_ID"`
	CreateTime string         `json:"create_time"`
	ISCSI      string         `json:"iSCSI"`
	Iqn        string         `json:"iqn"`
	MaxNsID    int            `json:"max_nsid"`
	Nqn        string         `json:"nqn"`
	SubsysName string         `json:"subsys_name"`
	VolList    []subsysVolume `json:"vol_list"`
}

type CreateSubsysResponse = subsysResponse

type ListSubsysResponse []subsysResponse

type RetrieveSubsysResponse subsysResponse

type volumeResponse struct {
	Capacity       int    `json:"capacity"`
	ReplicaNum     int    `json:"replica_num"`
	SectorSize     int    `json:"sector_size"`
	CloneStatus    string `json:"clone_status"`
	CreateTime     string `json:"create_time"`
	EcRatio        string `json:"ec_ratio"`
	Format         string `json:"format"`
	PoolName       string `json:"pool_name"`
	RedundancyType string `json:"redundancy_type"`
	Uuid           string `json:"uuid"`
	VolumeName     string `json:"volume_name"`
}

type RetrieveVolumeResponse = volumeResponse

type ListVolumeResponse []volumeResponse

type RetrieveSubsysAuthResponse struct {
	Auth string `json:"auth"`
}

type RetrieveSubsysChapResponse struct {
	Chap string `json:"chap"`
}

type flattenVolumeResponse struct {
	Status string `json:"status"`
	TaskId string `json:"task_id"`
}

func (resp *flattenVolumeResponse) IsDone() bool {
	if resp.Status == "success" {
		return true
	}
	return false
}

type FlattenVolumeResponse = flattenVolumeResponse

type FlattenVolumeProgressResponse = flattenVolumeResponse

type snapshotDesc struct {
	Size         int    `json:"size"`
	Used         int    `json:"used"`
	Reference    int    `json:"reference"`
	CreateTime   string `json:"create_time"`
	VolumeName   string `json:"volume_name"`
	SnapshotName string `json:"snapshot_name"`
}

type CreateSnapshotResponse snapshotDesc

type RetrieveSnapshotResponse snapshotDesc

type snapshotSummary struct {
	Size         int    `json:"size"`
	Reference    int    `json:"reference"`
	CreateTime   string `json:"create_time"`
	VolumeName   string `json:"volume_name"`
	SnapshotName string `json:"snapshot_name"`
}

type ListSnapshotResponse []snapshotSummary

type accountInfo struct {
	AccountName string `json:"account_name"`
	Password    string `json:"password"`
}

type ListAccountResponse []accountInfo

type RetrieveAccountResponse accountInfo

type CreateAccountResponse accountInfo

type groupInfo struct {
	GroupName string            `json:"group_name"`
	ISCSI     map[string]string `json:"iSCSI"`
	NVMeoF    map[string]string `json:"NVMeoF"`
}

type ListGroupResponse []groupInfo

type RetrieveGroupResponse groupInfo

type CreateGroupResponse groupInfo

type AddQualifierToGroupResponse groupInfo

type RemoveQualifierFromGroupResponse groupInfo
