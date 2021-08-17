# Базовые алгоритмы
## Алгоритмы поиска

### Бинарный поиск (binarySearch.go)
Идея: Бинарный поиск производится в упорядоченном массиве.

При бинарном поиске искомый ключ сравнивается с ключом среднего элемента в массиве. Если они равны, то поиск успешен. В противном случае поиск осуществляется аналогично в левой или правой частях массива.
Количество шагов поиска определится как **log<sub>2</sub>n↑**,

где **n**-количество элементов,
**↑** — округление в большую сторону до ближайшего целого числа.

На каждом шаге осуществляется поиск середины отрезка по формуле

**mid = (left + right)/2**

Если искомый элемент равен элементу с индексом mid, поиск завершается.
В случае если искомый элемент меньше элемента с индексом mid, на место mid перемещается правая граница рассматриваемого отрезка, в противном случае — левая граница.

Скорость работы алгоритма – **О(log<sub>2</sub>n)**

## Алгоритмы сортировки

### Быстрая сортировка (quickSort.go)
Суть алгоритма:

* Выбирается некоторый опорный элемент массива (mas[i])
* Запускается процедура разделения массива, которая перемещает все ключи, меньшие, либо равные a[i], влево от него, а все ключи, большие, либо равные a[i] - вправо
* Теперь массив состоит из двух подмножеств. Если в каком-то из подмассивов имеется более двух элементов, рекурсвино запускаем эту же процедуру

В конце получим полностью отсортированную последовательность.

Скорость работы алгоритма: **О(n log<sub>2</sub>n)** – средний случай, **O(n<sup>2</sup>)** – худший случай)  

### Сортировка слиянием (mergeSort.go)

Суть алгоритма: 
* Алгоритм рекурсивно делит массив на половины, пока не будет достигнут базовый случай с одним элементом
* Далее идет процедура слияния, которая объединяет два предварительно упорядоченных подмассива размерности n/2 в единый массив размерности n. 
* Начальные элементы предварительно упорядоченных массивов сравниваются между собой, и из них выбирается наименьший. Соответствующий указатель перемещается на следующий элемент.
* Процедура повторяется до тех пор, пока не достигнут конец одного из подмассивов. Оставшиеся элементы другого подмассива при этом передаются в итоговый массив в неизменном виде.

Скорость работы алгоритма: **О(n log<sub>2</sub>n)**

### Сортировка вставками (insertionSort.go)
Идея: алгоритм можно сравнить с сортированием колоды или одежды в шкафу. 

Суть алгоритма:
* Для сравнения в начале выбираем нулевой элемент массива и считаем его уже отсортированным.
* Далее, пробегаясь по элементам массива, начиная с первого, сравниваем элемент с другим элементом, который стоит слева. 
* Если элемент слева больше текущего, меняем элементы местами.
* Повторяем предыдущий шаг, пока элемент слева не будет меньше текущего, или же текущий элмент не встанет в самом начале массива.

Скорость работы алгоритма: **О(n<sup>2</sup>)** (худший случай), **О(n)** (лучший случай)