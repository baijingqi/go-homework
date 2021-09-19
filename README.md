# go-homework

sql.ErrNoRows 是Scan方法返回的错误，代表QueryRow方法没有返回有效的结果行。这代表没有符合条件的结果，不属于程序运行中的错误，且对业务没有影响，因此wrap该error并抛给上层是不合理的。
