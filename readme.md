### 扫码小程序

#### Redis

逻辑: 先去picture_book中寻找, 然后再去third_book中寻找

##### redis 绘本数据构成 

reids 结构: list 

picture_book:

key: wechat:picture_book:{isbn} 如: wechat:picture_book:9787544245906

value: {book_name}:{isbn}:{status}, 如: HANDA'S SURPRISE Read and Share:9780763608637:1

third_book:

key: wechat:third_book:{isbn} 如: wechat:third_book:9787544245906

value: {book_name}:{isbn}:{status}, 如: HANDA'S SURPRISE Read and Share:9780763608637:1

#### 部署



