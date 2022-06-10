package models

type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"` // gorm:"index" 声明这个字段是索引
	Tag        Tag    `json:"tag"`                 // 嵌套的 struct 用于关联查询
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// ExistsArticleByID 文章是否存在
func ExistsArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// GetArticleTotal 获取文章总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// GetArticles 获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// Preloads 是一个预加载器，gorm 在查询出结构之后，会在内部处理对应的映射逻辑，将其填充到 Article 的 Tag 中
	db.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// GetArticle 获取文章详情
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	// Article 有一个结构体成员是 TagID 也就是外键
	// gorm 会通过结构体+ID的方式寻找俩个结构体之间的关联关系
	// Article 中有一个结构体成员是 Tag
	db.Model(&article).Related(&article.Tag)
	return
}

// EditArticle 编辑文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}

// AddArticle 新增文章
func AddArticle(data map[string]interface{}) bool {
	// v.(I)
	// I 表示接口类型
	// 类型断言，用于判断一个接口的值实际类型是否为某个类型
	// 或者一个非接口值是否实现了某个接口类型
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

// DeleteArticle 刪除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}
