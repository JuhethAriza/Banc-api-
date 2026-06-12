# Estándares de Desarrollo

Este documento contiene las convenciones y mejores prácticas que todo el equipo de desarrollo debe seguir para mantener un código consistente, legible y mantenible.

---

## 1. Gestión de Ramas en Git

### 1.1 Estructura de Ramas

```
main/master     → Código en producción (solo merge desde develop)
develop         → Código de desarrollo estable
feature/        → Nuevas funcionalidades
bugfix/         → Corrección de bugs
hotfix/         → Correcciones urgentes en producción
release/        → Preparación de versiones
```

### 1.2 Convención de Nombres de Ramas

**Formato:** `tipo/#tarea-descripcion-corta`

**Ejemplos:**
- `feature/H01-autenticacion-usuario`
- `feature/H02-carrito-compras`
- `bugfix/B03-correccion-calculo-precio`
- `hotfix/B04-error-login-critico`
- `release/v1.2.0`

**Reglas:**
- Usar minúsculas
- Separar palabras con guiones (-)
- Ser descriptivo pero conciso
- No usar espacios ni caracteres especiales

### 1.3 Flujo de Trabajo

1. **Crear nueva rama desde develop:**
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b feature/nombre-funcionalidad
   ```

2. **Trabajar en la rama:**
   - Hacer commits frecuentes y descriptivos
   - Mantener la rama actualizada con develop

3. **Antes de mergear:**
   - Actualizar la rama con develop: `git rebase develop` o `git merge develop`
   - Resolver conflictos si existen
   - Verificar que el código compile y funcione

4. **Merge a develop:**
   - Crear Pull Request (PR) o Merge Request (MR)
   - Solicitar revisión de código
   - Una vez aprobado, mergear a develop

---

## 2. Convenciones de Commits

### 2.1 Formato de Mensajes

**Estructura:**
```
tipo: descripción corta (máximo 50 caracteres)

Descripción detallada (opcional, máximo 72 caracteres por línea)
- Punto adicional 1
- Punto adicional 2
```

### 2.2 Tipos de Commits

- `feat:` Nueva funcionalidad
- `fix:` Corrección de bug
- `docs:` Cambios en documentación
- `style:` Cambios de formato (espacios, comas, etc.)
- `refactor:` Refactorización de código
- `test:` Agregar o modificar tests
- `chore:` Tareas de mantenimiento (dependencias, configuraciones)
- `perf:` Mejoras de rendimiento
- `ci:` Cambios en CI/CD

### 2.3 Ejemplos de Commits

```
feat: agregar autenticación con JWT

- Implementar login de usuarios
- Agregar middleware de autenticación
- Crear endpoint de refresh token
```

```
fix: corregir cálculo de descuentos en carrito
```

```
refactor: mejorar estructura de servicios de productos

- Separar lógica de negocio de controladores
- Crear ProductService
- Actualizar tests relacionados
```

## 3. Nomenclatura de Código

### 3.1 Variables

**Reglas generales:**
- Nombres descriptivos y claros
- Evitar abreviaciones innecesarias
- Usar el idioma del proyecto (español o inglés, pero ser consistente)

**Convenciones por lenguaje:**

**PHP/Laravel:**
```php
// ✅ Correcto
$nombreUsuario = "Juan";
$totalProductos = 10;
$fechaCreacion = now();
$esActivo = true;

// ❌ Incorrecto
$n = "Juan";
$tp = 10;
$fc = now();
$act = true;
```

**JavaScript/TypeScript:**
```javascript
// ✅ Correcto
const nombreUsuario = "Juan";
const totalProductos = 10;
const fechaCreacion = new Date();
const esActivo = true;

// ❌ Incorrecto
const n = "Juan";
const tp = 10;
```

**Constantes:**
```php
// ✅ Correcto
const MAX_INTENTOS_LOGIN = 3;
const PRECIO_DEFAULT = 0.00;
const ESTADO_ACTIVO = 'activo';
```

### 3.2 Métodos/Funciones

**Reglas:**
- Nombres en modo imperativo (verbos)
- Descriptivos de lo que hacen
- Un método debe hacer una sola cosa

**Ejemplos:**

```php
// ✅ Correcto
public function calcularTotal()
public function obtenerUsuarioPorId($id)
public function validarEmail($email)
public function enviarNotificacion($usuario)
public function esValido()

// ❌ Incorrecto
public function calculo()
public function get($id)
public function validar()
public function hacer()
```

**Convenciones comunes:**
- `obtener*` / `get*` - Para obtener datos
- `crear*` / `create*` - Para crear registros
- `actualizar*` / `update*` - Para actualizar
- `eliminar*` / `delete*` - Para eliminar
- `validar*` / `validate*` - Para validaciones
- `es*` / `is*` - Para booleanos
- `tiene*` / `has*` - Para verificar existencia

### 3.3 Clases

**Reglas:**
- Sustantivos en singular
- PascalCase (primera letra de cada palabra en mayúscula)
- Descriptivos del propósito

**Ejemplos:**

```php
// ✅ Correcto
class UsuarioController
class ProductoService
class CalculadoraPrecios
class ValidadorEmail
class RepositorioOrdenes

// ❌ Incorrecto
class Usuario
class ProductoSvc
class Calc
class Val
```

**Sufijos comunes:**
- `Controller` - Controladores
- `Service` - Servicios de negocio
- `Repository` - Repositorios de datos
- `Model` - Modelos de datos
- `Helper` - Utilidades
- `Validator` - Validadores
- `Exception` - Excepciones personalizadas

---

## 4. Manejo de Errores

### 4.1 Principios Generales

- ✅ Siempre manejar errores explícitamente
- ✅ Proporcionar mensajes de error claros y útiles
- ✅ Registrar errores para debugging
- ✅ No exponer información sensible al usuario final
- ✅ Usar excepciones apropiadas


### 4.2 Códigos de Estado HTTP

- `200` - OK (éxito)
- `201` - Created (recurso creado)
- `400` - Bad Request (solicitud inválida)
- `401` - Unauthorized (no autenticado)
- `403` - Forbidden (sin permisos)
- `404` - Not Found (recurso no encontrado)
- `422` - Unprocessable Entity (error de validación)
- `500` - Internal Server Error (error del servidor)

### 4.3 Logging

**Reglas:**
- Registrar errores con contexto suficiente
- Usar niveles apropiados (error, warning, info, debug)
- No registrar información sensible (contraseñas, tokens)
- Incluir información útil para debugging

---

## 5. Nomenclatura de Base de Datos

### 5.1 Tablas

**Reglas:**
- Nombres en plural
- Minúsculas
- Separar palabras con guión bajo (_)
- Ser descriptivos

**Ejemplos:**
```
- users
- productos
- ordenes_compra
- detalles_orden
- tipos_documento
- metodos_pago

### 5.2 Campos/Columnas

**Reglas:**
- Minúsculas
- Separar palabras con guión bajo (_)
- Ser descriptivos
- Usar nombres consistentes

**Ejemplos:**
```
- id
- nombre
- apellido
- email
- fecha_creacion
- fecha_actualizacion
- usuario_id
- es_activo
- precio_unitario
- cantidad_disponible


### 5.3 Campos Comunes
**Timestamps:**
- `created_at` - Fecha de creación
- `updated_at` - Fecha de actualización
- `deleted_at` - Fecha de eliminación (soft delete)

**Foreign Keys:**
- Formato: `tabla_referenciada_id`
- Ejemplos: `usuario_id`, `producto_id`, `categoria_id`

**Booleanos:**
- Prefijo `es_` o `is_`: `es_activo`, `es_visible`, `es_publico`
- O usar nombres descriptivos: `habilitado`, `activo`

**Estados:**
- Campo `estado` o `status`
- Valores consistentes: `activo`, `inactivo`, `pendiente`, `completado`

### 5.4 Índices

**Nomenclatura:**
- Primary Key: `pk_tabla` (generalmente solo `id`)
- Foreign Key: `fk_tabla_campo`
- Unique: `uk_tabla_campo`
- Index: `idx_tabla_campo`

---

## 6. Comentarios y Documentación

### 6.1 Comentarios en Código

**Reglas:**
- Comentar el "por qué", no el "qué"
- Mantener comentarios actualizados
- Evitar comentarios obvios
- Usar comentarios para explicar lógica compleja

**Ejemplos:**

```php
// ✅ Buen comentario
// Aplicar descuento acumulativo: primero descuento por volumen,
// luego descuento por cliente frecuente
$precio = aplicarDescuentoVolumen($precio);
$precio = aplicarDescuentoClienteFrecuente($precio);
```

### 6.2 Documentación de Métodos

**PHP (PHPDoc):**
```php
/**
 * Calcula el precio total de una orden incluyendo impuestos y descuentos.
 *
 * @param int $ordenId ID de la orden
 * @param float $descuento Descuento a aplicar (0-100)
 * @return float Precio total calculado
 * @throws ModelNotFoundException Si la orden no existe
 */
public function calcularPrecioTotal(int $ordenId, float $descuento = 0): float
{
    // ...
}
```

---

## 7. Checklist Antes de Hacer Commit

- [ ] El código compila sin errores
- [ ] No hay código comentado innecesario
- [ ] Los nombres de variables/métodos/clases son descriptivos
- [ ] Se manejan los errores apropiadamente
- [ ] No hay información sensible (contraseñas, tokens) en el código
- [ ] El mensaje de commit sigue las convenciones
- [ ] Los cambios están relacionados (un commit = un cambio lógico)
- [ ] Se actualizó la documentación si es necesario

---

## 8. Recursos Adicionales

- Revisar código de compañeros antes de mergear
- Mantener comunicación constante con el equipo
- Preguntar si hay dudas sobre las convenciones
- Actualizar este documento si se acuerdan nuevas convenciones

---

**Última actualización:** [2025-11-09]
**Versión:** 1.0

