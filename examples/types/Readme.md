# Система типів

## 1. Системні типи

### 1.1 Примітиви

#### [1.1.1 Визначення змінних](primitives/definition/definition.go)

#### [1.1.2 Нульові значення](primitives/zero/zero.go)

#### [1.1.3 Розмір типів](primitives/size/size.go)

### 1.2 Перетворення типів

#### [1.2.1 Перетворення "подібних" типів](conversion/conversion.go)

#### [1.2.2 Перетворення примітивів](conversion/primitives/conversion.go)

#### [1.2.3 Невизначений тип](conversion/interface/interface.go)

### 1.3 Функції

#### [1.3.1 Визначення функцій](functions/functions.go)

#### [1.3.2 Затінення змінних](functions/shadowing/shadowing.go)

### 1.4 Покажчики

#### [1.4.1 Покажчики](pointers/pointers.go)

#### [1.4.2 Покажчики у функціях](pointers/in_functions/pointers_in_functions.go)

### 1.5 Константи

#### [1.5.1 Константи](constants/constants.go)

#### [1.5.2 Iota](constants/iota/iota.go)

### 1.6 rune

#### [1.6.1 rune](rune/rune.go)

## 2. Типи визначені користувачем

### 2.1 Структури

#### [2.1.1 Визначення структур](struct/definition.go)

#### [2.1.2 Приведення структур](struct/conversion/conversion.go)

#### [2.1.3 Розмір](struct/size/size.go)

### 2.2 Примітиви

#### [2.2.1 Нові типи на основі примітивів](custom_primitives/custom_primitives.go)

### 2.3 Методи

#### [2.3.1 Визначення методів](methods/methods.go)

#### [2.3.2 Покажчик vs значення у методах](methods/copy/copy.go)

#### [2.3.3 nil у методах](methods/nil/nil.go)

#### [2.3.4 Методи як функції](methods/functions/functions.go)

## 3 Колекції

### 3.1 Масиви

#### [3.1.1. Створення масиву](array/definition/definition.go)

#### [3.1.2. Масив з елементами різного типу](array/dynamic_type/dynamic_type.go)

#### [3.1.3. Розмір масиву](array/size/size.go)

#### [3.1.4. Прохід по масиву](array/loop/loop.go)

#### [3.1.5. Порівняння](array/compare/compare.go)

#### [3.1.6. Копіювання](array/copy/copy.go)

#### [3.1.7. Поверхневе копіювання](array/copy/shallow/shallow.go)

#### [3.1.8. Передача у функцію](array/functions/functions.go)

#### [3.1.9. Двовимірний масив](array/2d/2d.go)

### 3.2. Зрізи

#### [3.2.1. Створення зрізу](slice/definition/definition.go)

#### [3.2.2. Порожній зріз](slice/empty/empty.go)

#### [3.2.3. Копіювання](slice/copy/copy.go)

#### [3.2.4. Повне копіювання](slice/copy/deep/deep_copy.go)

#### [3.2.5. Порівняння](slice/compare/compare.go)

#### [3.2.6. Зріз зрізу](slice/slice/slice.go)

#### [3.2.7. Додавання нових елементів](slice/append/append.go)

#### [3.2.8. Додавання масиву нових елементів](slice/append/multiple/append_multiple.go)

#### [3.2.9. Додавання елементів в зріз зрізу](slice/append/multiple/append_multiple.go)

### 3.3 Мапи

#### [3.3.1. Створення мапи](map/definition.go)

#### [3.3.2. Основні операції](map/basic/basic.go)

#### [3.3.1. Прохід по мапі](map/range/range.go)
