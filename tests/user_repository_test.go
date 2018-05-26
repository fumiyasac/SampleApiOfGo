package test

import (
	"testing"
	"time"

	"github.com/fumiyasac/SampleApi/constants"
	"github.com/fumiyasac/SampleApi/factories"
	"github.com/fumiyasac/SampleApi/repositories"
)

// MockUserRepository ... UserRepositoryのMock実装
type MockUserRepository struct {
	repositories.UserRepository
}

// NewMockUserRepository ... 構造体の新規作成メソッド
func NewMockUserRepository() MockUserRepository {
	return MockUserRepository{}
}

// GetLists ... UserRepositoryのMock実装(UserRepositoryに定義したGetListsメソッドの肩代わり)
func (mock *MockUserRepository) GetLists() (factories.UserItemsFactory, bool) {

	var userFactories []factories.SingleUserFactory
	var userItemsFactory factories.UserItemsFactory

	userFactories = []factories.SingleUserFactory{
		factories.SingleUserFactory{
			ID:          1,
			Username:    "testuser01",
			MailAddress: "test01@example.com",
			Password:    "$2a$10$9KMWYY/m8IRzaU4FEQXY1OBaCJ0RsRWATsmGAk0JaniuWB11dveDe",
			UserStatus:  constants.GetUserStatusNameFromStatusCode(constants.UserSubscribed.GetRawValue()),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		factories.SingleUserFactory{
			ID:          2,
			Username:    "testuser02",
			MailAddress: "testaddress02@example.com",
			Password:    "$2a$10$QKSwHPYN0nSEBVLbDq1nIuoRLDlVrMidWloGh77.iZ1OioXnOXGPy",
			UserStatus:  constants.GetUserStatusNameFromStatusCode(constants.UserSubscribed.GetRawValue()),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	userItemsFactory = factories.UserItemsFactory{
		Users: userFactories,
	}

	return userItemsFactory, true
}

// TestCaseGetLists ... ユーザー一覧取得のテスト
func TestCaseGetLists(t *testing.T) {
	var mock = NewMockUserRepository()
	factory, actural := mock.GetLists()
	expected := true

	if actural != expected {
		t.Errorf("Case 期待する値:(%v), 実際の値:(%v), データ:(%v)", expected, actural, factory)
	}
}
