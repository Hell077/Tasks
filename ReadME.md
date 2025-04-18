
# Задания по Go (Golang)

## 1. Основы синтаксиса Go
- **Задание 1:** Напишите программу, которая выводит строку "Hello, World!".
- **Задание 2:** Создайте переменные различных типов (целые числа, строки, булевы значения) и выведите их на экран.
- **Задание 3:** Напишите программу, которая находит сумму чисел от 1 до 100 с помощью цикла.
- **Задание 4:** Напишите функцию, которая принимает два числа и возвращает их сумму.
- **Задание 5:** Напишите программу, которая выполняет операцию деления двух чисел с проверкой на деление на ноль.

## 2. Управляющие конструкции
- **Задание 1:** Напишите программу, которая проверяет, является ли число четным или нечетным.
- **Задание 2:** Напишите программу с использованием switch-case для вывода дня недели по числу (1 — Понедельник, 2 — Вторник и т.д.).
- **Задание 3:** Напишите программу, которая выводит таблицу умножения для числа от 1 до 10.

## 3. Массивы и срезы
- **Задание 1:** Напишите программу, которая создает массив из 5 чисел и выводит его.
- **Задание 2:** Напишите программу, которая находит максимальное значение в срезе чисел.
- **Задание 3:** Напишите программу, которая объединяет два среза в один.
- **Задание 4:** Напишите программу, которая сортирует срез чисел по возрастанию.

## 4. Функции
- **Задание 1:** Напишите функцию, которая находит факториал числа.
- **Задание 2:** Напишите программу, которая вычисляет число Фибоначчи для заданного числа.
- **Задание 3:** Напишите функцию, которая принимает строку и возвращает ее в обратном порядке.
- **Задание 4:** Напишите функцию, которая проверяет, является ли строка палиндромом.

## 5. Структуры и интерфейсы
- **Задание 1:** Создайте структуру `Person` с полями `Name` и `Age`, создайте экземпляр этой структуры и выведите его.
- **Задание 2:** Напишите метод для структуры `Person`, который выводит информацию о человеке.
- **Задание 3:** Создайте интерфейс `Shape` с методами `Area` и `Perimeter`. Реализуйте этот интерфейс для структур `Circle` и `Rectangle`.
- **Задание 4:** Напишите программу, которая проверяет, реализует ли структура интерфейс.

## 6. Пакеты и библиотеки
- **Задание 1:** Напишите программу, которая использует пакет `math` для вычисления квадратного корня числа.
- **Задание 2:** Создайте свой собственный пакет, который содержит функцию для вычисления площади прямоугольника.
- **Задание 3:** Напишите программу, которая использует стандартный пакет `time` для вывода текущего времени и даты.

## 7. Горутины и каналы
- **Задание 1:** Напишите программу, которая запускает две горутины, каждая из которых выводит число от 1 до 5.
- **Задание 2:** Напишите программу, которая использует каналы для передачи данных между горутинами.
- **Задание 3:** Реализуйте с помощью канала механизм синхронизации между горутинами.
- **Задание 4:** Напишите программу, которая использует буферизированный канал для передачи данных между горутинами.

## 8. Ошибки и обработка исключений
- **Задание 1:** Напишите программу, которая делит два числа, обрабатывая ошибку деления на ноль.
- **Задание 2:** Создайте функцию, которая возвращает ошибку, если строка пустая, и обрабатывает эту ошибку.
- **Задание 3:** Напишите функцию, которая проверяет, является ли файл доступным для чтения. Если файл не найден, верните ошибку.

## 9. Работа с файлами
- **Задание 1:** Напишите программу, которая читает содержимое файла и выводит его на экран.
- **Задание 2:** Напишите программу, которая записывает строку в файл.
- **Задание 3:** Напишите программу, которая копирует содержимое одного файла в другой.
- **Задание 4:** Напишите программу, которая создает и записывает структуру данных в файл с использованием `json` формата.

## 10. Работа с базами данных
- **Задание 1:** Напишите программу, которая подключается к базе данных SQLite и выполняет простой запрос.
- **Задание 2:** Напишите программу, которая подключается к базе данных PostgreSQL, выполняет запрос на выборку данных и выводит результат.
- **Задание 3:** Напишите программу, которая выполняет операцию вставки и удаления данных в таблице базы данных.

## 11. HTTP сервер и клиент
- **Задание 1:** Напишите простой HTTP сервер, который отвечает на запросы с текстом "Hello, World!".
- **Задание 2:** Напишите программу, которая делает HTTP GET запрос к публичному API и выводит результат.
- **Задание 3:** Напишите сервер с обработкой POST-запроса, который принимает JSON данные и возвращает их в ответе.
- **Задание 4:** Реализуйте авторизацию с использованием токенов (JWT) для защиты API.

## 12. Тестирование
- **Задание 1:** Напишите простую функцию и создайте тест для проверки ее работы с использованием пакета `testing`.
- **Задание 2:** Напишите тест, который проверяет, что функция правильно обрабатывает ошибку.
- **Задание 3:** Напишите тест для HTTP обработчика, который проверяет статус код ответа и содержание.

## 13. Продвинутые темы
- **Задание 1:** Реализуйте паттерн проектирования "Singleton" в Go.
- **Задание 2:** Реализуйте паттерн "Observer", который уведомляет подписчиков о событиях.
- **Задание 3:** Напишите программу, которая использует горутины и каналы для реализации пула соединений с базой данных.
- **Задание 4:** Напишите микросервис на Go, который использует gRPC для общения с другими микросервисами.
- **Задание 5:** Реализуйте HTTP сервер с использованием `gorilla/mux` и настройте маршруты для разных типов запросов.

## 14. Оптимизация и производительность
- **Задание 1:** Напишите программу, которая измеряет производительность обработки больших объемов данных с использованием различных структур данных (срезы, карты, массивы).
- **Задание 2:** Реализуйте оптимизацию для работы с большими данными с использованием параллельных вычислений (горутин и каналов).
- **Задание 3:** Напишите программу, которая профилирует использование памяти и времени выполнения с помощью стандартных инструментов Go.
