package user

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vmkevv/suprat-api/config"
	"github.com/vmkevv/suprat-api/ent/enttest"
	"github.com/vmkevv/suprat-api/internal/services"
)

var ctx context.Context = context.Background()

func TestSave(t *testing.T) {
	config.SetEnvs()
	conf, _ := config.GetConfig()
	entClent := enttest.Open(t, "postgres", conf.PostgresConn())
	defer entClent.Close()
	actions := actions{
		db: entClent,
	}
	t.Run("Should save a new user", func(t *testing.T) {
		savedUser, err := actions.Save(ctx, "Kevin", "Vargas", "kevin2@mail.com", "12345")
		if err != nil {
			t.Errorf("Expected err to be nil, but got: %v\n", err)
		}
		if savedUser.ID == 0 {
			t.Errorf("User ID should be other than 0")
		}
	})
}

// AeertErr check the code, message and type on an expected err
func AssertRespStatus(t *testing.T, err error, code int, msg string) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected an error  but got nil")
	}
	supratErr, ok := err.(services.SuprError)
	if !ok {
		t.Fatal("Expected an SuprError type but got another error")
	}
	if supratErr.Code != code {
		t.Errorf("Expected error code: \n\t%d\n But got: \n\t%d\n", code, supratErr.Code)
	}
	if supratErr.Message != msg {
		t.Errorf("Expected error message to be: \n\t%s\n But got: \n\t%s\n", msg, supratErr.Message)
	}
}
