package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/golang/mock/gomock"
	"github.com/gabao55/hexagonal-arch-go/application"
	"github.com/gabao55/hexagonal-arch-go/mock_application"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)
	
	service := application.ProductService{
		persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(nil, nil).AnyTimes()

	service := application.ProductService{
		persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(nil, nil).AnyTimes()

	service := application.ProductService{
		persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}