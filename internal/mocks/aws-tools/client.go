// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/tools/aws/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	acm "github.com/aws/aws-sdk-go/service/acm"
	gomock "github.com/golang/mock/gomock"
	aws "github.com/mattermost/mattermost-cloud/internal/tools/aws"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockAWS is a mock of AWS interface
type MockAWS struct {
	ctrl     *gomock.Controller
	recorder *MockAWSMockRecorder
}

// MockAWSMockRecorder is the mock recorder for MockAWS
type MockAWSMockRecorder struct {
	mock *MockAWS
}

// NewMockAWS creates a new mock instance
func NewMockAWS(ctrl *gomock.Controller) *MockAWS {
	mock := &MockAWS{ctrl: ctrl}
	mock.recorder = &MockAWSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAWS) EXPECT() *MockAWSMockRecorder {
	return m.recorder
}

// GetCertificateSummaryByTag mocks base method
func (m *MockAWS) GetCertificateSummaryByTag(key, value string, logger logrus.FieldLogger) (*acm.CertificateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateSummaryByTag", key, value, logger)
	ret0, _ := ret[0].(*acm.CertificateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateSummaryByTag indicates an expected call of GetCertificateSummaryByTag
func (mr *MockAWSMockRecorder) GetCertificateSummaryByTag(key, value, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateSummaryByTag", reflect.TypeOf((*MockAWS)(nil).GetCertificateSummaryByTag), key, value, logger)
}

// GetAndClaimVpcResources mocks base method
func (m *MockAWS) GetAndClaimVpcResources(clusterID, owner string, logger logrus.FieldLogger) (aws.ClusterResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndClaimVpcResources", clusterID, owner, logger)
	ret0, _ := ret[0].(aws.ClusterResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAndClaimVpcResources indicates an expected call of GetAndClaimVpcResources
func (mr *MockAWSMockRecorder) GetAndClaimVpcResources(clusterID, owner, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndClaimVpcResources", reflect.TypeOf((*MockAWS)(nil).GetAndClaimVpcResources), clusterID, owner, logger)
}

// ReleaseVpc mocks base method
func (m *MockAWS) ReleaseVpc(clusterID string, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseVpc", clusterID, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseVpc indicates an expected call of ReleaseVpc
func (mr *MockAWSMockRecorder) ReleaseVpc(clusterID, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseVpc", reflect.TypeOf((*MockAWS)(nil).ReleaseVpc), clusterID, logger)
}

// GetPrivateZoneDomainName mocks base method
func (m *MockAWS) GetPrivateZoneDomainName(logger logrus.FieldLogger) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateZoneDomainName", logger)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateZoneDomainName indicates an expected call of GetPrivateZoneDomainName
func (mr *MockAWSMockRecorder) GetPrivateZoneDomainName(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateZoneDomainName", reflect.TypeOf((*MockAWS)(nil).GetPrivateZoneDomainName), logger)
}

// CreatePrivateCNAME mocks base method
func (m *MockAWS) CreatePrivateCNAME(dnsName string, dnsEndpoints []string, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePrivateCNAME", dnsName, dnsEndpoints, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePrivateCNAME indicates an expected call of CreatePrivateCNAME
func (mr *MockAWSMockRecorder) CreatePrivateCNAME(dnsName, dnsEndpoints, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrivateCNAME", reflect.TypeOf((*MockAWS)(nil).CreatePrivateCNAME), dnsName, dnsEndpoints, logger)
}

// CreatePublicCNAME mocks base method
func (m *MockAWS) CreatePublicCNAME(dnsName string, dnsEndpoints []string, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePublicCNAME", dnsName, dnsEndpoints, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePublicCNAME indicates an expected call of CreatePublicCNAME
func (mr *MockAWSMockRecorder) CreatePublicCNAME(dnsName, dnsEndpoints, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePublicCNAME", reflect.TypeOf((*MockAWS)(nil).CreatePublicCNAME), dnsName, dnsEndpoints, logger)
}

// IsProvisionedPrivateCNAME mocks base method
func (m *MockAWS) IsProvisionedPrivateCNAME(dnsName string, logger logrus.FieldLogger) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProvisionedPrivateCNAME", dnsName, logger)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsProvisionedPrivateCNAME indicates an expected call of IsProvisionedPrivateCNAME
func (mr *MockAWSMockRecorder) IsProvisionedPrivateCNAME(dnsName, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProvisionedPrivateCNAME", reflect.TypeOf((*MockAWS)(nil).IsProvisionedPrivateCNAME), dnsName, logger)
}

// DeletePrivateCNAME mocks base method
func (m *MockAWS) DeletePrivateCNAME(dnsName string, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePrivateCNAME", dnsName, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePrivateCNAME indicates an expected call of DeletePrivateCNAME
func (mr *MockAWSMockRecorder) DeletePrivateCNAME(dnsName, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePrivateCNAME", reflect.TypeOf((*MockAWS)(nil).DeletePrivateCNAME), dnsName, logger)
}

// DeletePublicCNAME mocks base method
func (m *MockAWS) DeletePublicCNAME(dnsName string, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePublicCNAME", dnsName, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePublicCNAME indicates an expected call of DeletePublicCNAME
func (mr *MockAWSMockRecorder) DeletePublicCNAME(dnsName, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePublicCNAME", reflect.TypeOf((*MockAWS)(nil).DeletePublicCNAME), dnsName, logger)
}

// TagResource mocks base method
func (m *MockAWS) TagResource(resourceID, key, value string, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", resourceID, key, value, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagResource indicates an expected call of TagResource
func (mr *MockAWSMockRecorder) TagResource(resourceID, key, value, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockAWS)(nil).TagResource), resourceID, key, value, logger)
}

// UntagResource mocks base method
func (m *MockAWS) UntagResource(resourceID, key, value string, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", resourceID, key, value, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockAWSMockRecorder) UntagResource(resourceID, key, value, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockAWS)(nil).UntagResource), resourceID, key, value, logger)
}

// IsValidAMI mocks base method
func (m *MockAWS) IsValidAMI(AMIImage string, logger logrus.FieldLogger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidAMI", AMIImage, logger)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidAMI indicates an expected call of IsValidAMI
func (mr *MockAWSMockRecorder) IsValidAMI(AMIImage, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidAMI", reflect.TypeOf((*MockAWS)(nil).IsValidAMI), AMIImage, logger)
}
