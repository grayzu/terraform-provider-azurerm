package datashare

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DataSetMappingStatus enumerates the values for data set mapping status.
type DataSetMappingStatus string

const (
	// Broken ...
	Broken DataSetMappingStatus = "Broken"
	// Ok ...
	Ok DataSetMappingStatus = "Ok"
)

// PossibleDataSetMappingStatusValues returns an array of possible values for the DataSetMappingStatus const type.
func PossibleDataSetMappingStatusValues() []DataSetMappingStatus {
	return []DataSetMappingStatus{Broken, Ok}
}

// DataSetType enumerates the values for data set type.
type DataSetType string

const (
	// AdlsGen1File ...
	AdlsGen1File DataSetType = "AdlsGen1File"
	// AdlsGen1Folder ...
	AdlsGen1Folder DataSetType = "AdlsGen1Folder"
	// AdlsGen2File ...
	AdlsGen2File DataSetType = "AdlsGen2File"
	// AdlsGen2FileSystem ...
	AdlsGen2FileSystem DataSetType = "AdlsGen2FileSystem"
	// AdlsGen2Folder ...
	AdlsGen2Folder DataSetType = "AdlsGen2Folder"
	// Blob ...
	Blob DataSetType = "Blob"
	// BlobFolder ...
	BlobFolder DataSetType = "BlobFolder"
	// Container ...
	Container DataSetType = "Container"
	// KustoCluster ...
	KustoCluster DataSetType = "KustoCluster"
	// KustoDatabase ...
	KustoDatabase DataSetType = "KustoDatabase"
	// SQLDBTable ...
	SQLDBTable DataSetType = "SqlDBTable"
	// SQLDWTable ...
	SQLDWTable DataSetType = "SqlDWTable"
)

// PossibleDataSetTypeValues returns an array of possible values for the DataSetType const type.
func PossibleDataSetTypeValues() []DataSetType {
	return []DataSetType{AdlsGen1File, AdlsGen1Folder, AdlsGen2File, AdlsGen2FileSystem, AdlsGen2Folder, Blob, BlobFolder, Container, KustoCluster, KustoDatabase, SQLDBTable, SQLDWTable}
}

// InvitationStatus enumerates the values for invitation status.
type InvitationStatus string

const (
	// Accepted ...
	Accepted InvitationStatus = "Accepted"
	// Pending ...
	Pending InvitationStatus = "Pending"
	// Rejected ...
	Rejected InvitationStatus = "Rejected"
	// Withdrawn ...
	Withdrawn InvitationStatus = "Withdrawn"
)

// PossibleInvitationStatusValues returns an array of possible values for the InvitationStatus const type.
func PossibleInvitationStatusValues() []InvitationStatus {
	return []InvitationStatus{Accepted, Pending, Rejected, Withdrawn}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindAdlsGen1File ...
	KindAdlsGen1File Kind = "AdlsGen1File"
	// KindAdlsGen1Folder ...
	KindAdlsGen1Folder Kind = "AdlsGen1Folder"
	// KindAdlsGen2File ...
	KindAdlsGen2File Kind = "AdlsGen2File"
	// KindAdlsGen2FileSystem ...
	KindAdlsGen2FileSystem Kind = "AdlsGen2FileSystem"
	// KindAdlsGen2Folder ...
	KindAdlsGen2Folder Kind = "AdlsGen2Folder"
	// KindBlob ...
	KindBlob Kind = "Blob"
	// KindBlobFolder ...
	KindBlobFolder Kind = "BlobFolder"
	// KindContainer ...
	KindContainer Kind = "Container"
	// KindDataSet ...
	KindDataSet Kind = "DataSet"
	// KindKustoCluster ...
	KindKustoCluster Kind = "KustoCluster"
	// KindKustoDatabase ...
	KindKustoDatabase Kind = "KustoDatabase"
	// KindSQLDBTable ...
	KindSQLDBTable Kind = "SqlDBTable"
	// KindSQLDWTable ...
	KindSQLDWTable Kind = "SqlDWTable"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindAdlsGen1File, KindAdlsGen1Folder, KindAdlsGen2File, KindAdlsGen2FileSystem, KindAdlsGen2Folder, KindBlob, KindBlobFolder, KindContainer, KindDataSet, KindKustoCluster, KindKustoDatabase, KindSQLDBTable, KindSQLDWTable}
}

// KindBasicDataSetMapping enumerates the values for kind basic data set mapping.
type KindBasicDataSetMapping string

const (
	// KindBasicDataSetMappingKindAdlsGen2File ...
	KindBasicDataSetMappingKindAdlsGen2File KindBasicDataSetMapping = "AdlsGen2File"
	// KindBasicDataSetMappingKindAdlsGen2FileSystem ...
	KindBasicDataSetMappingKindAdlsGen2FileSystem KindBasicDataSetMapping = "AdlsGen2FileSystem"
	// KindBasicDataSetMappingKindAdlsGen2Folder ...
	KindBasicDataSetMappingKindAdlsGen2Folder KindBasicDataSetMapping = "AdlsGen2Folder"
	// KindBasicDataSetMappingKindBlob ...
	KindBasicDataSetMappingKindBlob KindBasicDataSetMapping = "Blob"
	// KindBasicDataSetMappingKindBlobFolder ...
	KindBasicDataSetMappingKindBlobFolder KindBasicDataSetMapping = "BlobFolder"
	// KindBasicDataSetMappingKindContainer ...
	KindBasicDataSetMappingKindContainer KindBasicDataSetMapping = "Container"
	// KindBasicDataSetMappingKindDataSetMapping ...
	KindBasicDataSetMappingKindDataSetMapping KindBasicDataSetMapping = "DataSetMapping"
	// KindBasicDataSetMappingKindKustoCluster ...
	KindBasicDataSetMappingKindKustoCluster KindBasicDataSetMapping = "KustoCluster"
	// KindBasicDataSetMappingKindKustoDatabase ...
	KindBasicDataSetMappingKindKustoDatabase KindBasicDataSetMapping = "KustoDatabase"
	// KindBasicDataSetMappingKindSQLDBTable ...
	KindBasicDataSetMappingKindSQLDBTable KindBasicDataSetMapping = "SqlDBTable"
	// KindBasicDataSetMappingKindSQLDWTable ...
	KindBasicDataSetMappingKindSQLDWTable KindBasicDataSetMapping = "SqlDWTable"
)

// PossibleKindBasicDataSetMappingValues returns an array of possible values for the KindBasicDataSetMapping const type.
func PossibleKindBasicDataSetMappingValues() []KindBasicDataSetMapping {
	return []KindBasicDataSetMapping{KindBasicDataSetMappingKindAdlsGen2File, KindBasicDataSetMappingKindAdlsGen2FileSystem, KindBasicDataSetMappingKindAdlsGen2Folder, KindBasicDataSetMappingKindBlob, KindBasicDataSetMappingKindBlobFolder, KindBasicDataSetMappingKindContainer, KindBasicDataSetMappingKindDataSetMapping, KindBasicDataSetMappingKindKustoCluster, KindBasicDataSetMappingKindKustoDatabase, KindBasicDataSetMappingKindSQLDBTable, KindBasicDataSetMappingKindSQLDWTable}
}

// KindBasicSourceShareSynchronizationSetting enumerates the values for kind basic source share synchronization
// setting.
type KindBasicSourceShareSynchronizationSetting string

const (
	// KindScheduleBased ...
	KindScheduleBased KindBasicSourceShareSynchronizationSetting = "ScheduleBased"
	// KindSourceShareSynchronizationSetting ...
	KindSourceShareSynchronizationSetting KindBasicSourceShareSynchronizationSetting = "SourceShareSynchronizationSetting"
)

// PossibleKindBasicSourceShareSynchronizationSettingValues returns an array of possible values for the KindBasicSourceShareSynchronizationSetting const type.
func PossibleKindBasicSourceShareSynchronizationSettingValues() []KindBasicSourceShareSynchronizationSetting {
	return []KindBasicSourceShareSynchronizationSetting{KindScheduleBased, KindSourceShareSynchronizationSetting}
}

// KindBasicSynchronizationSetting enumerates the values for kind basic synchronization setting.
type KindBasicSynchronizationSetting string

const (
	// KindBasicSynchronizationSettingKindScheduleBased ...
	KindBasicSynchronizationSettingKindScheduleBased KindBasicSynchronizationSetting = "ScheduleBased"
	// KindBasicSynchronizationSettingKindSynchronizationSetting ...
	KindBasicSynchronizationSettingKindSynchronizationSetting KindBasicSynchronizationSetting = "SynchronizationSetting"
)

// PossibleKindBasicSynchronizationSettingValues returns an array of possible values for the KindBasicSynchronizationSetting const type.
func PossibleKindBasicSynchronizationSettingValues() []KindBasicSynchronizationSetting {
	return []KindBasicSynchronizationSetting{KindBasicSynchronizationSettingKindScheduleBased, KindBasicSynchronizationSettingKindSynchronizationSetting}
}

// KindBasicTrigger enumerates the values for kind basic trigger.
type KindBasicTrigger string

const (
	// KindBasicTriggerKindScheduleBased ...
	KindBasicTriggerKindScheduleBased KindBasicTrigger = "ScheduleBased"
	// KindBasicTriggerKindTrigger ...
	KindBasicTriggerKindTrigger KindBasicTrigger = "Trigger"
)

// PossibleKindBasicTriggerValues returns an array of possible values for the KindBasicTrigger const type.
func PossibleKindBasicTriggerValues() []KindBasicTrigger {
	return []KindBasicTrigger{KindBasicTriggerKindScheduleBased, KindBasicTriggerKindTrigger}
}

// OutputType enumerates the values for output type.
type OutputType string

const (
	// Csv ...
	Csv OutputType = "Csv"
	// Parquet ...
	Parquet OutputType = "Parquet"
)

// PossibleOutputTypeValues returns an array of possible values for the OutputType const type.
func PossibleOutputTypeValues() []OutputType {
	return []OutputType{Csv, Parquet}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Moving ...
	Moving ProvisioningState = "Moving"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Deleting, Failed, Moving, Succeeded}
}

// RecurrenceInterval enumerates the values for recurrence interval.
type RecurrenceInterval string

const (
	// Day ...
	Day RecurrenceInterval = "Day"
	// Hour ...
	Hour RecurrenceInterval = "Hour"
)

// PossibleRecurrenceIntervalValues returns an array of possible values for the RecurrenceInterval const type.
func PossibleRecurrenceIntervalValues() []RecurrenceInterval {
	return []RecurrenceInterval{Day, Hour}
}

// ShareKind enumerates the values for share kind.
type ShareKind string

const (
	// CopyBased ...
	CopyBased ShareKind = "CopyBased"
	// InPlace ...
	InPlace ShareKind = "InPlace"
)

// PossibleShareKindValues returns an array of possible values for the ShareKind const type.
func PossibleShareKindValues() []ShareKind {
	return []ShareKind{CopyBased, InPlace}
}

// ShareSubscriptionStatus enumerates the values for share subscription status.
type ShareSubscriptionStatus string

const (
	// Active ...
	Active ShareSubscriptionStatus = "Active"
	// Revoked ...
	Revoked ShareSubscriptionStatus = "Revoked"
	// Revoking ...
	Revoking ShareSubscriptionStatus = "Revoking"
	// SourceDeleted ...
	SourceDeleted ShareSubscriptionStatus = "SourceDeleted"
)

// PossibleShareSubscriptionStatusValues returns an array of possible values for the ShareSubscriptionStatus const type.
func PossibleShareSubscriptionStatusValues() []ShareSubscriptionStatus {
	return []ShareSubscriptionStatus{Active, Revoked, Revoking, SourceDeleted}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusAccepted ...
	StatusAccepted Status = "Accepted"
	// StatusCanceled ...
	StatusCanceled Status = "Canceled"
	// StatusFailed ...
	StatusFailed Status = "Failed"
	// StatusInProgress ...
	StatusInProgress Status = "InProgress"
	// StatusSucceeded ...
	StatusSucceeded Status = "Succeeded"
	// StatusTransientFailure ...
	StatusTransientFailure Status = "TransientFailure"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusAccepted, StatusCanceled, StatusFailed, StatusInProgress, StatusSucceeded, StatusTransientFailure}
}

// SynchronizationMode enumerates the values for synchronization mode.
type SynchronizationMode string

const (
	// FullSync ...
	FullSync SynchronizationMode = "FullSync"
	// Incremental ...
	Incremental SynchronizationMode = "Incremental"
)

// PossibleSynchronizationModeValues returns an array of possible values for the SynchronizationMode const type.
func PossibleSynchronizationModeValues() []SynchronizationMode {
	return []SynchronizationMode{FullSync, Incremental}
}

// TriggerStatus enumerates the values for trigger status.
type TriggerStatus string

const (
	// TriggerStatusActive ...
	TriggerStatusActive TriggerStatus = "Active"
	// TriggerStatusInactive ...
	TriggerStatusInactive TriggerStatus = "Inactive"
	// TriggerStatusSourceSynchronizationSettingDeleted ...
	TriggerStatusSourceSynchronizationSettingDeleted TriggerStatus = "SourceSynchronizationSettingDeleted"
)

// PossibleTriggerStatusValues returns an array of possible values for the TriggerStatus const type.
func PossibleTriggerStatusValues() []TriggerStatus {
	return []TriggerStatus{TriggerStatusActive, TriggerStatusInactive, TriggerStatusSourceSynchronizationSettingDeleted}
}

// Type enumerates the values for type.
type Type string

const (
	// SystemAssigned ...
	SystemAssigned Type = "SystemAssigned"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{SystemAssigned}
}
