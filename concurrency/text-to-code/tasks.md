
# Задачи для практики

## 1 .Напиши две функции: generate(nums ...int) <-chan int — кладёт числа в канал, и square(in <-chan int) <-chan int — читает из канала и возвращает квадраты. Соедини их в pipeline.


## 2. Напиши функцию fanOut(in <-chan int, n int) []<-chan int — читает из одного канала и распределяет значения по n выходным каналам равномерно.

## 3. Напиши workerPool(jobs <-chan string, n int) <-chan string — запускает n воркеров которые читают задачи из jobs, обрабатывают через process(job string) string и пишут результаты в выходной канал.

## 4. Напиши функцию fetchWithTimeout(url string, timeout time.Duration) (string, error) — делает HTTP GET запрос и возвращает ошибку если не уложился в timeout. Используй context.

## 5. Напиши функцию rateLimit(in <-chan int, rps int) <-chan int — пропускает не более rps значений в секунду из входного канала.

## 6. Напиши функцию retry(fn func() error, attempts int, delay time.Duration) error — вызывает fn до attempts раз с задержкой delay между попытками. Возвращает последнюю ошибку если все попытки провалились.

## 7. Напиши функцию parallel(tasks []func(), limit int) — выполняет все tasks параллельно но не более limit одновременно. Дожидается завершения всех.

## 8. Напиши функцию first(ctx context.Context, fns ...func(ctx context.Context) (int, error)) (int, error) — запускает все функции параллельно и возвращает результат первой которая завершилась без ошибки. Остальные отменяет через context.

## 9. Напиши функцию batch(in <-chan int, size int, timeout time.Duration) <-chan []int — собирает значения из канала в батчи по size элементов или флашит батч если прошло timeout с последнего элемента.

## 10. Напиши структуру CircuitBreaker с методом Call(fn func() error) error — если fn падает более 3 раз подряд, следующие вызовы сразу возвращают ошибку не вызывая fn. Сбрасывается после 5 секунд.