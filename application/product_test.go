package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/gabao55/hexagonal-arch-go/application"
	"github.com/google/uuid"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product {}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product {}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product {}
	product.ID = uuid.New().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal to zero", err.Error())

	product.Price = 10
	product.ID = "invalid"
	_, err = product.IsValid()
	require.NotNil(t, err)
	
	product.ID = uuid.New().String()
	product.Name = ""
	_, err = product.IsValid()
	require.NotNil(t, err)
	
	product.Name = "Hello"
	_, err = product.IsValid()
	require.Nil(t, err)
}

func TestProduct_Getters(t *testing.T) {
	product := application.Product {}
	id := uuid.New().String()
	name := "Hello"
	price := 10.0
	status := application.ENABLED
	product.ID = id
	product.Name = name
	product.Price = price
	product.Status = status

	require.Equal(t, id, product.GetId())
	require.Equal(t, name, product.GetName())
	require.Equal(t, price, product.GetPrice())
	require.Equal(t, status, product.GetStatus())
}