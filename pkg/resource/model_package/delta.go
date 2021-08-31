// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package model_package

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customSetDefaults(a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.ApprovalDescription, b.ko.Spec.ApprovalDescription) {
		delta.Add("Spec.ApprovalDescription", a.ko.Spec.ApprovalDescription, b.ko.Spec.ApprovalDescription)
	} else if a.ko.Spec.ApprovalDescription != nil && b.ko.Spec.ApprovalDescription != nil {
		if *a.ko.Spec.ApprovalDescription != *b.ko.Spec.ApprovalDescription {
			delta.Add("Spec.ApprovalDescription", a.ko.Spec.ApprovalDescription, b.ko.Spec.ApprovalDescription)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CertifyForMarketplace, b.ko.Spec.CertifyForMarketplace) {
		delta.Add("Spec.CertifyForMarketplace", a.ko.Spec.CertifyForMarketplace, b.ko.Spec.CertifyForMarketplace)
	} else if a.ko.Spec.CertifyForMarketplace != nil && b.ko.Spec.CertifyForMarketplace != nil {
		if *a.ko.Spec.CertifyForMarketplace != *b.ko.Spec.CertifyForMarketplace {
			delta.Add("Spec.CertifyForMarketplace", a.ko.Spec.CertifyForMarketplace, b.ko.Spec.CertifyForMarketplace)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ClientToken, b.ko.Spec.ClientToken) {
		delta.Add("Spec.ClientToken", a.ko.Spec.ClientToken, b.ko.Spec.ClientToken)
	} else if a.ko.Spec.ClientToken != nil && b.ko.Spec.ClientToken != nil {
		if *a.ko.Spec.ClientToken != *b.ko.Spec.ClientToken {
			delta.Add("Spec.ClientToken", a.ko.Spec.ClientToken, b.ko.Spec.ClientToken)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.InferenceSpecification, b.ko.Spec.InferenceSpecification) {
		delta.Add("Spec.InferenceSpecification", a.ko.Spec.InferenceSpecification, b.ko.Spec.InferenceSpecification)
	} else if a.ko.Spec.InferenceSpecification != nil && b.ko.Spec.InferenceSpecification != nil {
		if !reflect.DeepEqual(a.ko.Spec.InferenceSpecification.Containers, b.ko.Spec.InferenceSpecification.Containers) {
			delta.Add("Spec.InferenceSpecification.Containers", a.ko.Spec.InferenceSpecification.Containers, b.ko.Spec.InferenceSpecification.Containers)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.InferenceSpecification.SupportedContentTypes, b.ko.Spec.InferenceSpecification.SupportedContentTypes) {
			delta.Add("Spec.InferenceSpecification.SupportedContentTypes", a.ko.Spec.InferenceSpecification.SupportedContentTypes, b.ko.Spec.InferenceSpecification.SupportedContentTypes)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.InferenceSpecification.SupportedRealtimeInferenceInstanceTypes, b.ko.Spec.InferenceSpecification.SupportedRealtimeInferenceInstanceTypes) {
			delta.Add("Spec.InferenceSpecification.SupportedRealtimeInferenceInstanceTypes", a.ko.Spec.InferenceSpecification.SupportedRealtimeInferenceInstanceTypes, b.ko.Spec.InferenceSpecification.SupportedRealtimeInferenceInstanceTypes)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.InferenceSpecification.SupportedResponseMIMETypes, b.ko.Spec.InferenceSpecification.SupportedResponseMIMETypes) {
			delta.Add("Spec.InferenceSpecification.SupportedResponseMIMETypes", a.ko.Spec.InferenceSpecification.SupportedResponseMIMETypes, b.ko.Spec.InferenceSpecification.SupportedResponseMIMETypes)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.InferenceSpecification.SupportedTransformInstanceTypes, b.ko.Spec.InferenceSpecification.SupportedTransformInstanceTypes) {
			delta.Add("Spec.InferenceSpecification.SupportedTransformInstanceTypes", a.ko.Spec.InferenceSpecification.SupportedTransformInstanceTypes, b.ko.Spec.InferenceSpecification.SupportedTransformInstanceTypes)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MetadataProperties, b.ko.Spec.MetadataProperties) {
		delta.Add("Spec.MetadataProperties", a.ko.Spec.MetadataProperties, b.ko.Spec.MetadataProperties)
	} else if a.ko.Spec.MetadataProperties != nil && b.ko.Spec.MetadataProperties != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.MetadataProperties.CommitID, b.ko.Spec.MetadataProperties.CommitID) {
			delta.Add("Spec.MetadataProperties.CommitID", a.ko.Spec.MetadataProperties.CommitID, b.ko.Spec.MetadataProperties.CommitID)
		} else if a.ko.Spec.MetadataProperties.CommitID != nil && b.ko.Spec.MetadataProperties.CommitID != nil {
			if *a.ko.Spec.MetadataProperties.CommitID != *b.ko.Spec.MetadataProperties.CommitID {
				delta.Add("Spec.MetadataProperties.CommitID", a.ko.Spec.MetadataProperties.CommitID, b.ko.Spec.MetadataProperties.CommitID)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.MetadataProperties.GeneratedBy, b.ko.Spec.MetadataProperties.GeneratedBy) {
			delta.Add("Spec.MetadataProperties.GeneratedBy", a.ko.Spec.MetadataProperties.GeneratedBy, b.ko.Spec.MetadataProperties.GeneratedBy)
		} else if a.ko.Spec.MetadataProperties.GeneratedBy != nil && b.ko.Spec.MetadataProperties.GeneratedBy != nil {
			if *a.ko.Spec.MetadataProperties.GeneratedBy != *b.ko.Spec.MetadataProperties.GeneratedBy {
				delta.Add("Spec.MetadataProperties.GeneratedBy", a.ko.Spec.MetadataProperties.GeneratedBy, b.ko.Spec.MetadataProperties.GeneratedBy)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.MetadataProperties.ProjectID, b.ko.Spec.MetadataProperties.ProjectID) {
			delta.Add("Spec.MetadataProperties.ProjectID", a.ko.Spec.MetadataProperties.ProjectID, b.ko.Spec.MetadataProperties.ProjectID)
		} else if a.ko.Spec.MetadataProperties.ProjectID != nil && b.ko.Spec.MetadataProperties.ProjectID != nil {
			if *a.ko.Spec.MetadataProperties.ProjectID != *b.ko.Spec.MetadataProperties.ProjectID {
				delta.Add("Spec.MetadataProperties.ProjectID", a.ko.Spec.MetadataProperties.ProjectID, b.ko.Spec.MetadataProperties.ProjectID)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.MetadataProperties.Repository, b.ko.Spec.MetadataProperties.Repository) {
			delta.Add("Spec.MetadataProperties.Repository", a.ko.Spec.MetadataProperties.Repository, b.ko.Spec.MetadataProperties.Repository)
		} else if a.ko.Spec.MetadataProperties.Repository != nil && b.ko.Spec.MetadataProperties.Repository != nil {
			if *a.ko.Spec.MetadataProperties.Repository != *b.ko.Spec.MetadataProperties.Repository {
				delta.Add("Spec.MetadataProperties.Repository", a.ko.Spec.MetadataProperties.Repository, b.ko.Spec.MetadataProperties.Repository)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelApprovalStatus, b.ko.Spec.ModelApprovalStatus) {
		delta.Add("Spec.ModelApprovalStatus", a.ko.Spec.ModelApprovalStatus, b.ko.Spec.ModelApprovalStatus)
	} else if a.ko.Spec.ModelApprovalStatus != nil && b.ko.Spec.ModelApprovalStatus != nil {
		if *a.ko.Spec.ModelApprovalStatus != *b.ko.Spec.ModelApprovalStatus {
			delta.Add("Spec.ModelApprovalStatus", a.ko.Spec.ModelApprovalStatus, b.ko.Spec.ModelApprovalStatus)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics, b.ko.Spec.ModelMetrics) {
		delta.Add("Spec.ModelMetrics", a.ko.Spec.ModelMetrics, b.ko.Spec.ModelMetrics)
	} else if a.ko.Spec.ModelMetrics != nil && b.ko.Spec.ModelMetrics != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Bias, b.ko.Spec.ModelMetrics.Bias) {
			delta.Add("Spec.ModelMetrics.Bias", a.ko.Spec.ModelMetrics.Bias, b.ko.Spec.ModelMetrics.Bias)
		} else if a.ko.Spec.ModelMetrics.Bias != nil && b.ko.Spec.ModelMetrics.Bias != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Bias.Report, b.ko.Spec.ModelMetrics.Bias.Report) {
				delta.Add("Spec.ModelMetrics.Bias.Report", a.ko.Spec.ModelMetrics.Bias.Report, b.ko.Spec.ModelMetrics.Bias.Report)
			} else if a.ko.Spec.ModelMetrics.Bias.Report != nil && b.ko.Spec.ModelMetrics.Bias.Report != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Bias.Report.ContentDigest, b.ko.Spec.ModelMetrics.Bias.Report.ContentDigest) {
					delta.Add("Spec.ModelMetrics.Bias.Report.ContentDigest", a.ko.Spec.ModelMetrics.Bias.Report.ContentDigest, b.ko.Spec.ModelMetrics.Bias.Report.ContentDigest)
				} else if a.ko.Spec.ModelMetrics.Bias.Report.ContentDigest != nil && b.ko.Spec.ModelMetrics.Bias.Report.ContentDigest != nil {
					if *a.ko.Spec.ModelMetrics.Bias.Report.ContentDigest != *b.ko.Spec.ModelMetrics.Bias.Report.ContentDigest {
						delta.Add("Spec.ModelMetrics.Bias.Report.ContentDigest", a.ko.Spec.ModelMetrics.Bias.Report.ContentDigest, b.ko.Spec.ModelMetrics.Bias.Report.ContentDigest)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Bias.Report.ContentType, b.ko.Spec.ModelMetrics.Bias.Report.ContentType) {
					delta.Add("Spec.ModelMetrics.Bias.Report.ContentType", a.ko.Spec.ModelMetrics.Bias.Report.ContentType, b.ko.Spec.ModelMetrics.Bias.Report.ContentType)
				} else if a.ko.Spec.ModelMetrics.Bias.Report.ContentType != nil && b.ko.Spec.ModelMetrics.Bias.Report.ContentType != nil {
					if *a.ko.Spec.ModelMetrics.Bias.Report.ContentType != *b.ko.Spec.ModelMetrics.Bias.Report.ContentType {
						delta.Add("Spec.ModelMetrics.Bias.Report.ContentType", a.ko.Spec.ModelMetrics.Bias.Report.ContentType, b.ko.Spec.ModelMetrics.Bias.Report.ContentType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Bias.Report.S3URI, b.ko.Spec.ModelMetrics.Bias.Report.S3URI) {
					delta.Add("Spec.ModelMetrics.Bias.Report.S3URI", a.ko.Spec.ModelMetrics.Bias.Report.S3URI, b.ko.Spec.ModelMetrics.Bias.Report.S3URI)
				} else if a.ko.Spec.ModelMetrics.Bias.Report.S3URI != nil && b.ko.Spec.ModelMetrics.Bias.Report.S3URI != nil {
					if *a.ko.Spec.ModelMetrics.Bias.Report.S3URI != *b.ko.Spec.ModelMetrics.Bias.Report.S3URI {
						delta.Add("Spec.ModelMetrics.Bias.Report.S3URI", a.ko.Spec.ModelMetrics.Bias.Report.S3URI, b.ko.Spec.ModelMetrics.Bias.Report.S3URI)
					}
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Explainability, b.ko.Spec.ModelMetrics.Explainability) {
			delta.Add("Spec.ModelMetrics.Explainability", a.ko.Spec.ModelMetrics.Explainability, b.ko.Spec.ModelMetrics.Explainability)
		} else if a.ko.Spec.ModelMetrics.Explainability != nil && b.ko.Spec.ModelMetrics.Explainability != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Explainability.Report, b.ko.Spec.ModelMetrics.Explainability.Report) {
				delta.Add("Spec.ModelMetrics.Explainability.Report", a.ko.Spec.ModelMetrics.Explainability.Report, b.ko.Spec.ModelMetrics.Explainability.Report)
			} else if a.ko.Spec.ModelMetrics.Explainability.Report != nil && b.ko.Spec.ModelMetrics.Explainability.Report != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest, b.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest) {
					delta.Add("Spec.ModelMetrics.Explainability.Report.ContentDigest", a.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest, b.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest)
				} else if a.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest != nil && b.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest != nil {
					if *a.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest != *b.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest {
						delta.Add("Spec.ModelMetrics.Explainability.Report.ContentDigest", a.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest, b.ko.Spec.ModelMetrics.Explainability.Report.ContentDigest)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Explainability.Report.ContentType, b.ko.Spec.ModelMetrics.Explainability.Report.ContentType) {
					delta.Add("Spec.ModelMetrics.Explainability.Report.ContentType", a.ko.Spec.ModelMetrics.Explainability.Report.ContentType, b.ko.Spec.ModelMetrics.Explainability.Report.ContentType)
				} else if a.ko.Spec.ModelMetrics.Explainability.Report.ContentType != nil && b.ko.Spec.ModelMetrics.Explainability.Report.ContentType != nil {
					if *a.ko.Spec.ModelMetrics.Explainability.Report.ContentType != *b.ko.Spec.ModelMetrics.Explainability.Report.ContentType {
						delta.Add("Spec.ModelMetrics.Explainability.Report.ContentType", a.ko.Spec.ModelMetrics.Explainability.Report.ContentType, b.ko.Spec.ModelMetrics.Explainability.Report.ContentType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.Explainability.Report.S3URI, b.ko.Spec.ModelMetrics.Explainability.Report.S3URI) {
					delta.Add("Spec.ModelMetrics.Explainability.Report.S3URI", a.ko.Spec.ModelMetrics.Explainability.Report.S3URI, b.ko.Spec.ModelMetrics.Explainability.Report.S3URI)
				} else if a.ko.Spec.ModelMetrics.Explainability.Report.S3URI != nil && b.ko.Spec.ModelMetrics.Explainability.Report.S3URI != nil {
					if *a.ko.Spec.ModelMetrics.Explainability.Report.S3URI != *b.ko.Spec.ModelMetrics.Explainability.Report.S3URI {
						delta.Add("Spec.ModelMetrics.Explainability.Report.S3URI", a.ko.Spec.ModelMetrics.Explainability.Report.S3URI, b.ko.Spec.ModelMetrics.Explainability.Report.S3URI)
					}
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality, b.ko.Spec.ModelMetrics.ModelDataQuality) {
			delta.Add("Spec.ModelMetrics.ModelDataQuality", a.ko.Spec.ModelMetrics.ModelDataQuality, b.ko.Spec.ModelMetrics.ModelDataQuality)
		} else if a.ko.Spec.ModelMetrics.ModelDataQuality != nil && b.ko.Spec.ModelMetrics.ModelDataQuality != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints) {
				delta.Add("Spec.ModelMetrics.ModelDataQuality.Constraints", a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints)
			} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest) {
					delta.Add("Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest", a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest)
				} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest != nil {
					if *a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest != *b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest {
						delta.Add("Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest", a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentDigest)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType) {
					delta.Add("Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType", a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType)
				} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType != nil {
					if *a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType != *b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType {
						delta.Add("Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType", a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.ContentType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI) {
					delta.Add("Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI", a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI)
				} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI != nil {
					if *a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI != *b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI {
						delta.Add("Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI", a.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI, b.ko.Spec.ModelMetrics.ModelDataQuality.Constraints.S3URI)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics) {
				delta.Add("Spec.ModelMetrics.ModelDataQuality.Statistics", a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics)
			} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest) {
					delta.Add("Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest", a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest)
				} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest != nil {
					if *a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest != *b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest {
						delta.Add("Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest", a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentDigest)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType) {
					delta.Add("Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType", a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType)
				} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType != nil {
					if *a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType != *b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType {
						delta.Add("Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType", a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.ContentType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI) {
					delta.Add("Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI", a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI)
				} else if a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI != nil && b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI != nil {
					if *a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI != *b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI {
						delta.Add("Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI", a.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI, b.ko.Spec.ModelMetrics.ModelDataQuality.Statistics.S3URI)
					}
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality, b.ko.Spec.ModelMetrics.ModelQuality) {
			delta.Add("Spec.ModelMetrics.ModelQuality", a.ko.Spec.ModelMetrics.ModelQuality, b.ko.Spec.ModelMetrics.ModelQuality)
		} else if a.ko.Spec.ModelMetrics.ModelQuality != nil && b.ko.Spec.ModelMetrics.ModelQuality != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Constraints, b.ko.Spec.ModelMetrics.ModelQuality.Constraints) {
				delta.Add("Spec.ModelMetrics.ModelQuality.Constraints", a.ko.Spec.ModelMetrics.ModelQuality.Constraints, b.ko.Spec.ModelMetrics.ModelQuality.Constraints)
			} else if a.ko.Spec.ModelMetrics.ModelQuality.Constraints != nil && b.ko.Spec.ModelMetrics.ModelQuality.Constraints != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest) {
					delta.Add("Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest", a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest)
				} else if a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest != nil && b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest != nil {
					if *a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest != *b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest {
						delta.Add("Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest", a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentDigest)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType) {
					delta.Add("Spec.ModelMetrics.ModelQuality.Constraints.ContentType", a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType)
				} else if a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType != nil && b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType != nil {
					if *a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType != *b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType {
						delta.Add("Spec.ModelMetrics.ModelQuality.Constraints.ContentType", a.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.ContentType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI) {
					delta.Add("Spec.ModelMetrics.ModelQuality.Constraints.S3URI", a.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI)
				} else if a.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI != nil && b.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI != nil {
					if *a.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI != *b.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI {
						delta.Add("Spec.ModelMetrics.ModelQuality.Constraints.S3URI", a.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI, b.ko.Spec.ModelMetrics.ModelQuality.Constraints.S3URI)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Statistics, b.ko.Spec.ModelMetrics.ModelQuality.Statistics) {
				delta.Add("Spec.ModelMetrics.ModelQuality.Statistics", a.ko.Spec.ModelMetrics.ModelQuality.Statistics, b.ko.Spec.ModelMetrics.ModelQuality.Statistics)
			} else if a.ko.Spec.ModelMetrics.ModelQuality.Statistics != nil && b.ko.Spec.ModelMetrics.ModelQuality.Statistics != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest) {
					delta.Add("Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest", a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest)
				} else if a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest != nil && b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest != nil {
					if *a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest != *b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest {
						delta.Add("Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest", a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentDigest)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType) {
					delta.Add("Spec.ModelMetrics.ModelQuality.Statistics.ContentType", a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType)
				} else if a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType != nil && b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType != nil {
					if *a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType != *b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType {
						delta.Add("Spec.ModelMetrics.ModelQuality.Statistics.ContentType", a.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.ContentType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI) {
					delta.Add("Spec.ModelMetrics.ModelQuality.Statistics.S3URI", a.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI)
				} else if a.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI != nil && b.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI != nil {
					if *a.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI != *b.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI {
						delta.Add("Spec.ModelMetrics.ModelQuality.Statistics.S3URI", a.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI, b.ko.Spec.ModelMetrics.ModelQuality.Statistics.S3URI)
					}
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelPackageDescription, b.ko.Spec.ModelPackageDescription) {
		delta.Add("Spec.ModelPackageDescription", a.ko.Spec.ModelPackageDescription, b.ko.Spec.ModelPackageDescription)
	} else if a.ko.Spec.ModelPackageDescription != nil && b.ko.Spec.ModelPackageDescription != nil {
		if *a.ko.Spec.ModelPackageDescription != *b.ko.Spec.ModelPackageDescription {
			delta.Add("Spec.ModelPackageDescription", a.ko.Spec.ModelPackageDescription, b.ko.Spec.ModelPackageDescription)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelPackageGroupName, b.ko.Spec.ModelPackageGroupName) {
		delta.Add("Spec.ModelPackageGroupName", a.ko.Spec.ModelPackageGroupName, b.ko.Spec.ModelPackageGroupName)
	} else if a.ko.Spec.ModelPackageGroupName != nil && b.ko.Spec.ModelPackageGroupName != nil {
		if *a.ko.Spec.ModelPackageGroupName != *b.ko.Spec.ModelPackageGroupName {
			delta.Add("Spec.ModelPackageGroupName", a.ko.Spec.ModelPackageGroupName, b.ko.Spec.ModelPackageGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelPackageName, b.ko.Spec.ModelPackageName) {
		delta.Add("Spec.ModelPackageName", a.ko.Spec.ModelPackageName, b.ko.Spec.ModelPackageName)
	} else if a.ko.Spec.ModelPackageName != nil && b.ko.Spec.ModelPackageName != nil {
		if *a.ko.Spec.ModelPackageName != *b.ko.Spec.ModelPackageName {
			delta.Add("Spec.ModelPackageName", a.ko.Spec.ModelPackageName, b.ko.Spec.ModelPackageName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SourceAlgorithmSpecification, b.ko.Spec.SourceAlgorithmSpecification) {
		delta.Add("Spec.SourceAlgorithmSpecification", a.ko.Spec.SourceAlgorithmSpecification, b.ko.Spec.SourceAlgorithmSpecification)
	} else if a.ko.Spec.SourceAlgorithmSpecification != nil && b.ko.Spec.SourceAlgorithmSpecification != nil {
		if !reflect.DeepEqual(a.ko.Spec.SourceAlgorithmSpecification.SourceAlgorithms, b.ko.Spec.SourceAlgorithmSpecification.SourceAlgorithms) {
			delta.Add("Spec.SourceAlgorithmSpecification.SourceAlgorithms", a.ko.Spec.SourceAlgorithmSpecification.SourceAlgorithms, b.ko.Spec.SourceAlgorithmSpecification.SourceAlgorithms)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ValidationSpecification, b.ko.Spec.ValidationSpecification) {
		delta.Add("Spec.ValidationSpecification", a.ko.Spec.ValidationSpecification, b.ko.Spec.ValidationSpecification)
	} else if a.ko.Spec.ValidationSpecification != nil && b.ko.Spec.ValidationSpecification != nil {
		if !reflect.DeepEqual(a.ko.Spec.ValidationSpecification.ValidationProfiles, b.ko.Spec.ValidationSpecification.ValidationProfiles) {
			delta.Add("Spec.ValidationSpecification.ValidationProfiles", a.ko.Spec.ValidationSpecification.ValidationProfiles, b.ko.Spec.ValidationSpecification.ValidationProfiles)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ValidationSpecification.ValidationRole, b.ko.Spec.ValidationSpecification.ValidationRole) {
			delta.Add("Spec.ValidationSpecification.ValidationRole", a.ko.Spec.ValidationSpecification.ValidationRole, b.ko.Spec.ValidationSpecification.ValidationRole)
		} else if a.ko.Spec.ValidationSpecification.ValidationRole != nil && b.ko.Spec.ValidationSpecification.ValidationRole != nil {
			if *a.ko.Spec.ValidationSpecification.ValidationRole != *b.ko.Spec.ValidationSpecification.ValidationRole {
				delta.Add("Spec.ValidationSpecification.ValidationRole", a.ko.Spec.ValidationSpecification.ValidationRole, b.ko.Spec.ValidationSpecification.ValidationRole)
			}
		}
	}

	return delta
}
