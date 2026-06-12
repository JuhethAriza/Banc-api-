# Arquitectura Limpia - Banc API

## Estructura del Proyecto

```
banc-api/
├── cmd/                          # Punto de entrada de la aplicación
│   └── main.go
├── internal/                     # Código privado de la aplicación
│   ├── domain/                   # CAPA DE DOMINIO (Enterprise Business Rules)
│   │   ├── entity/               # Entidades del negocio (User, Account, etc.)
│   │   ├── repository/           # Interfaces de repositorios (contratos)
│   │   └── valueobject/          # Objetos de valor (Role, TypeAccount)
│   │
│   ├── application/              # CAPA DE APLICACIÓN (Application Business Rules)
│   │   ├── usecase/              # Casos de uso (lógica de negocio específica)
│   │   └── dto/                  # Data Transfer Objects
│   │
│   ├── infrastructure/           # CAPA DE INFRAESTRUCTURA
│   │   └── persistence/          # Implementaciones de repositorios (GORM)
│   │
│   └── interface/                # CAPA DE INTERFAZ (Interface Adapters)
│       ├── http/                 # Handlers HTTP y rutas
│       └── middleware/           # Middleware (auth, logging, etc.)
│
├── pkg/                          # Código público/reutilizable
│   ├── database/                 # Configuración de base de datos
│   ├── response/                 # Respuestas HTTP estandarizadas
│   ├── logger/                   # Logger
│   └── validator/                # Validadores
│
├── config/                       # Configuración y variables de entorno
├── migrations/                   # Migraciones de base de datos
└── tests/                        # Tests
```

## Flujo de Dependencias

```
┌─────────────────────────────────────────────────────────────┐
│                        Interface Layer                       │
│                    (Handlers, Routes, DTOs)                  │
└────────────────────────────┬────────────────────────────────┘
                             │ depende de
                             ▼
┌─────────────────────────────────────────────────────────────┐
│                      Application Layer                       │
│                  (Use Cases, Application DTOs)               │
└────────────────────────────┬────────────────────────────────┘
                             │ depende de
                             ▼
┌─────────────────────────────────────────────────────────────┐
│                        Domain Layer                          │
│            (Entities, Repository Interfaces, Value Objects)  │
└────────────────────────────┬────────────────────────────────┘
                             ▲ depende de
                             │
┌────────────────────────────┴────────────────────────────────┐
│                   Infrastructure Layer                       │
│          (Repository Implementations, Database, External)    │
└─────────────────────────────────────────────────────────────┘
```

## Principios de Diseño

### 1. Inversión de Dependencias (DIP)
Las capas superiores dependen de abstracciones, no de implementaciones concretas.
```go
// En domain/repository/user_repository.go
type UserRepository interface {
    GetAll() ([]User, error)
    GetByID(id uint) (*User, error)
    Create(user *User) error
    Update(user *User) error
    Delete(id uint) error
}

// En infrastructure/persistence/user_repository_impl.go
type userRepositoryImpl struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
    return &userRepositoryImpl{db: db}
}
```

### 2. Separación de Responsabilidades
- **Domain**: Reglas de negocio empresariales (independientes de la aplicación)
- **Application**: Casos de uso específicos de la aplicación
- **Infrastructure**: Implementaciones técnicas (BD, APIs externas)
- **Interface**: Adaptadores HTTP/REST

### 3. Value Objects
Objetos inmutables que representan conceptos del dominio:
- `Role`: Roles de usuario con permisos
- `TypeAccount`: Tipos de cuenta bancaria

### 4. Entities
Entidades con identidad propia y lógica de negocio:
- `User`: Usuario del sistema
- `Account`: Cuenta bancaria

## Wire-up de Dependencias

El contenedor de dependencias en `internal/init.go` configura todas las capas:

```go
func NewAppContainer(db *database.DB) *AppContainer {
    // Infraestructura
    userRepo := persistence.NewUserRepository(db.Instance())
    
    // Aplicación
    userUseCase := usecase.NewUserUseCase(userRepo)
    
    // Interface
    userHandler := http.NewUserHandler(userUseCase)
    
    return &AppContainer{...}
}
```

## Ejemplo de Uso

### Crear un nuevo usuario

**Request:**
```http
POST /users
Content-Type: application/json

{
    "username": "juan",
    "email": "juan@example.com",
    "password": "secret123",
    "role": "user"
}
```

**Response:**
```json
{
    "message": "Usuario creado exitosamente",
    "data": {
        "id": 1,
        "username": "juan",
        "email": "juan@example.com",
        "role": "user",
        "fecha_creacion": "2026-06-11T10:00:00Z"
    }
}
```

## Beneficios de esta Arquitectura

1. **Testabilidad**: Cada capa puede testearse independientemente
2. **Mantenibilidad**: Cambios en una capa no afectan las demás
3. **Escalabilidad**: Fácil agregar nuevos casos de uso o endpoints
4. **Flexibilidad**: Cambiar de base de datos o framework HTTP no afecta el dominio
5. **Claridad**: Cada archivo tiene una responsabilidad bien definida
