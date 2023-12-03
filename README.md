## Apache Benchmark

- Get Order By ID - 2551 RPS   
  NOTICE: replace order_uid in URL on your.

```bash
 ab -n 5000 -c 10 http://0.0.0.0:8080/orders/baf99a3a-1142-4a1d-823f-6882fc71e8d3
 ```

![Снимок экрана 2023-12-03 в 17.12.10.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2F6z%2Frgws40695y35tnhg8gs54svm0000gn%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_xDA1PY%2F%D0%A1%D0%BD%D0%B8%D0%BC%D0%BE%D0%BA%20%D1%8D%D0%BA%D1%80%D0%B0%D0%BD%D0%B0%202023-12-03%20%D0%B2%2017.12.10.png)

- Create Order - 1854 RPS

```bash
ab -m POST -T application/json -c 10 -n 5000 http://0.0.0.0:8080/orders
```  

![Снимок экрана 2023-12-03 в 17.12.38.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2F6z%2Frgws40695y35tnhg8gs54svm0000gn%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_5z8CJS%2F%D0%A1%D0%BD%D0%B8%D0%BC%D0%BE%D0%BA%20%D1%8D%D0%BA%D1%80%D0%B0%D0%BD%D0%B0%202023-12-03%20%D0%B2%2017.12.38.png)