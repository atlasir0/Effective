package app

import (
	handlers "Effective_Mobile/internal/handlers_api"
	"Effective_Mobile/internal/services"
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	Log        *slog.Logger
	HTTPServer *http.Server
	Router     *mux.Router
}

func New(log *slog.Logger, port int, userService *services.UserService, worklogService *services.WorklogService) *App {
	app := &App{
		Log: log,
	}
	app.Router = mux.NewRouter()
	handlers.SetupRoutes(app.Router, userService, worklogService)
	app.HTTPServer = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: app.Router,
	}
	return app
}

func (a *App) MustRun() {
	a.Log.Info("starting HTTP server", slog.String("addr", a.HTTPServer.Addr))
	if err := a.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.Log.Error("failed to start HTTP server", slog.String("error", err.Error()))
		panic(err)
	}
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.HTTPServer.Shutdown(ctx); err != nil {
		a.Log.Error("failed to shutdown HTTP server", slog.String("error", err.Error()))
	} else {
		a.Log.Info("HTTP server stopped gracefully")
	}
}


package validate

import (
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestValidateForm1(t *testing.T) {
 // Позитивный тест
 t.Run("Positive Test for Form1", func(t *testing.T) {
  name := "John Doe"
  age := 30
  email := "john.doe@example.com"
  form := Form1{
   User: User{
    Name: &name,
    Age:  &age,
   },
   Email: &email,
  }
  assert.True(t, ValidateForm1(form), "ValidateForm1() should return true")
 })

 // Негативные тесты
 t.Run("Negative Test for Form1 - Missing Name", func(t *testing.T) {
  age := 30
  email := "john.doe@example.com"
  form := Form1{
   User: User{
    Name: nil,
    Age:  &age,
   },
   Email: &email,
  }
  assert.False(t, ValidateForm1(form), "ValidateForm1() should return false")
 })

 t.Run("Negative Test for Form1 - Missing Age", func(t *testing.T) {
  name := "John Doe"
  email := "john.doe@example.com"
  form := Form1{
   User: User{
    Name: &name,
    Age:  nil,
   },
   Email: &email,
  }
  assert.False(t, ValidateForm1(form), "ValidateForm1() should return false")
 })

 t.Run("Negative Test for Form1 - Missing Email", func(t *testing.T) {
  name := "John Doe"
  age := 30
  form := Form1{
   User: User{
    Name: &name,
    Age:  &age,
   },
   Email: nil,
  }
  assert.False(t, ValidateForm1(form), "ValidateForm1() should return false")
 })
}

func TestValidateForm2(t *testing.T) {
 // Позитивный тест
 t.Run("Positive Test for Form2", func(t *testing.T) {
  name := "John Doe"
  age := 30
  address := "123 Main St"
  form := Form2{
   User: User{
    Name: &name,
    Age:  &age,
   },
   Address: &address,
  }
  assert.True(t, ValidateForm2(form), "ValidateForm2() should return true")
 })

 // Негативные тесты
 t.Run("Negative Test for Form2 - Missing Name", func(t *testing.T) {
  age := 30
  address := "123 Main St"
  form := Form2{
   User: User{
    Name: nil,
    Age:  &age,
   },
   Address: &address,
  }
  assert.False(t, ValidateForm2(form), "ValidateForm2() should return false")
 })

 t.Run("Negative Test for Form2 - Missing Age", func(t *testing.T) {
  name := "John Doe"
  address := "123 Main St"
  form := Form2{
   User: User{
    Name: &name,
    Age:  nil,
   },
   Address: &address,
  }
  assert.False(t, ValidateForm2(form), "ValidateForm2() should return false")
 })

 t.Run("Negative Test for Form2 - Missing Address", func(t *testing.T) {
  name := "John Doe"
  age := 30
  form := Form2{
   User: User{
    Name: &name,
    Age:  &age,
   },
   Address: nil,
  }
  assert.False(t, ValidateForm2(form), "ValidateForm2() should return false")
 })
}

package validate

import (
 "syy"
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
 // Позитивный тест
 t.Run("Positive Test", func(t *testing.T) {
  seo := &syy.Seo{
   User: syy.User{
    Age: "30",
   },
  }
  err := Validate(seo)
  assert.NoError(t, err, "Validate() should not return an error")
 })

 // Негативный тест
 t.Run("Negative Test - Missing Age", func(t *testing.T) {
  seo := &syy.Seo{
   User: syy.User{
    Age: "",
   },
  }
  err := Validate(seo)
  assert.Error(t, err, "Validate() should return an error")
  assert.Equal(t, "wtf", err.Error(), "Error message should match")
 })
}

package validate

import (
 "syy"
 "testing"

 "github.com/stretchr/testify/assert"
)

// Функция-помощник для создания экземпляра Seo с корректными данными
func newValidSeo() *syy.Seo {
 return &syy.Seo{
  User: syy.User{
   Age: "30",
   // Другие поля
  },
  Address: syy.Address{
   Street: "123 Main St",
   // Другие поля
  },
  // Другие поля и структуры
 }
}

// Функция-помощник для создания экземпляра Seo с некорректными данными
func newInvalidSeo() *syy.Seo {
 return &syy.Seo{
  User: syy.User{
   Age: "",
   // Другие поля
  },
  Address: syy.Address{
   Street: "123 Main St",
   // Другие поля
  },
  // Другие поля и структуры
 }
}

func TestValidate(t *testing.T) {
 // Позитивный тест
 t.Run("Positive Test", func(t *testing.T) {
  seo := newValidSeo()
  err := Validate(seo)
  assert.NoError(t, err, "Validate() should not return an error")
 })

 // Негативный тест
 t.Run("Negative Test - Missing Age", func(t *testing.T) {
  seo := newInvalidSeo()
  err := Validate(seo)
  assert.Error(t, err, "Validate() should return an error")
  assert.Equal(t, "wtf", err.Error(), "Error message should match")
 })
}