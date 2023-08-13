package example

// tag
// -       ignore field in the struct.
// auto    当Field类型为int、int32、int64、uint、uint32或uint64时，可以将其设置为自动递增。如果模型定义中没有主键，则具有上述类型之一的字段Id将被视为自增键
// pk      设置为主键。用于使用其他类型字段作为主键。
// null    默认情况下字段不为 NULL。将 null 设置为 ALLOW NULL。
// index   为一个字段添加索引
// unique  为一个字段添加唯一键        `orm:"unique"`
// column  设置数据库表中字段的列名。   `orm:"column(Student_name)"`
// size    字符串字段的默认值为 varchar(255)  `orm:"size(60)"`
// digits / decimals  设置 float32 或 float64 的精度。   `orm:"digits(12);decimals(4)"` : 共12位，点后4位。例如：12345678.1234
// auto_now：每次保存都会更新时间。
// auto_now_add：设置第一次保存的时间
// type     `orm:"auto_now_add;type(date)"`
// default       `orm:"default(1);description(this is status)"`
// Comment       `orm:"description(this is status)"`
//   on_delete     设置删除关联关系后如何处理字段
//   cascade     级联删除（默认）
//   set_null    设置为 NULL。需要设置null=true
//   set_default 设置为默认值。需要设置默认值。
//   do_nothing  什么都不做。忽略。

type Student struct {
	Id      int      `orm:"column(id);auto"`
	Name    string   `orm:"column(name)"`
	Email   string   `orm:"column(email)"`
	Profile *Profile `orm:"rel(one);on_delete(cascade)"` // 一个学生对应一个属性   数据库要先插入profile，然后指定profile结构体
	Posts   []*Post  `orm:"reverse(many)"`               // 一个学生对应多篇文章
}

type Post struct {
	Id      int     `orm:"column(id);auto"`
	Content string   `orm:"column(content)"`
	Student *Student `orm:"rel(fk)"`  // 一个文章对应一个学生
	Tags    []*Tag   `orm:"rel(m2m)"` // 一篇文章对应多个tag
}

type Profile struct {
	Id      int     `orm:"column(id);auto"`
	Age     int      `orm:"column(age)"`
	Student *Student `orm:"reverse(one)"` // 一个属性对应一个学生
}

type Tag struct {
	Id    int     `orm:"column(id);auto"`
	Name  string  `orm:"column(name)"`
	Posts []*Post `orm:"reverse(many)"` // 一个tag对应多篇文章
}
