package training

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
)

// OrderBy enumerates the values for order by.
type OrderBy string

const (
	// Newest ...
	Newest OrderBy = "Newest"
	// Oldest ...
	Oldest OrderBy = "Oldest"
	// Suggested ...
	Suggested OrderBy = "Suggested"
)

// PossibleOrderByValues returns an array of possible values for the OrderBy const type.
func PossibleOrderByValues() []OrderBy {
	return []OrderBy{Newest, Oldest, Suggested}
}

// Platform enumerates the values for platform.
type Platform string

const (
	// CoreML ...
	CoreML Platform = "CoreML"
	// TensorFlow ...
	TensorFlow Platform = "TensorFlow"
)

// PossiblePlatformValues returns an array of possible values for the Platform const type.
func PossiblePlatformValues() []Platform {
	return []Platform{CoreML, TensorFlow}
}

// Status enumerates the values for status.
type Status string

const (
	// ErrorImageFormat ...
	ErrorImageFormat Status = "ErrorImageFormat"
	// ErrorImageSize ...
	ErrorImageSize Status = "ErrorImageSize"
	// ErrorLimitExceed ...
	ErrorLimitExceed Status = "ErrorLimitExceed"
	// ErrorSource ...
	ErrorSource Status = "ErrorSource"
	// ErrorStorage ...
	ErrorStorage Status = "ErrorStorage"
	// ErrorTagLimitExceed ...
	ErrorTagLimitExceed Status = "ErrorTagLimitExceed"
	// ErrorUnknown ...
	ErrorUnknown Status = "ErrorUnknown"
	// OK ...
	OK Status = "OK"
	// OKDuplicate ...
	OKDuplicate Status = "OKDuplicate"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{ErrorImageFormat, ErrorImageSize, ErrorLimitExceed, ErrorSource, ErrorStorage, ErrorTagLimitExceed, ErrorUnknown, OK, OKDuplicate}
}

// Status1 enumerates the values for status 1.
type Status1 string

const (
	// Done ...
	Done Status1 = "Done"
	// Exporting ...
	Exporting Status1 = "Exporting"
	// Failed ...
	Failed Status1 = "Failed"
)

// PossibleStatus1Values returns an array of possible values for the Status1 const type.
func PossibleStatus1Values() []Status1 {
	return []Status1{Done, Exporting, Failed}
}

// Account represents a user account
type Account struct {
	autorest.Response `json:"-"`
	// UserName - Gets the name of the account owner
	UserName *string `json:"UserName,omitempty"`
	// Email - Gets the email associated with this account
	Email *string `json:"Email,omitempty"`
	// Keys - Gets the api keys associated with this account
	Keys *APIKeys `json:"Keys,omitempty"`
	// Quotas - Gets the quotas associated with this account
	Quotas *AccountQuota `json:"Quotas,omitempty"`
}

// AccountQuota represents a set of quotas assocated with an account
type AccountQuota struct {
	// Tier - Gets the tier of user
	Tier *string `json:"Tier,omitempty"`
	// Projects - Gets the project quota
	Projects *Quota `json:"Projects,omitempty"`
	// Predictions - Gets the prediction quota
	Predictions *Quota `json:"Predictions,omitempty"`
	// PerProject - Gets a list of quotas that apply per-project for each project
	PerProject *[]PerProjectQuota `json:"PerProject,omitempty"`
}

// APIKeys ...
type APIKeys struct {
	TrainingKeys   *KeyPair `json:"TrainingKeys,omitempty"`
	PredictionKeys *KeyPair `json:"PredictionKeys,omitempty"`
}

// Domain ...
type Domain struct {
	autorest.Response `json:"-"`
	ID                *uuid.UUID `json:"Id,omitempty"`
	Name              *string    `json:"Name,omitempty"`
	Exportable        *bool      `json:"Exportable,omitempty"`
}

// Export ...
type Export struct {
	autorest.Response `json:"-"`
	// Platform - Possible values include: 'CoreML', 'TensorFlow'
	Platform Platform `json:"Platform,omitempty"`
	// Status - Possible values include: 'Exporting', 'Failed', 'Done'
	Status      Status1 `json:"Status,omitempty"`
	DownloadURI *string `json:"DownloadUri,omitempty"`
}

// Image image model to be sent as JSON
type Image struct {
	ID           *uuid.UUID       `json:"Id,omitempty"`
	Created      *date.Time       `json:"Created,omitempty"`
	Width        *int32           `json:"Width,omitempty"`
	Height       *int32           `json:"Height,omitempty"`
	ImageURI     *string          `json:"ImageUri,omitempty"`
	ThumbnailURI *string          `json:"ThumbnailUri,omitempty"`
	Tags         *[]ImageTag      `json:"Tags,omitempty"`
	Predictions  *[]PredictionTag `json:"Predictions,omitempty"`
}

// ImageCreateResult ...
type ImageCreateResult struct {
	SourceURL *string `json:"SourceUrl,omitempty"`
	// Status - Possible values include: 'OK', 'OKDuplicate', 'ErrorSource', 'ErrorImageFormat', 'ErrorImageSize', 'ErrorStorage', 'ErrorLimitExceed', 'ErrorTagLimitExceed', 'ErrorUnknown'
	Status Status `json:"Status,omitempty"`
	Image  *Image `json:"Image,omitempty"`
}

// ImageCreateSummary ...
type ImageCreateSummary struct {
	autorest.Response `json:"-"`
	IsBatchSuccessful *bool                `json:"IsBatchSuccessful,omitempty"`
	Images            *[]ImageCreateResult `json:"Images,omitempty"`
}

// ImageFileCreateBatch ...
type ImageFileCreateBatch struct {
	Images *[]ImageFileCreateEntry `json:"Images,omitempty"`
	TagIds *[]uuid.UUID            `json:"TagIds,omitempty"`
}

// ImageFileCreateEntry ...
type ImageFileCreateEntry struct {
	Name     *string      `json:"Name,omitempty"`
	Contents *[]byte      `json:"Contents,omitempty"`
	TagIds   *[]uuid.UUID `json:"TagIds,omitempty"`
}

// ImageIDCreateBatch ...
type ImageIDCreateBatch struct {
	Images *[]ImageIDCreateEntry `json:"Images,omitempty"`
	TagIds *[]uuid.UUID          `json:"TagIds,omitempty"`
}

// ImageIDCreateEntry ...
type ImageIDCreateEntry struct {
	ID     *uuid.UUID   `json:"Id,omitempty"`
	TagIds *[]uuid.UUID `json:"TagIds,omitempty"`
}

// ImagePredictionResult result of an image prediction request
type ImagePredictionResult struct {
	autorest.Response `json:"-"`
	ID                *uuid.UUID            `json:"Id,omitempty"`
	Project           *uuid.UUID            `json:"Project,omitempty"`
	Iteration         *uuid.UUID            `json:"Iteration,omitempty"`
	Created           *date.Time            `json:"Created,omitempty"`
	Predictions       *[]ImageTagPrediction `json:"Predictions,omitempty"`
}

// ImageTag ...
type ImageTag struct {
	TagID   *uuid.UUID `json:"TagId,omitempty"`
	Created *date.Time `json:"Created,omitempty"`
}

// ImageTagCreateBatch ...
type ImageTagCreateBatch struct {
	Tags *[]ImageTagCreateEntry `json:"Tags,omitempty"`
}

// ImageTagCreateEntry ...
type ImageTagCreateEntry struct {
	ImageID *uuid.UUID `json:"ImageId,omitempty"`
	TagID   *uuid.UUID `json:"TagId,omitempty"`
}

// ImageTagCreateSummary ...
type ImageTagCreateSummary struct {
	autorest.Response `json:"-"`
	Created           *[]ImageTagCreateEntry `json:"Created,omitempty"`
	Duplicated        *[]ImageTagCreateEntry `json:"Duplicated,omitempty"`
	Exceeded          *[]ImageTagCreateEntry `json:"Exceeded,omitempty"`
}

// ImageTagPrediction ...
type ImageTagPrediction struct {
	TagID       *uuid.UUID `json:"TagId,omitempty"`
	Tag         *string    `json:"Tag,omitempty"`
	Probability *float64   `json:"Probability,omitempty"`
}

// ImageURL ...
type ImageURL struct {
	URL *string `json:"Url,omitempty"`
}

// ImageURLCreateBatch ...
type ImageURLCreateBatch struct {
	Images *[]ImageURLCreateEntry `json:"Images,omitempty"`
	TagIds *[]uuid.UUID           `json:"TagIds,omitempty"`
}

// ImageURLCreateEntry ...
type ImageURLCreateEntry struct {
	URL    *string      `json:"Url,omitempty"`
	TagIds *[]uuid.UUID `json:"TagIds,omitempty"`
}

// Iteration iteration model to be sent over JSON
type Iteration struct {
	autorest.Response `json:"-"`
	// ID - Gets the id of the iteration
	ID *uuid.UUID `json:"Id,omitempty"`
	// Name - Gets or sets the name of the iteration
	Name *string `json:"Name,omitempty"`
	// IsDefault - Gets or sets a value indicating whether the iteration is the default iteration for the project
	IsDefault *bool `json:"IsDefault,omitempty"`
	// Status - Gets the current iteration status
	Status *string `json:"Status,omitempty"`
	// Created - Gets the time this iteration was completed
	Created *date.Time `json:"Created,omitempty"`
	// LastModified - Gets the time this iteration was last modified
	LastModified *date.Time `json:"LastModified,omitempty"`
	// TrainedAt - Gets the time this iteration was last modified
	TrainedAt *date.Time `json:"TrainedAt,omitempty"`
	// ProjectID - Gets the project id of the iteration
	ProjectID *uuid.UUID `json:"ProjectId,omitempty"`
	// Exportable - Whether the iteration can be exported to another format for download
	Exportable *bool `json:"Exportable,omitempty"`
	// DomainID - Get or sets a guid of the domain the iteration has been trained on
	DomainID *uuid.UUID `json:"DomainId,omitempty"`
}

// IterationPerformance represents the detailed performance data for a trained iteration
type IterationPerformance struct {
	autorest.Response `json:"-"`
	// PerTagPerformance - Gets the per-tag performance details for this iteration
	PerTagPerformance *[]TagPerformance `json:"PerTagPerformance,omitempty"`
	// Precision - Gets the precision
	Precision *float64 `json:"Precision,omitempty"`
	// PrecisionStdDeviation - Gets the standard deviation for the precision
	PrecisionStdDeviation *float64 `json:"PrecisionStdDeviation,omitempty"`
	// Recall - Gets the recall
	Recall *float64 `json:"Recall,omitempty"`
	// RecallStdDeviation - Gets the standard deviation for the recall
	RecallStdDeviation *float64 `json:"RecallStdDeviation,omitempty"`
}

// KeyPair ...
type KeyPair struct {
	PrimaryKey   *string `json:"PrimaryKey,omitempty"`
	SecondaryKey *string `json:"SecondaryKey,omitempty"`
}

// ListDomain ...
type ListDomain struct {
	autorest.Response `json:"-"`
	Value             *[]Domain `json:"value,omitempty"`
}

// ListExport ...
type ListExport struct {
	autorest.Response `json:"-"`
	Value             *[]Export `json:"value,omitempty"`
}

// ListImage ...
type ListImage struct {
	autorest.Response `json:"-"`
	Value             *[]Image `json:"value,omitempty"`
}

// ListIteration ...
type ListIteration struct {
	autorest.Response `json:"-"`
	Value             *[]Iteration `json:"value,omitempty"`
}

// ListProject ...
type ListProject struct {
	autorest.Response `json:"-"`
	Value             *[]Project `json:"value,omitempty"`
}

// PerProjectQuota represents a set of quotas for a given project
type PerProjectQuota struct {
	// ProjectID - Gets the project id of the project this quota applies to
	ProjectID *uuid.UUID `json:"ProjectId,omitempty"`
	// Iterations - Gets the iteration quota for the project
	Iterations *Quota `json:"Iterations,omitempty"`
	// Images - Gets the image quota for the project
	Images *Quota `json:"Images,omitempty"`
	// Tags - Gets the tag quota for the project
	Tags *Quota `json:"Tags,omitempty"`
}

// Prediction result of an image classification request
type Prediction struct {
	ID           *uuid.UUID       `json:"Id,omitempty"`
	Project      *uuid.UUID       `json:"Project,omitempty"`
	Iteration    *uuid.UUID       `json:"Iteration,omitempty"`
	Created      *date.Time       `json:"Created,omitempty"`
	Predictions  *[]PredictionTag `json:"Predictions,omitempty"`
	ImageURI     *string          `json:"ImageUri,omitempty"`
	ThumbnailURI *string          `json:"ThumbnailUri,omitempty"`
}

// PredictionQuery ...
type PredictionQuery struct {
	autorest.Response `json:"-"`
	Results           *[]Prediction         `json:"Results,omitempty"`
	Token             *PredictionQueryToken `json:"Token,omitempty"`
}

// PredictionQueryTag ...
type PredictionQueryTag struct {
	ID           *uuid.UUID `json:"Id,omitempty"`
	MinThreshold *float64   `json:"MinThreshold,omitempty"`
	MaxThreshold *float64   `json:"MaxThreshold,omitempty"`
}

// PredictionQueryToken ...
type PredictionQueryToken struct {
	Session      *string `json:"Session,omitempty"`
	Continuation *string `json:"Continuation,omitempty"`
	MaxCount     *int32  `json:"MaxCount,omitempty"`
	// OrderBy - Possible values include: 'Newest', 'Oldest', 'Suggested'
	OrderBy     OrderBy               `json:"OrderBy,omitempty"`
	Tags        *[]PredictionQueryTag `json:"Tags,omitempty"`
	IterationID *uuid.UUID            `json:"IterationId,omitempty"`
	StartTime   *date.Time            `json:"StartTime,omitempty"`
	EndTime     *date.Time            `json:"EndTime,omitempty"`
	Application *string               `json:"Application,omitempty"`
}

// PredictionTag ...
type PredictionTag struct {
	TagID       *uuid.UUID `json:"TagId,omitempty"`
	Tag         *string    `json:"Tag,omitempty"`
	Probability *float64   `json:"Probability,omitempty"`
}

// Project represents a project
type Project struct {
	autorest.Response `json:"-"`
	// ID - Gets the project id
	ID *uuid.UUID `json:"Id,omitempty"`
	// Name - Gets or sets the name of the project
	Name *string `json:"Name,omitempty"`
	// Description - Gets or sets the description of the project
	Description *string `json:"Description,omitempty"`
	// Settings - Gets or sets the project settings
	Settings *ProjectSettings `json:"Settings,omitempty"`
	// CurrentIterationID - Gets the current iteration id
	CurrentIterationID *uuid.UUID `json:"CurrentIterationId,omitempty"`
	// Created - Gets the date this project was created
	Created *date.Time `json:"Created,omitempty"`
	// LastModified - Gets the date this project was last modifed
	LastModified *date.Time `json:"LastModified,omitempty"`
	// ThumbnailURI - Gets the thumbnail url representing the image
	ThumbnailURI *string `json:"ThumbnailUri,omitempty"`
}

// ProjectSettings represents settings associated with a project
type ProjectSettings struct {
	// DomainID - Gets or sets the id of the Domain to use with this project
	DomainID *uuid.UUID `json:"DomainId,omitempty"`
}

// Quota represents a quota
type Quota struct {
	// Total - The total allowable amount in the quota
	Total *int32 `json:"Total,omitempty"`
	// Used - The amount of quota that has currently been used
	Used *int32 `json:"Used,omitempty"`
	// TimeUntilReset - Gets the time remaining until the quota resets. Null if this quota does not reset.
	TimeUntilReset *string `json:"TimeUntilReset,omitempty"`
}

// Tag represents a Tag
type Tag struct {
	autorest.Response `json:"-"`
	// ID - Gets the Tag ID
	ID *uuid.UUID `json:"Id,omitempty"`
	// Name - Gets or sets the name of the tag
	Name *string `json:"Name,omitempty"`
	// Description - Gets or sets the description of the tag
	Description *string `json:"Description,omitempty"`
	// ImageCount - Gets the number of images with this tag
	ImageCount *int32 `json:"ImageCount,omitempty"`
}

// TagList ...
type TagList struct {
	autorest.Response   `json:"-"`
	Tags                *[]Tag `json:"Tags,omitempty"`
	TotalTaggedImages   *int32 `json:"TotalTaggedImages,omitempty"`
	TotalUntaggedImages *int32 `json:"TotalUntaggedImages,omitempty"`
}

// TagPerformance represents performance data for a particular tag in a trained iteration
type TagPerformance struct {
	ID   *uuid.UUID `json:"Id,omitempty"`
	Name *string    `json:"Name,omitempty"`
	// Precision - Gets the precision
	Precision *float64 `json:"Precision,omitempty"`
	// PrecisionStdDeviation - Gets the standard deviation for the precision
	PrecisionStdDeviation *float64 `json:"PrecisionStdDeviation,omitempty"`
	// Recall - Gets the recall
	Recall *float64 `json:"Recall,omitempty"`
	// RecallStdDeviation - Gets the standard deviation for the recall
	RecallStdDeviation *float64 `json:"RecallStdDeviation,omitempty"`
}
