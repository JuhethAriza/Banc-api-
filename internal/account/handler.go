package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service ServiceAccount
}

func (a *AccountHandler) PostAccount(c *gin.Context) {
	var account Account
	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de cuenta inválidos"})
		return
	}

	account = Account{
		UserID:       account.UserID,
		TipoCuenta:   account.TipoCuenta,
		NumeroCuenta: account.NumeroCuenta,
		Token:        account.Token,
	}

	if account.TipoCuenta == "" || account.NumeroCuenta == "" || account.UserID == 0 || account.Token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de cuenta incompletos"})
		return
	} else {
		resp := ResponseAccount{
			ID:         account.ID,
			UserID:     account.UserID,
			TipoCuenta: account.TipoCuenta,
			Saldo:      account.Saldo,
			Estado:     account.Estado,
		}
		c.JSON(http.StatusOK, gin.H{"message": "Cuenta creada exitosamente", "account": resp})
		return
	}
}

// putch account

func (a *AccountHandler) PutchAccount(c *gin.Context, id uint) {
	var account Account
	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de cuenta inválidos"})
		return
	}

	account = Account{
		UserID:       account.UserID,
		TipoCuenta:   account.TipoCuenta,
		NumeroCuenta: account.NumeroCuenta,
		Token:        account.Token,
	}
}
