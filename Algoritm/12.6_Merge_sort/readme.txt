//ЗАДАНИЕ 12.6.1
//Реализуйте сортировку слиянием, Merge sort
https://play.golang.org/p/UiNZ2qSRdhd


Использование
Это действительно быстрая сортировка, которая имеет одинаковую асимптотику на любых входных данных. Но у него есть два недостатка:

Использование дополнительной памяти.
Рекурсивность: требуется n рекурсивных вызовов функции, что может вызвать ошибку переполнения стека для больших массивов.


Сложность
В лучшем случае: O(n log(n)).
В среднем случае: O(n log(n)).
В худшем случае: O(n log(n)).
Емкостная, в худшем: O(n).
