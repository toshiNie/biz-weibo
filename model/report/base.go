package report

type Metrics struct {
	Date                       string  `json:"date"`
	Hour                       int     `json:"hour"`
	Consume                    float64 `json:"consume"`                       // 消耗
	Pv                         int     `json:"pv"`                            // 曝光量
	Ecpm                       float64 `json:"ecpm"`                          // 千次曝光成本
	Bhv                        int     `json:"bhv"`                           // 互动数
	Ctr                        float64 `json:"ctr"`                           // 互动率小数点后4位
	Acpe                       float64 `json:"acpe"`                          // 单次互动成本
	TrafficBhv                 int     `json:"traffic_bhv"`                   // 导流数
	TrafficBhvRate             float64 `json:"traffic_bhv_rate"`              // 导流率小数点后4位
	TrafficBhvCost             float64 `json:"traffic_bhv_cost"`              // 单次导流成本
	Transform                  int     `json:"transform"`                     // 应用转化数
	TransformCost              float64 `json:"transform_cost"`                // 应用转化成本
	Bhv70000001                int     `json:"bhv_70000001"`                  // 激活数
	Bhv70000001Cost            float64 `json:"bhv_70000001_cost"`             // 单次激活成本
	Bhv70000004                int     `json:"bhv_70000004"`                  // 注册数
	Bhv70000004Cost            float64 `json:"bhv_70000004_cost"`             // 注册成本
	FirstTimePay               int     `json:"first_time_pay"`                // 首次付费数
	FirstTimePayCost           float64 `json:"first_time_pay_cost"`           // 首次付费成本
	Bhv70000005                int     `json:"bhv_70000005"`                  // 付费数
	DeepTransferCost           float64 `json:"deep_transfer_cost"`            // 付费金额
	Bhv70000005Cost            float64 `json:"bhv_70000005_cost"`             // 单次付费成本
	TansferCostFirstday        float64 `json:"tansfer_cost_firstday"`         // 首日付费金额
	TansferCostSevenday        float64 `json:"tansfer_cost_sevenday"`         // 七日付费金额
	TansferCostFirstdayRoi     float64 `json:"tansfer_cost_firstday_roi"`     // 首日ROI小数点后4位
	TransferNumberFirstday     int     `json:"transfer_number_firstday"`      // 首日付费次数
	TransferNumberThreeday     int     `json:"transfer_number_threeday"`      // 3日付费次数
	TransferCostThreeday       float64 `json:"transfer_cost_threeday"`        // 3日付费金额
	TransferCostThreedayRoi    float64 `json:"transfer_cost_threeday_roi"`    // 3日ROI小数点后4位
	TansferCostSevendayRoi     float64 `json:"tansfer_cost_sevenday_roi"`     // 七日ROI小数点后4位
	Bhv70000003                int     `json:"bhv_70000003"`                  // 下单购买数
	Bhv70000003Cost            float64 `json:"bhv_70000003_cost"`             // 下单购买成本
	Bhv70000006                int     `json:"bhv_70000006"`                  // 次日留存数
	Bhv70000006Cost            float64 `json:"bhv_70000006_cost"`             // 次日留存成本
	Bhv70000010                int     `json:"bhv_70000010"`                  // 3日留存数
	Bhv70000011                int     `json:"bhv_70000011"`                  // 7日留存数
	Bhv70000006Rcv             int     `json:"bhv_70000006_rcv"`              // 次留回传数
	Bhv70000010Rcv             int     `json:"bhv_70000010_rcv"`              // 3日留存回传数
	Bhv70000011Rcv             int     `json:"bhv_70000011_rcv"`              // 7日留存回传数
	Bhv70000007                int     `json:"bhv_70000007"`                  // 首次登录
	Bhv70000008                int     `json:"bhv_70000008"`                  // 首次创角
	FirstGameAveragePlayTime   int     `json:"first_game_average_play_time"`  // 首次游戏平均时长
	Bhv70000012                int     `json:"bhv_70000012"`                  // app内访问次数
	Bhv70000012Cost            float64 `json:"bhv_70000012_cost"`             // app内访问成本
	SkanPv                     int     `json:"skan_pv,omitempty"`             // SKAN曝光量
	SkanClick                  int     `json:"skan_click,omitempty"`          // SKAN点击数
	Bhv70000013                int     `json:"bhv_70000013,omitempty"`        // SKAN转化数
	BookingTransform           int     `json:"booking_transform"`             // 预约转化数
	BookingTransformCost       float64 `json:"booking_transform_cost"`        // 预约转化成本
	Bhv70000001Pvdate          int     `json:"bhv_70000001_pvdate"`           // 激活数（计费时间）
	Bhv70000001CostPvdate      float64 `json:"bhv_70000001_cost_pvdate"`      // 单次激活成本（计费时间)
	TransferNumberSevenday     int     `json:"transfer_number_sevenday"`      // 7日付费次数（激活时间）
	TransferNumberSevendayCost float64 `json:"transfer_number_sevenday_cost"` // 7日付费成本（激活时间）
	FloatFirstPayNum           int     `json:"float_first_pay_num"`           // 首日首次付费数
	FloatFirstPayNumCost       float64 `json:"float_first_pay_num_cost"`      // 首日首次付费成本
}

type SumMetrics struct {
	Consume float64 `json:"consume"`
	Ctr     float64 `json:"ctr"`
	Date    string  `json:"date"`
	Hour    string  `json:"hour"`
	Pv      int     `json:"pv"`
}

type PageInfo struct {
	Page     int `json:"page"`
	Pages    int `json:"pages"`
	PageSize int `json:"page_size"`
	TotalNum int `json:"total_num"`
}
