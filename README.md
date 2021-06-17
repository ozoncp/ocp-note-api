# Note API for Ozon Code Platform
![Build & Test](https://github.com/ozoncp/ocp-note-api/actions/workflows/golangci-lint.yml/badge.svg?branch=main)
<!-- ![Build & Test](https://github.com/ozoncp/ocp-note-api/actions/workflows/go.yml/badge.svg?branch=main) -->
<!-- [![codecov](https://codecov.io/gh/ozoncp/ocp-note-api/branch/main/graph/badge.svg)](https://codecov.io/gh/ozoncp/ocp-note-api) -->

# Сборка и запуск сервиса и его окружения


https://user-images.githubusercontent.com/13212921/122275099-fe0f8b00-ceeb-11eb-9843-394ccc5b558c.mp4

# Ручки сервиса 
## CreateNoteV1
Добавляет заметку в базу данных, присваивая ей определенный `Id`.
### Входные параметры
* `UserId` - идентификатор пользователя.
* `ClassroomId` - идентификатор классной комнаты.
* `DocumentId` - идентификатор документа.  

## MultiCreateNotesV1
Добавляет сразу несколько заметок в базу данных, присваивая им определенные `Id`.
### Входные параметры
Входным параметром является массив, каждый элемент которого содержит в себе 3 параметра:
* `UserId` - идентификатор пользователя.
* `ClassroomId` - идентификатор классной комнаты.
* `DocumentId` - идентификатор документа.  

## UpdateNoteV1
Обновляет данные указанной заметки.
### Входные параметры
* `Id` - идентификатор целевой заметки.
* `UserId` - новый идентификатор пользователя, которым хотим обновить старое значение.
* `ClassroomId` - новый идентификатор классной комнаты, которым хотим обновить старое значение.
* `DocumentId` - новый идентификатор документа, которым хотим обновить старое значение.  

## DescribeNoteV1
Предоставляет информацию (`UserId`, `ClassroomId`, `DocumentId) об указанной заметке.
### Входные параметры
* `Id` - идентификатор целевой заметки.

## ListNotesV1
Возвращает указанное количество заметок начиная с указанного отступа.
### Входные параметры
* `Limit` - количество заметок, которое нужно получить.
* `Offset` - отступ от начала таблицы заметок.  

## RemoveNoteV1
Удаляет указанную заметку.
### Входные параметры
* `Id` - идентификатор целевой заметки.
