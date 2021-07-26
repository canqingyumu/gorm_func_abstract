
// DriverInfo [...]
type DriverInfo struct {
	ID          int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`                             // 主键id
	UId         int64     `gorm:"unique_index:uniq_uid_role_channel;column:uid;type:bigint(20) unsigned;not null" json:"uid"`   // 用户id
}


func tableName() string {
	return "driver_info"
}


//所有db 特殊操作的闭包
type Option func() *gorm.DB

func UpdateSelect(db *gorm.DB, fields []string) Option {
	return func() *gorm.DB {
		return db.Select(strings.Join(fields, ","))
	}
}

func UpdateOmit(db *gorm.DB, fields []string) Option {
	return func() *gorm.DB {
		return db.Omit(strings.Join(fields, ","))
	}
}

func OrderBy(db *gorm.DB, field string, sort string) Option {
	return func() *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", field, sort))
	}
}



//更新
func (d *DriverInfo) Updates(ctx context.Context, db *gorm.DB, where map[string]interface{}, updateInfo *DriverInfo, options []Option) error {
	operationName := "dao.DriverInfo.update"
	span, ctx := jaeger.StartSpanFromContext(ctx, operationName)
	defer jaeger.Finish(ctx, span, operationName)

	// init
	if db == nil {
		db = dao.GetDB()
	}

	//表名
	db = db.Table(tableName())

	//各种其他函数
	if len(options) > 0 {
		for _, option := range options {
			db = option()
		}
	}

	//查询条件
	if len(where) > 0 {
		for key, val := range where {
			db = db.Where(key, val)
		}
	}

	//更新数据
	err := db.Updates(updateInfo).Error
	if err != nil {
		log.New().WithContext(ctx).Named(operationName).Error(
			"db update err",
			zap.Error(err),
			zap.Any("where", where),
		)
		return err
	}

	return nil
}

//获取信息
func (d *DriverInfo) First(ctx context.Context, db *gorm.DB, where map[string]interface{}, options []Option) (*DriverInfo, error) {
	operationName := "dao.DriverInfo.first"
	span, ctx := jaeger.StartSpanFromContext(ctx, operationName)
	defer jaeger.Finish(ctx, span, operationName)

	// init
	if db == nil {
		db = dao.GetDB()
	}

	//表名
	db = db.Table(tableName())

	//各种其他函数
	if len(options) > 0 {
		for _, option := range options {
			db = option()
		}
	}

	//查询条件
	if len(where) > 0 {
		for key, val := range where {
			db = db.Where(key, val)
		}
	}

	result := &DriverInfo{}
	//判断是否为空
	if db.First(result).RecordNotFound() {
		return result, nil
	}

	// 只取一条
	err := db.First(result).Error
	if err != nil {
		log.New().WithContext(ctx).Named(operationName).Error(
			"db query first err",
			zap.Error(err),
			zap.Any("where", where),
		)
		return result, err
	}
	return result, nil
}

//插入
func (d *DriverInfo) Insert(ctx context.Context, db *gorm.DB, data *DriverInfo) error {
	operationName := "dao.DriverInfo.insert"
	span, ctx := jaeger.StartSpanFromContext(ctx, operationName)
	defer jaeger.Finish(ctx, span, operationName)

	// init
	if db == nil {
		db = dao.GetDB()
	}

	//表名
	db = db.Table(tableName())

	//insert
	err := db.Create(data).Error
	if err != nil {
		log.New().WithContext(ctx).Named(operationName).Error(
			"db insert err",
			zap.Error(err),
			zap.Any("insert_data", data),
		)
		return err
	}

	return nil
}

//获取列表
func (d *DriverInfo) List(ctx context.Context, db *gorm.DB, where map[string]interface{}, options []Option) ([]DriverInfo, error) {
	operationName := "dao.DriverInfo.get"
	span, ctx := jaeger.StartSpanFromContext(ctx, operationName)
	defer jaeger.Finish(ctx, span, operationName)
	// init
	if db == nil {
		db = dao.GetDB()
	}

	//表名
	db = db.Table(tableName())

	//各种其他函数
	if len(options) > 0 {
		for _, option := range options {
			db = option()
		}
	}

	//查询条件
	if len(where) > 0 {
		for key, val := range where {
			db = db.Where(key, val)
		}
	}

	//是否为空
	result := make([]DriverInfo, 0)
	if db.Find(&result).RecordNotFound() {
		return result, nil
	}

	// 取列表
	err := db.Find(&result).Error
	if err != nil {
		log.New().WithContext(ctx).Named(operationName).Error(
			"db query list err",
			zap.Error(err),
			zap.Any("where", where),
		)
		return result, err
	}
	return result, err
}
