// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// Code generated by MockGen. DO NOT EDIT.
// Source: ../interfaces.go

// Package mock_internal is a generated GoMock package.
package mock_internal

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
)

// MockKubeClient is a mock of KubeClient interface
type MockKubeClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubeClientMockRecorder
}

// MockKubeClientMockRecorder is the mock recorder for MockKubeClient
type MockKubeClientMockRecorder struct {
	mock *MockKubeClient
}

// NewMockKubeClient creates a new mock instance
func NewMockKubeClient(ctrl *gomock.Controller) *MockKubeClient {
	mock := &MockKubeClient{ctrl: ctrl}
	mock.recorder = &MockKubeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeClient) EXPECT() *MockKubeClientMockRecorder {
	return m.recorder
}

// ListPods mocks base method
func (m *MockKubeClient) ListPods(namespace string, opts v11.ListOptions) (*v10.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPods", namespace, opts)
	ret0, _ := ret[0].(*v10.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPods indicates an expected call of ListPods
func (mr *MockKubeClientMockRecorder) ListPods(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPods", reflect.TypeOf((*MockKubeClient)(nil).ListPods), namespace, opts)
}

// ListNodes mocks base method
func (m *MockKubeClient) ListNodes() (*v10.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes")
	ret0, _ := ret[0].(*v10.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes
func (mr *MockKubeClientMockRecorder) ListNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockKubeClient)(nil).ListNodes))
}

// ListServiceAccounts mocks base method
func (m *MockKubeClient) ListServiceAccounts(namespace string, opts v11.ListOptions) (*v10.ServiceAccountList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceAccounts", namespace, opts)
	ret0, _ := ret[0].(*v10.ServiceAccountList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServiceAccounts indicates an expected call of ListServiceAccounts
func (mr *MockKubeClientMockRecorder) ListServiceAccounts(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceAccounts", reflect.TypeOf((*MockKubeClient)(nil).ListServiceAccounts), namespace, opts)
}

// ListDeployments mocks base method
func (m *MockKubeClient) ListDeployments(namespace string, opts v11.ListOptions) (*v1.DeploymentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployments", namespace, opts)
	ret0, _ := ret[0].(*v1.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployments indicates an expected call of ListDeployments
func (mr *MockKubeClientMockRecorder) ListDeployments(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockKubeClient)(nil).ListDeployments), namespace, opts)
}

// ListDaemonSets mocks base method
func (m *MockKubeClient) ListDaemonSets(namespace string, opts v11.ListOptions) (*v1.DaemonSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDaemonSets", namespace, opts)
	ret0, _ := ret[0].(*v1.DaemonSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDaemonSets indicates an expected call of ListDaemonSets
func (mr *MockKubeClientMockRecorder) ListDaemonSets(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDaemonSets", reflect.TypeOf((*MockKubeClient)(nil).ListDaemonSets), namespace, opts)
}

// GetDeployment mocks base method
func (m *MockKubeClient) GetDeployment(namespace, name string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", namespace, name)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment
func (mr *MockKubeClientMockRecorder) GetDeployment(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockKubeClient)(nil).GetDeployment), namespace, name)
}

// PatchDeployment mocks base method
func (m *MockKubeClient) PatchDeployment(namespace, name, jsonPatch string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchDeployment", namespace, name, jsonPatch)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchDeployment indicates an expected call of PatchDeployment
func (mr *MockKubeClientMockRecorder) PatchDeployment(namespace, name, jsonPatch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDeployment", reflect.TypeOf((*MockKubeClient)(nil).PatchDeployment), namespace, name, jsonPatch)
}

// PatchDaemonSet mocks base method
func (m *MockKubeClient) PatchDaemonSet(namespace, name, jsonPatch string) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchDaemonSet", namespace, name, jsonPatch)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchDaemonSet indicates an expected call of PatchDaemonSet
func (mr *MockKubeClientMockRecorder) PatchDaemonSet(namespace, name, jsonPatch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDaemonSet", reflect.TypeOf((*MockKubeClient)(nil).PatchDaemonSet), namespace, name, jsonPatch)
}

// DeletePods mocks base method
func (m *MockKubeClient) DeletePods(namespace string, opts v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePods", namespace, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePods indicates an expected call of DeletePods
func (mr *MockKubeClientMockRecorder) DeletePods(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePods", reflect.TypeOf((*MockKubeClient)(nil).DeletePods), namespace, opts)
}

// DeleteServiceAccount mocks base method
func (m *MockKubeClient) DeleteServiceAccount(secret *v10.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceAccount", secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccount indicates an expected call of DeleteServiceAccount
func (mr *MockKubeClientMockRecorder) DeleteServiceAccount(secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccount", reflect.TypeOf((*MockKubeClient)(nil).DeleteServiceAccount), secret)
}

// DeleteSecret mocks base method
func (m *MockKubeClient) DeleteSecret(secret *v10.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockKubeClientMockRecorder) DeleteSecret(secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockKubeClient)(nil).DeleteSecret), secret)
}

// MockARMClient is a mock of ARMClient interface
type MockARMClient struct {
	ctrl     *gomock.Controller
	recorder *MockARMClientMockRecorder
}

// MockARMClientMockRecorder is the mock recorder for MockARMClient
type MockARMClientMockRecorder struct {
	mock *MockARMClient
}

// NewMockARMClient creates a new mock instance
func NewMockARMClient(ctrl *gomock.Controller) *MockARMClient {
	mock := &MockARMClient{ctrl: ctrl}
	mock.recorder = &MockARMClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockARMClient) EXPECT() *MockARMClientMockRecorder {
	return m.recorder
}

// RestartVirtualMachine mocks base method
func (m *MockARMClient) RestartVirtualMachine(resourceGroup, vmName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartVirtualMachine", resourceGroup, vmName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestartVirtualMachine indicates an expected call of RestartVirtualMachine
func (mr *MockARMClientMockRecorder) RestartVirtualMachine(resourceGroup, vmName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartVirtualMachine", reflect.TypeOf((*MockARMClient)(nil).RestartVirtualMachine), resourceGroup, vmName)
}

// RestartVirtualMachineScaleSets mocks base method
func (m *MockARMClient) RestartVirtualMachineScaleSets(resourceGroup, vmssName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartVirtualMachineScaleSets", resourceGroup, vmssName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestartVirtualMachineScaleSets indicates an expected call of RestartVirtualMachineScaleSets
func (mr *MockARMClientMockRecorder) RestartVirtualMachineScaleSets(resourceGroup, vmssName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartVirtualMachineScaleSets", reflect.TypeOf((*MockARMClient)(nil).RestartVirtualMachineScaleSets), resourceGroup, vmssName)
}

// GetVirtualMachinePowerState mocks base method
func (m *MockARMClient) GetVirtualMachinePowerState(resourceGroup, vmName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachinePowerState", resourceGroup, vmName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachinePowerState indicates an expected call of GetVirtualMachinePowerState
func (mr *MockARMClientMockRecorder) GetVirtualMachinePowerState(resourceGroup, vmName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachinePowerState", reflect.TypeOf((*MockARMClient)(nil).GetVirtualMachinePowerState), resourceGroup, vmName)
}

// GetVirtualMachineScaleSetInstancePowerState mocks base method
func (m *MockARMClient) GetVirtualMachineScaleSetInstancePowerState(resourceGroup, vmssName, instanceID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineScaleSetInstancePowerState", resourceGroup, vmssName, instanceID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineScaleSetInstancePowerState indicates an expected call of GetVirtualMachineScaleSetInstancePowerState
func (mr *MockARMClientMockRecorder) GetVirtualMachineScaleSetInstancePowerState(resourceGroup, vmssName, instanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineScaleSetInstancePowerState", reflect.TypeOf((*MockARMClient)(nil).GetVirtualMachineScaleSetInstancePowerState), resourceGroup, vmssName, instanceID)
}