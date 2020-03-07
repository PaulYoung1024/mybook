package models

type Category struct {
	Id     int
	Pid    int    //分类id
	Title  string `orm:"size(30);unique"`
	Intro  string //介绍
	Icon   string
	Cnt    int  //分类下的图书数量
	Sort   int  //排序
	Status bool //状态，true显示
}

func (m *Category) TableName() string {
	return TNCategory()
}

//查询表md_category里某个pid下的记录
func (m *Category) GetCates(pid int, status int) (cates []Category, err error) {
	qs := GetOrm("r").QueryTable(TNCategory())
	if pid > -1 {
		qs = qs.Filter("pid", pid)
	}
	if status == 0 || status == 1 {
		qs = qs.Filter("status", status)
	}
	_, err = qs.OrderBy("-status", "sort", "title").All(&cates)
	return
}

//查询表md_category里某个id下的一条记录
func (m *Category) Find(id int) (cate Category) {
	cate.Id = id
	GetOrm("r").Read(&cate)
	return cate
}
