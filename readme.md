### 扫码小程序

主要使用gin框架实现, 利用jwt. 



存储:

#### Redis

##### 数据结构: list



##### 逻辑: 

##### 先去picture_book中寻找, 然后再去third_book中寻找



##### 数据构成  

picture_book:

* key: wechat:picture_book:{isbn} 如: wechat:picture_book:9787544245906
* value: {book_name}:{isbn}:{status}, 如: HANDA'S SURPRISE Read and Share:9780763608637:1

third_book:

* key: wechat:third_book:{isbn} 如: wechat:third_book:9787544245906
* value: {book_name}:{isbn}:{status}, 如: HANDA'S SURPRISE Read and Share:9780763608637:1



