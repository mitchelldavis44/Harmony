// infrastructure/mock_infrastructure.go
package infrastructure

import "fmt"

type Infrastructure interface {
  CreateResource(name string, instanceType string, imageID string, securityGroupId string, keyPairName string, subnetId string, iamInstanceProfile string, vpcId string) (string, error)
  DeleteResource(name string) error
}

type MockInfrastructure struct {
	Resources map[string]bool
}

func NewMockInfrastructure() *MockInfrastructure {
	return &MockInfrastructure{
		Resources: make(map[string]bool),
	}
}

// Update function signature to match interface
func (m *MockInfrastructure) CreateResource(name string, instanceType string, imageID string, securityGroupId string, keyPairName string, subnetId string, iamInstanceProfile string, vpcId string) (string, error) {
	if _, ok := m.Resources[name]; ok {
		return "", fmt.Errorf("resource already exists: %s", name)
	}
	m.Resources[name] = true
	return name, nil
}

func (m *MockInfrastructure) DeleteResource(name string) error {
	if _, ok := m.Resources[name]; !ok {
		return fmt.Errorf("resource does not exist: %s", name)
	}
	delete(m.Resources, name)
	return nil
}