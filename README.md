# Алгоритм Левита

Реализация алгоритма Левита для поиска кратчайших путей в графе на языке Go.

## Описание

Алгоритм Левита - это улучшенная версия алгоритма Дейкстры для поиска кратчайших путей в графе. Он использует двунаправленную очередь (дек) для оптимизации процесса поиска и может работать быстрее алгоритма Дейкстры на некоторых типах графов.

### Особенности реализации

- Поддержка неориентированных взвешенных графов
- Использование списков смежности для представления графа
- Эффективная реализация с использованием двунаправленной очереди
- Возможность найти кратчайшие пути от одной вершины до всех остальных

## Структура проекта

- `levitts_algorithm.go` - основная реализация алгоритма
- `levitts_algorithm_test.go` - тесты
- `go.mod` - файл модуля Go

## Использование

```go
// Создание графа
g := NewGraph(5)

// Добавление рёбер
g.AddEdge(0, 1, 5)  // ребро между вершинами 0 и 1 весом 5
g.AddEdge(1, 2, 2)
// ...

// Поиск кратчайших путей от вершины 0
distances := g.Levit(0)
```

## Сложность алгоритма

- Временная сложность: O(M * N), где M - количество рёбер, N - количество вершин
- Пространственная сложность: O(N)

## Преимущества алгоритма Левита

1. Может работать быстрее алгоритма Дейкстры на разреженных графах
2. Эффективен для графов с небольшим количеством отрицательных рёбер
3. Использует оптимизированную структуру данных (дек)

## Требования

- Go 1.15 или выше

## Запуск тестов

```bash
go test
```