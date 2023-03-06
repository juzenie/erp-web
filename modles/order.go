package modles

import (
	sqldriver "database/sql/driver"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

const (
	defaultDateTimeFormat = `"2006-01-02 15:04:05"`
)

type Order struct {
	Acct            string  `gorm:"column:acct"             json:"acct"`             //会计科目
	Acctclose       string  `gorm:"column:acctclose"        json:"acctclose"`        //标识
	Addcost         float64 `gorm:"column:addcost"          json:"addcost"`          //数量（保留十位）
	Addcostx        float64 `gorm:"column:addcostx"         json:"addcostx"`         //数量（保留十位）
	Allowsplit      string  `gorm:"column:allowsplit"       json:"allowsplit"`       //标识
	Assaypercentage float64 `gorm:"column:assaypercentage"  json:"assaypercentage"`  //百分比
	Autoplan        string  `gorm:"column:autoplan"         json:"autoplan"`         //标识
	Baseid          string  `gorm:"column:baseid"           json:"baseid"`           //ID
	Batch           string  `gorm:"column:batch"            json:"batch"`            //批次号
	Bdnup           string  `gorm:"column:bdnup"            json:"bdnup"`            //标识
	Bdnvar          float64 `gorm:"column:bdnvar"           json:"bdnvar"`           //金额（保留十位）
	Bomcode         string  `gorm:"column:bomcode"          json:"bomcode"`          //物料清单代码
	Bookedqty       float64 `gorm:"column:bookedqty"        json:"bookedqty"`        //数量（保留十位）
	Boqty           float64 `gorm:"column:boqty"            json:"boqty"`            //数量（保留十位）
	Erpbuffer       Time    `gorm:"column:erpbuffer"        json:"erpbuffer"`        //日期时间
	Chooseprint     string  `gorm:"column:chooseprint"      json:"chooseprint"`      //标识
	Commitdate      Time    `gorm:"column:commitdate"       json:"commitdate"`       //日期
	Compqty         float64 `gorm:"column:compqty"          json:"compqty"`          //数量（保留十位）
	Costcenter      string  `gorm:"column:costcenter"       json:"costcenter"`       //成本中心
	CreateTime      Time    `gorm:"column:create_time"      json:"create_time"`      //创建日期(系统)
	CreateUser      string  `gorm:"column:create_user"      json:"create_user"`      //创建人(系统)
	Cumlotid        string  `gorm:"column:cumlotid"         json:"cumlotid"`         //ID
	Docid           string  `gorm:"column:docid"            json:"docid"`            //ID
	Erpdomain       string  `gorm:"column:erpdomain"        json:"erpdomain"`        //域
	Dr              string  `gorm:"column:dr"               json:"dr"`               //删除标识(系统)
	Draw            string  `gorm:"column:draw"             json:"draw"`             //图纸
	Dueenddate      string  `gorm:"column:dueenddate"       json:"dueenddate"`       //标识
	Enddate         Time    `gorm:"column:enddate"          json:"enddate"`          //日期
	Engcode         string  `gorm:"column:engcode"          json:"engcode"`          //说明
	Erpprechar1     string  `gorm:"column:erpprechar1"      json:"erpprechar1"`      //预留字段长文本
	Erpprechar2     string  `gorm:"column:erpprechar2"      json:"erpprechar2"`      //预留字段长文本
	Erpprechar3     string  `gorm:"column:erpprechar3"      json:"erpprechar3"`      //预留字段长文本
	Erpprechar4     string  `gorm:"column:erpprechar4"      json:"erpprechar4"`      //预留字段长文本
	Erppredate1     Time    `gorm:"column:erppredate1"      json:"erppredate1"`      //预留日期
	Erppredate2     Time    `gorm:"column:erppredate2"      json:"erppredate2"`      //预留日期
	Erppredec1      float64 `gorm:"column:erppredec1"       json:"erppredec1"`       //预留数值型
	Erppredec2      float64 `gorm:"column:erppredec2"       json:"erppredec2"`       //预留数值型
	Erppredec3      float64 `gorm:"column:erppredec3"       json:"erppredec3"`       //预留数值型
	Erppreint1      int64   `gorm:"column:erppreint1"       json:"erppreint1"`       //预留字段整型
	Erppreint2      int64   `gorm:"column:erppreint2"       json:"erppreint2"`       //预留字段整型
	Expiredate      Time    `gorm:"column:expiredate"       json:"expiredate"`       //日期
	Ficlosedate     Time    `gorm:"column:ficlosedate"      json:"ficlosedate"`      //日期
	Flccostcenter   string  `gorm:"column:flccostcenter"    json:"flccostcenter"`    //成本中心
	Flracct         string  `gorm:"column:flracct"          json:"flracct"`          //会计科目
	Flrsub          string  `gorm:"column:flrsub"           json:"flrsub"`           //细分科目
	Fsmtype         string  `gorm:"column:fsmtype" json:"fsmtype"`                   //现场服务类型
	Glbdnsign       string  `gorm:"column:glbdnsign" json:"glbdnsign"`               //标识
	Gllbrsign       string  `gorm:"column:gllbrsign" json:"gllbrsign"`               //标识
	Grade           string  `gorm:"column:grade" json:"grade"`                       //物料等级
	ID              string  `gorm:"column:id;primary_key" json:"id"`                 //id(系统)
	Isssite         string  `gorm:"column:isssite" json:"isssite"`                   //地点
	Jointtype       string  `gorm:"column:jointtype" json:"jointtype"`               //复合/副产品线型
	Laborcost       float64 `gorm:"column:laborcost" json:"laborcost"`               //数量（保留十位）
	Laborcostx      float64 `gorm:"column:laborcostx" json:"laborcostx"`             //数量（保留十位）
	LastModified    Time    `gorm:"column:last_modified" json:"last_modified"`       //修改时间
	LastModifyUser  string  `gorm:"column:last_modify_user" json:"last_modify_user"` //修改人(系统)
	Lastrcddate     Time    `gorm:"column:lastrcddate" json:"lastrcddate"`           //日期
	Lbrup           string  `gorm:"column:lbrup" json:"lbrup"`                       //标识
	Lbrvar          float64 `gorm:"column:lbrvar" json:"lbrvar"`                     //金额（保留十位）
	Leadtime        int64   `gorm:"column:leadtime" json:"leadtime"`                 //整数
	Linessqe        float64 `gorm:"column:linessqe" json:"linessqe"`                 //数量（保留两位）
	Erplocation     string  `gorm:"column:erplocation" json:"erplocation"`           //库位
	Lotnext         string  `gorm:"column:lotnext" json:"lotnext"`                   //批次号
	Memo            string  `gorm:"column:memo" json:"memo"`                         //备注
	Mixvar          float64 `gorm:"column:mixvar" json:"mixvar"`                     //金额（保留十位）
	Mthdvar         float64 `gorm:"column:mthdvar" json:"mthdvar"`                   //金额（保留十位）
	Mtlcost         float64 `gorm:"column:mtlcost" json:"mtlcost"`                   //数量（保留十位）
	Mtlcostx        float64 `gorm:"column:mtlcostx" json:"mtlcostx"`                 //数量（保留十位）
	Mtlno           string  `gorm:"column:mtlno" json:"mtlno"`                       //物料
	Mtlvar          float64 `gorm:"column:mtlvar" json:"mtlvar"`                     //金额（保留十位）
	Mvaracct        string  `gorm:"column:mvaracct" json:"mvaracct"`                 //会计科目
	Mvarcc          string  `gorm:"column:mvarcc" json:"mvarcc"`                     //成本中心
	Mvarsubacct     string  `gorm:"column:mvarsubacct" json:"mvarsubacct"`           //细分科目
	Mvrracct        string  `gorm:"column:mvrracct" json:"mvrracct"`                 //会计科目
	Mvrrcc          string  `gorm:"column:mvrrcc" json:"mvrrcc"`                     //成本中心
	Mvrrsubacct     string  `gorm:"column:mvrrsubacct" json:"mvrrsubacct"`           //细分科目
	Myldvar         float64 `gorm:"column:myldvar" json:"myldvar"`                   //金额（保留十位）
	Needdate        Time    `gorm:"column:needdate" json:"needdate"`                 //日期
	Needtime        int64   `gorm:"column:needtime" json:"needtime"`                 //整数
	Oldid           string  `gorm:"column:oldid" json:"oldid"`                       //ID
	Orddate         Time    `gorm:"column:orddate" json:"orddate"`                   //日期
	Orderno         string  `gorm:"column:orderno" json:"orderno"`                   //工单
	Orderprint      string  `gorm:"column:orderprint" json:"orderprint"`             //标识
	Ordid           string  `gorm:"column:ordid" json:"ordid"`                       //ID
	Ordqty          float64 `gorm:"column:ordqty" json:"ordqty"`                     //数量（保留十位）
	Ordtype         string  `gorm:"column:ordtype" json:"ordtype"`                   //订单类型
	Erpoutput       float64 `gorm:"column:erpoutput" json:"erpoutput"`               //数量（保留十位）
	Overheadcost    float64 `gorm:"column:overheadcost" json:"overheadcost"`         //数量（保留十位）
	Overheadcostx   float64 `gorm:"column:overheadcostx" json:"overheadcostx"`       //数量（保留十位）
	Perdate         Time    `gorm:"column:perdate" json:"perdate"`                   //日期
	Prechar1        string  `gorm:"column:prechar1" json:"prechar1"`                 //预留字段长文本
	Prechar2        string  `gorm:"column:prechar2" json:"prechar2"`                 //预留字段长文本
	Prechar3        string  `gorm:"column:prechar3" json:"prechar3"`                 //预留字段长文本
	Prechar4        string  `gorm:"column:prechar4" json:"prechar4"`                 //预留字段长文本
	Prechar5        string  `gorm:"column:prechar5" json:"prechar5"`                 //预留字段长文本
	Precompqty      float64 `gorm:"column:precompqty" json:"precompqty"`             //数量（保留十位）
	Predate1        Time    `gorm:"column:predate1" json:"predate1"`                 //预留日期
	Predate2        Time    `gorm:"column:predate2" json:"predate2"`                 //预留日期
	Predec1         float64 `gorm:"column:predec1" json:"predec1"`                   //预留数值型
	Predec2         float64 `gorm:"column:predec2" json:"predec2"`                   //预留数值型
	Predec3         float64 `gorm:"column:predec3" json:"predec3"`                   //预留数值型
	Predec4         float64 `gorm:"column:predec4" json:"predec4"`                   //预留数值型
	Prescrapqty     float64 `gorm:"column:prescrapqty" json:"prescrapqty"`           //数量（保留十位）
	Priority        string  `gorm:"column:priority" json:"priority"`                 //优先级
	Process         string  `gorm:"column:process" json:"process"`                   //处理
	Prodline        string  `gorm:"column:prodline" json:"prodline"`                 //生产线
	Productrate     float64 `gorm:"column:productrate" json:"productrate"`           //百分比
	Projectcode     string  `gorm:"column:projectcode" json:"projectcode"`           //项目代码
	Qtytype         string  `gorm:"column:qtytype" json:"qtytype"`                   //数量类型
	Queuepercentage float64 `gorm:"column:queuepercentage" json:"queuepercentage"`   //百分比
	Receiveqty      float64 `gorm:"column:receiveqty" json:"receiveqty"`             //数量（保留十位）
	Recpstatus      string  `gorm:"column:recpstatus" json:"recpstatus"`             //状态
	Erpreference    string  `gorm:"column:erpreference" json:"erpreference"`         //BOM参考号
	Rejectcost      float64 `gorm:"column:rejectcost" json:"rejectcost"`             //数量（保留十位）
	Reldate         Time    `gorm:"column:reldate" json:"reldate"`                   //日期
	Rev             string  `gorm:"column:rev" json:"rev"`                           //版本
	Rjctqty         float64 `gorm:"column:rjctqty" json:"rjctqty"`                   //数量（保留十位）
	Schedulecode    string  `gorm:"column:schedulecode" json:"schedulecode"`         //计划代码
	Seqno           int64   `gorm:"column:seqno" json:"seqno"`                       //序号
	Serial          string  `gorm:"column:serial" json:"serial"`                     //批次号
	Serno           string  `gorm:"column:serno" json:"serno"`                       //序列号
	Setuptime       int64   `gorm:"column:setuptime" json:"setuptime"`               //整数
	Shift           float64 `gorm:"column:shift" json:"shift"`                       //班次
	Singlesign      string  `gorm:"column:singlesign" json:"singlesign"`             //标识
	Site            string  `gorm:"column:site" json:"site"`                         //地点
	Sojob           string  `gorm:"column:sojob" json:"sojob"`                       //说明
	Staclosedate    Time    `gorm:"column:staclosedate" json:"staclosedate"`         //日期
	Stacloseuserid  string  `gorm:"column:stacloseuserid" json:"stacloseuserid"`     //ID
	Startdate       Time    `gorm:"column:startdate" json:"startdate"`               //日期
	Status          string  `gorm:"column:status" json:"status"`                     //状态
	Subacct         string  `gorm:"column:subacct" json:"subacct"`                   //细分科目
	Subvar          float64 `gorm:"column:subvar" json:"subvar"`                     //金额（保留十位）
	Svaracct        string  `gorm:"column:svaracct" json:"svaracct"`                 //会计科目
	Svarcc          string  `gorm:"column:svarcc" json:"svarcc"`                     //成本中心
	Svarsubacct     string  `gorm:"column:svarsubacct" json:"svarsubacct"`           //细分科目
	Svrracct        string  `gorm:"column:svrracct" json:"svrracct"`                 //会计科目
	Svrrcc          string  `gorm:"column:svrrcc" json:"svrrcc"`                     //成本中心
	Svrrsubacct     string  `gorm:"column:svrrsubacct" json:"svrrsubacct"`           //细分科目
	Techcode        string  `gorm:"column:techcode" json:"techcode"`                 //工艺流程代码
	TenantID        string  `gorm:"column:tenant_id" json:"tenant_id"`               //租户ID(系统)
	Toolcode        string  `gorm:"column:toolcode" json:"toolcode"`                 //工具代码
	Transfercost    float64 `gorm:"column:transfercost" json:"transfercost"`         //数量（保留十位）
	Transfercostx   float64 `gorm:"column:transfercostx" json:"transfercostx"`       //数量（保留十位）
	Unitcost        float64 `gorm:"column:unitcost" json:"unitcost"`                 //数量（保留十位）
	Updatecost      float64 `gorm:"column:updatecost" json:"updatecost"`             //数量（保留十位）
	Userid1         string  `gorm:"column:userid1" json:"userid1"`                   //预留字段长文本
	Userid2         string  `gorm:"column:userid2" json:"userid2"`                   //预留字段长文本
	Valid           string  `gorm:"column:valid" json:"valid"`                       //标识
	Varsign         string  `gorm:"column:varsign" json:"varsign"`                   //标识
	Vendorno        string  `gorm:"column:vendorno" json:"vendorno"`                 //供应商
	Version         int64   `gorm:"column:version" json:"version"`                   //版本(系统)
	Wipcost         float64 `gorm:"column:wipcost" json:"wipcost"`                   //数量（保留十位）
	Xvaracct        string  `gorm:"column:xvaracct" json:"xvaracct"`                 //会计科目
	Xvarcc          string  `gorm:"column:xvarcc" json:"xvarcc"`                     //成本中心
	Xvarsubacct     string  `gorm:"column:xvarsubacct" json:"xvarsubacct"`           //细分科目
	Summrbnumber    float64 `gorm:"column:summrbnumber" json:"summrbnumber"`         //累计已形成领料单数量
	Depcode         string  `gorm:"column:depcode" json:"depcode"`                   //部门
	Amountmcv       float64 `gorm:"column:amountmcv" json:"amountmcv"`               //尾差
}

// TableName sets the insert table name for this struct type
func (o *Order) TableName() string {
	return "pp_sfc_orderheader"
}

//type Time struct{ time.Time }

type Time time.Time

// MarshalJSON on Time format Time field with %Y-%m-%d %H:%M:%S.
func (st Time) MarshalJSON() ([]byte, error) {
	t := time.Time(st)
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	// 注意 `"2006-01-02 15:04:05"`。因为是 JSON，双引号不能少
	return []byte(t.Format(defaultDateTimeFormat)), nil
}

// Value insert timestamp into mysql need this function.
func (st Time) Value() (sqldriver.Value, error) {
	t := time.Time(st)
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t, nil
}

// Scan valueof time.Time.
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		//*t = Time{Time: value}
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
