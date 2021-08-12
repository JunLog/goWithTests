package maps

type Dictionary map[string]string

// 不能直接定义一个空的 map 变量
// var m map[string]string 这会 造成 nil 指针异常
// 初始化一个 空的 map 变量，应该
// dictionary = map[string][string]
// 或者 dictionary = make(map[string]string)

// var (
// 	ErrNotFound   = errors.New("找不到你要找的词")
// 	ErrWordExists = errors.New("不能添加，因为已经存在")
// )

const (
	ErrNotFound         = DictionaryErr("找不到你要找的词")
	ErrWordExists       = DictionaryErr("不能添加，因为已经存在")
	ErrWordDoedNotExist = DictionaryErr("不能更新，因为本身就不存在")
)

type DictionaryErr string

// 类型实现 error 接口
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		// errors.New() 构造一个新的 error 值
		// return "", errors.New("找不到您要找的词")
		// 将 错误 提取成 变量，摆脱 魔术错误(magic error)?
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	// d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoedNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	// d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	// map 的内置函数 delete ，两个参数 第一个是 map ，第二个是 键
	delete(d, word)
}
